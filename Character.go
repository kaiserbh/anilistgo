package anilistgo

import (
	"encoding/json"
	"fmt"
)

// Character Query
type Character struct {
	ID                 int             `json:"id"`
	Name               CharacterName   `json:"name"`
	Image              CharacterImage  `json:"image"`
	Description        string          `json:"description"`
	Gender             string          `json:"gender"`
	DateOfBirth        FuzzyDate       `json:"dateOfBirth"`
	Age                string          `json:"age"`
	IsFavourite        bool            `json:"isFavourite"`
	IsFavouriteBlocked bool            `json:"isFavouriteBlocked"`
	SiteURL            string          `json:"siteUrl"`
	Media              MediaConnection `json:"media"`
	Favourites         int             `json:"favourites"`
	ModNotes           string          `json:"modNotes"`
}

// CharacterImage of the characters Large or Medium
type CharacterImage struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

// CharacterName The names of the character
type CharacterName struct {
	// First The character's given name
	First string `json:"first"`
	// Middle The character's middle name
	Middle string `json:"middle"`
	// Last The character's surname
	Last string `json:"last"`
	// Full The character's first and last name
	Full string `json:"full"`
	// Native The character's full name in their native language
	Native string `json:"native"`
	// Alternative Other names the character might be referred to as
	Alternative []string `json:"alternative"`
	// AlternativeSpoiler Other names the character might be referred to as but are spoilers
	AlternativeSpoiler []string `json:"alternativeSpoiler"`
	// UserPreferred The currently authenticated users preferred name language. Default romaji for non-authenticated
	UserPreferred string `json:"userPreferred"`
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
