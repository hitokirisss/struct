package models

import (
	"fmt"

	"github.com/hitokirisss/struct/models/human"
)

type Worker struct {
	data   human.Human
	ID     int
	Groups []string
}



func (w *Worker) SetNewData(newName string, newSurname string) {
	w.data.SetNewNameAndSurname(newName, newSurname)
}

func (w Worker) GetData() string {
	return w.data.GetData()
}

func (w Worker) GetGroupID() string {
	for _, group := range w.Groups {
		return group
	}
	return ""
}

func (w *Worker) DeleteGroup(groupID string) error {
	for i, group := range w.Groups {
		if group == groupID {
			w.Groups = append(w.Groups[0:i], w.Groups[i+1:]...)
		}
	}
	return fmt.Errorf("группа с таким id %s не найдена", groupID)
}

func (w *Worker) AddGroup(groupID string) error {
	for _, group := range w.Groups {
		if group == groupID {
			return fmt.Errorf("группа с таким id %s уже есть", groupID)
		}
	}
	w.Groups = append(w.Groups, groupID)
	return nil
}

func GetWorkers() []Worker{
	workers := []Worker{
		{
			data: human.Human{
				Name: "Олег",
				Surname: "Евгеев",
				Age: 21,
			},
			ID:     1,
			Groups: []string{"1", "2"},
		},
		{
			data: human.Human{
				Name: "Айта",
				Surname: "Оконов",
				Age: 21,
			},
			ID: 2,
			Groups: []string{"1", "3"},
		},
		{
			data: human.Human{
				Name: "Дмитрий",
				Surname: "Иванов",
				Age: 23,
			},
			ID: 3,
			Groups: []string{"2","3"},
		}, 
	}
	return workers
}



