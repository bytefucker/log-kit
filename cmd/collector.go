package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
)

var collectorCmd = &cobra.Command{
	Use:   "collector",
	Short: "collector",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var cfgFile string

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
		viper.AddConfigPath(path.Join(pwd, "config"))
		viper.SetConfigType("yaml")
		viper.SetConfigName("collector")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Fprintln(os.Stderr, "reda config file error:", err)
	}
}
