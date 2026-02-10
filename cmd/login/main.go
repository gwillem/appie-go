package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"

	appie "github.com/gwillem/appie-go"
)

const defaultConfigPath = ".appie.json"

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Where to write tokens? [%s]: ", defaultConfigPath)
	pathInput, _ := reader.ReadString('\n')
	configPath := strings.TrimSpace(pathInput)
	if configPath == "" {
		configPath = defaultConfigPath
	}

	client := appie.New(appie.WithConfigPath(configPath))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	fmt.Println("Opening browser for AH login...")

	if err := client.Login(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Login failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Login successful! Tokens saved to %s\n", configPath)
}
