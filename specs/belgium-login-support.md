# Spec: Belgium (ah.be) support

## Context

appie-go targets the Albert Heijn Netherlands market. Belgium (ah.be) runs on the
same hosts/backend, but two request parameters select the market — and appie-go
hardcoded the Dutch values, so Belgian members got NL login rejections and the NL
assortment/pricing.

Root causes, both verified live (login page + a capture of the real ah.be iOS app):

1. **Login** — the market is selected by the OAuth **`client_id`**. appie-go sent
   `appie-ios` (Dutch app), so AH validated credentials against the **Netherlands**
   member directory and rejected Belgian accounts ("niet geldig voor Albert Heijn
   Nederland"). The Belgian app uses **`appie-be-ios`**.
2. **Assortment/pricing** — selected by the **`x-application`** header. appie-go
   sent `AHWEBSHOP` (NL), so search/bonus returned the Dutch catalog even for a
   Belgian member. The Belgian app sends **`AHBEWEBSHOP`**. Verified: `roggebrood`
   returns 9 NL hits under AHWEBSHOP vs 4 BE hits under AHBEWEBSHOP, matching ah.be
   exactly.

The hosts are irrelevant (`login.ah.nl?client_id=appie-be-ios` serves the Belgian
login; `api.ah.nl` + `AHBEWEBSHOP` returns the Belgian catalog), so only these two
parameters change; API/login hosts stay on the default backend.

## What

- Add `WithCountry(country string) Option` selecting the market by both
  parameters: `"nl"` (default) -> `appie-ios` / `AHWEBSHOP`; `"be"` ->
  `appie-be-ios` / `AHBEWEBSHOP`. Case-insensitive.
- Route the `x-application` header and the bonus `application` query param through
  the configured value instead of the hardcoded `AHWEBSHOP`.
- Persist the chosen country in the config file so it need not be repeated;
  restore it on load. An explicit `WithCountry` still wins over the stored value
  (explicit > stored > nl default).
- Add a `--country` flag to the CLI wiring the option into every subcommand.

## Out of scope

- Switching API/login hosts per country (unnecessary — the two params are the switch).
- Automated checkout / delivery slots / payment.

## Done when

- `New(WithCountry("be"))` sets `clientID == "appie-be-ios"` and
  `application == "AHBEWEBSHOP"`; `"nl"` -> `appie-ios` / `AHWEBSHOP`; hosts unchanged.
- Requests carry the country's `x-application` header; `loginURL()` carries its client_id.
- `SearchProducts` under `be` returns the Belgian assortment/pricing.
- Country round-trips through the config; an explicit country beats the stored one.
- `appie --country be login` reaches the Belgian login; existing suite stays green.
