/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package yield

import (
	"os"

	"github.com/spf13/cobra"
)

// yieldCmd represents the yield command
var YieldCmd = &cobra.Command{
	Use:   "yield",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// yieldCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// yieldCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
