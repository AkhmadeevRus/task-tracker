package main

import (
	cmd "github.com/AkhmadeevRus/task-tracker/cmd/tracker"
	"github.com/sirupsen/logrus"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("execute error: %s", err)
	}
}
