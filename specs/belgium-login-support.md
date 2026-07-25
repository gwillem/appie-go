# Spec: Belgium (ah.be) support

## Context

appie-go targets the Albert Heijn Netherlands market. The data API works for
Belgium unmodified — the same backend serves both markets. The gap was **login**:
a Belgian member could not authenticate.

Root cause (verified against the live login page): the market is selected by the
OAuth **`client_id`**, not by the host. appie-go sent `client_id=appie-ios` (the
Dutch app), so AH validated credentials against the **Netherlands** member
directory and rejected Belgian accounts with "niet geldig voor Albert Heijn
Nederland". The Belgian app uses **`client_id=appie-be-ios`**.

Confirmed the host is irrelevant: `login.ah.nl?client_id=appie-be-ios` serves the
Belgian login, and `api.ah.nl` issues tokens for `appie-be-ios`. So only the
client_id needs to change; the API and login hosts stay on the default backend.

## What

- Add `WithCountry(country string) Option` selecting the market by client_id:
  `"nl"` -> `appie-ios` (default), `"be"` -> `appie-be-ios`. Case-insensitive.
- Persist the chosen country in the config file so it need not be repeated;
  restore it on load. An explicit `WithCountry` still wins over the stored value
  (explicit > stored > nl default).
- Add a `--country` flag to the CLI wiring the option into every subcommand.

## Out of scope

- Switching API/login hosts per country (unnecessary — client_id is the switch).
- Automated checkout / delivery slots / payment.

## Done when

- `New(WithCountry("be"))` sets `clientID == "appie-be-ios"`; `"nl"` -> `appie-ios`;
  the API/login hosts are unchanged.
- `loginURL()` carries the country's client_id.
- Country round-trips through the config; an explicit country beats the stored one.
- `appie --country be login` reaches the Belgian login; existing suite stays green.
