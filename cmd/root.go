package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "dots",
	Short: "Dots is a minimalistic dotfiles manager",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func setupCommands() {
	rootCMD.AddCommand(initCMD)
	rootCMD.AddCommand(addCMD)
}

func SetupAndExecute() {
	setupCommands()

	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
