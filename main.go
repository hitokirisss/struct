package main

import (
	"fmt"

	"github.com/hitokirisss/struct/models"
)

func main() {
	workers := models.GetWorkers()
	fmt.Println(workers)
	fmt.Println()
	students := models.GetStudents()
	fmt.Println(students)
	err:= Submit("2", 1, students)
	if err != nil {
		panic(err)
	}
	fmt.Println(students)
}



func Submit(groupID string, taskID int, students []models.Student) error {
	for i, student := range students {
		if student.GetGroupID() == groupID {
			err := students[i].SubmitTask(taskID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

