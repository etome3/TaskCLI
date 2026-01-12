package cmd

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli/v3"
)

type Task struct {
	Name  string    `json:"name"`
	Added time.Time `json:"added"`
}

const (
	dataDir      = "data"
	dataFileName = "tasks.json"
)

var dataFile = filepath.Join(dataDir, dataFileName)

var root = cli.Command{
	Name:  "TaskCLI",
	Usage: "A CLI application for managing tasks",
	Commands: []*cli.Command{
		&AddCmd,
		&CompleteCmd,
		&ListCmd,
	},
}

func Run() {
	initData()
	err := root.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
