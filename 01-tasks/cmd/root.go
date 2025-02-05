package cmd

import (
	// "fmt"
	// "os"
	"github.com/spf13/cobra"
)

var (
	dataFile string
)

var rootCmd = &cobra.Command{}

func Excute() error {
	return rootCmd.Execute()
}
