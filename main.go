package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/hitokirisss/struct/Handlers"
	student "github.com/hitokirisss/struct/repository/Student"
	worker "github.com/hitokirisss/struct/repository/Worker"
)

func main() {
	workerRepo := worker.New()
	workerHandler := handlers.NewWorkerHandler(workerRepo)

	studentRepo := student.New()
	studentHandler := handlers.NewStudentHandler(studentRepo)

	r := mux.NewRouter()

	r.HandleFunc("/get-workers", workerHandler.GetWorkers).Methods("GET")	

	r.HandleFunc("/get-students", studentHandler.GetStudents).Methods("GET")

	r.HandleFunc("/get-student/{id}", studentHandler.GetStudent).Methods("GET")

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", r)
}
