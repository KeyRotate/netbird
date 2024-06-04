package main

import (
	"github.com/keyrotate/netbird/signal/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
