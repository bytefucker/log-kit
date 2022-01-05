package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var collectorCmd = &cobra.Command{
	Use:   "collector",
	Short: "collector",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}
