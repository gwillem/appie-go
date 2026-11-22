package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	appie "github.com/gwillem/appie-go"
)

type loginCommand struct {
	Port int `long:"port" env:"APPIE_LOGIN_PORT" default:"0" description:"Fixed port for the local login page (http://localhost:<port>); 0 picks a random port"`
}

func (cmd *loginCommand) Execute(args []string) error {
	configPath := globalOpts.Config

	if err := os.MkdirAll(filepath.Dir(configPath), 0o700); err != nil {
		return fmt.Errorf("create config dir: %w", err)
	}

	client := appie.New(append(clientOpts(), appie.WithLoginPort(cmd.Port))...)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	fmt.Println("Opening browser for AH login. If it doesn't open, visit:")

	if err := client.Login(ctx); err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	fmt.Printf("Login successful! Tokens saved to %s\n", configPath)
	fmt.Println("After you have been authorized, the access keys will be automatically refreshed.")
	return nil
}
