package appie

import (
	"context"
	"fmt"
	"strconv"
)

const fetchMemberQuery = `query FetchMember {
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
}`

// memberResponse matches the GraphQL response for FetchMember.
type memberResponse struct {
	Member struct {
		ID           int    `json:"id"`
		EmailAddress string `json:"emailAddress"`
		Gender       string `json:"gender"`
		DateOfBirth  string `json:"dateOfBirth"`
		PhoneNumber  string `json:"phoneNumber"`
		IsB2B        bool   `json:"isB2B"`
		Name         struct {
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Address struct {
			Street           string `json:"street"`
			HouseNumber      int    `json:"houseNumber"`
			HouseNumberExtra string `json:"houseNumberExtra"`
			PostalCode       string `json:"postalCode"`
			City             string `json:"city"`
			CountryCode      string `json:"countryCode"`
		} `json:"address"`
		Cards struct {
			Bonus    string `json:"bonus"`
			Gall     string `json:"gall"`
			Airmiles string `json:"airmiles"`
		} `json:"cards"`
		CustomerProfileAudiences   []string `json:"customerProfileAudiences"`
		CustomerProfileProperties  []struct {
			Key   string `json:"key"`
			Value any    `json:"value"`
		} `json:"customerProfileProperties"`
	} `json:"member"`
}

// GetMember retrieves the member profile including address,
// loyalty cards (Bonus, Gall & Gall), and customer segmentation data.
func (c *Client) GetMember(ctx context.Context) (*Member, error) {
	var resp memberResponse
	if err := c.doGraphQL(ctx, fetchMemberQuery, nil, &resp); err != nil {
		return nil, fmt.Errorf("get member failed: %w", err)
	}

	m := resp.Member
	return &Member{
		ID:          strconv.Itoa(m.ID),
		FirstName:   m.Name.First,
		LastName:    m.Name.Last,
		Email:       m.EmailAddress,
		Gender:      m.Gender,
		DateOfBirth: m.DateOfBirth,
		PhoneNumber: m.PhoneNumber,
		Address: Address{
			Street:           m.Address.Street,
			HouseNumber:      m.Address.HouseNumber,
			HouseNumberExtra: m.Address.HouseNumberExtra,
			PostalCode:       m.Address.PostalCode,
			City:             m.Address.City,
			CountryCode:      m.Address.CountryCode,
		},
		BonusCardNumber: m.Cards.Bonus,
		GallCardNumber:  m.Cards.Gall,
		Audiences:       m.CustomerProfileAudiences,
	}, nil
}

