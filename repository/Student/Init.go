package student

import "github.com/hitokirisss/struct/models"

type StudentMemmoryRepo struct {
	students []models.Student
}

func New() *StudentMemmoryRepo {
	return &StudentMemmoryRepo{
		students: models.GetStudents(),
	}
}
