package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yihongzhi/log-kit/analyzer"
	"os"
)

var analyzerCmd = &cobra.Command{
	Use:   "analyzer",
	Short: "log analyzer",
	Run: func(cmd *cobra.Command, args []string) {
		if appConfig.LogLevel != "" {
			log.SetLevel(appConfig.LogLevel)
		}
		analyzer, err := analyzer.NewLogAnalyzer(&appConfig)
		if err != nil {
			log.Errorln("init log analyzer error", err)
			os.Exit(1)
			return
		}
		if analyzer.Start(); err != nil {
			log.Errorln("start log analyzer error", err)
			os.Exit(1)
			return
		}
	},
}
