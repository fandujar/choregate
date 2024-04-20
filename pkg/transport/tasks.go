package transport

import (
	"encoding/json"
	"net/http"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/services"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Task represents a task in the API.
type Task struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

// TaskHandler represents the HTTP handler for tasks.
type TaskHandler struct {
	service services.TaskService
}

// NewTaskHandler creates a new instance of TaskHandler.
func NewTaskHandler(service services.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

// GetTasksHandler handles the GET /tasks endpoint.
func (h *TaskHandler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

// GetTaskHandler handles the GET /tasks/{id} endpoint.
func (h *TaskHandler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	taskID := uuid.MustParse(id)
	if taskID == uuid.Nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.service.FindByID(r.Context(), taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// CreateTaskHandler handles the POST /tasks endpoint.
func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskConfig := entities.TaskConfig{
		ID:    task.ID,
		Title: task.Title,
	}

	t, err := entities.NewTask(&taskConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.Create(r.Context(), t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// RunTaskHandler handles the POST /tasks/{id}/run endpoint.
func (h *TaskHandler) RunTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	taskID := uuid.MustParse(id)
	if taskID == uuid.Nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	err := h.service.Run(r.Context(), taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// RegisterTasksRoutes registers the routes for the tasks API.
func RegisterTasksRoutes(r chi.Router, service services.TaskService) {
	handler := NewTaskHandler(service)

	r.Get("/tasks", handler.GetTasksHandler)
	r.Get("/tasks/{id}", handler.GetTaskHandler)
	r.Post("/tasks/{id}/run", handler.RunTaskHandler)
	r.Post("/tasks", handler.CreateTaskHandler)
}
