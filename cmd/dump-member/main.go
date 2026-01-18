package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	appie "github.com/gwillem/appie-go"
)

// memberProfileQuery fetches member profile with analytics and segmentation data.
const memberProfileQuery = `query FetchMember {
  member {
    id
    emailAddress
    name { first last }
    analytics { digimon idmon idsas batch firebase sitespect }
    customerProfileAudiences
    customerProfileProperties { key value }
  }
}`

// lifestyleScoreQuery fetches lifestyle check scores.
const lifestyleScoreQuery = `query LifestyleCheckScore {
  lifestyleCheckScore {
    chosenGoal
    goalImageUrl
    totalAverageScore
    chapterScores {
      chapterId
      name
      iconType
      averageScore
    }
  }
}`

// budgetQuery fetches the user's spending budget via GraphQL.
const budgetQuery = `query OrderBudget {
  orderBudgetV2 {
    result {
      id
      memberId
      proposition
      type
      amount { amount }
      carriedOver { amount }
      realizedCosts { amount }
      unRealizedCosts { amount }
      availableToSpend { amount }
      createdAt
      startAt
    }
  }
}`

// favoriteStoreQuery fetches the user's favorite store with geolocation.
const favoriteStoreQuery = `query FavoriteStore {
  storesFavouriteStore {
    id
    name
    address {
      street
      houseNumber
      postalCode
      city
    }
    geoLocation {
      latitude
      longitude
    }
  }
}`

type memberProfileResponse struct {
	Member struct {
		ID           int    `json:"id"`
		EmailAddress string `json:"emailAddress"`
		Name         struct {
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Analytics struct {
			Digimon   string `json:"digimon"`
			Idmon     string `json:"idmon"`
			Idsas     string `json:"idsas"`
			Batch     string `json:"batch"`
			Firebase  string `json:"firebase"`
			Sitespect string `json:"sitespect"`
		} `json:"analytics"`
		CustomerProfileAudiences  []string `json:"customerProfileAudiences"`
		CustomerProfileProperties []struct {
			Key   string `json:"key"`
			Value any    `json:"value"`
		} `json:"customerProfileProperties"`
	} `json:"member"`
}

type lifestyleScoreResponse struct {
	LifestyleCheckScore *struct {
		ChosenGoal        string `json:"chosenGoal"`
		GoalImageURL      string `json:"goalImageUrl"`
		TotalAverageScore int    `json:"totalAverageScore"`
		ChapterScores     []struct {
			ChapterID    string `json:"chapterId"`
			Name         string `json:"name"`
			IconType     string `json:"iconType"`
			AverageScore int    `json:"averageScore"`
		} `json:"chapterScores"`
	} `json:"lifestyleCheckScore"`
}

type favoriteStoreResponse struct {
	StoresFavouriteStore *struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Address struct {
			Street      string `json:"street"`
			HouseNumber string `json:"houseNumber"`
			PostalCode  string `json:"postalCode"`
			City        string `json:"city"`
		} `json:"address"`
		GeoLocation *struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"geoLocation"`
	} `json:"storesFavouriteStore"`
}

type budgetResponse struct {
	OrderBudgetV2 struct {
		Result *struct {
			ID               string `json:"id"`
			MemberID         int    `json:"memberId"`
			Proposition      string `json:"proposition"`
			Type             string `json:"type"`
			Amount           money  `json:"amount"`
			CarriedOver      money  `json:"carriedOver"`
			RealizedCosts    money  `json:"realizedCosts"`
			UnRealizedCosts  money  `json:"unRealizedCosts"`
			AvailableToSpend money  `json:"availableToSpend"`
			CreatedAt        string `json:"createdAt"`
			StartAt          string `json:"startAt"`
		} `json:"result"`
	} `json:"orderBudgetV2"`
}

type money struct {
	Amount float64 `json:"amount"`
}

func main() {
	client, err := appie.NewWithConfig(".appie.json")
	if err != nil {
		log.Fatal(err)
	}

	if !client.IsAuthenticated() {
		log.Fatal("Not authenticated. Run login first.")
	}

	ctx := context.Background()

	fmt.Println("# Member Profile Data")
	fmt.Println()

	// 1. Fetch member profile with audiences and properties
	fmt.Println("## Customer Profile")
	fmt.Println()
	fetchMemberProfile(ctx, client)

	// 2. Fetch budget
	fmt.Println()
	fmt.Println("## Current Budget")
	fmt.Println()
	fetchBudget(ctx, client)

	// 3. Fetch lifestyle score
	fmt.Println()
	fmt.Println("## Lifestyle Check Score")
	fmt.Println()
	fetchLifestyleScore(ctx, client)

	// 4. Fetch favorite store with geolocation
	fmt.Println()
	fmt.Println("## Favorite Store (GeoLocation)")
	fmt.Println()
	fetchFavoriteStore(ctx, client)
}

func fetchMemberProfile(ctx context.Context, client *appie.Client) {
	var resp memberProfileResponse
	if err := doGraphQL(ctx, client, memberProfileQuery, nil, &resp); err != nil {
		fmt.Printf("Error fetching member profile: %v\n", err)
		return
	}

	m := resp.Member
	fmt.Printf("Member: %s %s (ID: %d)\n", m.Name.First, m.Name.Last, m.ID)
	fmt.Printf("Email: %s\n", m.EmailAddress)

	fmt.Println()
	fmt.Println("### Analytics IDs")
	fmt.Printf("  - digimon:   %s\n", m.Analytics.Digimon)
	fmt.Printf("  - idmon:     %s\n", m.Analytics.Idmon)
	fmt.Printf("  - idsas:     %s\n", m.Analytics.Idsas)
	fmt.Printf("  - batch:     %s\n", m.Analytics.Batch)
	fmt.Printf("  - firebase:  %s\n", m.Analytics.Firebase)
	fmt.Printf("  - sitespect: %s\n", m.Analytics.Sitespect)

	fmt.Println()
	fmt.Printf("### Customer Profile Audiences (%d)\n", len(m.CustomerProfileAudiences))
	if len(m.CustomerProfileAudiences) == 0 {
		fmt.Println("  (none)")
	} else {
		for _, audience := range m.CustomerProfileAudiences {
			fmt.Printf("  - %s\n", audience)
		}
	}

	fmt.Println()
	fmt.Printf("### Customer Profile Properties (%d)\n", len(m.CustomerProfileProperties))
	if len(m.CustomerProfileProperties) == 0 {
		fmt.Println("  (none)")
	} else {
		for _, prop := range m.CustomerProfileProperties {
			fmt.Printf("  - %s: %v\n", prop.Key, prop.Value)
		}
	}
}

func fetchBudget(ctx context.Context, client *appie.Client) {
	var resp budgetResponse
	if err := doGraphQL(ctx, client, budgetQuery, nil, &resp); err != nil {
		fmt.Printf("Error fetching budget: %v\n", err)
		return
	}

	if resp.OrderBudgetV2.Result == nil {
		fmt.Println("No budget configured")
		return
	}

	b := resp.OrderBudgetV2.Result
	fmt.Printf("Budget ID: %s\n", b.ID)
	fmt.Printf("Type: %s\n", b.Type)
	fmt.Printf("Proposition: %s\n", b.Proposition)
	fmt.Printf("Amount: €%.2f\n", b.Amount.Amount)
	fmt.Printf("Carried Over: €%.2f\n", b.CarriedOver.Amount)
	fmt.Printf("Realized Costs: €%.2f\n", b.RealizedCosts.Amount)
	fmt.Printf("Unrealized Costs: €%.2f\n", b.UnRealizedCosts.Amount)
	fmt.Printf("Available to Spend: €%.2f\n", b.AvailableToSpend.Amount)
	fmt.Printf("Created: %s\n", b.CreatedAt)
	fmt.Printf("Start: %s\n", b.StartAt)
}

func fetchLifestyleScore(ctx context.Context, client *appie.Client) {
	var resp lifestyleScoreResponse
	if err := doGraphQL(ctx, client, lifestyleScoreQuery, nil, &resp); err != nil {
		fmt.Printf("Error fetching lifestyle score: %v\n", err)
		return
	}

	if resp.LifestyleCheckScore == nil {
		fmt.Println("No lifestyle score available (questionnaire not completed)")
		return
	}

	score := resp.LifestyleCheckScore
	fmt.Printf("Total Average Score: %d\n", score.TotalAverageScore)
	fmt.Printf("Chosen Goal: %s\n", score.ChosenGoal)

	if len(score.ChapterScores) > 0 {
		fmt.Println()
		fmt.Println("### Chapter Scores")
		for _, ch := range score.ChapterScores {
			fmt.Printf("  - %s (%s): %d\n", ch.Name, ch.ChapterID, ch.AverageScore)
		}
	}
}

func fetchFavoriteStore(ctx context.Context, client *appie.Client) {
	var resp favoriteStoreResponse
	if err := doGraphQL(ctx, client, favoriteStoreQuery, nil, &resp); err != nil {
		fmt.Printf("Error fetching favorite store: %v\n", err)
		return
	}

	if resp.StoresFavouriteStore == nil {
		fmt.Println("No favorite store set")
		return
	}

	store := resp.StoresFavouriteStore
	fmt.Printf("Store: %s (ID: %d)\n", store.Name, store.ID)
	fmt.Printf("Address: %s %s, %s %s\n",
		store.Address.Street,
		store.Address.HouseNumber,
		store.Address.PostalCode,
		store.Address.City)

	if store.GeoLocation != nil {
		fmt.Printf("Latitude: %f\n", store.GeoLocation.Latitude)
		fmt.Printf("Longitude: %f\n", store.GeoLocation.Longitude)
		fmt.Printf("Google Maps: https://www.google.com/maps?q=%f,%f\n",
			store.GeoLocation.Latitude, store.GeoLocation.Longitude)
	}
}

// doGraphQL is a helper to make GraphQL requests using the client's token.
func doGraphQL(ctx context.Context, client *appie.Client, query string, variables map[string]any, result any) error {
	type graphQLRequest struct {
		Query     string         `json:"query"`
		Variables map[string]any `json:"variables,omitempty"`
	}
	type graphQLError struct {
		Message string `json:"message"`
	}
	type graphQLResponse struct {
		Data   json.RawMessage `json:"data"`
		Errors []graphQLError  `json:"errors"`
	}

	reqBody := graphQLRequest{
		Query:     query,
		Variables: variables,
	}

	var resp graphQLResponse
	if err := doRequest(ctx, client, http.MethodPost, "/graphql", reqBody, &resp); err != nil {
		return err
	}

	if len(resp.Errors) > 0 {
		return fmt.Errorf("graphql error: %s", resp.Errors[0].Message)
	}

	if result != nil {
		if err := json.Unmarshal(resp.Data, result); err != nil {
			return fmt.Errorf("failed to decode graphql response: %w", err)
		}
	}

	return nil
}

// doRequest is a helper to make HTTP requests using the client's token.
func doRequest(ctx context.Context, client *appie.Client, method, path string, body, result any) error {
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, "https://api.ah.nl"+path, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "Appie/9.28 (iPhone17,3; iPhone; CPU OS 26_1 like Mac OS X)")
	req.Header.Set("x-client-name", "appie-ios")
	req.Header.Set("x-client-version", "9.28")
	req.Header.Set("x-application", "AHWEBSHOP")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.AccessToken())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %d %s", resp.StatusCode, string(respBody))
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
