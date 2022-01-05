package cmd

import (
	"fmt"
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
		collector, err := collector.NewLogCollector(&collectorConfig)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Init LogCollector error", err)
			os.Exit(1)
		}
		err = collector.Start()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Start LogCollector error", err)
			os.Exit(1)
		}
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
