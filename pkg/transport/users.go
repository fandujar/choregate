package transport

import (
	"encoding/json"
	"net/http"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/rbac"
	"github.com/fandujar/choregate/pkg/services"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// User represents a user in the API.
type User struct {
	ID         uuid.UUID `json:"id"`
	Slug       string    `json:"slug"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	SystemRole string    `json:"system_role"`
	Password   string    `json:"password"`
}

type UsersHandler struct {
	service services.OrganizationService
}

func RegisterUsersRoutes(r chi.Router, service services.OrganizationService) {
	handler := &UsersHandler{service: service}
	roles := rbac.SetupRoles()

	r.Route("/users", func(r chi.Router) {
		// GET /users
		r.Group(func(r chi.Router) {
			r.Use(rbac.PermissionInjectorMiddleware(rbac.Permission{Scope: "users", Action: "read"}))
			r.Use(rbac.RBAC(roles))
			r.Get("/", handler.GetUsersHandler)
			r.Get("/{id}", handler.GetUserHandler)
		})

		// POST /users
		r.Group(func(r chi.Router) {
			r.Use(rbac.PermissionInjectorMiddleware(rbac.Permission{Scope: "users", Action: "create"}))
			r.Use(rbac.RBAC(roles))
			r.Post("/users", handler.CreateUserHandler)
		})

		// PUT /users
		r.Group(func(r chi.Router) {
			r.Use(rbac.PermissionInjectorMiddleware(rbac.Permission{Scope: "users", Action: "update"}))
			r.Use(rbac.RBAC(roles))
			r.Put("/users/{id}", handler.UpdateUserHandler)
		})

		// DELETE /users
		r.Group(func(r chi.Router) {
			r.Use(rbac.PermissionInjectorMiddleware(rbac.Permission{Scope: "users", Action: "delete"}))
			r.Use(rbac.RBAC(roles))
			r.Delete("/users/{id}", handler.DeleteUserHandler)
		})
	})
}

// GetUsersHandler handles the GET /users endpoint.
func (h *UsersHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// GetUserHandler handles the GET /users/{id} endpoint.
func (h *UsersHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	uid, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(r.Context(), uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// CreateUserHandler handles the POST /users endpoint.
func (h *UsersHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userConfig := &entities.UserConfig{
		ID:         u.ID,
		Slug:       u.Slug,
		Name:       u.Name,
		Email:      u.Email,
		SystemRole: u.SystemRole,
		Password:   u.Password,
	}

	user, err := entities.NewUser(userConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateUserHandler handles the PUT /users/{id} endpoint.
func (h *UsersHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	userID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var u User
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userConfig := &entities.UserConfig{
		ID:         userID,
		Slug:       u.Slug,
		Name:       u.Name,
		Email:      u.Email,
		SystemRole: u.SystemRole,
		Password:   u.Password,
	}

	user, err := entities.NewUser(userConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.UpdateUser(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUserHandler handles the DELETE /users/{id} endpoint.
func (h *UsersHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	uid, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteUser(r.Context(), uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
