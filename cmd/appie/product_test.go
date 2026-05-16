package main

import (
	"slices"
	"testing"

	appie "github.com/gwillem/appie-go"
)

func TestMissingIDs(t *testing.T) {
	tests := []struct {
		name     string
		ids      []int
		products []appie.Product
		want     []int
	}{
		{
			name:     "all present",
			ids:      []int{1, 2, 3},
			products: []appie.Product{{ID: 1}, {ID: 2}, {ID: 3}},
			want:     nil,
		},
		{
			name:     "some missing",
			ids:      []int{1, 2, 3, 4},
			products: []appie.Product{{ID: 1}, {ID: 3}},
			want:     []int{2, 4},
		},
		{
			name:     "all missing",
			ids:      []int{1, 2},
			products: nil,
			want:     []int{1, 2},
		},
		{
			name:     "preserves input order",
			ids:      []int{9, 5, 7},
			products: []appie.Product{{ID: 5}},
			want:     []int{9, 7},
		},
		{
			name:     "duplicate input id reported each time",
			ids:      []int{1, 1, 2},
			products: []appie.Product{{ID: 2}},
			want:     []int{1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := missingIDs(tt.ids, tt.products)
			if !slices.Equal(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}
