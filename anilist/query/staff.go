package query

import (
	"encoding/json"
	"fmt"
)

// Staff Voice actors or production staff
type Staff struct {
	ID               int            `json:"id"`
	Name             Name           `json:"name"`
	Language         string         `json:"language"`
	Image            Image          `json:"image"`
	Description      string         `json:"description"`
	IsFavourite      bool           `json:"isFavourite"`
	SiteURL          string         `json:"siteUrl"`
	StaffMedia       StaffMedia     `json:"staffMedia"`
	Characters       Characters     `json:"characters"`
	CharacterMedia   CharacterMedia `json:"characterMedia"`
	Staff            string         `json:"staff"`
	Submitter        string         `json:"submitter"`
	SubmissionStatus int            `json:"submissionStatus"`
	Favourites       int            `json:"favourites"`
	ModNotes         string         `json:"modNotes"`
	Edges            []Edges        `json:"edges"`
}

// Name staff name first, last, full name, and native
type Name struct {
	First  string `json:"first"`
	Last   string `json:"last"`
	Full   string `json:"full"`
	Native string `json:"native"`
}

// Image of the staff Large or Medium
type Image struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

// StaffMedia Media where the staff member has a production role
type StaffMedia struct {
	Edges []Edges `json:"edges"`
}

// CharacterMedia Media the actor voiced characters in. (Same data as characters with media as node instead of characters)
type CharacterMedia struct {
	Edges []Edges `json:"edges"`
}

const staffQuery = `id,
					name {
						first,
						last,
						full,
						native,
					},
					language,
					image {
						large,
						medium,
					},
					description,
					isFavourite,
					siteUrl,
					staffMedia{
						edges {
							id,
						},
					},
					characters {
						edges{
							id,
						},
					},
					characterMedia {
						edges{
							id,
						},
					},
					staff {
						id,
					},
					submitter{
						id,
					},
					submissionStatus,
					submissionNotes,
					favourites,
					modNotes,
				`

// NewStaff Create new Staff Object
func NewStaff() *Staff {
	s := Staff{}
	return &s
}

// FilterStaffByID Search Anilist Staff by it's ID
func (s *Staff) FilterStaffByID(id int) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Staff(id: %v) {
				%s
			  }
		}
	`, id, staffQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	cleanData := CleanStaffJSON(PostRequest(jsonValue))
	if err := json.Unmarshal(cleanData, &s); err != nil {
		panic(err)
	}

}

// FilterStaffByName Search Anilist Staff by it's ID
func (s *Staff) FilterStaffByName(name string) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Staff(search: "%v") {
				%s
			  }
		}
	`, name, staffQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	cleanData := CleanStaffJSON(PostRequest(jsonValue))
	if err := json.Unmarshal(cleanData, &s); err != nil {
		panic(err)
	}

}
