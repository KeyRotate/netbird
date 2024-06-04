package main

import (
	"github.com/keyrotate/netbird/management/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
