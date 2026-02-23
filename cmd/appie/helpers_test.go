package main

import (
	"testing"

	appie "github.com/gwillem/appie-go"
)

func TestFindList(t *testing.T) {
	lists := []appie.ShoppingList{
		{ID: "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", Name: "Boodschappen"},
		{ID: "11111111-2222-3333-4444-555555555555", Name: "Weekmenu"},
	}

	t.Run("exact match", func(t *testing.T) {
		got, err := findList(lists, "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Name != "Boodschappen" {
			t.Fatalf("got %q, want %q", got.Name, "Boodschappen")
		}
	})

	t.Run("not found", func(t *testing.T) {
		_, err := findList(lists, "00000000-0000-0000-0000-000000000000")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("prefix does not match", func(t *testing.T) {
		_, err := findList(lists, "aaaaaaaa")
		if err == nil {
			t.Fatal("expected error for prefix match, got nil")
		}
	})
}
