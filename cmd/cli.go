/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	app "github.com/ThailanTec/cli-api/internal"
	"github.com/spf13/cobra"
)

var a, b float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Soma ou some bb",
	Long:  `Soma com nós truta`,
	Run: func(cmd *cobra.Command, args []string) {
		calc := app.NewCalc()

		calc.A = a
		calc.B = b

		fmt.Println(calc.Sum())
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().Float64VarP(&a, "a", "a", 0, "A")
	cliCmd.Flags().Float64VarP(&b, "b", "b", 0, "B")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
