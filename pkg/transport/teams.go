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

// Team represents a team in the API.
type Team struct {
	ID      string        `json:"id"`
	Name    string        `json:"name"`
	Members []*TeamMember `json:"members"`
}

type TeamMember struct {
	ID     string `json:"id"`
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type TeamsHandler struct {
	service services.TeamService
}

func RegisterTeamsRoutes(r chi.Router, service services.TeamService) {
	handler := &TeamsHandler{service: service}
	roles := rbac.SetupRoles()

	r.Route("/teams", func(r chi.Router) {
		// GET /teams
		r.Group(func(r chi.Router) {
			r.Use(rbac.PermissionInjectorMiddleware(rbac.Permission{Scope: "teams", Action: "read"}))
			r.Use(rbac.RBAC(roles))
			r.Get("/", handler.GetTeamsHandler)
			r.Get("/{id}", handler.GetTeamHandler)
		})

		// POST /teams
		r.Group(func(r chi.Router) {
			r.Use(rbac.PermissionInjectorMiddleware(rbac.Permission{Scope: "teams", Action: "create"}))
			r.Use(rbac.RBAC(roles))
			r.Post("/", handler.CreateTeamHandler)
		})

		// PUT /teams
		r.Group(func(r chi.Router) {
			r.Use(rbac.PermissionInjectorMiddleware(rbac.Permission{Scope: "teams", Action: "update"}))
			r.Use(rbac.RBAC(roles))
			r.Put("/{id}", handler.UpdateTeamHandler)
		})

		// DELETE /teams
		r.Group(func(r chi.Router) {
			r.Use(rbac.PermissionInjectorMiddleware(rbac.Permission{Scope: "teams", Action: "delete"}))
			r.Use(rbac.RBAC(roles))
			r.Delete("/{id}", handler.DeleteTeamHandler)
		})
	})
}

func (h *TeamsHandler) GetTeamsHandler(w http.ResponseWriter, r *http.Request) {
	teams, err := h.service.GetTeams(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(teams)
}

func (h *TeamsHandler) GetTeamHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	teamID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	team, err := h.service.GetTeam(r.Context(), teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(team)
}

func (h *TeamsHandler) CreateTeamHandler(w http.ResponseWriter, r *http.Request) {
	var t Team
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	teamConfig := &entities.TeamConfig{
		Name: t.Name,
	}

	team, err := entities.NewTeam(teamConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.CreateTeam(r.Context(), team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *TeamsHandler) UpdateTeamHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	teamID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var t Team
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	teamConfig := &entities.TeamConfig{
		ID:   teamID,
		Name: t.Name,
	}

	team, err := entities.NewTeam(teamConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.UpdateTeam(r.Context(), team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *TeamsHandler) DeleteTeamHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	teamID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTeam(r.Context(), teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
