package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "pangud serve",
	Long:  `start pangud`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server start")
		serve()
	},
}

func init() {
	serveCmd.PersistentFlags().StringVarP(&conf, "conf", "c", ".",
		"default is .")

	rootCmd.AddCommand(serveCmd)
}
func serve() {
	endpoint := initEndpoint()
	endpoint.Serve()
}
