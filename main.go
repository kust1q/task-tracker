package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func Usage() {
	fmt.Println("Usage: ./task-tracker OPTION [ID] [TEXT]")
}

func Help() {
	Usage()
	fmt.Println("Your personal task tracker\nOPTIONS\nadd - adding a new task\ndelete - deleting tasks\nupdate - updating task")
	fmt.Println("mark-in-progress - marking a task as in progress\nmark-done - marking a task as done")
	fmt.Println("list [done / todo / in-progress] - listing all tasks / by status")
}

func ReadTasks(filename string) ([]Task, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		log.Println("Error with open file")
		os.Exit(1)
	}
	defer file.Close()
	tasks := make([]Task, 0)
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	if len(bytes.TrimSpace(data)) == 0 {
		return []Task{}, nil
	}
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func LoadTasks(tasks []Task, filename string) error {
	for i := range tasks {
		tasks[i].Id = i + 1
	}
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0666)
}

func AddTask(tasks []Task, description string) []Task {
	var task Task
	task.Id = len(tasks)
	task.Description = description
	task.Status = "todo"
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	tasks = append(tasks, task)
	return tasks
}

func UpdateTask(tasks []Task, id int, description string) {
	tasks[id-1].Description = description
}

func DeleteTasks(tasks []Task, id int) []Task {
	tasks = append(tasks[:id-1], tasks[id:]...)
	return tasks
}

func MarkInProgress(tasks []Task, id int) {
	tasks[id-1].Status = "in-progress"
}

func MarkDone(tasks []Task, id int) {
	tasks[id-1].Status = "done"
}

func List(tasks []Task) {
	for _, task := range tasks {
		fmt.Printf("%d. %s\ncreated - %s\nupdated - %s\nstatus - %s\n", task.Id, task.Description, task.CreatedAt, task.UpdatedAt, task.Status)
	}
}

func ListToDo(tasks []Task) {
	for _, task := range tasks {
		if task.Status == "todo" {
			fmt.Printf("%d. %s\ncreated - %s\nupdated - %s\nstatus - %s\n", task.Id, task.Description, task.CreatedAt, task.UpdatedAt, task.Status)
		}
	}
}

func ListDone(tasks []Task) {
	for _, task := range tasks {
		if task.Status == "done" {
			fmt.Printf("%d. %s\ncreated - %s\nupdated - %s\nstatus - %s\n", task.Id, task.Description, task.CreatedAt, task.UpdatedAt, task.Status)
		}
	}
}

func ListInProgress(tasks []Task) {
	for _, task := range tasks {
		if task.Status == "in-progress" {
			fmt.Printf("%d. %s\ncreated - %s\nupdated - %s\nstatus - %s\n", task.Id, task.Description, task.CreatedAt, task.UpdatedAt, task.Status)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		Help()
		os.Exit(0)
	}

	filename := "tasks.json"
	tasks, err := ReadTasks(filename)
	if err != nil {
		log.Println("Error with decode file")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) != 3 {
			Help()
			os.Exit(0)
		}
		tasks = AddTask(tasks, os.Args[2])
	case "update":
		if len(os.Args) != 4 {
			Help()
			os.Exit(0)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Wrong task id")
			Help()
			os.Exit(0)
		}
		UpdateTask(tasks, id, os.Args[3])
	case "delete":
		if len(os.Args) != 3 {
			Help()
			os.Exit(0)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Wrong task id")
			Help()
			os.Exit(0)
		}
		if id > len(tasks) {
			fmt.Println("Wrong index")
			os.Exit(0)
		}
		tasks = DeleteTasks(tasks, id)
	case "mark-in-progress":
		if len(os.Args) != 3 {
			Help()
			os.Exit(0)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Wrong task id")
			Help()
			os.Exit(0)
		}
		if id > len(tasks) {
			fmt.Println("Wrong index")
			os.Exit(0)
		}
		MarkInProgress(tasks, id)
	case "mark-done":
		if len(os.Args) != 3 {
			Help()
			os.Exit(0)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Wrong task id")
			Help()
			os.Exit(0)
		}
		if id > len(tasks) {
			fmt.Println("Wrong index")
			os.Exit(0)
		}
		MarkDone(tasks, id)
	case "list":
		if len(os.Args) == 2 {
			List(tasks)
		} else if len(os.Args) == 3 {
			switch os.Args[2] {
			case "done":
				ListDone(tasks)
			case "in-progress":
				ListInProgress(tasks)
			case "todo":
				ListToDo(tasks)
			}
		} else {
			Help()
			os.Exit(0)
		}
	default:
		Help()
		return
	}

	err = LoadTasks(tasks, filename)
	if err != nil {
		log.Println("Error with update file with tasks")
		os.Exit(1)
	}
}
