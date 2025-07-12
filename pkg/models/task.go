package models

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "todo"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in-progress"
	TASK_STATUS_DONE        TaskStatus = "done"
)

type Task struct {
	ID          int64      `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func NewTask(id int64, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      TASK_STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ListTasks(status TaskStatus) error {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return err
	}

	filteredTasks := filterTasks(tasks, status)
	fmt.Println()
	for _, task := range filteredTasks {
		formattedId := lipgloss.NewStyle().
			Bold(true).
			Width(5).
			Render(fmt.Sprintf("ID:%d", task.ID))
		formattedStatus := lipgloss.NewStyle().
			Bold(true).
			Width(12).
			Foreground(lipgloss.Color(statusColor(task.Status))).
			Render(string(task.Status))

		relativeUpdatedTime := task.UpdatedAt.Format("2006-01-02 15:04:05")

		taskStyle := lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(lipgloss.Color("#3C3C3C")).
			Render(fmt.Sprintf("%s %s %s (%s)", formattedId, formattedStatus, task.Description, relativeUpdatedTime))
		fmt.Println(taskStyle)
	}
	fmt.Println()

	return nil
}

func statusColor(status TaskStatus) string {
	switch status {
	case TASK_STATUS_TODO:
		return "#3C3C3C"
	case TASK_STATUS_IN_PROGRESS:
		return "202"
	case TASK_STATUS_DONE:
		return "#04B575"
	default:
		return "#3C3C3C"
	}
}

func filterTasks(tasks []Task, status TaskStatus) []Task {
	if status == "all" {
		return tasks
	}
	var result []Task
	for _, task := range tasks {
		if task.Status == status {
			result = append(result, task)
		}
	}
	return result
}

func AddTask(description string) error {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return err
	}

	var newTaskId int64
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		newTaskId = lastTask.ID + 1
	} else {
		newTaskId = 1
	}

	task := NewTask(newTaskId, description)
	tasks = append(tasks, *task)

	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFCC66"))

	formattedId := style.Render(fmt.Sprintf("(ID: %d)", task.ID))
	fmt.Printf("\nTask added successfully: %s\n\n", formattedId)
	return WriteTasksToFile(tasks)
}

func DeleteTaskById(id int64) error {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return err
	}

	var newTasks []Task
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}

	if len(newTasks) == len(tasks) {
		fmt.Printf("\ntask for (ID:%d) not found\n", id)
		return nil
	}

	fmt.Println("\nTask deleted successfully")
	return WriteTasksToFile(newTasks)
}
