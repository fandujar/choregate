package transport

import (
	"encoding/json"
	"net/http"

	"github.com/fandujar/choregate/pkg/entities"
	"github.com/fandujar/choregate/pkg/rbac"
	"github.com/fandujar/choregate/pkg/services"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

// Task represents a task in the API.
type Task struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
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
	ctx := r.Context()
	scope := h.service.GetTaskScopeFromContext(ctx)

	tasks, err := h.service.FindAll(r.Context(), scope)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

// GetTaskHandler handles the GET /tasks/{id} endpoint.
func (h *TaskHandler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	scope := h.service.GetTaskScopeFromContext(r.Context())
	if scope == nil {
		http.Error(w, "invalid task scope", http.StatusBadRequest)
		return
	}

	taskID, err := uuid.Parse(id)
	if taskID == uuid.Nil || err != nil {
		log.Error().Err(err).Msgf("invalid task ID: %s", id)
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.service.FindByID(r.Context(), taskID, scope)
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

	ctx := r.Context()
	scope := h.service.GetTaskScopeFromContext(ctx)
	if scope == nil {
		http.Error(w, "invalid task scope", http.StatusBadRequest)
		return
	}

	taskConfig := entities.TaskConfig{
		ID:   task.ID,
		Name: task.Name,
		TaskScope: &entities.TaskScope{
			Owner: scope.Owner,
		},
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
	scope := h.service.GetTaskScopeFromContext(r.Context())
	if scope == nil {
		http.Error(w, "invalid task scope", http.StatusBadRequest)
		return
	}

	taskID := uuid.MustParse(id)
	if taskID == uuid.Nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	var taskRunID uuid.UUID

	if runID == "" {
		taskRunID = uuid.Nil
	} else {
		taskRunID = uuid.MustParse(runID)
		if taskRunID == uuid.Nil {
			http.Error(w, "invalid task run ID", http.StatusBadRequest)
			return
		}
	}

	err := h.service.Run(r.Context(), taskID, taskRunID, scope)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateStepsHandler handles the PUT /tasks/{id}/steps endpoint.
func (h *TaskHandler) UpdateStepsHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	scope := h.service.GetTaskScopeFromContext(r.Context())
	if scope == nil {
		http.Error(w, "invalid task scope", http.StatusBadRequest)
		return
	}

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

	task, err := h.service.FindByID(r.Context(), taskID, scope)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	task.TaskSpec.Steps = steps
	log.Debug().Msgf("updating task %s with steps %v", taskID, steps)
	err = h.service.Update(r.Context(), task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetTaskStepsHandler handles the GET /tasks/{id}/steps endpoint.
func (h *TaskHandler) GetTaskStepsHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	scope := h.service.GetTaskScopeFromContext(r.Context())
	if scope == nil {
		http.Error(w, "invalid task scope", http.StatusBadRequest)
		return
	}

	taskID := uuid.MustParse(id)
	if taskID == uuid.Nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.service.FindByID(r.Context(), taskID, scope)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if task.TaskSpec.Steps == nil {
		task.TaskSpec.Steps = []tekton.Step{}
	}

	json.NewEncoder(w).Encode(task.TaskSpec.Steps)
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

	if taskRuns == nil {
		taskRuns = []*entities.TaskRun{}
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
		taskRunID = uuid.MustParse(runID)
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

// GetTaskRunLogsHandler handles the GET /tasks/{id}/runs/{runID}/logs endpoint.
func (h *TaskHandler) GetTaskRunLogsHandler(w http.ResponseWriter, r *http.Request) {
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
		taskRunID = uuid.MustParse(runID)
		if taskRunID == uuid.Nil {
			http.Error(w, "invalid task run ID", http.StatusBadRequest)
			return
		}
	}

	logs, err := h.service.FindTaskRunLogs(r.Context(), taskID, taskRunID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(logs)
}

// GetTaskRunStatusHandler handles the GET /tasks/{id}/runs/{runID}/status endpoint.
func (h *TaskHandler) GetTaskRunStatusHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	runID := chi.URLParam(r, "runID")

	taskID := uuid.MustParse(id)
	if taskID == uuid.Nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	taskRunID := uuid.MustParse(runID)
	if taskRunID == uuid.Nil {
		http.Error(w, "invalid task run ID", http.StatusBadRequest)
		return
	}

	status, err := h.service.FindTaskRunStatus(r.Context(), taskID, taskRunID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(status)
}

// FakeHandler is a placeholder for future use
func (h *TaskHandler) FakeHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

// RegisterTasksRoutes registers the routes for the tasks API.
func RegisterTasksRoutes(r chi.Router, service services.TaskService) {
	handler := NewTaskHandler(service)
	roles := rbac.SetupRoles()

	r.Route("/tasks", func(r chi.Router) {
		// GET /tasks
		r.Group(func(r chi.Router) {
			r.Use(rbac.PermissionInjectorMiddleware(rbac.Permission{Scope: "tasks", Action: "read"}))
			r.Use(rbac.RBAC(roles))
			r.Get("/", handler.GetTasksHandler)
			r.Get("/{id}", handler.GetTaskHandler)
			r.Get("/{id}/steps", handler.GetTaskStepsHandler)
			r.Get("/{id}/runs", handler.GetTaskRunsHandler)
			r.Get("/{id}/runs/{runID}", handler.GetTaskRunHandler)
			r.Get("/{id}/runs/{runID}/logs", handler.GetTaskRunLogsHandler)
			r.Get("/{id}/runs/{runID}/status", handler.GetTaskRunStatusHandler)
		})

		// POST /tasks
		r.Group(func(r chi.Router) {
			r.Use(rbac.PermissionInjectorMiddleware(rbac.Permission{Scope: "tasks", Action: "create"}))
			r.Use(rbac.RBAC(roles))
			r.Post("/", handler.CreateTaskHandler)
			r.Post("/{id}/runs", handler.RunTaskHandler)
		})

		// PUT /tasks
		r.Group(func(r chi.Router) {
			r.Use(rbac.PermissionInjectorMiddleware(rbac.Permission{Scope: "tasks", Action: "update"}))
			r.Use(rbac.RBAC(roles))
			r.Put("/{id}/steps", handler.UpdateStepsHandler)
		})
	})

	// TODO: implement task params handler
	// r.Put("/tasks/{id}/params", handler.FakeHandler)

}
