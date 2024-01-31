package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "query",
	Short: "Perform a Prom query",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hi range")
	},
}

func perform_queries() {
	//
}
