/*
Copyright © 2022 BROCHIER MAXENCE maxence@brochier.xyz

*/
package cmd

import (
	"github.com/Maxoulfou/rdvm/execution"
	"github.com/spf13/cobra"
)

// rebootCmd represents the reboot command
var rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Reboot your VPS",
	Long:  `Restart your VPS, it will automatically turn off and on within 5 minutes`,
	Run: func(cmd *cobra.Command, args []string) {
		InitReboot()
	},
}

func init() {
	rootCmd.AddCommand(rebootCmd)
}

func InitReboot() {
	execution.Reboot(execution.Execute("reboot", Configuration))
}
