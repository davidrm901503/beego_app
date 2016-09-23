package models

import (
	"fmt"
)

var DefaultTaskList *TaskManager

type Task struct {
	ID    int64
	Title string 
}

func New(title string) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("empty title")
	}
	return &Task{0, title}, nil
}


type TaskManager struct {
	tasks  []*Task
	lastID int64
}


func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (m *TaskManager) Save(task *Task) error {
	if task.ID == 0 {
		m.lastID++
		task.ID = m.lastID
		m.tasks = append(m.tasks, cloneTask(task))
		return nil
	}

	for i, t := range m.tasks {
		if t.ID == task.ID {
			m.tasks[i] = cloneTask(task)
			return nil
		}
	}
	return fmt.Errorf("unknown task")
}


func cloneTask(t *Task) *Task {
	c := *t
	return &c
}


func (m *TaskManager) All() []*Task {
	return m.tasks
}

func init() {
	DefaultTaskList = NewTaskManager()
}
