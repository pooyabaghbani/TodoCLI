/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/pooyabaghbani/TodoCLI/src/todo"
	"github.com/spf13/cobra"
)

var returnAll bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Return a list of tasks ",
	Long:  `Return a list of all of the uncompleted tasks, with the option to return all tasks regardless of whether or not they are completed.`,
	Run: func(cmd *cobra.Command, args []string) {
		todo.List(returnAll)
	},
	// DisableFlagParsing: true,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.
	listCmd.Flags().BoolVarP(&returnAll, "all", "a", false, "return all tasks")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
