package appie

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// personalBonusJSON is an anonymized sample of the
// /mobile-services/bonuspage/v1/personal response. It uses the same envelope as
// the national bonus section, so the existing productResponse/collectBonusProducts
// decode it; the personal-only fields (activationStatus, promotionType, …) are
// present to prove they are tolerated and ignored.
const personalBonusJSON = `{
  "sectionType": "PERSONAL",
  "sectionDescription": "Bonus Box",
  "bonusGroupOrProducts": [
    {"product": {
      "webshopId": 111111, "title": "Test Bio Yoghurt", "brand": "AH Biologisch",
      "salesUnitSize": "500 g", "unitPriceDescription": "per kg 3.98",
      "currentPrice": 1.49, "priceBeforeBonus": 1.99, "isBonus": true,
      "bonusMechanism": "25% KORTING", "mainCategory": "Zuivel", "subCategory": "Yoghurt",
      "availableOnline": true, "isOrderable": true, "isPreviouslyBought": true,
      "propertyIcons": ["biologisch"],
      "activationStatus": "ACTIVATED", "promotionType": "PERSONAL", "segmentType": "BBOX",
      "images": [{"url": "https://example.invalid/img.jpg", "width": 800, "height": 800}]
    }},
    {"product": {
      "webshopId": 222222, "title": "Test Koffiebonen", "brand": "Perla",
      "currentPrice": 4.99, "priceBeforeBonus": 6.49, "isBonus": true,
      "bonusMechanism": "1 + 1 GRATIS", "mainCategory": "Koffie",
      "availableOnline": true, "isOrderable": true,
      "activationStatus": "NOT_ACTIVATED", "promotionType": "PERSONAL", "segmentType": "BBOX"
    }}
  ]
}`

func TestGetPersonalBonus(t *testing.T) {
	var gotPath, gotDate string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotDate = r.URL.Query().Get("bonusStartDate")
		w.Write([]byte(personalBonusJSON))
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))
	client.accessToken = "test"

	products, err := client.GetPersonalBonus(context.Background(), "2026-07-13")
	if err != nil {
		t.Fatal(err)
	}

	if gotPath != "/mobile-services/bonuspage/v1/personal" {
		t.Errorf("unexpected path %q", gotPath)
	}
	if gotDate != "2026-07-13" {
		t.Errorf("expected bonusStartDate 2026-07-13, got %q", gotDate)
	}
	if len(products) != 2 {
		t.Fatalf("expected 2 products, got %d", len(products))
	}

	p := products[0]
	if p.ID != 111111 {
		t.Errorf("expected ID 111111, got %d", p.ID)
	}
	if p.Title != "Test Bio Yoghurt" {
		t.Errorf("expected title Test Bio Yoghurt, got %q", p.Title)
	}
	if p.Price.Now != 1.49 {
		t.Errorf("expected price now 1.49, got %v", p.Price.Now)
	}
	if p.Price.Was != 1.99 {
		t.Errorf("expected price was 1.99, got %v", p.Price.Was)
	}
	if p.BonusMechanism != "25% KORTING" {
		t.Errorf("expected bonus mechanism 25%% KORTING, got %q", p.BonusMechanism)
	}
	if !p.IsBonus {
		t.Error("expected IsBonus true")
	}
	if p.Brand != "AH Biologisch" {
		t.Errorf("expected brand AH Biologisch, got %q", p.Brand)
	}
}

func TestGetPersonalBonusDefaultsToCurrentWeek(t *testing.T) {
	const metaJSON = `{"periods":[{"bonusStartDate":"2026-07-06","bonusEndDate":"2026-07-12"}]}`
	const personalJSON = `{"sectionType":"PERSONAL","bonusGroupOrProducts":[{"product":{"webshopId":333333,"title":"X","currentPrice":1,"isBonus":true}}]}`

	var personalDate string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/mobile-services/bonuspage/v3/metadata":
			w.Write([]byte(metaJSON))
		case "/mobile-services/bonuspage/v1/personal":
			personalDate = r.URL.Query().Get("bonusStartDate")
			w.Write([]byte(personalJSON))
		default:
			t.Errorf("unexpected path %q", r.URL.Path)
		}
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))
	client.accessToken = "test"

	products, err := client.GetPersonalBonus(context.Background(), "")
	if err != nil {
		t.Fatal(err)
	}
	if personalDate != "2026-07-06" {
		t.Errorf("expected empty date to default to current week 2026-07-06, got %q", personalDate)
	}
	if len(products) != 1 {
		t.Fatalf("expected 1 product, got %d", len(products))
	}
}

func TestGetBonusPeriods(t *testing.T) {
	const metaJSON = `{"periods":[
	  {"bonusStartDate":"2026-07-06","bonusEndDate":"2026-07-12"},
	  {"bonusStartDate":"2026-07-13","bonusEndDate":"2026-07-19"}
	]}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mobile-services/bonuspage/v3/metadata" {
			t.Errorf("unexpected path %q", r.URL.Path)
		}
		w.Write([]byte(metaJSON))
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))
	client.accessToken = "test"

	periods, err := client.GetBonusPeriods(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(periods) != 2 {
		t.Fatalf("expected 2 periods, got %d", len(periods))
	}
	if periods[0].StartDate != "2026-07-06" || periods[0].EndDate != "2026-07-12" {
		t.Errorf("unexpected current period: %+v", periods[0])
	}
	if periods[1].StartDate != "2026-07-13" || periods[1].EndDate != "2026-07-19" {
		t.Errorf("unexpected next-week period: %+v", periods[1])
	}
}
