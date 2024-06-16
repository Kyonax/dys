/*
Copyright © 2024 Kyonax <kyonax_on_tech>

Built by Kyonax, the code junkie who thrives on simplicity.
Use it, just give credit where it's due.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Kyonax/dys/cmd/check"
	"github.com/Kyonax/dys/cmd/desync"
	"github.com/Kyonax/dys/cmd/sync"
	"github.com/Kyonax/dys/cmd/yield"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "dys",
	Short: "Sync your dot-files, zero-friction",
	Long: `Dot-files Yield Script is a CLI tool for all of us who love to tweak and personalize our setups to perfection. Who’s got time to waste on manual configs? Let dys streamline the process.

With a single command, you can download your dot-files and integrate them into any device.
Forget the tedious setup rituals; dys makes it seamless:

- Deploy your optimized dev environment to a fresh rig in seconds.
- Maintain uniform configs across all your devices.
- Simplify your dot-file management and deployment like a pro.

The go-to for code junkies who value their time. Use dys, or DIY!`,

	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommands() {
	rootCmd.AddCommand(check.CheckCmd)
	rootCmd.AddCommand(desync.DesyncCmd)
	rootCmd.AddCommand(sync.SyncCmd)
	rootCmd.AddCommand(yield.YieldCmd)
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dys.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommands()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".dys" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".dys")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
