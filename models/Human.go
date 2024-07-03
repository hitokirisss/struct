package models

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     int
}

func (h *Human) SetNewNameAndSurname(newName string, NewSurname string) {
	h.Name = newName
	h.Surname = NewSurname
}

func (h Human) GetData() string {
	return fmt.Sprintf("%s %s %d", h.Name, h.Surname, h.Age)
}
