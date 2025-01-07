package handlers

import (
	"Project1/internal/taskService" // Импортируем наш сервис
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service *taskService.TaskService
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}

func (h *Handler) PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID задачи из параметров URL
	id := mux.Vars(r)["id"]

	// Преобразуем строку ID в int
	idTask, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Декодируем обновленную задачу
	var updatedTask taskService.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Обновляем задачу и возвращаем результат
	task, err := h.Service.UpdateTaskByID(idTask, updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID задачи из параметров URL
	id := mux.Vars(r)["id"]

	// Преобразуем строку ID в int
	idTask, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Удаляем задачу
	if err := h.Service.DeleteTaskByID(idTask); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Возвращаем код 204
	w.WriteHeader(http.StatusNoContent)
}
