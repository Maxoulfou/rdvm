/*
Copyright © 2022 BROCHIER MAXENCE maxence@brochier.xyz

*/
package cmd

import (
	"github.com/Maxoulfou/rdvm/execution"
	"github.com/spf13/cobra"
)

// bootCmd represents the boot command
var bootCmd = &cobra.Command{
	Use:   "boot",
	Short: "Turn your VPS in \"unrompiche\" mode",
	Long:  `Launch your VPS (not through the window) but to access it and its services`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("boot called")
		InitBoot()
	},
}

func init() {
	rootCmd.AddCommand(bootCmd)
}

func InitBoot() {
	execution.Boot(execution.Execute("boot", Configuration))
}
