package anilistgo

import (
	"encoding/json"
	"fmt"
)

// User query for Anilist
type User struct {
	// ID The id of the user
	ID int64 `json:"id"`
	// Name The name of the user
	Name string `json:"name"`
	// About The bio written by user (Markdown)
	About string `json:"about"`
	// Avatar The user's avatar images
	Avatar UserAvatar `json:"avatar"`
	// BannerImage The user's banner images
	BannerImage string `json:"bannerImage"`
	// IsFollowing If the authenticated user if following this user
	IsFollowing bool `json:"isFollowing"`
	// IsFollower If this user if following the authenticated user
	IsFollower bool `json:"isFollower"`
	// IsBlocked If the user is blocked by the authenticated user
	IsBlocked bool `json:"isBlocked"`
	// Bans
	Bans []map[string]string `json:"bans"`
	// Options The user's general options
	Options UserOptions `json:"options"`
	// MediaListOptions The user's media list options
	MediaListOptions MediaListOptions `json:"mediaListOptions"`
	// Favourites The users favourites
	Favourites Favourites `json:"favourites"`
	// UnreadNotificationCount The number of unread notifications the user has
	UnreadNotificationCount int `json:"unreadNotificationCount"`
	// SiteURL The url for the user page on the AniList website
	SiteURL string `json:"siteUrl"`
	// Statistics The users anime & manga list statistics
	Statistics UserStatisticTypes `json:"statistics"`
	// DonatorTier The donation tier of the user
	DonatorTier int `json:"donatorTier"`
	// DonatorBadge Custom donation badge text
	DonatorBadge string `json:"donatorBadge"`
	// ModeratorRoles The user's moderator roles if they are a site moderator
	ModeratorRoles string `json:"moderatorRoles"`
	// CreatedAt When the user's account was created. (Does not exist for accounts created before 2020)
	CreatedAt int `json:"createdAt"`
	// UpdatedAt When the user's data was last updated
	UpdatedAt int `json:"updatedAt"`
}

// UserAvatar Images Large and medium size
type UserAvatar struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

// UserOptions A user's general options
type UserOptions struct {
	// TitleLanguage The language the user wants to see media titles in
	TitleLanguage string `json:"titleLanguage"`
	// DisplayAdultContent Whether the user has enabled viewing of 18+ content
	DisplayAdultContent bool `json:"displayAdultContent"`
	// AiringNotification Whether the user receives notifications when a show they are watching aires
	AiringNotification bool `json:"airingNotification"`
	// ProfileColor Profile highlight color (blue, purple, pink, orange, red, green, gray)
	ProfileColor string `json:"profileColor"`
	// NotificationOptions Notification options
	NotificationOptions []NotificationOption `json:"notificationOptions"`
	// Timezone The user's timezone offset (Auth user only)
	Timezone string `json:"timezone"`
	// ActivityMergeTime Minutes between activity for them to be merged together. 0 is Never, Above 2 weeks (20160 mins) is Always.
	ActivityMergeTime int `json:"activityMergeTime"`
	// StaffNameLanguage The language the user wants to see staff and character names in
	StaffNameLanguage string `json:"staffNameLanguage"`
}

// MediaListOptions options for MediaList
type MediaListOptions struct {
	ScoreFormat string               `json:"scoreFormat"`
	RowOrder    string               `json:"rowOrder"`
	AnimeList   MediaListTypeOptions `json:"animeList"`
	MangaList   MediaListTypeOptions `json:"mangaList"`
}

// MediaListTypeOptions A user's list options for anime or manga lists
type MediaListTypeOptions struct {
	SectionOrder                  []string `json:"sectionOrder"`
	SplitCompletedSectionByFormat bool     `json:"splitCompletedSectionByFormat"`
	CustomLists                   []string `json:"customLists"`
	AdvancedScoring               []string `json:"advancedScoring"`
	AdvancedScoringEnabled        bool     `json:"advancedScoringEnabled"`
}

type Favourites struct {
	Anime      MediaConnection     `json:"anime"`
	Manga      MediaConnection     `json:"manga"`
	Characters CharacterConnection `json:"characters"`
	Staff      StaffConnection     `json:"staff"`
	Studios    StudioConnection    `json:"studios"`
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
func (u *User) FilterUserByName(name string) (bool, error) {
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
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := CleanUserJSON(request)
	if err := json.Unmarshal(cleanData, &u); err != nil {
		return false, err
	}

	return true, nil
}

// FilterUserByID Search Anilist User by it's ID
func (u *User) FilterUserByID(ID int) (bool, error) {
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
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := CleanUserJSON(request)

	if err := json.Unmarshal(cleanData, &u); err != nil {
		return false, err
	}

	return true, nil
}
