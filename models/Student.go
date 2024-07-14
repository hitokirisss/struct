package models

type Student struct {
	Data    Human 
	GroupID string
	ID      int
	Tasks   []Task
}

func GetStudents() []Student {
	students := []Student{
		{
			Data: Human{
				Name:    "Давид",
				Surname: "Петров",
				Age:     19,
			},
			GroupID: "1",
			ID:      1,
			Tasks: []Task{
				{ID: 1},
				{ID: 10},
			},
		},
		{
			Data: Human{
				Name:    "A",
				Surname: "Петров",
				Age:     19,
			},
			GroupID: "2",
			ID:      2,
			Tasks: []Task{
				{ID: 2},
				{ID: 9},
			},
		},
		{
			Data: Human{
				Name:    "B",
				Surname: "Петров",
				Age:     19,
			},
			GroupID: "3",
			ID:      3,
			Tasks: []Task{
				{ID: 3},
				{ID: 8},
			},
		},
		{
			Data: Human{
				Name:    "C",
				Surname: "Петров",
				Age:     19,
			},
			GroupID: "1",
			ID:      4,
			Tasks: []Task{
				{ID: 4},
				{ID: 7},
			},
		},
		{
			Data: Human{
				Name:    "D",
				Surname: "Петров",
				Age:     19,
			},
			GroupID: "2",
			ID:      5,
			Tasks: []Task{
				{ID: 5},
				{ID: 6},
			},
		},
		{
			Data: Human{
				Name:    "E",
				Surname: "Петров",
				Age:     19,
			},
			GroupID: "3",
			ID:      6,
			Tasks: []Task{
				{ID: 6},
				{ID: 5},
			},
		},
		{
			Data: Human{
				Name:    "D",
				Surname: "Петров",
				Age:     19,
			},
			GroupID: "1",
			ID:      7,
			Tasks: []Task{
				{ID: 7},
				{ID: 4},
			},
		},
		{
			Data: Human{
				Name:    "E",
				Surname: "Петров",
				Age:     19,
			},
			GroupID: "2",
			ID:      8,
			Tasks: []Task{
				{ID: 8},
				{ID: 3},
			},
		},
		{
			Data: Human{
				Name:    "F",
				Surname: "Петров",
				Age:     19,
			},
			GroupID: "3",
			ID:      9,
			Tasks: []Task{
				{ID: 9},
				{ID: 2},
			},
		},
		{
			Data: Human{
				Name:    "G",
				Surname: "Петров",
				Age:     19,
			},
			GroupID: "1",
			ID:      10,
			Tasks: []Task{
				{ID: 10},
				{ID: 1},
			},
		},
	}

	return students
}
