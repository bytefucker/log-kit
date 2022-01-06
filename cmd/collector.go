package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yihongzhi/log-kit/collector"
	"github.com/yihongzhi/log-kit/config"
	"os"
)

var (
	cfgFile         string
	collectorConfig config.CollectorConfig
)

var collectorCmd = &cobra.Command{
	Use:   "collector",
	Short: "log collector",
	Run: func(cmd *cobra.Command, args []string) {
		level, _ := log.ParseLevel(collectorConfig.LogLevel)
		log.SetLevel(level)
		log.SetOutput(os.Stdout)
		collector, err := collector.NewLogCollector(&collectorConfig)
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

func init() {
	cobra.OnInitialize(initConfig)
	collectorCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config/collector.yaml)")
	rootCmd.AddCommand(collectorCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		pwd, err := os.Getwd()
		cobra.CheckErr(err)
		viper.AddConfigPath(pwd)
		viper.SetConfigType("yaml")
		viper.SetConfigName("collector")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stdout, "load config file:", viper.ConfigFileUsed())
		viper.Unmarshal(&collectorConfig)
	} else {
		fmt.Fprintln(os.Stderr, "load config file error:", err)
	}
}
