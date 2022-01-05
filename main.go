package main

import (
	"github.com/yihongzhi/log-kit/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
