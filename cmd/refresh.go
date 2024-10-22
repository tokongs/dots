package cmd

import "github.com/spf13/cobra"

var refreshCMD = &cobra.Command{
	Use:   "refresh",
	Short: "Refreshes all tracked files into the Dots directory.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return client.Refresh()
	},
}
