/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/Maxoulfou/rdvm/execution"
	"github.com/spf13/cobra"
)

// shutdownCmd represents the shutdown command
var shutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "Turn your VPS in \"rompiche\" mode",
	Long:  `Simply turn off your VPS`,
	Run: func(cmd *cobra.Command, args []string) {
		InitShutdown()
	},
}

func init() {
	rootCmd.AddCommand(shutdownCmd)
}

func InitShutdown() {
	execution.Shutdown(execution.Execute("shutdown", Configuration))
}
