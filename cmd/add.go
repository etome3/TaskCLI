package cmd

import (
	"context"
	"fmt"
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
	tasks, err := openJson()
	if err != nil {
		return err
	}
	tasks = append(tasks, task)
	err = writeJson(tasks)
	if err != nil {
		return err
	}
	fmt.Printf("Added task: \"%s\"\nAt: %v", task.Name, task.Added.Format("Mon Jan 2 15:04:05 MST 2006"))
	return nil
}
