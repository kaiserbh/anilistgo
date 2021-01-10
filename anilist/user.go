package anilist

import (
	"encoding/json"
	"fmt"
)

// User query for Anilist
type User struct {
	ID               int64            `json:"id"`
	Name             string           `json:"name"`
	About            string           `json:"about"`
	Avatar           Avatar           `json:"avatar"`
	BannerImage      string           `json:"bannerImage"`
	IsFollowing      bool             `json:"isFollowing"`
	IsFollower       bool             `json:"isFollower"`
	IsBlocked        bool             `json:"isBlocked"`
	Bans             []string         `json:"bans"`
	Options          Options          `json:"options"`
	MediaListOptions MediaListOptions `json:"mediaListOptions"`
}

// Avatar Images Large and medium size
type Avatar struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

// Options User Options
type Options struct {
	TitleLanguage       string `json:"titleLanguage"`
	DisplayAdultContent bool   `json:"displayAdultContent"`
	AiringNotification  bool   `json:"airingNotification"`
	ProfileColor        string `json:"profileColor"`
}

// MediaListOptions options for MediaList
type MediaListOptions struct {
	ScoreFormat string    `json:"scoreFormat"`
	RowOrder    string    `json:"rowOrder"`
	AnimeList   AnimeList `json:"animeList"`
	MangaList   MangaList `json:"mangaList"`
}

// AnimeList part of MediaList
type AnimeList struct {
	SectionOrder                  []string `json:"sectionOrder"`
	SplitCompletedSectionByFormat bool     `json:"splitCompletedSectionByFormat"`
	Theme                         Theme    `json:"theme"`
	CustomLists                   []string `json:"customLists"`
	AdvancedScoring               []string `json:"advancedScoring"`
	AdvancedScoringEnabled        bool     `json:"advancedScoringEnabled"`
}

// MangaList part of MediaList
type MangaList struct {
	SectionOrder                  []string `json:"sectionOrder"`
	SplitCompletedSectionByFormat bool     `json:"splitCompletedSectionByFormat"`
	Theme                         Theme    `json:"theme"`
	CustomLists                   []string `json:"customLists"`
	AdvancedScoring               []string `json:"advancedScoring"`
	AdvancedScoringEnabled        bool     `json:"advancedScoringEnabled"`
}

// Theme user theme types
type Theme struct {
	ThemeType   string `json:"themeType"`
	Theme       string `json:"theme"`
	CoverImages string `json:"coverImages"`
}

const userQuery = `
					id,
    				name,
    				about,
    				avatar {
      					large,
      					medium,
					},
					bannerImage,
    				isFollowing,
    				isFollower,
    				isBlocked,
					bans,
					options {
      					titleLanguage,
      					displayAdultContent,
      					airingNotifications,
      					profileColor,
    				},
					mediaListOptions {
      					scoreFormat,
						rowOrder,
						animeList {
							sectionOrder,
							splitCompletedSectionByFormat,
							theme,
							customLists,
							advancedScoring,
							advancedScoringEnabled,
						  }
						mangaList {
							sectionOrder,
							splitCompletedSectionByFormat,
							theme,
							customLists,
							advancedScoring,
							advancedScoringEnabled,
						},
					},
					`

// FilterUserByName Search Anilist User by it's userName
func (u *User) FilterUserByName(name string) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			User(name: "%v") {
				%s
			  }
		}
	`, name, userQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	cleanData := CleanUserJSON(PostRequest(jsonValue))
	if err := json.Unmarshal(cleanData, &u); err != nil {
		panic(err)
	}

}
