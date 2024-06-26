package student

import (
	"errors"
	"fmt"
)

var (
	errStudentNotFound = errors.New("Student not found")
	errGroupNotFound   = errors.New("Group not found")
	errTaskNotFound    = errors.New("Task not found")
)

type StudentMemmoryRepo struct {
	students []Student
}

func New() *StudentMemmoryRepo {
	return &StudentMemmoryRepo{
		students: GetStudents(),
	}
}

func (s *Student) SetNewData(newName string, newSurname string) {
	s.data.SetNewNameAndSurname(newName, newSurname)
}

func (s *Student) SetgroupID(newGroupID string) {
	s.groupID = newGroupID
}

func (s *Student) GetData() string {
	return s.data.GetData()
}

func (s *Student) GetGroupID() string {
	return s.groupID
}

func (s *Student) SubmitTask(taskID int) error {
	flag := false

	for i, task := range s.Tasks {
		if task.ID == taskID {
			s.Tasks[i].Submit()
			flag = true
		}
	}

	if !flag {
		return fmt.Errorf("task with id %d not found", taskID)
	}

	return nil
}
