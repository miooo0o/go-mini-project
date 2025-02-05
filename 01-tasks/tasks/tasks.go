package tasks

import (
	// "encoding/csv"
	// "fmt"
	// "os"
	// "strconv"
	// "syscall"
	// "github.com/spf13/cobra"
	"time"
)

type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	IsComplete  bool
	CompletedAt time.Time
}

// func loadFile()
// func closeFile()
// func LoadTasks()
// func AddTask()
