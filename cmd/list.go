package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

var ListCmd = cli.Command{
	Name:    "list",
	Aliases: []string{"l", "ls"},
	Usage:   "List uncompleted tasks",
	Action:  List,
}

func List(ctx context.Context, c *cli.Command) error {
	tasks, err := readJson(dataFile)
	if err != nil {
		return err
	}
	taskList := "Uncompleted Tasks:\n\n"
	for _, task := range tasks {
		taskList = taskList + fmt.Sprintf("Task: %v\nAdded at: %v\n\n", task.Name, task.Added.Format("Mon Jan 2 15:04:05 MST 2006"))
	}
	fmt.Print(taskList)
	return nil
}
