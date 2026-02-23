package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	appie "github.com/gwillem/appie-go"
)

type shoppingListCommand struct {
	Show shoppingListShowCommand `command:"show" description:"Show items in a list"`
	Add  shoppingListAddCommand  `command:"add" description:"Add a product to a list"`
	Rm   shoppingListRmCommand   `command:"rm" description:"Remove an item from a list"`
}

func (cmd *shoppingListCommand) Execute(args []string) error {
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	lists, err := client.GetShoppingLists(ctx, 0)
	if err != nil {
		return fmt.Errorf("failed to get shopping lists: %w", err)
	}

	if len(lists) == 0 {
		fmt.Println("No shopping lists")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, l := range lists {
		fmt.Fprintf(w, "  %s\t%s\t%d items\n", l.ID, l.Name, l.ItemCount)
	}
	return w.Flush()
}

// findList finds a list by exact UUID match.
func findList(lists []appie.ShoppingList, id string) (*appie.ShoppingList, error) {
	for i, l := range lists {
		if l.ID == id {
			return &lists[i], nil
		}
	}
	return nil, fmt.Errorf("list %q not found", id)
}

// show subcommand

type shoppingListShowCommand struct {
	Args struct {
		ListID string `positional-arg-name:"list-id" required:"true"`
	} `positional-args:"yes"`
}

func (cmd *shoppingListShowCommand) Execute(args []string) error {
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	lists, err := client.GetShoppingLists(ctx, 0)
	if err != nil {
		return fmt.Errorf("failed to get shopping lists: %w", err)
	}

	list, err := findList(lists, cmd.Args.ListID)
	if err != nil {
		return err
	}

	items, err := client.GetShoppingListItems(ctx, list.ID)
	if err != nil {
		return fmt.Errorf("failed to get list items: %w", err)
	}

	fmt.Printf("%s (%d items)\n\n", list.Name, len(items))

	if len(items) == 0 {
		fmt.Println("No items")
		return nil
	}

	// Enrich items with product details
	var productIDs []int
	for _, item := range items {
		if item.ProductID > 0 {
			productIDs = append(productIDs, item.ProductID)
		}
	}

	products := make(map[int]*appie.Product)
	if len(productIDs) > 0 {
		ps, err := client.GetProductsByIDs(ctx, productIDs)
		if err != nil {
			return fmt.Errorf("failed to fetch product details: %w", err)
		}
		for i := range ps {
			products[ps[i].ID] = &ps[i]
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, item := range items {
		p := products[item.ProductID]
		title := item.Name
		unitSize := ""
		price := ""
		if p != nil {
			title = p.Title
			unitSize = p.UnitSize
			if p.Price.Now > 0 {
				price = fmt.Sprintf("€%.2f", p.Price.Now)
			}
		}
		if title == "" {
			title = fmt.Sprintf("(product %d)", item.ProductID)
		}
		fmt.Fprintf(w, "  %d\t%s\t%s\t%d\t%s\n", item.ProductID, title, unitSize, item.Quantity, price)
	}
	return w.Flush()
}

// add subcommand

type shoppingListAddCommand struct {
	Args struct {
		ListID  string `positional-arg-name:"list-id" required:"true"`
		Product string `positional-arg-name:"product" required:"true"`
	} `positional-args:"yes"`
	Quantity int `short:"n" long:"quantity" default:"1" description:"Quantity to add"`
}

func (cmd *shoppingListAddCommand) Execute(args []string) error {
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	lists, err := client.GetShoppingLists(ctx, 0)
	if err != nil {
		return fmt.Errorf("failed to get shopping lists: %w", err)
	}

	list, err := findList(lists, cmd.Args.ListID)
	if err != nil {
		return err
	}

	product := cmd.Args.Product

	// If numeric, use as product ID directly
	productID, err := strconv.Atoi(product)
	if err != nil {
		// Search for the product
		products, err := client.SearchProducts(ctx, product, 15)
		if err != nil {
			return fmt.Errorf("search failed: %w", err)
		}
		if len(products) == 0 {
			return fmt.Errorf("no products found for %q", product)
		}
		if len(products) > 1 {
			printProducts(products)
			return fmt.Errorf("multiple matches for %q, specify product ID", product)
		}
		productID = products[0].ID
		fmt.Printf("Found: %s\n", products[0].Title)
	}

	if err := client.AddToFavoriteList(ctx, list.ID, []appie.ListItem{{ProductID: productID, Quantity: cmd.Quantity}}); err != nil {
		return err
	}

	fmt.Printf("Added %dx %d to %s\n", cmd.Quantity, productID, list.Name)
	return nil
}

// rm subcommand

type shoppingListRmCommand struct {
	Args struct {
		ListID    string `positional-arg-name:"list-id" required:"true"`
		ProductID int    `positional-arg-name:"product-id" required:"true"`
	} `positional-args:"yes"`
}

func (cmd *shoppingListRmCommand) Execute(args []string) error {
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	lists, err := client.GetShoppingLists(ctx, 0)
	if err != nil {
		return fmt.Errorf("failed to get shopping lists: %w", err)
	}

	list, err := findList(lists, cmd.Args.ListID)
	if err != nil {
		return err
	}

	if err := client.RemoveFromFavoriteList(ctx, list.ID, []int{cmd.Args.ProductID}); err != nil {
		return err
	}

	fmt.Printf("Removed %d from %s\n", cmd.Args.ProductID, list.Name)
	return nil
}
