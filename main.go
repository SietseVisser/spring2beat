package main

import (
	"os"

	"github.com/SietseVisser/spring2beat/cmd"

	_ "github.com/SietseVisser/spring2beat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
