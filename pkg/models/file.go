package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

func getTaskFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		logrus.Fatalf("error getting current working derictory %s", err)
	}
	return path.Join(cwd, "task.json")
}

func ReadTaskFromFile() ([]Task, error) {
	filePath := getTaskFilePath()
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		fmt.Println("file dosen't exist. Creating new file...")
		file, err := os.Create(filePath)
		os.WriteFile(filePath, []byte("[]"), os.ModeAppend.Perm())

		if err != nil {
			logrus.Fatalf("error crating file: %s", err)
		}
		defer file.Close()
		return []Task{}, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		logrus.Fatalf("error opening file: %s", err)
	}
	defer file.Close()

	tasks := []Task{}
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Printf("error decoding file: %s\n", err)
		return nil, err
	}

	return tasks, nil
}

func WriteTasksToFile(tasks []Task) error {
	filepath := getTaskFilePath()
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("error creating file: ", err)
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		fmt.Println("error encoding file: ", err)
		return err
	}

	return nil
}
