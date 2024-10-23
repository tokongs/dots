package cmd

import "github.com/spf13/cobra"

var applyCMD = &cobra.Command{
	Use:   "apply",
	Short: "Copies the dotfiles from the repo to their correct location",
	RunE: func(cmd *cobra.Command, args []string) error {
		return client.Apply()
	},
}
