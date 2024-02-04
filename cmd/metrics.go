/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"promctl/exporter"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/spf13/cobra"
)

// metricsCmd represents the metrics command
var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Prints the available metrics for the configured Prometheus db",

	Run: func(cmd *cobra.Command, args []string) {
		output_file, _ := cmd.Flags().GetString("output")

		server, _ := getDerverFromConfig()
		client, err := api.NewClient(api.Config{
			Address: server,
		})
		if err != nil {
			fmt.Printf("Error creating client: %v\n", err)
		}

		api := v1.NewAPI(client)

		result, _, err := api.Query(
			context.TODO(), "{__name__!=\"\"}", time.Now(), v1.WithTimeout(10*time.Second),
		)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if result != nil && result.Type() == model.ValVector {
			if output_file == "stdout" {
				for _, elem := range result.(model.Vector) {
					fmt.Println(elem)
				}
			} else {
				exporter.Export_raw(result, output_file)
				fmt.Println("Metrics exported to ", output_file)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(metricsCmd)
	metricsCmd.Flags().StringP("output", "o", "stdout", "Export metrics output to file")
}
