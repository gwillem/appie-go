package main

import (
	"fmt"
	"runtime"

	selfupdate "github.com/gwillem/go-selfupdate"
)

type updateCommand struct{}

func (c *updateCommand) Execute(args []string) error {
	url := fmt.Sprintf(
		"https://github.com/gwillem/appie-go/releases/latest/download/appie-%s-%s",
		runtime.GOOS, runtime.GOARCH,
	)

	fmt.Printf("Current version: %s\n", version)
	fmt.Printf("Checking for updates...\n")

	updated, err := selfupdate.Update(url)
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	if updated {
		fmt.Println("Updated successfully!")
	} else {
		fmt.Println("Already up to date.")
	}

	return nil
}
