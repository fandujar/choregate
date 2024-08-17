package transport

import (
	"encoding/json"
	"net/http"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/services"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// User represents a user in the API.
type User struct {
	ID       uuid.UUID `json:"id"`
	Slug     string    `json:"slug"`
	Name     string    `json:"name"`
	Roles    []string  `json:"roles"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type UsersHandler struct {
	service services.UserService
}

func RegisterUsersRoutes(r chi.Router, service services.UserService) {
	handler := &UsersHandler{service: service}

	r.Get("/users", handler.GetUsersHandler)
	r.Get("/users/{id}", handler.GetUserHandler)
	r.Post("/users", handler.CreateUserHandler)
	r.Put("/users/{id}", handler.UpdateUserHandler)
	r.Delete("/users/{id}", handler.DeleteUserHandler)
}

// GetUsersHandler handles the GET /users endpoint.
func (h *UsersHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.FindAll(r.Context())
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
		ID:       u.ID,
		Slug:     u.Slug,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
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

	uid, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &entities.User{}
	err = json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = uid

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