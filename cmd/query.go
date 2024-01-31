package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "query",
	Short: "Perform a Prom query",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hi")
	},
}

func query(query string) model.Value {
	client, err := api.NewClient(api.Config{
		Address: "http://localhost:9090",
	})

	if err != nil {
		fmt.Println("Error creating client")
	}

	v1api := v1.NewAPI(client)

	result, _, err := v1api.Query(
		context.TODO(), query, time.Now(), v1.WithTimeout(10*time.Second),
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	return result
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(fmt.Errorf("error"))
	}
}
