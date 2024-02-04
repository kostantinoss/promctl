package cmd

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/spf13/cobra"
)

var queryCommand = &cobra.Command{
	Use:   "query",
	Short: "Perform a simple query",
	Long:  "",
	Run:   Query,
}

func Query(cmd *cobra.Command, args []string) {
	query, _ := cmd.Flags().GetString("query")
	output, _ := cmd.Flags().GetString("output")

	fmt.Println(output)

	api := v1.NewAPI(nil)

	result, _, err := api.Query(
		context.TODO(), query, time.Now(), v1.WithTimeout(10*time.Second),
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	if output == "file" {
		// exporter.Export_raw()
	} else {
		for _, elem := range result.(model.Vector) {
			fmt.Println(elem)
		}
	}
}

func init() {
	rootCmd.AddCommand(queryCommand)
	queryCommand.Flags().StringP("server", "s", "", "")
	queryCommand.Flags().StringP("query", "q", "", "")
	queryCommand.Flags().StringP("output", "o", "", "")

}
