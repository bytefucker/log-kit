package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yihongzhi/log-kit/collector"
	"os"
)

var collectorCmd = &cobra.Command{
	Use:   "collector",
	Short: "log collector",
	Run: func(cmd *cobra.Command, args []string) {
		if appConfig.LogLevel != "" {
			log.SetLevel(appConfig.LogLevel)
		}
		collector, err := collector.NewCollector(&appConfig)
		if err != nil {
			log.Errorln("init log collector error", err)
			os.Exit(1)
		}
		if err := collector.ListenAndServe(); err != nil {
			log.Errorln("start log collector error", err)
			os.Exit(1)
		}
	},
}
