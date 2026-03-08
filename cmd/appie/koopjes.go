package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	appie "github.com/gwillem/appie-go"
)

type koopjesCommand struct {
	Args struct {
		PostalCode string `positional-arg-name:"postcode" required:"true" description:"Postal code (e.g. 3521GZ)"`
	} `positional-args:"yes"`
}

func (cmd *koopjesCommand) Execute(args []string) error {
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	stores, err := client.SearchStores(ctx, cmd.Args.PostalCode)
	if err != nil {
		return fmt.Errorf("store search failed: %w", err)
	}
	if len(stores) == 0 {
		return fmt.Errorf("no stores found near %q", cmd.Args.PostalCode)
	}

	store := stores[0]
	fmt.Fprintf(os.Stderr, "%s %s, %s\n\n",
		store.Address.Street, store.Address.HouseNumber, store.Address.City)

	bargains, err := client.GetBargains(ctx, store.ID)
	if err != nil {
		return fmt.Errorf("failed to get bargains: %w", err)
	}

	if len(bargains) == 0 {
		fmt.Println("No bargains available at this store right now.")
		return nil
	}

	printBargains(bargains)
	return nil
}

func printBargains(bargains []appie.Bargain) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, b := range bargains {
		fmt.Fprintf(w, "  %s\t%s\t%s → %s\t-%g%%\t×%d\t%s\n",
			b.Product.Title,
			b.Product.UnitSize,
			b.PriceWas, b.PriceNow,
			b.MarkdownPercentage,
			b.Stock,
			b.ExpirationDate,
		)
	}
	w.Flush()
}
