/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// altCmd represents the alt command
var altCmd = &cobra.Command{
	Use:   "alt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("alt called")
	},
	//Deprecated: "[newalt] deprecated",
}

func init() {
	//rootCmd.Flags().MarkDeprecated("alt", "Deprecated flag config")
	gopherCmd.AddCommand(altCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// altCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// altCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
