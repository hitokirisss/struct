package worker

import (
	"errors"

	"github.com/hitokirisss/struct/models"
)

var (
	errWorkerNotFound = errors.New("Worker not found")
	errGroupNotFound  = errors.New("Group not found")
)

func (w *WorkerMemmoryRepo) GetWorkers() []models.Worker {
	return w.workers
}

func (w *WorkerMemmoryRepo) SetNewData(newName string, newSurname string, workerID int) error {
	for i, worker := range w.workers {
		if worker.ID == workerID {
			w.workers[i].Data.SetNewNameAndSurname(newName, newSurname)

			return nil
		}
	}

	return errWorkerNotFound
}

func (w WorkerMemmoryRepo) GetData(workerID int) (string, error) {
	for _, worker := range w.workers {
		if worker.ID == workerID {
			return worker.Data.GetData(), nil
		}
	}

	return "", errWorkerNotFound
}

func (w WorkerMemmoryRepo) GetGroupID(workerID int) ([]string, error) {
	for _, worker := range w.workers {
		if worker.ID == workerID {
			return worker.Groups, nil
		}
	}

	return nil, errWorkerNotFound
}

func (w *WorkerMemmoryRepo) DeleteGroup(groupID string, workerID int) error {
	for i, worker := range w.workers {
		if worker.ID == workerID {
			for j, group := range worker.Groups {
				if group == groupID {
					w.workers[i].Groups = append(w.workers[i].Groups[:j], w.workers[i].Groups[j+1:]...)

					return nil
				}
			}

			return errGroupNotFound
		}
	}

	return errWorkerNotFound
}

func (w *WorkerMemmoryRepo) AddGroup(groupID string, workerID int) error {
	for i, worker := range w.workers {
		if worker.ID == workerID {
			for _, group := range worker.Groups {
				if group == groupID {

					return nil
				}
			}

			w.workers[i].Groups = append(w.workers[i].Groups, groupID)

			return nil
		}
	}

	return errWorkerNotFound
}

func (w WorkerMemmoryRepo) GetWorker(workerID int) (models.Worker, error) {
	for _, worker := range w.workers {
		if worker.ID == workerID {
			return worker, nil
		}
	}

	return models.Worker{}, errWorkerNotFound

}
 