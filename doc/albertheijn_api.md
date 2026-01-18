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
| | Checkout | |
| | Details by taxonomy | |
| **Shopping Lists** | Get lists (v3) | ✅ |
| | Get items (v2) | |
| **Member** | FetchMember (GraphQL) | ✅ |
| **Stores** | FetchStore | |
| | GetFavoriteStore | |
| **Recommendations** | Crosssells | |
| | Don't forget | |
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

### Get Order Details (grouped by taxonomy)

```
GET /mobile-services/order/v1/<orderId>/details-grouped-by-taxonomy
```

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

### Get List Items (v2)

```
GET /mobile-services/shoppinglist/v2/items
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

### Get "Don't Forget" Lane

```
POST /mobile-services/v2/recommendations/dontforgetlane
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
| FetchEntryPoints | Home screen entry points | |
| FetchCuratedLists | Curated shopping lists | |
| FetchNBACard | Next best action card | |
| FetchPageEntries | Page entries | |
| FetchPageTemplate | Page templates | |
| FetchPurchaseStampServerTime | Purchase stamp time | |
| FetchSmartLane | Smart suggestions lane | |
| FetchStore | Store details | |
| GetFavoriteStore | User's favorite store | |
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
    ...StoreFragment
  }
}
```

#### FetchStore

```graphql
query FetchStore($storeId: Int!) {
  store(storeId: $storeId) {
    __typename
    ...StoreFragment
  }
}
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
