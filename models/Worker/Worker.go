package worker

import (
	"github.com/hitokirisss/struct/models/human"
)

type Worker struct {
	data   human.Human
	ID     int
	Groups []string
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



