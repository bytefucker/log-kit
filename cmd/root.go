package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "log-kit",
	Short: "日志收集组件",
}

func Execute() error {
	return rootCmd.Execute()
}
