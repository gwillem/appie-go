# Albert Heijn Mobile API Documentation

Base URL: `https://api.ah.nl`

## Implementation Status

| Category | Endpoint | Status |
|----------|----------|--------|
| **Auth** | Anonymous token | ✅ |
| | Exchange code | ✅ |
| | Refresh token | ✅ |
| | Logout | |
| | Federate code | |
| **Products** | Search | ✅ |
| | Get by ID | ✅ |
| | Get by IDs | ✅ |
| | Categories | |
| **Bonus** | By category | ✅ |
| | Spotlight | ✅ |
| | Previously bought | |
| | Metadata | |
| **Orders** | Get active | ✅ |
| | Add/update items | ✅ |
| | Update state | ✅ |
| | Details by taxonomy | ✅ |
| | Download invoice (PDF) | ✅ |
| | Checkout | |
| **Shopping Lists** | Get lists (v3) | ✅ |
| | Get items (v2) | ✅ |
| | Add items (v2 PATCH) | ✅ |
| | Add to favorite list (GraphQL) | ✅ |
| **Member** | FetchMember (GraphQL) | ✅ |
| **Stores** | FetchStore (GraphQL) | |
| | GetFavoriteStore (GraphQL) | |
| **Recipes** | FetchRecipes (GraphQL) | |
| **Pricing** | FetchTotalPrice (GraphQL) | |
| **Recommendations** | Crosssells | |
| | Don't forget | |
| **Receipts** | Get all receipts | ✅ |
| | Get receipt by ID | ✅ |
| **Config** | Feature flags | ✅ |
| | Version check | |

---

## Required Headers

All requests require these headers:

```
User-Agent: Appie/9.28 (iPhone17,3; iPhone; CPU OS 26_1 like Mac OS X)
x-clientname: ipad
x-clientversion: 9.28
x-application: AHWEBSHOP
x-accept-language: nl-NL
x-fraud-detection-installation-id: <uuid>
x-correlation-id: <uuid>
Content-Type: application/json
Accept: application/json
```

For authenticated requests, add:
```
Authorization: Bearer <access_token>
```

---

## Authentication

### Get Anonymous Token

```
POST /mobile-auth/v1/auth/token/anonymous
```

**Request:**
```json
{"clientId": "appie-ios"}
```

**Response:**
```json
{
  "access_token": "27993385_xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  "refresh_token": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  "expires_in": 604798
}
```

### Exchange Auth Code for Token

```
POST /mobile-auth/v1/auth/token
```

**Request:**
```json
{
  "clientId": "appie-ios",
  "code": "<auth_code>"
}
```

### Refresh Token

```
POST /mobile-auth/v1/auth/token/refresh
```

**Request:**
```json
{
  "clientId": "appie-ios",
  "refreshToken": "<refresh_token>"
}
```

### Federate Code (for webviews)

```
POST /mobile-auth/v1/auth/federate/code
```

### Logout

```
POST /mobile-auth/v1/auth/token/logout
```

**Request:**
```json
{
  "clientId": "appie-ios",
  "refreshToken": "<refresh_token>"
}
```

---

## Products

### Search Products

```
GET /mobile-services/product/search/v2?page=0&size=30&sortOn=RELEVANCE&taxonomyId=<id>
GET /mobile-services/product/search/v2?query=<search_term>&page=0&size=30&sortOn=RELEVANCE
```

**Query Parameters:**
- `query` - Search term
- `page` - Page number (0-based)
- `size` - Results per page
- `sortOn` - `RELEVANCE`, `PRICE_ASC`, `PRICE_DESC`
- `taxonomyId` - Category ID
- `adType` - `TAXONOMY` for category browsing

**Response Product:**
```json
{
  "webshopId": 436752,
  "hqId": 60727,
  "title": "AH Biologisch Rundergehakt",
  "salesUnitSize": "300 g",
  "unitPriceDescription": "normale prijs per kg €20.97",
  "images": [
    {"width": 800, "height": 800, "url": "https://static.ah.nl/..."}
  ],
  "currentPrice": 5.66,
  "priceBeforeBonus": 6.29,
  "isBonus": true,
  "bonusMechanism": "10% KORTING",
  "mainCategory": "Vlees",
  "subCategory": "Rundergehakt",
  "brand": "AH Biologisch",
  "nutriscore": "D",
  "availableOnline": true,
  "isPreviouslyBought": true,
  "orderAvailabilityStatus": "IN_ASSORTMENT",
  "isOrderable": true,
  "propertyIcons": ["biologisch"],
  "discountLabels": [
    {"code": "DISCOUNT_PERCENTAGE", "defaultDescription": "10% korting", "percentage": 10}
  ]
}
```

### Get Products by IDs

```
GET /mobile-services/product/search/v2/products?ids=603740&ids=603734&sortOn=INPUT_PRODUCT_IDS
```

### Get Product Detail

```
GET /mobile-services/product/detail/v4/fir/<webshopId>
```

**Response:**
```json
{
  "productId": 415761,
  "productCard": {
    "webshopId": 415761,
    "hqId": 123456,
    "title": "Product Title",
    "brand": "Brand",
    "salesUnitSize": "500 g",
    "mainCategory": "Category",
    "subCategory": "Subcategory",
    "images": [...],
    "isBonus": false,
    "isFavorite": false,
    "isPreviouslyBought": true,
    "isOrderable": true,
    "availableOnline": true,
    "nutriscore": "A",
    "descriptionHighlights": "<html>...",
    "propertyIcons": ["vegetarisch"],
    "discountLabels": []
  },
  "properties": [...],
  "tradeItem": {...},
  "disclaimerText": "..."
}
```

### Get Category Sub-categories

```
GET /mobile-services/v1/product-shelves/categories/<categoryId>/sub-categories
```

---

## Orders

### Get Active Order Summary

```
GET /mobile-services/order/v1/summaries/active?sortBy=DEFAULT
```

**Response:**
```json
{
  "id": 316501042,
  "state": "REOPENED",
  "shoppingType": "DELIVERY",
  "totalPrice": {
    "priceBeforeDiscount": 72.69,
    "priceAfterDiscount": 60.52,
    "priceDiscount": 12.17,
    "priceTotalPayable": 60.52
  },
  "deliveryInformation": {
    "deliveryDate": "2026-01-20",
    "deliveryStartTime": "18:00",
    "deliveryEndTime": "20:00",
    "address": {
      "street": "...",
      "houseNumber": 39,
      "zipCode": 3522,
      "city": "UTRECHT"
    }
  },
  "orderedProducts": [
    {
      "amount": 4,
      "quantity": 4,
      "product": {
        "webshopId": 199922,
        "title": "...",
        "brand": "...",
        "images": [...]
      }
    }
  ]
}
```

### Add/Update Items in Order

```
PUT /mobile-services/order/v1/items?sortBy=DEFAULT
```

**Request:**
```json
{
  "items": [
    {
      "productId": 553353,
      "quantity": 1,
      "originCode": "PRD",
      "description": "",
      "strikethrough": false
    }
  ]
}
```

### Update Order State

```
PUT /mobile-services/order/v1/<orderId>/state?orderBy=DEFAULT
```

**Request:** Plain text body (e.g., `RESET`)

**Response (409 Conflict):**
```json
{
  "status": 409,
  "message": "{\"success\":false,\"errorCode\":\"409\",\"errorMessage\":\"Unable to update, closing time has passed for order\"}",
  "correlationId": "...",
  "timestamp": "2026-01-23T10:58:33.222Z"
}
```

### Get Order Details (grouped by taxonomy)

```
GET /mobile-services/order/v1/<orderId>/details-grouped-by-taxonomy
```

**Response:**
```json
{
  "orderId": 229775812,
  "deliveryDate": "2025-12-09",
  "orderState": "DELIVERED",
  "closingTime": "2025-12-08T22:59:00Z",
  "deliveryType": "HOME",
  "deliveryTimePeriod": {
    "startDateTime": "2025-12-09T18:00:00",
    "endDateTime": "2025-12-09T20:00:00",
    "startTimeUtc": "2025-12-09T17:00:00Z",
    "endTimeUtc": "2025-12-09T19:00:00Z"
  },
  "groupedProductsInTaxonomy": [
    {
      "taxonomyName": "Groente, aardappelen",
      "orderedProducts": [
        {
          "amount": 1,
          "quantity": 1,
          "allocatedQuantity": 1,
          "product": {
            "webshopId": 164358,
            "title": "AH Oranje zoete aardappel",
            "brand": "AH",
            "salesUnitSize": "1 kg",
            "priceBeforeBonus": 3.79,
            "isBonus": false,
            "isPreviouslyBought": true
          },
          "position": 6
        }
      ]
    }
  ],
  "invoiceId": "2294567-00199",
  "cancellable": false,
  "orderPayments": [],
  "address": {
    "street": "...",
    "houseNumber": 39,
    "zipCode": 3522,
    "zipCodeExtra": "JS",
    "city": "UTRECHT",
    "countryCode": "NLD",
    "type": "H"
  },
  "orderMethod": "PLANSERVICE",
  "reopenable": false
}
```

### Download Invoice (PDF)

```
GET /mobile-services/order/v1/invoice/download?invoiceId=<invoiceId>
```

**Query Parameters:**
- `invoiceId` - Invoice ID from order details (e.g., "2294567-00199")

**Response:** Binary PDF file

### Get Checkout Info

```
GET /mobile-services/order/v1/<orderId>/checkout
```

**Response:**
```json
{
  "kassaKoopjes": [...],
  "missingBonus": [...],
  "nonChosen": [...],
  "nonDeliverables": [...],
  "recommendedProducts": [...],
  "samples": [...],
  "showMakeCompleet": true
}
```

---

## Shopping Lists

### Get All Lists (v3)

```
GET /mobile-services/lists/v3/lists?productId=<id>
```

Note: The `productId` parameter is required but returns all lists regardless of value.

**Response:**
```json
[
  {
    "id": "305e6a50-a970-457b-8831-409f572832d4",
    "description": "My List",
    "itemCount": 4,
    "hasFavoriteProduct": false,
    "productImages": [
      [{"width": 80, "height": 80, "url": "https://static.ah.nl/..."}]
    ]
  }
]
```

### Add Items to Shopping List (v2)

```
PATCH /mobile-services/shoppinglist/v2/items
```

**Request:**
```json
{
  "items": [
    {
      "description": "AH Woksaus teriyaki",
      "strikeThrough": false,
      "originCode": "PRD",
      "productId": 482500,
      "type": "SHOPPABLE",
      "searchTerm": "AH Woksaus teriyaki",
      "quantity": 1
    }
  ]
}
```

**Fields:**
- `type`: `SHOPPABLE` for products
- `originCode`: `PRD` for product-linked items, `TXT` for free-text items
- `description` / `searchTerm`: product title

**Response:** Full shopping list state (same format as GET below).

### Get List Items (v2)

```
GET /mobile-services/shoppinglist/v2/items
```

**Response:**
```json
{
  "id": "943e873f-9a64-403b-becb-c92f0978eb5d",
  "items": [
    {
      "listItemId": 0,
      "strikedthrough": false,
      "quantity": 1,
      "description": "350 g koelverse vegan shoarma",
      "type": "SHOPPABLE",
      "originCode": "TXT",
      "position": 32,
      "sorting": {
        "position": 1
      },
      "vagueTermDetails": {
        "searchTermValue": "350 g koelverse vegan shoarma",
        "bonus": false
      }
    }
  ],
  "dateLastSynced": "2026-01-23T11:59:19",
  "dateLastSyncedMillis": 1769165959010,
  "activeSorting": "E",
  "storeNumber": 0
}
```

---

## Receipts (Kassabonnen)

In-store purchase receipts from physical AH stores.

### Get All Receipts

```
GET /mobile-services/v1/receipts
```

**Response:**
```json
{
  "receipts": [
    {
      "transactionId": "1234567890",
      "datetime": "2026-01-20T14:30:00",
      "storeId": 1234,
      "storeName": "AH Utrecht Centrum",
      "total": 45.67
    }
  ]
}
```

### Get Receipt Details

```
GET /mobile-services/v2/receipts/<transactionId>
```

**Response:**
```json
{
  "transactionId": "1234567890",
  "datetime": "2026-01-20T14:30:00",
  "storeId": 1234,
  "storeName": "AH Utrecht Centrum",
  "total": 45.67,
  "receiptItems": [
    {
      "description": "AH Halfvolle melk",
      "quantity": 2,
      "amount": 2.58,
      "unitPrice": 1.29,
      "productId": 12345
    }
  ]
}
```

---

## Bonus / Promotions

### Get Bonus Metadata

```
GET /mobile-services/bonuspage/v3/metadata
```

### Get Bonus Section

```
GET /mobile-services/bonuspage/v2/section?application=AHWEBSHOP&category=<category>&date=<YYYY-MM-DD>&promotionType=NATIONAL
```

### Get Previously Bought Bonus Products

```
GET /mobile-services/bonuspage/v2/section/previously-bought?application=AHWEBSHOP&date=<YYYY-MM-DD>
```

### Get Spotlight Bonus

```
GET /mobile-services/bonuspage/v2/section/spotlight?application=AHWEBSHOP&date=<YYYY-MM-DD>
```

---

## Recommendations

### Get Cross-sells

```
POST /mobile-services/v2/recommendations/crosssells
```

**Request:**
```json
{
  "limit": 6,
  "experimentId": "exp-var-with-category",
  "productId": 165625,
  "propensityFilter": true,
  "basketItems": [
    {
      "productId": 165625,
      "position": 2,
      "quantity": 1
    }
  ]
}
```

**Response:** Same format as Don't Forget Lane (see below).

### Get "Don't Forget" Lane

```
POST /mobile-services/v2/recommendations/dontforgetlane
```

**Request:**
```json
{
  "positiveCmOnly": true,
  "offset": 0,
  "usecaseId": "app_mylist",
  "basketItems": [
    {
      "quantity": 1,
      "position": 1,
      "productId": 482500
    }
  ],
  "limit": 7
}
```

**Response:**
```json
{
  "dataLakeModel": {
    "name": "TransformerModelWithCategory",
    "requestId": "",
    "version": "2026_02_01-11_50"
  },
  "productCards": [
    {
      "webshopId": 2600,
      "title": "AH Goudse jong 48+ plakken",
      "brand": "AH",
      "salesUnitSize": "190 g",
      "images": [...],
      "isBonus": true,
      "bonusMechanism": "2e halve prijs",
      "priceBeforeBonus": 2.95,
      "currentPrice": 2.21
    }
  ]
}
```

---

## Configuration

### Get Feature Flags

```
GET /mobile-services/config/v1/features/ios?version=9.28
```

### Version Check

```
GET /mobile-services/versioncheck/v3/ipad/9.28/check
```

### Get Webflow Config

```
GET /mobile-services/v2/webflow
```

---

## GraphQL

```
POST /graphql
```

All GraphQL requests use the same endpoint with different queries.

**Required Headers for GraphQL:**
```
x-apollo-operation-name: <OperationName>
x-apollo-operation-type: query
apollographql-client-name: nl.ah.Appie-apollo-ios
apollographql-client-version: 9.28-260102201630
```

### Known Operations

| Operation | Description | Status |
|-----------|-------------|--------|
| FetchMember | Get member profile, address, cards | ✅ Implemented |
| FetchProduct (nutrition) | Get nutritional info via tradeItem | ✅ Implemented |
| AddProductsToFavoriteList | Add products to named list | ✅ Implemented |
| FetchOrderTrackTrace | Track and trace info for order | ✅ Implemented |
| FetchStore | Store details by ID | |
| GetFavoriteStore | User's favorite store | |
| FetchRecipes | Recipe search | |
| FetchTotalPrice | Calculate order total price | |
| SearchProducts | Product search with facets/variants | |
| FetchEntryPoints | Home screen entry points | |
| FetchCuratedLists | Curated shopping lists | |
| FetchNBACard | Next best action card | |
| FetchPageEntries | Page entries | |
| FetchPageTemplate | Page templates | |
| FetchPersonalizedAdvertisementV2 | Targeted promotions | |
| FetchPurchaseStampServerTime | Purchase stamp time | |
| FetchSmartLane | Smart suggestions lane | |
| MessageCenterGetUnreadMessagesInfo | Unread messages count | |

#### FetchMember

Fetches member profile with address, cards, and customer segments.

```graphql
query FetchMember {
  member {
    __typename
    ...MemberFragment
  }
}
fragment MemberAddressFragment on MemberAddress {
  __typename
  street
  houseNumber
  houseNumberExtra
  postalCode
  city
  countryCode
}
fragment MemberFragment on Member {
  __typename
  address { __typename ...MemberAddressFragment }
  analytics { __typename digimon idmon idsas batch firebase sitespect }
  cards { __typename airmiles bonus gall }
  company { __typename id name addressInvoice { __typename ...MemberAddressFragment } customOffersAllowed }
  contactSubscriptions
  dateOfBirth
  emailAddress
  gender
  id
  isB2B
  memberships
  name { __typename first last }
  phoneNumber
  customerProfileAudiences
  customerProfileProperties { __typename key value }
}
```

**Response:**
```json
{
  "data": {
    "member": {
      "id": 27993385,
      "emailAddress": "user@example.com",
      "gender": "MALE",
      "dateOfBirth": "1985-03-15",
      "phoneNumber": "+31612345678",
      "isB2B": false,
      "name": {
        "first": "Jan",
        "last": "de Vries"
      },
      "address": {
        "street": "Hoofdstraat",
        "houseNumber": 42,
        "houseNumberExtra": "A",
        "postalCode": "1234AB",
        "city": "AMSTERDAM",
        "countryCode": "NL"
      },
      "cards": {
        "bonus": "2620123456789",
        "gall": null,
        "airmiles": null
      },
      "customerProfileAudiences": ["segment1", "segment2"]
    }
  }
}
```

#### FetchOrderTrackTrace

Fetches track and trace information for an order.

```graphql
query FetchOrderTrackTrace($orderId: Int!) {
  order(id: $orderId) {
    __typename
    delivery {
      __typename
      trackAndTraceV2 {
        __typename
        orderId
        type
        orderType
        message
        etaBlock {
          __typename
          range {
            __typename
            start
            end
          }
        }
      }
    }
  }
}
```

**Variables:**
```json
{
  "orderId": 229775812
}
```

**Response:**
```json
{
  "data": {
    "order": {
      "__typename": "Order",
      "delivery": {
        "__typename": "Delivery",
        "trackAndTraceV2": {
          "__typename": "TrackAndTraceV2",
          "orderId": 229775812,
          "type": "DELIVERED",
          "orderType": "HOME_DELIVERY",
          "message": "Je boodschappen zijn bezorgd.",
          "etaBlock": null
        }
      }
    }
  }
}
```

**Track Types:** `DELIVERED`, `IN_TRANSIT`, `PREPARING`, etc.

#### AddProductsToFavoriteList

Adds products to a named favorite list (v3).

```graphql
mutation AddProductsToFavoriteList($favoriteListId: String!, $products: [FavoriteListProductMutation!]!) {
  favoriteListProductsAddV2(id: $favoriteListId, products: $products) {
    __typename
    status
    errorMessage
  }
}
```

**Variables:**
```json
{
  "favoriteListId": "181CBC7B-0088-4EAF-9E7F-8EEF8F8BBDBA",
  "products": [
    { "productId": 482500 }
  ]
}
```

Note: `favoriteListId` must be uppercase UUID.

**Response:**
```json
{
  "data": {
    "favoriteListProductsAddV2": {
      "status": "SUCCESS",
      "errorMessage": null
    }
  }
}
```

#### FetchRecipes

Search for recipes with filters and pagination.

```graphql
query FetchRecipes($searchText: String, $start: Int, $size: PageSize, $sortBy: RecipeSearchSortOption, $filters: [RecipeSearchQueryFilter!], $priorityRecipeIds: [Int!], $boostFavoriteRecipeIds: [Int!], $recipeIds: [Int!]!, $ingredients: [String!]) {
  recipeSearchV2(searchText: $searchText, start: $start, size: $size, sortBy: $sortBy, filters: $filters, priorityRecipeIds: $priorityRecipeIds, boostFavoriteRecipeIds: $boostFavoriteRecipeIds, recipeIds: $recipeIds, ingredients: $ingredients) {
    __typename
    correctedSearchTerm
    page { __typename total hasNextPage }
    filters {
      __typename label name
      filters { __typename name label group count selected }
    }
    result { __typename ...RecipeSummaryFragment }
  }
}
fragment RecipeImageFragment on RecipeImage { __typename width height url }
fragment RecipeSummaryFragment on RecipeSummary {
  __typename id title
  time { __typename cook }
  images(renditions: [S, M, D445X297, D890X594]) { __typename ...RecipeImageFragment }
  author { __typename origin { __typename type hostName url } }
  flags nutriScore
  tags { __typename key value }
}
```

**Variables:**
```json
{
  "searchText": "pasta",
  "start": 0,
  "size": 7,
  "sortBy": "MOST_RELEVANT",
  "filters": [],
  "priorityRecipeIds": [],
  "boostFavoriteRecipeIds": [],
  "recipeIds": [],
  "ingredients": []
}
```

**Response:**
```json
{
  "data": {
    "recipeSearchV2": {
      "correctedSearchTerm": null,
      "page": { "total": 245, "hasNextPage": true },
      "result": [
        {
          "id": 12345,
          "title": "Pasta carbonara",
          "time": { "cook": 25 },
          "images": [{ "width": 890, "height": 594, "url": "https://..." }],
          "nutriScore": "B",
          "tags": [{ "key": "cuisine", "value": "Italiaans" }]
        }
      ]
    }
  }
}
```

#### FetchTotalPrice

Calculate total price for a set of products (includes bonus discounts).

```graphql
query FetchTotalPrice($products: [PriceLineItem!]!) {
  totalPrice(products: $products) {
    __typename
    withoutDiscount { __typename amount }
    discount { __typename amount }
  }
}
```

**Variables:**
```json
{
  "products": [
    { "id": 165625, "quantity": 1 },
    { "id": 482500, "quantity": 1 }
  ]
}
```

**Response:**
```json
{
  "data": {
    "totalPrice": {
      "withoutDiscount": { "amount": 5.58 },
      "discount": { "amount": 0.38 }
    }
  }
}
```

The final price is `withoutDiscount.amount - discount.amount`.

#### FetchProduct (Nutritional Info)

Fetch nutritional info for a product via its tradeItem.

```graphql
query FetchProduct($productId: Int!) {
  product(id: $productId) {
    id
    tradeItem {
      nutritions {
        nutrients { type name value }
      }
    }
  }
}
```

**Variables:**
```json
{ "productId": 436752 }
```

**Response:**
```json
{
  "data": {
    "product": {
      "id": 436752,
      "tradeItem": {
        "nutritions": [{
          "nutrients": [
            { "type": "ENER-", "name": "Energie", "value": "1011 kJ (243 kcal)" },
            { "type": "FAT", "name": "Vet", "value": "18 g" },
            { "type": "PRO-", "name": "Eiwitten", "value": "20 g" },
            { "type": "SALTEQ", "name": "Zout", "value": "0.33 g" }
          ]
        }]
      }
    }
  }
}
```

#### FetchEntryPoints

Fetches UI entry points for the home screen.

```graphql
query FetchEntryPoints($name: String!, $version: String) {
  entryPointComponent(name: $name, version: $version) {
    name
    content { ... }
    entryPoints {
      name
      contentVariant { ... }
      metadata { group, dismissible }
    }
  }
}
```

**Variables:**
```json
{
  "name": "my-ah-lane-nl",
  "version": "9.28"
}
```

#### GetFavoriteStore

```graphql
query GetFavoriteStore {
  storesFavouriteStore {
    __typename
    ...StoresFragment
  }
}
fragment GeolocationFragment on GeoLocation { __typename latitude longitude }
fragment StoresFragment on Stores {
  __typename id
  address { __typename street houseNumber houseNumberExtra postalCode city countryCode }
  openingDays { __typename date openingHour { __typename openFrom openUntil } }
  geoLocation { __typename ...GeolocationFragment }
  phone storeType
  services { __typename code }
}
```

**Response:**
```json
{
  "data": {
    "storesFavouriteStore": {
      "id": 1527,
      "address": {
        "street": "Nachtegaalstraat",
        "houseNumber": 48,
        "postalCode": "3581AD",
        "city": "Utrecht"
      },
      "openingDays": [
        { "date": "2026-02-03", "openingHour": { "openFrom": "08:00", "openUntil": "22:00" } }
      ],
      "geoLocation": { "latitude": 52.09, "longitude": 5.12 },
      "phone": "030-2345678",
      "storeType": "AH",
      "services": [{ "code": "PICKUP" }, { "code": "SELFSERVICE" }]
    }
  }
}
```

#### FetchStore

```graphql
query FetchStore($storeId: Int!) {
  storesInformation(id: $storeId) {
    __typename
    ...StoresFragment
  }
}
```

Uses same `StoresFragment` as GetFavoriteStore above.

**Variables:**
```json
{ "storeId": 1527 }
```

---

## Analytics

### Send Bulk Analytics

```
POST /mobile-services/v3/analytics/bulk
```

**Status:** 202 Accepted

---

## Error Responses

```json
{
  "code": "ERROR_CODE",
  "message": "Human readable message"
}
```

Common error codes:
- `SESSION_EXPIRED` - Token expired, need to refresh
- `INVALID_CAPTCHA` - Captcha required for login

---

## Notes

- All prices are in EUR
- Product IDs: `webshopId` is the primary product identifier
- Order IDs are numeric (e.g., 316501042)
- Shopping list IDs are UUIDs
- Images are available in multiple resolutions (48, 80, 200, 400, 800px)
