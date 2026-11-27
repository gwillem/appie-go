package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"text/tabwriter"

	appie "github.com/gwillem/appie-go"
)

type basketCommand struct {
	Show    basketShowCommand    `command:"show" description:"Show items in basket"`
	Add     basketAddCommand     `command:"add" description:"Add an item to the basket"`
	Rm      basketRmCommand      `command:"rm" description:"Remove an item from the basket"`
	Check   basketCheckCommand   `command:"check" description:"Cross an item off your list"`
	Uncheck basketUncheckCommand `command:"uncheck" description:"Un-cross an item on your list"`
}

// Executing 'appie basket' directly will run this default 'Show' logic
func (cmd *basketCommand) Execute(args []string) error {
	return (&basketShowCommand{}).Execute(args)
}

// ============================================================================
// SHOW COMMAND
// ============================================================================

type basketShowCommand struct{}

func (cmd *basketShowCommand) Execute(args []string) error {
	client, err := appie.NewWithConfig(globalOpts.Config, clientOpts()...)
	if err != nil {
		return err
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	const query = `
	query FetchMyBasket {
		basket {
			items {
				id
				quantity
				isStrikethrough
				product {
					title
					price { now { amount } }
				}
			}
		}
	}`

	var response struct {
		Data struct {
			Basket struct {
				Items []struct {
					ID              int  `json:"id"`
					Quantity        int  `json:"quantity"`
					IsStrikethrough bool `json:"isStrikethrough"`
					Product         struct {
						Title string `json:"title"`
						Price struct {
							Now struct {
								Amount float64 `json:"amount"`
							} `json:"now"`
						} `json:"price"`
					} `json:"product"`
				} `json:"items"`
			} `json:"basket"`
		} `json:"data"`
	}

	if err := client.DoGraphQL(ctx, query, nil, &response.Data); err != nil {
		return fmt.Errorf("failed to fetch basket: %w", err)
	}

	items := response.Data.Basket.Items
	if len(items) == 0 {
		fmt.Println("Basket is empty")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, item := range items {
		checkedMarker := " "
		if item.IsStrikethrough {
			checkedMarker = "X"
		}

		price := item.Product.Price.Now.Amount * float64(item.Quantity)
		fmt.Fprintf(w, "[%s] %d\t%s\tx%d\t€%.2f\n", checkedMarker, item.ID, item.Product.Title, item.Quantity, price)
	}
	return w.Flush()
}

// ============================================================================
// CHECK / UNCHECK COMMANDS & HELPER
// ============================================================================

// Helper function to handle the GraphQL mutation for both Check and Uncheck
func setStrikethroughState(ctx context.Context, client *appie.Client, productID int, isChecked bool) error {
	// 1. Fetch current basket to get the item's current quantity
	const query = `
	query FetchMyBasketForToggle {
		basket {
			items {
				id
				quantity
			}
		}
	}`

	var queryResponse struct {
		Data struct {
			Basket struct {
				Items []struct {
					ID       int `json:"id"`
					Quantity int `json:"quantity"`
				} `json:"items"`
			} `json:"basket"`
		} `json:"data"`
	}

	if err := client.DoGraphQL(ctx, query, nil, &queryResponse.Data); err != nil {
		return fmt.Errorf("failed to fetch basket for current quantity: %w", err)
	}

	// 2. Find the product's current quantity
	var currentQuantity int
	var found bool
	for _, item := range queryResponse.Data.Basket.Items {
		if item.ID == productID {
			currentQuantity = item.Quantity
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("product %d not found in your basket", productID)
	}

	// 3. Perform the mutation using the fetched quantity
	const mutation = `
	mutation UpdateMyListBasket($items: [BasketMutation!]!, $input: BasketInput) {
		basketItemsUpdate(items: $items, input: $input) {
			status
		}
	}`

	variables := map[string]interface{}{
		"input": nil,
		"items": []map[string]interface{}{
			{
				"id":              productID,
				"quantity":        currentQuantity,
				"isStrikethrough": isChecked,
			},
		},
	}

	return client.DoGraphQL(ctx, mutation, variables, nil)
}

// Check subcommand
type basketCheckCommand struct {
	Args struct {
		ProductID int `positional-arg-name:"product-id" required:"true"`
	} `positional-args:"yes"`
}

func (cmd *basketCheckCommand) Execute(args []string) error {
	client, err := appie.NewWithConfig(globalOpts.Config, clientOpts()...)
	if err != nil {
		return err
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := setStrikethroughState(ctx, client, cmd.Args.ProductID, true); err != nil {
		return err
	}

	fmt.Printf("Checked item %d\n", cmd.Args.ProductID)
	return nil
}

// Uncheck subcommand
type basketUncheckCommand struct {
	Args struct {
		ProductID int `positional-arg-name:"product-id" required:"true"`
	} `positional-args:"yes"`
}

func (cmd *basketUncheckCommand) Execute(args []string) error {
	client, err := appie.NewWithConfig(globalOpts.Config, clientOpts()...)
	if err != nil {
		return err
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := setStrikethroughState(ctx, client, cmd.Args.ProductID, false); err != nil {
		return err
	}

	fmt.Printf("Unchecked item %d\n", cmd.Args.ProductID)
	return nil
}

// ============================================================================
// ADD / RM COMMANDS
// ============================================================================

// Add subcommand
type basketAddCommand struct {
	Args struct {
		ProductID int `positional-arg-name:"product-id" required:"true"`
		Quantity  int `positional-arg-name:"quantity" default:"1"`
	} `positional-args:"yes"`
}

func (cmd *basketAddCommand) Execute(args []string) error {
	client, err := appie.NewWithConfig(globalOpts.Config, clientOpts()...)
	if err != nil {
		return err
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	const mutation = `
	mutation UpdateMyListBasket($items: [BasketMutation!]!) {
		basketItemsUpdate(items: $items) { status }
	}`

	variables := map[string]interface{}{
		"items": []map[string]interface{}{{"id": cmd.Args.ProductID, "quantity": cmd.Args.Quantity}},
	}

	if err := client.DoGraphQL(ctx, mutation, variables, nil); err != nil {
		return err
	}
	fmt.Printf("Added %dx %d to basket\n", cmd.Args.Quantity, cmd.Args.ProductID)
	return nil
}

// Rm subcommand
type basketRmCommand struct {
	Args struct {
		ProductID int `positional-arg-name:"product-id" required:"true"`
	} `positional-args:"yes"`
}

func (cmd *basketRmCommand) Execute(args []string) error {
	client, err := appie.NewWithConfig(globalOpts.Config, clientOpts()...)
	if err != nil {
		return err
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	const mutation = `
	mutation UpdateMyListBasket($items: [BasketMutation!]!) {
		basketItemsUpdate(items: $items) { status }
	}`

	variables := map[string]interface{}{
		"items": []map[string]interface{}{{"id": cmd.Args.ProductID, "quantity": 0}},
	}

	if err := client.DoGraphQL(ctx, mutation, variables, nil); err != nil {
		return err
	}
	fmt.Printf("Removed %d from basket\n", cmd.Args.ProductID)
	return nil
}
