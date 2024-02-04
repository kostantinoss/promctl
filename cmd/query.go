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
	"gopkg.in/yaml.v3"
)

var queryCommand = &cobra.Command{
	Use:   "query",
	Short: "Perform a simple query",
	Long:  "",
	Run:   Query,
}

func init() {
	rootCmd.AddCommand(queryCommand)
	queryCommand.Flags().StringP("query", "q", "", "Provide query string")
	queryCommand.Flags().StringP("output", "o", "stdout", "Write query result to file")

}

func Query(cmd *cobra.Command, args []string) {
	output_file, _ := cmd.Flags().GetString("output")
	query, _ := cmd.Flags().GetString("query")
	if query == "" {
		fmt.Println("Error, query empty ")
		os.Exit(1)
	}

	server, _ := getDerverFromConfig()
	client, err := api.NewClient(api.Config{
		Address: server,
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
	}

	api := v1.NewAPI(client)

	result, _, err := api.Query(
		context.TODO(), query, time.Now(), v1.WithTimeout(10*time.Second),
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	if output_file != "stdout" {
		if isFilepath := checkFilename(output_file); isFilepath {
			exporter.Export_raw(result, output_file)
			fmt.Println("results exported at " + output_file)
		}
	} else {
		for _, elem := range result.(model.Vector) {
			fmt.Println(elem)
		}
	}
}

func getDerverFromConfig() (string, error) {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	config := PromConfig{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
	}

	return config.PromServer, err
}

func checkFilename(filename string) bool {
	// TODO: implement
	return true
}
