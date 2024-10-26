package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var addCMD = &cobra.Command{
	Use:   "add [files]",
	Short: "Add files to track given a list of files",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("must provide at least one path")
		}

		return client.Add(args)
	},
}
