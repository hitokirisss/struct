package student

import (
	"github.com/hitokirisss/struct/models/human"
	 "github.com/hitokirisss/struct/models/Task"
)

type Student struct {
		data    human.Human
		groupID string
		ID      int
		Tasks   []task.Task
	}

	func GetStudents() []Student {
		students:= []Student{
			{
				data: human.Human{
					Name: "Давид",
					Surname: "Петров",
					Age: 19,
				},
				groupID: "1",
				ID: 1,
				Tasks: []task.Task{
					{ID: 1},
					{ID: 10},
				},
			},
			{
				data: human.Human{
					Name: "A",
					Surname: "Петров",
					Age: 19,
				},
				groupID: "2",
				ID: 2,
				Tasks: []task.Task{
					{ID: 2},
					{ID: 9},
				},
			},
			{
				data: human.Human{
					Name: "B",
					Surname: "Петров",
					Age: 19,
				},
				groupID: "3",
				ID: 3,
				Tasks: []task.Task{
					{ID: 3},
					{ID: 8},
				},
			},
			{
				data: human.Human{
					Name: "C",
					Surname: "Петров",
					Age: 19,
				},
				groupID: "1",
				ID: 4,
				Tasks: []task.Task{
					{ID: 4},
					{ID: 7},
				},
			},
			{
				data: human.Human{
					Name: "D",
					Surname: "Петров",
					Age: 19,
				},
				groupID: "2",
				ID: 5,
				Tasks: []task.Task{
					{ID: 5},
					{ID: 6},
				},
			},
			{
				data: human.Human{
					Name: "E",
					Surname: "Петров",
					Age: 19,
				},
				groupID: "3",
				ID: 6,
				Tasks: []task.Task{
					{ID: 6},
					{ID: 5},
				},
			},
			{
				data: human.Human{
					Name: "D",
					Surname: "Петров",
					Age: 19,
				},
				groupID: "1",
				ID: 7,
				Tasks: []task.Task{
					{ID: 7},
					{ID: 4},
				},
			},
			{
				data: human.Human{
					Name: "E",
					Surname: "Петров",
					Age: 19,
				},
				groupID: "2",
				ID: 8,
				Tasks: []task.Task{
					{ID: 8},
					{ID: 3},
				},
			},
			{
				data: human.Human{
					Name: "F",
					Surname: "Петров",
					Age: 19,
				},
				groupID: "3",
				ID: 9,
				Tasks: []task.Task{
					{ID: 9},
					{ID: 2},
				},
			},
			{
				data: human.Human{
					Name: "G",
					Surname: "Петров",
					Age: 19,
				},
				groupID: "1",
				ID: 10,
				Tasks: []task.Task{
					{ID: 10},
					{ID: 1},
				},
			},
		}
	
		return students
	}
