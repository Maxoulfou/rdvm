/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/Maxoulfou/rdvm/config"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rdvm",
	Short: "Remote Dyjix VPS Management",
	Long: `rdvm is a simple CLI to manage your Dyjix VPS 
easily without the need to log into the management 
panel (vps.dyjix.eu)`,
}

var Configuration config.JsonConfig

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	Configuration.Load()
}
