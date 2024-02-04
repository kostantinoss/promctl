/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type PromConfig struct {
	PromServer string `yaml:"promServer"`
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")
		serverConfig, err := configPrometheus(server)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Prometheus server: ", serverConfig)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringP("server", "s", "localhost:9090", "Set Prometheus server")
	configCmd.Flags().StringP("file", "f", "", "Provide a custom config file")
}

func configPrometheus(server string) (string, error) {
	conf := PromConfig{
		PromServer: "http://" + server,
	}

	config, err := yaml.Marshal(&conf)
	if err != nil {
		fmt.Println(err.Error())
	}

	file, err := os.Create("config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	str_config := string(config)
	file.WriteString(str_config)

	return str_config, err
}
