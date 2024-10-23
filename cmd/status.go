package cmd

import "github.com/spf13/cobra"

var statusCMD = &cobra.Command{
	Use:   "status",
	Short: "Print the status of the dotfiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		return client.Status()
	},
}
