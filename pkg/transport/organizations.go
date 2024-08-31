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

type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type OrganizationMember struct {
	ID     string `json:"id"`
	OrgID  string `json:"org_id"`
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type OrganizationTeam struct {
	ID     string `json:"id"`
	OrgID  string `json:"org_id"`
	TeamID string `json:"team_id"`
}

type OrganizationsHandler struct {
	service services.OrganizationService
}

func RegisterOrganizationsRoutes(r chi.Router, service services.OrganizationService) {
	handler := &OrganizationsHandler{service: service}
	roles := rbac.SetupRoles()

	r.Route("/organizations", func(r chi.Router) {
		// GET /organizations
		r.Group(func(r chi.Router) {
			r.Use(rbac.PermissionInjectorMiddleware(rbac.Permission{Scope: "organizations", Action: "read"}))
			r.Use(rbac.RBAC(roles))
			r.Get("/", handler.GetOrganizationsHandler)
			r.Get("/{id}", handler.GetOrganizationHandler)
		})

		// POST /organizations
		r.Group(func(r chi.Router) {
			r.Use(rbac.PermissionInjectorMiddleware(rbac.Permission{Scope: "organizations", Action: "create"}))
			r.Use(rbac.RBAC(roles))
			r.Post("/", handler.CreateOrganizationHandler)
		})

		// PUT /organizations
		r.Group(func(r chi.Router) {
			r.Use(rbac.PermissionInjectorMiddleware(rbac.Permission{Scope: "organizations", Action: "update"}))
			r.Use(rbac.RBAC(roles))
			r.Put("/{id}", handler.UpdateOrganizationHandler)
		})

		// DELETE /organizations
		r.Group(func(r chi.Router) {
			r.Use(rbac.PermissionInjectorMiddleware(rbac.Permission{Scope: "organizations", Action: "delete"}))
			r.Use(rbac.RBAC(roles))
			r.Delete("/{id}", handler.DeleteOrganizationHandler)
		})
	})
}

func (h *OrganizationsHandler) GetOrganizationsHandler(w http.ResponseWriter, r *http.Request) {
	organizations, err := h.service.GetOrganizations(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(organizations)
}

func (h *OrganizationsHandler) GetOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	organizationID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	organization, err := h.service.GetOrganization(r.Context(), organizationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(organization)
}

func (h *OrganizationsHandler) CreateOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	var org Organization
	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	organizationConfig := &entities.OrganizationConfig{
		Name: org.Name,
	}

	organization, err := entities.NewOrganization(organizationConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.service.CreateOrganization(r.Context(), organization); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *OrganizationsHandler) UpdateOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	organizationID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var org Organization
	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	organizationConfig := &entities.OrganizationConfig{
		ID:   organizationID,
		Name: org.Name,
	}

	organization, err := entities.NewOrganization(organizationConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.service.UpdateOrganization(r.Context(), organization); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *OrganizationsHandler) DeleteOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	organizationID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteOrganization(r.Context(), organizationID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
