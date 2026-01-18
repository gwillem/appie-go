package main

import (
	"context"
	"fmt"
	"log"
	"sort"

	appie "github.com/gwillem/appie-go"
)

func main() {
	client, err := appie.NewWithConfig(".appie.json")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	flags, err := client.GetFeatureFlags(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Sort flags by name
	names := make([]string, 0, len(flags))
	for name := range flags {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Printf("Feature Flags (%d total):\n\n", len(flags))

	enabled := 0
	disabled := 0
	partial := 0

	for _, name := range names {
		pct := flags[name]
		status := ""
		if pct >= 100 {
			status = "ON"
			enabled++
		} else if pct == 0 {
			status = "OFF"
			disabled++
		} else {
			status = fmt.Sprintf("%d%%", pct)
			partial++
		}
		fmt.Printf("  %-55s %s\n", name, status)
	}

	fmt.Printf("\nSummary: %d enabled, %d disabled, %d partial rollout\n", enabled, disabled, partial)
}
