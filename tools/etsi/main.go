package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "etsi",
	Short: "etsi helps you write your ETSI documents in markdown.",
	Long:  `etsi is a CLI tool that helps you write your ETSI documents in markdown.`,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
