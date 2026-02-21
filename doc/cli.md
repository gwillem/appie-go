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

#### `order add <product-id> [quantity]`

Add a product to the active order. Default quantity is 1.

```
$ appie order add 371880
Added 1x Optimel Drinkyoghurt aardbei (€1.59)

$ appie order add 371880 3
Added 3x Optimel Drinkyoghurt aardbei (€1.59)
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

List shopping lists, or show items for a specific list.

```
  -i, --id ID    Show items for a specific list
```

```
$ appie list
ID                                    Name              Items
a1b2c3d4-e5f6-7890-abcd-ef1234567890  Boodschappen         12
f9e8d7c6-b5a4-3210-fedc-ba0987654321  Weekmenu              5

$ appie list -i a1b2c3d4-e5f6-7890-abcd-ef1234567890
Boodschappen (12 items)

   Product                         Qty  Checked
   AH Verse halfvolle melk 2L       1
   AH Pindakaas naturel             2
   Brinta Origineel                  1   [x]
   kaas (vrije tekst)               1
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

Search products by name. Needed to find product IDs for `order add`.

```
  -n, --limit NUM    Max results (default: 10)
```

```
$ appie search "pindakaas"
ID       Product                          Price  Bonus
371880   AH Pindakaas naturel 600g        3.59
204835   Calvé Pindakaas 650g             4.49   35% korting
192450   AH Biologisch pindakaas 350g     2.89
...
```
