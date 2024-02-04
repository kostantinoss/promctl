package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "promctl",
	Short: "promctl is simple query tool for Prometheus",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hi this is promctl handler function")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
