package student

import (
	"errors"
	"fmt"

	"github.com/hitokirisss/struct/models"
)

var (
	errStudentNotFound = errors.New("Student not found")
)

func (s *StudentMemmoryRepo) SetNewData(newName string, newSurname string, studentID int) error {
	for i, student := range s.students {
		if student.ID == studentID {
			s.students[i].Data.SetNewNameAndSurname(newName, newSurname)

			return nil
		}
	}

	return errStudentNotFound
}

func (s *StudentMemmoryRepo) SetGroupID(newGroupID string, studentID int) error {
	for i, student := range s.students {
		if studentID == student.ID {
			s.students[i].GroupID = newGroupID

			return nil
		}
	}

	return errStudentNotFound
}

func (s *StudentMemmoryRepo) GetData(studentID int) (string, error) {
	for _, student := range s.students {
		if student.ID == studentID {
			return student.Data.GetData(), nil
		}
	}

	return "", errStudentNotFound
}

func (s *StudentMemmoryRepo) GetGroupID(studentID int) (string, error) {
	for _, student := range s.students {
		if student.ID == studentID {
			return student.GroupID, nil
		}
	}

	return "", errStudentNotFound
}

func (s *StudentMemmoryRepo) SubmitTask(taskID int, studentID int) error {
	for i, student := range s.students {
		if student.ID == studentID {
			flag := false

			for j, task := range student.Tasks {
				if task.ID == taskID {
					s.students[i].Tasks[j].Submit()
					flag = true
				}
			}

			if !flag {
				return fmt.Errorf("task with id %d not found", taskID)
			}

			return nil
		}
	}

	return errStudentNotFound
}

func (s *StudentMemmoryRepo) GetStudents() []models.Student {
	return s.students
}