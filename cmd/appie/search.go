package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"text/tabwriter"

	appie "github.com/gwillem/appie-go"
)

type searchCommand struct {
	Args struct {
		Query string `positional-arg-name:"query" required:"true"`
	} `positional-args:"yes"`
	Limit int `short:"n" long:"limit" default:"20" description:"Max number of results"`
}

func (cmd *searchCommand) Execute(args []string) error {
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	products, err := client.SearchProducts(ctx, cmd.Args.Query, cmd.Limit)
	if err != nil {
		return fmt.Errorf("search failed: %w", err)
	}
	if len(products) == 0 {
		return fmt.Errorf("no products found for %q", cmd.Args.Query)
	}

	printProducts(products)
	return nil
}

func printProducts(products []appie.Product) {
	slices.SortFunc(products, func(a, b appie.Product) int {
		return cmp.Compare(a.Price.Now, b.Price.Now)
	})
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, p := range products {
		fmt.Fprintf(w, "  %d\t%s\t%s\t€%.2f\n", p.ID, p.Title, p.UnitSize, p.Price.Now)
	}
	w.Flush()
}
