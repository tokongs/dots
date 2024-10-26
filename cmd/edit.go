package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

var ApplyAfterEdit bool

func init() {
	editCMD.Flags().BoolVarP(&ApplyAfterEdit, "apply", "a", true, "Apply the edited files after editing.")
}

var editCMD = &cobra.Command{
	Use:   "edit [path]",
	Short: "Opens up your editor to edit a file or directory",
	Long: `Open the $EDITOR with the given path as the sole argument.
					The path is optional. The editor is launched with the Dots 
					directory as the working directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		editor := os.Getenv("EDITOR")

		path := ""
		if len(args) > 0 {
			path = args[0]
		}

		if err := client.Edit(context.Background(), editor, path); err != nil {
			return err
		}

		if !ApplyAfterEdit {
			return nil
		}

		return client.Apply(cmd.Context(), false)
	},
}
