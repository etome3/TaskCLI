package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func initData() {
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

func writeJson(path string, tasks []Task) error {
	output, err := json.MarshalIndent(&tasks, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, output, 0777)
	if err != nil {
		return err
	}
	return nil
}

func readJson(path string) ([]Task, error) {
	data, err := os.ReadFile(path)
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
