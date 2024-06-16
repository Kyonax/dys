/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package check

import (
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		cobra.CheckErr(err)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
