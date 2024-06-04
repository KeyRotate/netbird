package main

import (
	"os"

	"github.com/keyrotate/netbird/client/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
