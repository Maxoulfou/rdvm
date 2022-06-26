/*
Copyright © 2022 BROCHIER MAXENCE maxence@brochier.xyz

*/
package cmd

import (
	"github.com/Maxoulfou/rdvm/execution"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get actual status of your VPS",
	Long: `Get the current status of your VPS as well as its statusMSG, 
its status vm stat, its hostname and its IP`,
	Run: func(cmd *cobra.Command, args []string) {
		InitStatus()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func InitStatus() {
	execution.Status(execution.Execute("status", Configuration))
}
