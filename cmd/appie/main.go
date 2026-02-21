package main

import (
	"log"
	"os"
	"path/filepath"

	appie "github.com/gwillem/appie-go"
	"github.com/jessevdk/go-flags"
)

var globalOpts struct {
	Config  string         `short:"c" long:"config" description:"Path to config file"`
	Verbose bool           `short:"v" long:"verbose" description:"Verbose output"`
	Login   loginCommand   `command:"login" description:"Login to Albert Heijn"`
	Receipt receiptCommand `command:"receipt" description:"List recent receipts"`
	Order   orderCommand   `command:"order" subcommands-optional:"true" description:"Show active order contents"`
}

func clientOpts() []appie.Option {
	opts := []appie.Option{appie.WithConfigPath(globalOpts.Config)}
	if globalOpts.Verbose {
		opts = append(opts, appie.WithLogger(log.New(os.Stderr, "", log.Ltime)))
	}
	return opts
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
		os.Exit(1)
	}
}
