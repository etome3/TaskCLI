package cmd

import (
	"context"
	"fmt"
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

var dataDir = "data"
var dataFile = filepath.Join(dataDir, "tasks.json")

var root = cli.Command{
	Name:  "TaskCLI",
	Usage: "A CLI application for managing tasks",
	Commands: []*cli.Command{
		&AddCmd,
	},
}

func Run() {
	err := root.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	err := os.MkdirAll(dataDir, 0777)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile(dataFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening or creating file:", err)
	}
	defer file.Close()
}
