package cmd

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

var Pull bool

func init() {
	applyCMD.Flags().BoolVarP(&Pull, "pull", "p", true, "Pull from repo before applying files.")
}

var applyCMD = &cobra.Command{
	Use:   "apply",
	Short: "Copies the dotfiles from the repo to their correct location",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
		defer cancel()

		return client.Apply(ctx, Pull)
	},
}
