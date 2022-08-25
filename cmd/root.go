package cmd

import (
	"fmt"
	"os"

	"github.com/eaglerayp/echoserver/server"
	"github.com/spf13/cobra"
)

var (
	port int
)

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().IntVar(&port, "port", 8080, "default is 8080")
}

var rootCmd = &cobra.Command{
	Use:   "echoserver",
	Short: "echoserver is a simple api server for local testing purposes",
	Long: `echoserver is a simple api server for local testing purposes 
		implemented by go gin server print and echo whole input`,
}

var serverCmd = &cobra.Command{
	Use:   "http",
	Short: "start http server",
	Long:  `start http server with given port`,
	Run: func(cmd *cobra.Command, args []string) {
		server.StartGinServer(port)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
