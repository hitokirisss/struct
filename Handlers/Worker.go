package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hitokirisss/struct/models"
)

type WorkerHandler struct {
	repo WorkerRepo
}

type WorkerRepo interface {
	GetWorkers() []models.Worker
	SetNewData(newName string, newSurname string, workerID int) error
	GetData(workerID int) (string, error)
	GetGroupID(workerID int) ([]string, error)
	DeleteGroup(groupID string, workerID int) error
	AddGroup(groupID string, workerID int) error
	GetWorker(workerID int) (models.Worker, error)
}

func NewWorkerHandler(repo WorkerRepo) *WorkerHandler {
	return &WorkerHandler{
		repo: repo,
	}
}

func (handler *WorkerHandler) GetWorkers(w http.ResponseWriter, r *http.Request) {
	workers := handler.repo.GetWorkers()
	res, err := json.MarshalIndent(workers, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	w.Write(res)
}

func (handler *WorkerHandler) GetWorker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	workerID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	worker, err := handler.repo.GetWorker(workerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	res, err := json.MarshalIndent(worker, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	w.Write(res)
}
