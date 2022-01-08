package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yihongzhi/log-kit/config"
	"os"
)

var (
	cfgFile   string
	appConfig config.AppConfig
)

var rootCmd = &cobra.Command{
	Use:   "log-kit",
	Short: "日志收集组件",
}

func init() {
	cobra.OnInitialize(initConfig)
	collectorCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.yaml)")
	rootCmd.AddCommand(collectorCmd)
	rootCmd.AddCommand(analyzerCmd)
	rootCmd.AddCommand(managerCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		pwd, err := os.Getwd()
		cobra.CheckErr(err)
		viper.AddConfigPath(pwd)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stdout, "load config file success:", viper.ConfigFileUsed())
		viper.Unmarshal(&appConfig)
	} else {
		fmt.Fprintln(os.Stderr, "load config file error:", err)
	}
}

func Execute() error {
	return rootCmd.Execute()
}
