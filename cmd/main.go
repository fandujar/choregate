package main

import (
	"net/http"

	"github.com/fandujar/choregate/pkg/repositories/memory"
	"github.com/fandujar/choregate/pkg/services"
	"github.com/fandujar/choregate/pkg/transport"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	taskRepository := memory.NewInMemoryTaskRepository()
	taskService := services.NewTaskService(taskRepository)
	transport.RegisterTasksRoutes(r, *taskService)

	http.ListenAndServe(":8080", r)
}
