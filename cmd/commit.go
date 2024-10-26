package cmd

import (
	"github.com/spf13/cobra"
)

var Glob string

func init() {
	commitCMD.Flags().StringVarP(&Glob, "glob", "g", "*", "Glob to select files for commit")
}

var commitCMD = &cobra.Command{
	Use:   "commit [msg]",
	Short: "commit one or more files with a message.",
	Long: `Commit one or more files with a message. The glob is relative 
					to to the Dots directory. By default all files are commited.`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		return client.Commit(cmd.Context(), Glob, args[0])
	},
}
