# CLI Interface

## Global options

```
appie [options] <command>

  -c, --config PATH    Config file (default: ~/.config/appie/config.json)
  -v, --verbose        Verbose output
```

## Command conventions

All resource commands follow a consistent pattern:

```
appie <resource>                              Overview (list all)
appie <resource> show <target>                Detail view for one item
appie <resource> add <target> <item> [flags]  Add item to target
appie <resource> rm <target> <item>           Remove item from target
```

- Bare command always shows an overview/list, including IDs for use with subcommands.
- Mutations always require an explicit target ID as the first positional arg.
- `<target>` IDs must be specified in full (order IDs are integers, list IDs are UUIDs).
- `<item>` can be a numeric product ID or a search term (resolved to a single match).

## Commands

### `login`

Open browser for AH OAuth login.

### `search <query>`

Search products by name. Results sorted by price (low to high). Use product IDs with `order add` or `list add`.

```
  -n, --limit NUM    Max results (default: 20)
```

```
$ appie search "pindakaas"
  192450  AH Biologisch pindakaas          350 g  €2.89
  371880  AH Pindakaas naturel             600 g  €3.59
  204835  Calvé Pindakaas                  650 g  €4.49
  ...
```

### `bonus`

List bonus (promotion) products. Defaults to today.

```
  -d, --date DATE    Bonus date in YYYY-MM-DD format (default: today)
  -s, --spotlight    Show only spotlight/featured deals
```

```
$ appie bonus
Bonus week 21 feb - 27 feb

Product                          Was    Now    Discount
AH Verse pizza               →  4.99   3.49   30% korting
Alle Hak 370-720ml           →  2.19   1.09   2e halve prijs
Dove Douchegel 250ml         →  3.29   1.99   1+1 gratis
...
(247 products)
```

### `receipt`

List recent receipts.

```
  -n NUM    Number of recent receipts (default: 20)
```

```
$ appie receipt
2025-02-21T14:30:00  TXN-ABC123   42.50
2025-02-18T10:15:00  TXN-DEF456   87.30
...
```

#### `receipt show <transaction-id>`

Show items for a specific receipt.

```
$ appie receipt show TXN-ABC123
Receipt TXN-ABC123
Date:  2025-02-21T14:30:00

  2x AH Halfvolle melk               3.78
     AH Pindakaas naturel            3.59
     Brinta Origineel                 2.49

     Bonus korting                   -1.20
                                     ------
     PIN                             42.50
```

### `order`

List all open/scheduled orders (fulfillments).

```
$ appie order
  Order      Status     Delivery                         Total
  1234567  SUBMITTED  dinsdag 25 feb  18:00-20:00       87.30
  1234590  SUBMITTED  vrijdag 28 feb  08:00-10:00       42.15
```

#### `order show <order-id>`

Show contents of a specific order.

```
$ appie order show 1234567
Order 1234567  SUBMITTED
Delivery: dinsdag 25 feb  18:00-20:00

 1  AH Verse halfvolle melk 2L       1    1.89
 2  AH Pindakaas naturel             2    3.58
 3  Brinta Origineel                  1    2.49
                                    ──────
                            3 items   7.96
```

#### `order add <order-id> <product> [-n quantity]`

Add a product to an order. `<product>` can be a numeric product ID or a search term. Reopens the order if submitted.

```
  -n, --quantity NUM    Quantity to add (default: 1)
```

```
$ appie order add 1234567 371880
Added 1x 371880 to order 1234567

$ appie order add 1234567 371880 -n 3
Added 3x 371880 to order 1234567

$ appie order add 1234567 "halfvolle melk"
Found: AH Halfvolle melk
Added 1x 12345 to order 1234567

$ appie order add 1234567 "hagelslag puur"
  32786   AH Hagelslag puur                            250 g    €2.59
  55732   AH Chocoladepasta puur                       400 g    €2.69
  465752  AH Hagelslag xl puur                         600 g    €2.99
  1608    AH Hagelslag puur                            600 g    €2.99
  ...
multiple matches for "hagelslag puur", specify product ID
```

#### `order rm <order-id> <product-id>`

Remove a product from an order. Reopens the order if submitted.

```
$ appie order rm 1234567 371880
Removed Optimel Drinkyoghurt aardbei from order 1234567
```

### `list`

Show shopping lists overview.

```
$ appie list
  a1b2c3d4-e5f6-7890-abcd-ef1234567890  Boodschappen  12 items
  e5f6a7b8-c9d0-1234-5678-abcdef012345  Weekmenu       5 items
```

#### `list show <list-id>`

Show items in a list. Items are numbered for use with `list rm`.

```
$ appie list show a1b2c3d4-e5f6-7890-abcd-ef1234567890
Boodschappen (12 items)

 1  AH Halfvolle melk       1 L    1  €1.89
 2  AH Pindakaas naturel    600 g  2  €3.59
 3  kaas                           1
```

#### `list add <list-id> <product> [-n quantity]`

Add a product to a list. `<product>` can be a numeric product ID or a search term. It will overwrite any existing product quantities.

```
  -n, --quantity NUM    Quantity to add (default: 1)
```

```
$ appie list add a1b2c3d4-e5f6-7890-abcd-ef1234567890 371880
Added 1x 371880 to Boodschappen

$ appie list add a1b2c3d4-e5f6-7890-abcd-ef1234567890 "pindakaas"
Found: AH Pindakaas naturel
Added 1x 371880 to Boodschappen
```

#### `list rm <list-id> <product>`

Remove an item by its line number from `list show` output. If there are more products, remove all of them.

```
$ appie list rm a1b2c3d4-e5f6-7890-abcd-ef1234567890 3
Removed #3 (kaas) from Boodschappen
```
