package models


type Worker struct {
	Data   Human
	ID     int
	Groups []string
}

func GetWorkers() []Worker{
	workers := []Worker{
		{
			Data: Human{
				Name: "Олег",
				Surname: "Евгеев",
				Age: 21,
			},
			ID:     1,
			Groups: []string{"1", "2"},
		},
		{
			Data: Human{
				Name: "Айта",
				Surname: "Оконов",
				Age: 21,
			},
			ID: 2,
			Groups: []string{"1", "3"},
		},
		{
			Data: Human{
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