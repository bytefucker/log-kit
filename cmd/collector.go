package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yihongzhi/log-kit/collector"
	"os"
)

var collectorCmd = &cobra.Command{
	Use:   "collector",
	Short: "log collector",
	Run: func(cmd *cobra.Command, args []string) {
		level, _ := log.ParseLevel(appConfig.LogLevel)
		log.SetLevel(level)
		log.SetOutput(os.Stdout)
		collector, err := collector.NewLogCollector(&appConfig)
		if err != nil {
			log.Errorln("Init LogCollector error", err)
			os.Exit(1)
		}
		if err := collector.Start(); err != nil {
			log.Errorln("Start LogCollector error", err)
			os.Exit(1)
		}
		log.Infoln("Start LogCollector Success")
	},
}
