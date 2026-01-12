package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"
)

var CompleteCmd = cli.Command{
	Name:    "complete",
	Aliases: []string{"c", "done"},
	Usage:   "Complete a task",
	Action:  Complete,
	Arguments: []cli.Argument{
		&cli.StringArgs{
			Name: "completedTask",
			Max:  -1,
			Min:  1,
		},
	},
}

func Complete(ctx context.Context, c *cli.Command) error {
	taskName := strings.Join(c.StringArgs("completedTask"), " ")
	tasks, err := readJson(dataFile)
	if err != nil {
		return err
	}
	var newTasks []Task
	for _, task := range tasks {
		if task.Name != taskName {
			newTasks = append(newTasks, task)
		}
	}
	err = writeJson(dataFile, newTasks)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully completed task: %v", taskName)
	return nil
}
