package cmd

import (
	"github.com/spf13/cobra"
)

var initCMD = &cobra.Command{
	Use:   "init [directory]",
	Short: "Initializes a dots directory in the given path",
	RunE: func(cmd *cobra.Command, args []string) error {
		client.Init()
		return nil
	},
}
