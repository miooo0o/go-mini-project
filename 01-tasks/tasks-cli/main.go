package main

// "encoding/csv"
// "fmt"
// "os"
// "strconv"
// "syscall"
// "time"

import (
	"fmt"
	"os"
	"tasks-cli/config"
	"tasks-cli/install"
)

func getConfig() config.Config {
	configPath := os.Getenv("TASKS_CONFIG_PATH")
	if configPath == "" {
		fmt.Errorf("TASKS_CONFIG_PATH environment variable is not set.")
		os.Exit(1)
	}
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		os.Exit(1)
	}
	return cfg
}

/*
Prerequisite:
It is assumed that `config.json` and `tasks.csv` have been properly created
using `go run install.go` before running this program.
  - TASKS_CONFIG_PATH environment variable
*/
func main() {
	install.Run()
	cfg := getConfig()
}
