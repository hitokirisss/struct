package models

type Task struct{
	ID int
	IsSubmited bool
}

func (t *Task) Submit() {
	t.IsSubmited = true
}