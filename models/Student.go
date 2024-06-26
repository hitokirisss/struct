package models

import (
	"fmt"

	"github.com/hitokirisss/struct/models/human"
)

// type Student struct {
// 	data    human.Human
// 	groupID string
// 	ID      int
// 	Tasks   []Task
// }

// func (s *Student) SetNewData(newName string, newSurname string) {
// 	s.data.SetNewNameAndSurname(newName, newSurname)
// }

// func (s *Student) SetgroupID(newGroupID string) {
// 	s.groupID = newGroupID
// }

// func (s *Student) GetData() string {
// 	return s.data.GetData()
// }

// func (s *Student) GetGroupID() string {
// 	return s.groupID
// }

// func (s *Student) SubmitTask(taskID int) error {
// 	flag := false

// 	for i, task := range s.Tasks {
// 		if task.ID == taskID {
// 			s.Tasks[i].Submit()
// 			flag = true
// 		}
// 	}

// 	if !flag {
// 		return fmt.Errorf("task with id %d not found", taskID)
// 	}

// 	return nil
// }

// func GetStudents() []Student {
// 	students:= []Student{
// 		{
// 			data: human.Human{
// 				Name: "Давид",
// 				Surname: "Петров",
// 				Age: 19,
// 			},
// 			groupID: "1",
// 			ID: 1,
// 			Tasks: []Task{
// 				{ID: 1},
// 				{ID: 10},
// 			},
// 		},
// 		{
// 			data: human.Human{
// 				Name: "A",
// 				Surname: "Петров",
// 				Age: 19,
// 			},
// 			groupID: "2",
// 			ID: 2,
// 			Tasks: []Task{
// 				{ID: 2},
// 				{ID: 9},
// 			},
// 		},
// 		{
// 			data: human.Human{
// 				Name: "B",
// 				Surname: "Петров",
// 				Age: 19,
// 			},
// 			groupID: "3",
// 			ID: 3,
// 			Tasks: []Task{
// 				{ID: 3},
// 				{ID: 8},
// 			},
// 		},
// 		{
// 			data: human.Human{
// 				Name: "C",
// 				Surname: "Петров",
// 				Age: 19,
// 			},
// 			groupID: "1",
// 			ID: 4,
// 			Tasks: []Task{
// 				{ID: 4},
// 				{ID: 7},
// 			},
// 		},
// 		{
// 			data: human.Human{
// 				Name: "D",
// 				Surname: "Петров",
// 				Age: 19,
// 			},
// 			groupID: "2",
// 			ID: 5,
// 			Tasks: []Task{
// 				{ID: 5},
// 				{ID: 6},
// 			},
// 		},
// 		{
// 			data: human.Human{
// 				Name: "E",
// 				Surname: "Петров",
// 				Age: 19,
// 			},
// 			groupID: "3",
// 			ID: 6,
// 			Tasks: []Task{
// 				{ID: 6},
// 				{ID: 5},
// 			},
// 		},
// 		{
// 			data: human.Human{
// 				Name: "D",
// 				Surname: "Петров",
// 				Age: 19,
// 			},
// 			groupID: "1",
// 			ID: 7,
// 			Tasks: []Task{
// 				{ID: 7},
// 				{ID: 4},
// 			},
// 		},
// 		{
// 			data: human.Human{
// 				Name: "E",
// 				Surname: "Петров",
// 				Age: 19,
// 			},
// 			groupID: "2",
// 			ID: 8,
// 			Tasks: []Task{
// 				{ID: 8},
// 				{ID: 3},
// 			},
// 		},
// 		{
// 			data: human.Human{
// 				Name: "F",
// 				Surname: "Петров",
// 				Age: 19,
// 			},
// 			groupID: "3",
// 			ID: 9,
// 			Tasks: []Task{
// 				{ID: 9},
// 				{ID: 2},
// 			},
// 		},
// 		{
// 			data: human.Human{
// 				Name: "G",
// 				Surname: "Петров",
// 				Age: 19,
// 			},
// 			groupID: "1",
// 			ID: 10,
// 			Tasks: []Task{
// 				{ID: 10},
// 				{ID: 1},
// 			},
// 		},
// 	}

// 	return students
// }