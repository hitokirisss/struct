package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hitokirisss/struct/models/Worker"
)

type WorkerHandler struct {
	repo WorkerRepo
}

type WorkerRepo interface {
	GetWorkers() []worker.Worker
	SetNewData(newName string, newSurname string, workerID int) error
	GetData(workerID int) (string, error)
	GetGroupID(workerID int) ([]string, error)
	DeleteGroup(groupID string, workerID int) error
	AddGroup(groupID string, workerID int) error
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