/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package desync

import (
	"github.com/spf13/cobra"
)

// desyncCmd represents the desync command
var DesyncCmd = &cobra.Command{
	Use:   "desync",
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
	// desyncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// desyncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
