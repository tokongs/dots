package cmd

import (
	"github.com/spf13/cobra"
)

var ApplyAfterInit bool

func init() {
	initCMD.Flags().BoolVarP(&ApplyAfterInit, "apply", "a", true, "Apply dotsfiles after initialization.")
}

var initCMD = &cobra.Command{
	Use:   "init [repo-url]",
	Short: "Initializes a dots directory in the given path",
	Long: `Initializes the Dots repo with the provided repo.
					The dotfiles will also be applied.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := client.Clone(cmd.Context(), args[0]); err != nil {
			return err
		}
		if !ApplyAfterInit {
			return nil
		}

		return client.Apply(cmd.Context(), false)
	},
}
