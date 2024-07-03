package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hitokirisss/struct/models"
)

type StudentHandler struct {
	repo StudentRepo
}

type StudentRepo interface {
	SetNewData(newName string, newSurname string, studentID int) error
	SetGroupID(newGroupID string, studentID int) error
	GetData(studentID int) (string, error)
	GetGroupID(studentID int) (string, error)
	SubmitTask(taskID int, studentID int) error
	GetStudents() []models.Student
	GetStudent(studentID int) (models.Student, error)
}

func NewStudentHandler(repo StudentRepo) *StudentHandler {
	return &StudentHandler{
		repo: repo,
	}
}

func (handler *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request) {
	student := handler.repo.GetStudents()
	res, err := json.MarshalIndent(student, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	w.Write(res)

}

func (handler *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	studentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	student, err := handler.repo.GetStudent(studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	res, err := json.MarshalIndent(student, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	w.Write(res)
}
