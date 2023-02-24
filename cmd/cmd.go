package cmd

import "github.com/spf13/cobra"

var version = "0.0.1"

var RootCmd = &cobra.Command{
	Use:     "ula",
	Version: version,
	Short:   "Ula is a CLI tool to send emails",
	Long:    `Ula is a CLI tool to send emails`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

func Execute() error {
	RootCmd.AddCommand(CmdEvent)
	RootCmd.AddCommand(CmdServer)
	return RootCmd.Execute()
}
