/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/Maxoulfou/rdvm/execution"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get actual status, HN and IP about your VPS",
	Long:  `Get the current status of your VPS as well as its hostname and IP`,
	Run: func(cmd *cobra.Command, args []string) {
		InitInfo()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func InitInfo() {
	execution.Info(execution.Execute("info", Configuration))
}
