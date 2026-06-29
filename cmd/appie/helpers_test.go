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

func TestOrderCommandListStatus(t *testing.T) {
	tests := []struct {
		name      string
		cmd       orderCommand
		want      appie.FulfillmentStatus
		wantLabel string
		wantErr   bool
	}{
		{
			name:      "default open",
			want:      appie.FulfillmentStatusOpen,
			wantLabel: "open",
		},
		{
			name:      "closed",
			cmd:       orderCommand{Closed: true},
			want:      appie.FulfillmentStatusClosed,
			wantLabel: "closed",
		},
		{
			name:      "all",
			cmd:       orderCommand{All: true},
			want:      appie.FulfillmentStatusAll,
			wantLabel: "all",
		},
		{
			name:    "closed and all conflict",
			cmd:     orderCommand{Closed: true, All: true},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotLabel, err := tt.cmd.listStatus()
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("status = %q, want %q", got, tt.want)
			}
			if gotLabel != tt.wantLabel {
				t.Fatalf("label = %q, want %q", gotLabel, tt.wantLabel)
			}
		})
	}
}
