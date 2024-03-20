/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ThailanTec/cli-api/webserver"
	"github.com/spf13/cobra"
)

var port string

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Inicia o servidor HTTP irmão",
	Long:  `Longa historia`,
	Run: func(cmd *cobra.Command, args []string) {
		server := webserver.Server{Port: port}
		server.Server()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.Flags().StringVarP(&port, "port", "p", ":4040", "Porta usada no servidor http")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
