package transport

import (
	"encoding/json"
	"net/http"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/services"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type TriggerHandler struct {
	service services.TriggerService
}

// GetTriggersHandler handles the GET /triggers endpoint.
func (h *TriggerHandler) GetTriggersHandler(w http.ResponseWriter, r *http.Request) {
	triggers, err := h.service.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(triggers)
}

// GetTriggerHandler handles the GET /triggers/{id} endpoint.
func (h *TriggerHandler) GetTriggerHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	triggerID := uuid.MustParse(id)
	if triggerID == uuid.Nil {
		http.Error(w, "invalid trigger ID", http.StatusBadRequest)
		return
	}

	trigger, err := h.service.FindByID(r.Context(), triggerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(trigger)
}

// CreateTriggerHandler handles the POST /triggers endpoint.
func (h *TriggerHandler) CreateTriggerHandler(w http.ResponseWriter, r *http.Request) {
	var trigger *entities.Trigger
	if err := json.NewDecoder(r.Body).Decode(&trigger); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Create(r.Context(), trigger); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// NewTriggerHandler creates a new instance of TaskHandler.
func NewTriggerHandler(service services.TriggerService) *TriggerHandler {
	return &TriggerHandler{
		service: service,
	}
}

// RegisterTriggerRoutes registers the trigger routes with the router.
func RegisterTriggersRoutes(r chi.Router, service services.TriggerService) {
	handler := NewTriggerHandler(service)

	r.Get("/triggers", handler.GetTriggersHandler)
	r.Get("/triggers/{id}", handler.GetTriggerHandler)
	r.Post("/triggers", handler.CreateTriggerHandler)
}
