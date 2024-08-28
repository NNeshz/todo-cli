package main

import "fmt"

type Task struct {
	ID        int
	Name      string
	Completed bool
}

func (t *Task) String() string {
	status := "❌"
	if t.Completed {
		status = "✅"
	}
	return fmt.Sprintf("%d. %s %s", t.ID, status, t.Name)
}
