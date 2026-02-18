package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	appie "github.com/gwillem/appie-go"
)

type loginCommand struct{}

func (cmd *loginCommand) Execute(args []string) error {
	configPath := globalOpts.Config

	if err := os.MkdirAll(filepath.Dir(configPath), 0o700); err != nil {
		return fmt.Errorf("create config dir: %w", err)
	}

	client := appie.New(appie.WithConfigPath(configPath))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	fmt.Println("Opening browser for AH login...")

	if err := client.Login(ctx); err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	fmt.Printf("Login successful! Tokens saved to %s\n", configPath)
	return nil
}
