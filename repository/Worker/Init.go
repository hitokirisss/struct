package worker

import "github.com/hitokirisss/struct/models"

type WorkerMemmoryRepo struct {
	workers []models.Worker
}

func New() *WorkerMemmoryRepo {
	return &WorkerMemmoryRepo{
		workers: models.GetWorkers(),
	}
}
