package main

import (
	"fmt"
	"log"
	"os"

	appie "github.com/gwillem/appie-go"
	"github.com/jessevdk/go-flags"
)

// version is set at build time via -ldflags "-X main.version=..."
var version = "dev"

var globalOpts struct {
	Config  string `short:"c" long:"config" description:"Path to config file"`
	Verbose bool   `short:"v" long:"verbose" description:"Verbose output"`
	Country string `long:"country" description:"Albert Heijn country: nl or be (default: nl, or last used at login)"`

	Login   loginCommand        `command:"login" description:"Login to Albert Heijn"`
	Search  searchCommand       `command:"search" description:"Search for products"`
	Receipt receiptCommand      `command:"receipt" subcommands-optional:"true" description:"List recent receipts"`
	Order   orderCommand        `command:"order" subcommands-optional:"true" description:"List open orders"`
	List    shoppingListCommand `command:"list" subcommands-optional:"true" description:"Show shopping lists"`
	Koopjes koopjesCommand      `command:"koopjes" description:"Show last-chance bargains at a store"`
	Update  updateCommand       `command:"update" description:"Update appie to the latest version"`
}

func clientOpts() []appie.Option {
	opts := []appie.Option{appie.WithConfigPath(globalOpts.Config)}
	if globalOpts.Country != "" {
		opts = append(opts, appie.WithCountry(globalOpts.Country))
	}
	if globalOpts.Verbose {
		opts = append(opts, appie.WithLogger(log.New(os.Stderr, "", log.Ltime)))
	}
	return opts
}

func main() {
	if globalOpts.Config == "" {
		globalOpts.Config = appie.DefaultConfigPath()
	}

	p := flags.NewParser(&globalOpts, flags.Default)
	if _, err := p.Parse(); err != nil {
		if flags.WroteHelp(err) {
			os.Exit(0)
		}
		os.Exit(1)
	}
}

func init() {
	for _, arg := range os.Args[1:] {
		if arg == "--version" || arg == "-V" {
			fmt.Println("appie", version)
			os.Exit(0)
		}
	}
}
