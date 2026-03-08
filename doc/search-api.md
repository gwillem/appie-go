# Product Search API

The AH API offers two search mechanisms: a REST endpoint and a GraphQL query. The REST endpoint is simpler but limited; the GraphQL endpoint supports full faceted filtering.

## REST Search (`/mobile-services/product/search/v2`)

### Supported Parameters

| Parameter    | Description                          | Example             |
|-------------|--------------------------------------|---------------------|
| `query`     | Search term (required)               | `kaas`              |
| `page`      | Page number (0-based)                | `0`                 |
| `size`      | Results per page                     | `30`                |
| `sortOn`    | Sort order (see below)               | `RELEVANCE`         |
| `taxonomyId`| Filter by category ID                | `8444`              |

### Sort Options

`RELEVANCE`, `PRICEHIGHLOW`, `PRICELOWHIGH`, `TAXONOMY`, `PURCHASE_FREQUENCY`, `PURCHASE_DATE`, `PURCHASE_DEPARTMENT`, `NUTRISCORE`

### Limitations

- **No facet filtering:** The REST response includes rich filter metadata (`filters`, `taxonomyNodes`) but **ignores all filter query parameters** except `taxonomyId`. Parameters like `bonus=true`, `brand=AH`, `nutriscore=a` are silently ignored — the total result count does not change.
- **No hard `size` limit:** The `size` parameter has no enforced maximum — tested up to 1000 and it returns all matching products (capped at `totalElements`). Note: the actual count returned may slightly exceed the requested `size` due to injected sponsored/ad products (e.g., requesting 50 may return 53).

### Response Structure

Top-level keys: `products`, `page`, `filters`, `taxonomyNodes`, `sortOn`, `configuration`, `links`, `ads`

---

## GraphQL Search (`SearchProducts`)

The website uses this for filtered search. Supports full faceted filtering with combinable constraints.

### Query

```graphql
query SearchProducts($input: SearchProductsInput!) {
  searchProducts(input: $input) {
    totalFound
    products {
      id title brand category salesUnitSize icons
      priceV2 {
        now { amount }
        was { amount }
        discount { description promotionType segmentType }
        promotionLabels { topText centerText bottomText }
      }
      imagePack { small { url width height } }
      availability { isOrderable }
    }
    facets {
      facets {
        label name type
        options { name value matches selected }
      }
      quickFilters {
        label name type
        options { name value matches selected }
      }
    }
  }
}
```

### Input Structure

```json
{
  "input": {
    "query": "kaas",
    "sortType": "RELEVANCE",
    "searchInput": {
      "facets": [
        {"label": "bonus", "values": {"values": [true]}},
        {"label": "product.brand.name", "values": {"values": ["old-amsterdam"]}}
      ],
      "page": {"number": 0, "size": 36}
    }
  }
}
```

### Sort Types

`RELEVANCE`, `PRICE_LOW_HIGH`, `PRICE_HIGH_LOW`

Note: `NUTRI_SCORE` is NOT accepted as a sortType (despite being a valid REST sort).

### Available Facets

Facet values use **slug-style IDs** (e.g., `old-amsterdam` not `Old Amsterdam`). The `bonus` facet is the only one that takes a **boolean** `true`; all others take strings.

| Facet Label                           | Name              | Type     | Example Values                        |
|---------------------------------------|-------------------|----------|---------------------------------------|
| `bonus`                               | Bonus             | -        | `true` (boolean)                      |
| `product.brand.name`                  | Merk              | SINGLE   | `ah`, `old-amsterdam`, `beemster`     |
| `product.properties.nutri-score`      | Nutri-Score       | MULTIPLE | `a`, `b`, `c`, `d`, `e`              |
| `product.properties.dieet`            | Dieet             | MULTIPLE | `vegetarisch`, `veganistisch`, `biologisch`, `laag-suiker`, `laag-vet`, `laag-zout` |
| `product.properties.allergie`         | Allergie          | MULTIPLE | `zonder-lactose`, `zonder-gluten`, `zonder-noten`, `zonder-soja`, etc. |
| `product.properties.rijping`          | Rijping           | SINGLE   | `jong`, `jong-belegen`, `belegen`, `extra-belegen`, `oud` |
| `product.properties.vetgehalte`       | Vetgehalte        | MULTIPLE | `48`, `30`, `50`, `45`                |
| `product.properties.smaak`            | Smaak             | SINGLE   | `kaas`, `komijn`, `truffel`           |
| `product.properties.diepvries`        | Diepvries         | MULTIPLE | `y`                                   |
| `product.properties.keuken`           | Keuken            | MULTIPLE | `gouds`, `hollands`, `italiaans`      |
| `product.properties.keurmerk`         | Keurmerk          | SINGLE   | `weidemelk`, `beterleven1ster`        |
| `product.properties.bereidingswijze`  | Bereidingswijze   | MULTIPLE | `oven`, `airfryer`, `pan`, `magnetron`|
| `product.properties.uitgelicht`       | Prijsfavoriet     | -        | `Prijsfavoriet`, `Uit Nederland`      |
| `taxonomy.nodes`                      | Soort             | SINGLE   | `8444` (Plakken kaas), `8572` (Stukken kaas) |
| `price`                               | Prijs             | RANGE    | **Not working via mobile API** (400 error) |

### Combining Facets

Multiple facets can be combined (AND logic between different facets). Multiple values within one MULTIPLE-type facet use OR logic.

```json
{
  "facets": [
    {"label": "bonus", "values": {"values": [true]}},
    {"label": "product.brand.name", "values": {"values": ["ah"]}},
    {"label": "product.properties.nutri-score", "values": {"values": ["a", "b"]}}
  ]
}
```

### Discovering Facets

Every search response includes a `facets` field listing all available facets with their options and match counts for the current query. Use this to dynamically discover valid filter values.

### Bonus Mechanism Text

**The GraphQL search endpoint does NOT return bonus mechanism text** (e.g., "2 VOOR 5.50", "25% korting") for multi-buy deals. The `priceV2.discount` field is null and `priceV2.promotionLabel` is null for these products. Only simple percentage discounts sometimes populate `discount.description`.

This is confirmed by mitmproxy capture of the iOS Appie app (v9.31, March 2026): the app's own GraphQL search query requests `promotionLabel { tiers { mechanism description ... } }` but receives null for all products. The app uses the same `priceV2` without any special parameters (`forcePromotionVisibility`, `periodStart`/`periodEnd`).

**Tested and failed approaches:**
- `priceV2(forcePromotionVisibility: true)` — still null
- `priceV2(periodStart: "...", periodEnd: "...", forcePromotionVisibility: true)` — still null
- `promotionLabels { topText centerText bottomText }` (plural form) — empty array
- `extraConsents: ["NECESSARY"]` / `["ANALYSIS", "ADS_INSIDE", "ADS_OUTSIDE"]` — no effect
- All combinations of the above — no effect

**Working approach:** Enrich search results via the REST batch endpoint `GET /mobile-services/product/search/v2/products?ids=...` which returns `bonusMechanism` on every product. The iOS app does the same (fetches individual products via REST when tapped).

### iOS App Search Behavior (observed via mitmproxy)

The iOS app (v9.31) uses exactly the same `SearchProducts` GraphQL query for search, with these differences from our implementation:
- Passes `extraConsents: ["ANALYSIS", "ADS_INSIDE", "ADS_OUTSIDE"]`
- Passes `searchInput.intent: { intent: "ONLINE", orderId: <active_order_id> }`
- Passes `searchInput.page.merged: false`
- Passes `$includesVariants: true` and `$includesNutritionalInfo: true` conditional fragments
- Requests `imagePack(angles: [HERO, ANGLE_2D1])` for multiple product angles
- None of these affect bonus/promotion data availability

### Notes

- GraphQL introspection is disabled on the AH API
- The `price` facet returns 400 errors via the mobile API — likely requires the `x-client-platform-type: web` header
- The `previouslyBought` facet also returns 400 via mobile API
- Taxonomy node IDs are numeric but passed as strings in facet values
- The website sends `extraConsents: ["NECESSARY"]` and `page.merged: true` but these are not required
