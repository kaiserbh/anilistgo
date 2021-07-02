package anilistgo

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

// CharacterQuery graphql constant for CharacterQuery
const CharacterQuery = `id,
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

// NewCharacterQuery creates Character objects
func NewCharacterQuery() *Character {
	n := Character{}

	return &n
}

// FilterCharacterByName filters the characters from aniList by search query or name of the character.
func (c *Character) FilterCharacterByName(search string) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Character(search: "%v") {
				%s
			  }
		}
	`, search, CharacterQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := CleanCharacterJSON(request)
	if err := json.Unmarshal(cleanData, &c); err != nil {
		return false, err
	}

	return true, nil
}

// FilterCharacterID search the character by it's ID
func (c *Character) FilterCharacterID(ID int) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Character(id: %v) {
				%s
			  }
		}
	`, ID, CharacterQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := CleanCharacterJSON(request)
	if err := json.Unmarshal(cleanData, &c); err != nil {
		return false, err
	}

	return true, nil

}
