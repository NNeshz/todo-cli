package main

import (
	"encoding/json"
	"os"
)

func SaveTasksToFile(tl *TaskList, filename string) error {
	data, err := json.MarshalIndent(tl.Tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func LoadTasksFromFile(tl *TaskList, filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &tl.Tasks)
	if err != nil {
		return err
	}

	for _, task := range tl.Tasks {
		if task.ID >= tl.NextID {
			tl.NextID = task.ID + 1
		}
	}

	return nil
}
