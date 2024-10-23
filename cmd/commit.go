package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var Glob string

var commitCMD = &cobra.Command{
	Use:   "commit [msg]",
	Short: "commit one or more files with a message.",
	Long: "Commit one or more files with a message." +
		" The glob is relative to to the Dots directory." +
		" By default all files are commited.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("commit takes 1 argument")
		}

		return client.Commit(Glob, args[0])
	},
}
