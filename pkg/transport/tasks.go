package transport

import (
	"encoding/json"
	"net/http"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/services"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
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

	taskID, err := uuid.Parse(id)
	if taskID == uuid.Nil || err != nil {
		log.Error().Err(err).Msgf("invalid task ID: %s", id)
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
	runID := chi.URLParam(r, "taskRunID")

	taskID := uuid.MustParse(id)
	if taskID == uuid.Nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	var taskRunID uuid.UUID

	if runID == "" {
		taskRunID = uuid.Nil
	} else {
		taskRunID := uuid.MustParse(runID)
		if taskRunID == uuid.Nil {
			http.Error(w, "invalid task run ID", http.StatusBadRequest)
			return
		}
	}

	err := h.service.Run(r.Context(), taskID, taskRunID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateStepsHandler handles the PUT /tasks/{id}/steps endpoint.
func (h *TaskHandler) UpdateStepsHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	taskID := uuid.MustParse(id)
	if taskID == uuid.Nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	var steps []tekton.Step
	err := json.NewDecoder(r.Body).Decode(&steps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.service.FindByID(r.Context(), taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	task.Steps = steps
	log.Debug().Msgf("updating task %s with steps %v", taskID, steps)
	err = h.service.Update(r.Context(), task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetTaskRunsHandler handles the GET /tasks/{id}/runs endpoint.
func (h *TaskHandler) GetTaskRunsHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	taskID := uuid.MustParse(id)
	if taskID == uuid.Nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	taskRuns, err := h.service.FindTaskRuns(r.Context(), taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(taskRuns)
}

// GetTaskRunHandler handles the GET /tasks/{id}/runs/{runID} endpoint.
func (h *TaskHandler) GetTaskRunHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	runID := chi.URLParam(r, "runID")

	taskID := uuid.MustParse(id)
	if taskID == uuid.Nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	var taskRunID uuid.UUID

	if runID == "" {
		taskRunID = uuid.Nil
	} else {
		taskRunID := uuid.MustParse(runID)
		if taskRunID == uuid.Nil {
			http.Error(w, "invalid task run ID", http.StatusBadRequest)
			return
		}
	}

	taskRun, err := h.service.FindTaskRunByID(r.Context(), taskID, taskRunID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(taskRun)
}

// FakeHandler is a placeholder for future use
func (h *TaskHandler) FakeHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

// RegisterTasksRoutes registers the routes for the tasks API.
func RegisterTasksRoutes(r chi.Router, service services.TaskService) {
	handler := NewTaskHandler(service)

	r.Get("/tasks", handler.GetTasksHandler)
	r.Post("/tasks", handler.CreateTaskHandler)

	r.Get("/tasks/{id}", handler.GetTaskHandler)
	r.Post("/tasks/{id}/runs", handler.RunTaskHandler)
	r.Get("/tasks/{id}/runs", handler.GetTaskRunsHandler)
	r.Put("/tasks/{id}/steps", handler.UpdateStepsHandler)
	// TODO: implement task params handler
	r.Put("/tasks/{id}/params", handler.FakeHandler)

	r.Get("/tasks/{id}/runs/{runID}", handler.GetTaskRunHandler)
	r.Post("/tasks/{id}/runs/{runID}/retry", handler.RunTaskHandler)
	// TODO: implement task run logs handler
	r.Get("/tasks/{id}/runs/{runID}/logs", handler.FakeHandler)
}
