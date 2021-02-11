package query

import (
	"encoding/json"
	"fmt"
)

// Character Query
type Character struct {
	ID          int    `json:"id"`
	Name        Name   `json:"name"`
	Image       Image  `json:"image"`
	Description string `json:"description"`
	IsFavourite bool   `json:"isFavourite"`
	SiteURL     string `json:"siteUrl"`
	Favourites  int    `json:"favourites"`
	ModNotes    string `json:"modNotes"`
}

const characterQuery = `id,
					name {
						first,
						last,
						full,
						native,
					},
					image {
						large,
						medium,
					},
					description,
					isFavourite,
					siteUrl,
					favourites,
					modNotes,
				`

func newCharacterQuery() *Character {
	n := Character{}

	return &n
}

// FilterCharacterByName filters the characters from aniList by search query or name of the character.
func (c *Character) FilterCharacterByName(search string) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Character(search: "%v") {
				%s
			  }
		}
	`, search, characterQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	cleanData := CleanCharacterJSON(PostRequest(jsonValue))
	if err := json.Unmarshal(cleanData, &c); err != nil {
		panic(err)
	}

}

// FilterCharacterID search the character by it's ID
func (c *Character) FilterCharacterID(ID int) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Character(id: %v) {
				%s
			  }
		}
	`, ID, characterQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	cleanData := CleanCharacterJSON(PostRequest(jsonValue))
	if err := json.Unmarshal(cleanData, &c); err != nil {
		panic(err)
	}

}
