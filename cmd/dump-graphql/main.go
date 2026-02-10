package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"cmp"
	"os"
	"slices"

	appie "github.com/gwillem/appie-go"
)

const introspectionQuery = `{
  __schema {
    queryType { name }
    mutationType { name }
    subscriptionType { name }
    types {
      name
      kind
      description
      fields {
        name
        description
        args { name type { name kind } }
        type { name kind ofType { name kind } }
      }
    }
  }
}`

type introspectionResponse struct {
	Data struct {
		Schema struct {
			QueryType        *typeName   `json:"queryType"`
			MutationType     *typeName   `json:"mutationType"`
			SubscriptionType *typeName   `json:"subscriptionType"`
			Types            []gqlType   `json:"types"`
		} `json:"__schema"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type typeName struct {
	Name string `json:"name"`
}

type gqlType struct {
	Name        string     `json:"name"`
	Kind        string     `json:"kind"`
	Description string     `json:"description"`
	Fields      []gqlField `json:"fields"`
}

type gqlField struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Args        []struct {
		Name string `json:"name"`
		Type struct {
			Name string `json:"name"`
			Kind string `json:"kind"`
		} `json:"type"`
	} `json:"args"`
	Type struct {
		Name   string `json:"name"`
		Kind   string `json:"kind"`
		OfType *struct {
			Name string `json:"name"`
			Kind string `json:"kind"`
		} `json:"ofType"`
	} `json:"type"`
}

func main() {
	client, err := appie.NewWithConfig(".appie.json")
	if err != nil {
		log.Fatal(err)
	}

	if !client.IsAuthenticated() {
		log.Fatal("Not authenticated. Run login first.")
	}

	// Build request
	reqBody, _ := json.Marshal(map[string]string{"query": introspectionQuery})
	req, err := http.NewRequest("POST", "https://api.ah.nl/graphql", bytes.NewReader(reqBody))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Appie/9.28 (iPhone17,3; iPhone; CPU OS 26_1 like Mac OS X)")
	req.Header.Set("x-client-name", "appie-ios")
	req.Header.Set("x-client-version", "9.28")
	req.Header.Set("Authorization", "Bearer "+client.AccessToken())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		fmt.Printf("HTTP %d: %s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}

	var result introspectionResponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Failed to parse response: %v\nBody: %s", err, string(body))
	}

	if len(result.Errors) > 0 {
		fmt.Println("GraphQL Errors:")
		for _, e := range result.Errors {
			fmt.Printf("  - %s\n", e.Message)
		}
		fmt.Println("\nIntrospection may be disabled on this endpoint.")
		os.Exit(1)
	}

	// Print schema info
	fmt.Println("# Albert Heijn GraphQL Schema")
	fmt.Println()

	if result.Data.Schema.QueryType != nil {
		fmt.Printf("Query Type: %s\n", result.Data.Schema.QueryType.Name)
	}
	if result.Data.Schema.MutationType != nil {
		fmt.Printf("Mutation Type: %s\n", result.Data.Schema.MutationType.Name)
	}
	if result.Data.Schema.SubscriptionType != nil {
		fmt.Printf("Subscription Type: %s\n", result.Data.Schema.SubscriptionType.Name)
	}
	fmt.Println()

	// Group types by kind
	typesByKind := make(map[string][]gqlType)
	for _, t := range result.Data.Schema.Types {
		// Skip internal types
		if len(t.Name) > 0 && t.Name[0] == '_' {
			continue
		}
		typesByKind[t.Kind] = append(typesByKind[t.Kind], t)
	}

	// Print types by category
	kinds := []string{"OBJECT", "INPUT_OBJECT", "ENUM", "INTERFACE", "UNION", "SCALAR"}
	for _, kind := range kinds {
		types := typesByKind[kind]
		if len(types) == 0 {
			continue
		}

		slices.SortFunc(types, func(a, b gqlType) int {
			return cmp.Compare(a.Name, b.Name)
		})

		fmt.Printf("## %s Types (%d)\n\n", kind, len(types))
		for _, t := range types {
			fmt.Printf("### %s\n", t.Name)
			if t.Description != "" {
				fmt.Printf("%s\n", t.Description)
			}
			if len(t.Fields) > 0 {
				fmt.Println()
				for _, f := range t.Fields {
					typeName := f.Type.Name
					if typeName == "" && f.Type.OfType != nil {
						typeName = f.Type.OfType.Name
					}
					fmt.Printf("- `%s`: %s\n", f.Name, typeName)
				}
			}
			fmt.Println()
		}
	}
}
