package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tokongs/dots/dots"
)

var client = dots.Dots{
	Directory:  filepath.Join(os.Getenv("HOME"), ".dots"),
	RelativeTo: os.Getenv("HOME"),
}

func init() {
	rootCMD.AddCommand(initCMD)
	rootCMD.AddCommand(statusCMD)
	rootCMD.AddCommand(addCMD)
	rootCMD.AddCommand(commitCMD)
	rootCMD.AddCommand(applyCMD)
	rootCMD.AddCommand(editCMD)

	rootCMD.PersistentFlags().StringVarP(
		&client.Directory,
		"directory",
		"d",
		filepath.Join(os.Getenv("HOME"), ".dots"),
		"Directory to keeep as staging area for Dots.",
	)
	rootCMD.MarkPersistentFlagDirname("directory")

	rootCMD.PersistentFlags().StringVarP(
		&client.RelativeTo,
		"relative-to",
		"r",
		os.Getenv("HOME"),
		"Where to store files relative to.",
	)
	rootCMD.MarkPersistentFlagDirname("relative-to")
}

var rootCMD = &cobra.Command{
	Use:   "dots",
	Short: "Dots is a minimalistic dotfiles manager",
}

func SetupAndExecute() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := rootCMD.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
