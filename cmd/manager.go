package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yihongzhi/log-kit/manager"
	"os"
)

var managerCmd = &cobra.Command{
	Use:   "manager",
	Short: "log web manager",
	Run: func(cmd *cobra.Command, args []string) {
		level, _ := log.ParseLevel(appConfig.LogLevel)
		log.SetLevel(level)
		log.SetOutput(os.Stdout)
		server, err := manager.NewManagerServer(&appConfig)
		if err != nil {
			log.Errorln("init manager server error", err)
			os.Exit(1)
			return
		}
		if err = server.Start(); err != nil {
			log.Errorln("start manager server error", err)
			os.Exit(1)
			return
		}
	},
}
