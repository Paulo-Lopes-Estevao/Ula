package cmd

import (
	"github.com/ebizno/Ula/server"
	"github.com/spf13/cobra"
)

var portServer string

func init() {
	CmdServer.Flags().StringVarP(&portServer, "port", "p", "8080", "Port")
}

var CmdServer = &cobra.Command{
	Use:   "server",
	Short: "Ula server",
	Long:  `Ula server`,
	RunE:  ExecuteServer,
}

func ExecuteServer(cmd *cobra.Command, args []string) error {
	server.NewServer(portServer).Start()
	return nil
}
