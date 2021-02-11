package query

import (
	"encoding/json"
	"fmt"
)

// User query for Anilist
type User struct {
	ID                      int64            `json:"id"`
	Name                    string           `json:"name"`
	About                   string           `json:"about"`
	Avatar                  Avatar           `json:"avatar"`
	BannerImage             string           `json:"bannerImage"`
	IsFollowing             bool             `json:"isFollowing"`
	IsFollower              bool             `json:"isFollower"`
	IsBlocked               bool             `json:"isBlocked"`
	Bans                    []string         `json:"bans"`
	Options                 Options          `json:"options"`
	MediaListOptions        MediaListOptions `json:"mediaListOptions"`
	Favourites              Favourites       `json:"favourites"`
	UnreadNotificationCount int              `json:"unreadNotificationCount"`
	SiteURL                 string           `json:"siteUrl"`
	Statistics              Statistics       `json:"statistics"`
	DonatorTier             int              `json:"donatorTier"`
	DonatorBadge            string           `json:"donatorBadge"`
	ModeratorStatus         string           `json:"moderatorStatus"`
	UpdatedAt               int              `json:"updatedAt"`
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

// Favourites User Favourites
type Favourites struct {
	Anime      Anime      `json:"anime"`
	Manga      Manga      `json:"manga"`
	Characters Characters `json:"characters"`
	Staff      StaffU     `json:"staff"`
	Studios    Studio     `json:"studios"`
}

// Anime id of use favourites
type Anime struct {
	Edges             []Edges        `json:"edges"`
	Count             int            `json:"count"`
	MeanScore         float64        `json:"meanScore"`
	StandardDeviation float64        `json:"standardDeviation"`
	MinutesWatched    int            `json:"minutesWatched"`
	EpisodesWatched   int            `json:"episodesWatched"`
	Formats           []Formats      `json:"formats"`
	Statuses          []Statuses     `json:"statuses"`
	Scores            []Scores       `json:"scores"`
	Lengths           []Lengths      `json:"lengths"`
	ReleaseYears      []ReleaseYears `json:"releaseYears"`
	StartYears        []StartYears   `json:"startYears"`
	Genres            []Genres       `json:"genres"`
	Tags              []TagsUser     `json:"tags"`
	Countries         []Countries    `json:"countries"`
}

// Formats user statistic formats
type Formats struct {
	Format string `json:"format"`
}

// Statuses user statistic status
type Statuses struct {
	Status string `json:"status"`
}

// Scores user statistic scores
type Scores struct {
	Score string `json:"status"`
}

// Lengths user statistic lengths
type Lengths struct {
	Length string `json:"length"`
}

// ReleaseYears user statistic ReleaseYear
type ReleaseYears struct {
	ReleaseYear int `json:"releaseYear"`
}

// StartYears user statistic StartYears
type StartYears struct {
	StartYear int `json:"startYear"`
}

// Genres user statistic Genres
type Genres struct {
	Genre string `json:"genre"`
}

// TagsUser A tag that describes a theme or element of the media
type TagsUser struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int64   `json:"minutesWatched"`
	MediaIds       []int   `json:"mediaIds"`
	Tag            Tag     `json:"tag"`
}

// Tag A tag that describes a theme or element of the media\
type Tag struct {
	ID int `json:"id"`
}

// Countries of the media that the user watched/read from.
type Countries struct {
	Country string `json:"country"`
}

// Manga id of use favourites
type Manga struct {
	Edges        []Edges        `json:"edges"`
	Count        int            `json:"count"`
	MeanScore    float64        `json:"meanScore"`
	ChapterRead  int            `json:"chapterRead"`
	Formats      []Formats      `json:"formats"`
	Statuses     []Statuses     `json:"statuses"`
	Scores       []Scores       `json:"scores"`
	ReleaseYears []ReleaseYears `json:"releaseYears"`
	StartYears   []StartYears   `json:"startYears"`
	Genres       []Genres       `json:"genres"`
	Tags         []TagsUser     `json:"tags"`
	Countries    []Countries    `json:"countries"`
}

// Statistics User statistics anime or scores
type Statistics struct {
	Anime Anime `json:"anime"`
	Manga Manga `json:"manga"`
}

// StaffU staff favourite on users profile
type StaffU struct {
	Edges []Edges `json:"edges"`
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
					favourites {
						anime {
							edges {
								id,
							},
						},
						manga {
							edges {
								id,
							},
						},
						characters {
							edges {
								id,
							},
						},
						staff {
							edges {
								id,
							},
						},
						studios {
							edges {
								id,
							},
						},
					},
					statistics {
						anime {
							count,
							meanScore,
							standardDeviation,
							minutesWatched,
							episodesWatched,
							formats {
								format,
							},
							statuses {
								status,
							},
							scores {
								score,
							},
							lengths {
								length,
							},
							releaseYears {
								releaseYear,
							},
							startYears {
								startYear,
							},
							genres {
								genre,
							},
							tags {
								count,
								meanScore,
								minutesWatched,
								mediaIds,
								tag {
									id,
								},
							},
							countries {
          						country,
       						},
						},
						manga {
							count,
							meanScore,
							chaptersRead,
							formats {
								format,
							},
							statuses {
								status,
							},
							scores {
								score,
							},
							lengths {
								length,
							},
							releaseYears {
								releaseYear,
							},
							startYears {
								startYear,
							},
							genres {
								genre,
							},
							tags {
								count,
								meanScore,
								minutesWatched,
								mediaIds,
								tag {
									id,
								},
							},
							countries {
          						country,
       						},
						},	
					},
					unreadNotificationCount,
					siteUrl,
					donatorTier,
					donatorBadge,
					moderatorStatus,
					updatedAt,
					`

// NewUserQuery Create new User Object
func NewUserQuery() *User {
	u := User{}
	return &u
}

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

// FilterUserByID Search Anilist User by it's ID
func (u *User) FilterUserByID(ID int) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			User(id: %v) {
				%s
			  }
		}
	`, ID, userQuery),
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
