package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCMD = &cobra.Command{
	Use:   "dots",
	Short: "Dots is a minimalistic dotfiles manager",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("error")
	},
}
