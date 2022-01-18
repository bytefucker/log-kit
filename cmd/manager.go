package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yihongzhi/log-kit/manager"
	"os"
)

var managerCmd = &cobra.Command{
	Use:   "manager",
	Short: "log web manager",
	Run: func(cmd *cobra.Command, args []string) {
		if appConfig.LogLevel != "" {
			log.SetLevel(appConfig.LogLevel)
		}
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
