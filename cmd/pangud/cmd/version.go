package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "v",
	Long:  `pangud version number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version v0.1.0")
	},
}
