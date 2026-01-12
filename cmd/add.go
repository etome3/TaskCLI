package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v3"
)

var AddCmd = cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "Add a task",
	Action:  Add,
	Arguments: []cli.Argument{
		&cli.StringArgs{
			Name: "addedTask",
			Max:  -1,
			Min:  1,
		},
	},
}

func Add(ctx context.Context, c *cli.Command) error {
	taskName := strings.Join(c.StringArgs("addedTask"), " ")
	task := Task{
		Name:  taskName,
		Added: time.Now(),
	}
	err := writeJson(task)
	if err != nil {
		return err
	}
	fmt.Printf("Added task: \"%s\"\nAt: %v", task.Name, task.Added.Format("Mon Jan 2 15:04:05 MST 2006"))
	return nil
}

func writeJson(task Task) error {
	tasks, err := openJson()
	if err != nil {
		return err
	}
	tasks = append(tasks, task)
	output, err := json.MarshalIndent(&tasks, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(dataFile, output, 0777)
	if err != nil {
		return err
	}
	return nil
}

func openJson() ([]Task, error) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		data = []byte("[]")
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
