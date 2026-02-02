package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
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
	ctx := context.Background()

	fmt.Println()
	fmt.Println("=== Albert Heijn Login ===")
	fmt.Println()
	fmt.Println("1. Open this URL in your browser:")
	fmt.Println()
	fmt.Printf("   %s\n", client.LoginURL())
	fmt.Println()
	fmt.Println("2. Login with your credentials")
	fmt.Println("3. After login, browser will try to open 'appie://login-exit?code=...'")
	fmt.Println("4. Copy the 'code' value from the URL (or the full URL)")
	fmt.Println()
	fmt.Print("Paste code here: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	code := extractCode(strings.TrimSpace(input))
	if code == "" {
		fmt.Fprintf(os.Stderr, "Could not extract code from input\n")
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("Exchanging code for tokens...")

	if err := client.ExchangeCode(ctx, code); err != nil {
		fmt.Fprintf(os.Stderr, "Token exchange failed: %v\n", err)
		os.Exit(1)
	}

	if err := client.SaveConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to save config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("=== LOGIN SUCCESSFUL ===")
	fmt.Println()
	fmt.Printf("Tokens saved to %s\n", configPath)
}

func extractCode(input string) string {
	if !strings.Contains(input, "=") && !strings.Contains(input, "?") {
		return input
	}

	if idx := strings.Index(input, "code="); idx != -1 {
		code := input[idx+5:]
		if ampIdx := strings.Index(code, "&"); ampIdx != -1 {
			code = code[:ampIdx]
		}
		return code
	}

	return ""
}
