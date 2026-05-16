package main

import (
	"fmt"
	"os"

	appie "github.com/gwillem/appie-go"
)

type productCommand struct {
	Args struct {
		IDs []int `positional-arg-name:"id" required:"1"`
	} `positional-args:"yes"`
	Detail bool `short:"d" long:"detail" description:"Show full detail (incl. nutrition) for each product"`
}

func (cmd *productCommand) Execute(args []string) error {
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	ids := cmd.Args.IDs

	products, err := client.GetProductsByIDs(ctx, ids)
	if err != nil {
		return fmt.Errorf("get products failed: %w", err)
	}

	for _, id := range missingIDs(ids, products) {
		fmt.Fprintf(os.Stderr, "warning: product %d not found\n", id)
	}

	if len(products) == 0 {
		return fmt.Errorf("no products resolved")
	}

	if !cmd.Detail {
		printProducts(products)
		return nil
	}

	for i, p := range products {
		if i > 0 {
			fmt.Println()
		}
		full, err := client.GetProductFull(ctx, p.ID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: nutrition fetch for %d failed: %v\n", p.ID, err)
			printProductDetail(&p)
			continue
		}
		printProductDetail(full)
	}
	return nil
}

// missingIDs returns the input IDs that are not present in products.
// Order matches the input order; duplicates in input are reported once each.
func missingIDs(ids []int, products []appie.Product) []int {
	got := make(map[int]struct{}, len(products))
	for _, p := range products {
		got[p.ID] = struct{}{}
	}
	var missing []int
	for _, id := range ids {
		if _, ok := got[id]; !ok {
			missing = append(missing, id)
		}
	}
	return missing
}

func printProductDetail(p *appie.Product) {
	fmt.Printf("%s\n", p.Title)
	if p.Brand != "" {
		fmt.Printf("  Brand:       %s\n", p.Brand)
	}
	fmt.Printf("  ID:          %d\n", p.ID)
	if p.UnitSize != "" {
		fmt.Printf("  Unit size:   %s\n", p.UnitSize)
	}
	if p.Price.Was > 0 && p.Price.Was != p.Price.Now {
		fmt.Printf("  Price:       €%.2f (was €%.2f)\n", p.Price.Now, p.Price.Was)
	} else {
		fmt.Printf("  Price:       €%.2f\n", p.Price.Now)
	}
	if p.UnitPriceDescription != "" {
		fmt.Printf("  Unit price:  %s\n", p.UnitPriceDescription)
	}
	if p.BonusMechanism != "" {
		fmt.Printf("  Bonus:       %s\n", p.BonusMechanism)
	}
	if p.Category != "" {
		cat := p.Category
		if p.SubCategory != "" {
			cat += " / " + p.SubCategory
		}
		fmt.Printf("  Category:    %s\n", cat)
	}
	if p.NutriScore != "" {
		fmt.Printf("  Nutri-Score: %s\n", p.NutriScore)
	}
	fmt.Printf("  Available:   %t\n", p.IsAvailable)
	fmt.Printf("  Orderable:   %t\n", p.IsOrderable)
	if p.ShortDescription != "" {
		fmt.Printf("\n  %s\n", p.ShortDescription)
	}
	if len(p.NutritionalInfo) > 0 {
		fmt.Println("\n  Nutrition:")
		var lastPer string
		headerPrinted := false
		for _, n := range p.NutritionalInfo {
			if !headerPrinted || n.Per != lastPer {
				if headerPrinted {
					fmt.Println()
				}
				header := n.Per
				if header == "" {
					header = "(unspecified basis)"
				} else {
					header = "per " + header
				}
				fmt.Printf("    %s:\n", header)
				lastPer = n.Per
				headerPrinted = true
			}
			fmt.Printf("      %-30s %s\n", n.Name, n.Value)
		}
	}
}
