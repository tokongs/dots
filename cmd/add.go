package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/tokongs/dots/dots"
)

var client = dots.Dots{
	Directory:  "/home/tokongs/.dots",
	RelativeTo: "/home/tokongs",
}

var addCMD = &cobra.Command{
	Use:   "add [files]",
	Short: "Add files or directories to be tracked",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("must provide at least one path")
		}

		return client.Add(args[0])
	},
}
