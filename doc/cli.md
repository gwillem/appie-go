# CLI Interface

## Global options

```
appie [options] <command>

  -c, --config PATH    Config file (default: ~/.config/appie/config.json)
  -v, --verbose        Verbose output
```

## Commands

### `login`

Open browser for AH OAuth login.

### `receipt [transaction-id]`

List recent receipts, or show details for a specific transaction.

```
  -n NUM    Number of recent receipts (default: 20)
```

### `order`

Show contents of the active order. All order subcommands default to the active order.

```
$ appie order
Order 1234567  SUBMITTED
Delivery: dinsdag 25 feb  18:00-20:00

 1  AH Verse halfvolle melk 2L       1    1.89
 2  AH Pindakaas naturel             2    3.58
 3  Brinta Origineel                  1    2.49
                                    ──────
                            3 items   7.96
```

#### `order list`

List all open/scheduled orders (fulfillments). Active order marked with `*`.

```
$ appie order list
   Order      Status     Delivery                           Total
*  1234567  SUBMITTED  dinsdag 25 feb  18:00-20:00         87.30
   1234590  SUBMITTED  vrijdag 28 feb  08:00-10:00         42.15
```

#### `order add <product> [-n quantity]`

Add a product to the active order. `<product>` can be a numeric product ID or a search term. Default quantity is 1.

```
  -n, --quantity NUM    Quantity to add (default: 1)
```

```
$ appie order add 371880
Added 1x 371880

$ appie order add 371880 -n 3
Added 3x 371880

$ appie order add "halfvolle melk"
Found: AH Halfvolle melk
Added 1x 12345

$ appie order add "hagelslag puur"
  32786   AH Hagelslag puur                            250 g    €2.59
  55732   AH Chocoladepasta puur                       400 g    €2.69
  465752  AH Hagelslag xl puur                         600 g    €2.99
  1608    AH Hagelslag puur                            600 g    €2.99
  ...
multiple matches for "hagelslag puur", specify product ID
```

#### `order rm <product-id>`

Remove a product from the active order.

```
$ appie order rm 371880
Removed Optimel Drinkyoghurt aardbei
```

#### `order use <order-id>`

Switch the active order to a different fulfillment. Reopens if submitted.

```
$ appie order use 1234590
Reopened order 1234590 (was SUBMITTED)
Active order: 1234590
```

### `list`

Show shopping lists overview. All list subcommands accept a list name by prefix match.

```
$ appie list
  Boodschappen  12 items
  Weekmenu       5 items
```

#### `list show <name>`

Show items in a list by name (case-insensitive prefix match). Items are numbered for use with `list rm`.

```
$ appie list show Boodschappen
Boodschappen (12 items)

 1  AH Halfvolle melk       1 L    1  €1.89
 2  AH Pindakaas naturel    600 g  2  €3.59
 3  kaas                           1
```

#### `list add <product> [-n quantity] [-l name]`

Add a product to a list. `<product>` can be a numeric product ID or a search term. Uses `-l` to select the list by name (optional if only one list exists).

```
  -n, --quantity NUM    Quantity to add (default: 1)
  -l, --list NAME       List name (prefix match)
```

```
$ appie list add 371880 -l Boodschappen
Added 1x 371880 to Boodschappen

$ appie list add "pindakaas"
Found: AH Pindakaas naturel
Added 1x 371880 to Boodschappen
```

#### `list rm <number> [-l name]`

Remove an item by its line number from `list show` output.

```
  -l, --list NAME    List name (prefix match)
```

```
$ appie list rm 3 -l Boodschappen
Removed #3 (kaas) from Boodschappen
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

$ appie bonus -s
Bonus week 21 feb - 27 feb (spotlight)

Product                          Was    Now    Discount
AH Verse pizza               →  4.99   3.49   30% korting
...
(12 products)

$ appie bonus -d 2026-02-28
Bonus week 28 feb - 6 mrt
...
```

### `search <query>`

Search products by name. Results sorted by price (low to high). Use product IDs with `order add`.

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
