package rbac

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

type Permission struct {
	Scope  string
	Action string
}

type Role struct {
	Name        string
	Permissions []Permission
}

func SetupRoles() []Role {
	return []Role{
		{
			Name: "admin",
			Permissions: []Permission{
				{Scope: "users", Action: "create"},
				{Scope: "users", Action: "read"},
				{Scope: "users", Action: "update"},
				{Scope: "users", Action: "delete"},

				{Scope: "teams", Action: "create"},
				{Scope: "teams", Action: "read"},
				{Scope: "teams", Action: "update"},
				{Scope: "teams", Action: "delete"},

				{Scope: "organizations", Action: "create"},
				{Scope: "organizations", Action: "read"},
				{Scope: "organizations", Action: "update"},
				{Scope: "organizations", Action: "delete"},

				{Scope: "tasks", Action: "create"},
				{Scope: "tasks", Action: "read"},
				{Scope: "tasks", Action: "update"},
				{Scope: "tasks", Action: "delete"},
			},
		},
		{
			Name: "user",
			Permissions: []Permission{
				{Scope: "tasks", Action: "read"},
				{Scope: "tasks", Action: "create"},
				{Scope: "tasks", Action: "update"},
				{Scope: "tasks", Action: "delete"},

				{Scope: "teams", Action: "read"},
				{Scope: "organizations", Action: "read"},
			},
		},
		{
			Name: "viewer",
			Permissions: []Permission{
				{Scope: "tasks", Action: "read"},
				{Scope: "teams", Action: "read"},
				{Scope: "organizations", Action: "read"},
			},
		},
	}
}

func (r *Role) HasPermission(p Permission) bool {
	for _, perm := range r.Permissions {
		if perm.Scope == p.Scope && perm.Action == p.Action {
			return true
		}
	}
	return false
}

func (r *Role) HasPermissions(perms []Permission) bool {
	for _, perm := range perms {
		if !r.HasPermission(perm) {
			return false
		}
	}
	return true
}

func FindRole(roles []Role, roleName string) *Role {
	for _, role := range roles {
		if role.Name == roleName {
			return &role
		}
	}
	return nil
}

func PermissionFromContext(ctx context.Context) (Permission, error) {
	perm, ok := ctx.Value("permission").(Permission)
	if !ok {
		return Permission{}, errors.New("missing permission in context")
	}
	return perm, nil
}

// RBAC is a middleware that checks if the user has the required permission
func RBAC(roles []Role) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			permission, err := PermissionFromContext(ctx)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			token, tokenMap, err := jwtauth.FromContext(ctx)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if token == nil {
				http.Error(w, "missing token", http.StatusBadRequest)
				return
			}

			systemRole, ok := tokenMap["system_role"].(string)
			if !ok {
				http.Error(w, "missing system_role", http.StatusBadRequest)
				return
			}

			role := FindRole(roles, systemRole)
			if role == nil {
				http.Error(w, "invalid role", http.StatusForbidden)
				return
			}

			if !role.HasPermission(permission) {
				http.Error(w, "permission denied", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// PermissionInjectorMiddleware injects the permission into the context
func PermissionInjectorMiddleware(permission Permission) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "permission", permission)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
