package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/hitokirisss/struct/Handlers"
	"github.com/hitokirisss/struct/models"
	worker "github.com/hitokirisss/struct/models/Worker"
)

func main() {
	workerRepo := worker.New()
	workerHandler := handlers.NewWorkerHandler(workerRepo)

	r := mux.NewRouter()
	r.HandleFunc("/get-workers", workerHandler.GetWorkers).Methods("GET")
	
	// r.HandleFunc("/get-students")

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", r)
}

func Submit(groupID string, taskID int, students []models.Student) error {
	for i, student := range students {
		if student.GetGroupID() == groupID {
			err := students[i].SubmitTask(taskID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
