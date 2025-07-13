package tasks

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type taskItem struct {
	Description string
	Done bool
	CreatedAt time.Time
	DoneAt time.Time
}

type List []taskItem

func (self *List) Add(task string) {
	item := taskItem {
		Description: task,
		Done: false,
		CreatedAt: time.Now(),
		DoneAt: time.Time{},
	}

	*self = append(*self, item)
}

func (self *List) Complete(index int) error {
	list := *self
	if index <= 0 || index > len(list) {
		return fmt.Errorf("Task item %d does not exist!", index)
	}

	list[index - 1].Done = true
	list[index - 1].DoneAt = time.Now()

	return nil
}

func (self *List) Delete(index int) error {
	list := *self
	if index <= 0 || index > len(list) {
		return fmt.Errorf("Task item %d does not exist!", index)
	}

	*self = append(list[:index - 1], list[index:]...)

	return nil
}

func (self *List) Save(filename string) error {
	data, err := json.Marshal(self)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (self *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		} else {
			return err
		}
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, self)
}
