package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jessevdk/go-flags"
)

var globalOpts struct {
	Config  string         `short:"c" long:"config" description:"Path to config file"`
	Login   loginCommand   `command:"login" description:"Login to Albert Heijn"`
	Receipt receiptCommand `command:"receipt" description:"List recent receipts"`
}

func defaultConfigPath() string {
	dir := os.Getenv("XDG_CONFIG_HOME")
	if dir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return ".appie.json"
		}
		dir = filepath.Join(home, ".config")
	}
	return filepath.Join(dir, "appie", "config.json")
}

func main() {
	if globalOpts.Config == "" {
		globalOpts.Config = defaultConfigPath()
	}

	p := flags.NewParser(&globalOpts, flags.Default)
	if _, err := p.Parse(); err != nil {
		if flags.WroteHelp(err) {
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
