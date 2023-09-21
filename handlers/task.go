package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unawaretub86/task-rest/data"
	"github.com/unawaretub86/task-rest/entities"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	setJSONResponse(w, http.StatusOK, "Pong")
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	setJSONResponse(w, http.StatusOK, data.TasksData)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task entities.Task

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		setJSONResponse(w, http.StatusBadRequest, "Insert valid Task")
		return
	}

	json.Unmarshal(reqBody, &task)

	task.ID = len(data.TasksData) + 1
	data.TasksData = append(data.TasksData, task)

	setJSONResponse(w, http.StatusCreated, task)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		setJSONResponse(w, http.StatusNotFound, "Invalid ID")
		return
	}

	for _, task := range data.TasksData {
		if task.ID == id {
			setJSONResponse(w, http.StatusOK, task)
			return
		}
	}

	setJSONResponse(w, http.StatusNotFound, "Task not found")
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		setJSONResponse(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	for i, task := range data.TasksData {
		if task.ID == id {
			data.TasksData = append(data.TasksData[:i], data.TasksData[i+1:]...)
			setJSONResponse(w, http.StatusOK, "Task removed successfully")
			return
		}
	}

	setJSONResponse(w, http.StatusNotFound, "Task not found")
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		setJSONResponse(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var updatedTask entities.Task

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		setJSONResponse(w, http.StatusBadRequest, "Error reading request body")
		return
	}

	json.Unmarshal(reqBody, &updatedTask)

	for i, task := range data.TasksData {
		if task.ID == id {
			data.TasksData = append(data.TasksData[:i], data.TasksData[i+1:]...)
			updatedTask.ID = id
			data.TasksData = append(data.TasksData, updatedTask)
			setJSONResponse(w, http.StatusOK, updatedTask)
			return
		}
	}

	setJSONResponse(w, http.StatusNotFound, "Task not found")
}

func setJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
