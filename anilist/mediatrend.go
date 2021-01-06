package anilist

import (
	"encoding/json"
	"fmt"
)

// MediaTrend Daily Media Statistics
type MediaTrend struct {
	MediaID      int64 `json:"mediaID"`
	Date         int64 `json:"date"`
	Trending     int   `json:"Trending"`
	AverageScore int   `json:"averageScore"`
	Popularity   int   `json:"popularity"`
	InProgress   int   `json:"inProgress"`
	Releasing    int   `json:"releasing"`
	Episode      int   `json:"episode"`
	Media        Media `json:"media"`
}

const mediaTrendQueries = `mediaId,
						   date,
						   trending,
						   averageScore,
						   inProgress,
						   releasing,
						   episode,
						   media{
							   id,
							idMal,
							title {
								romaji,
								english,
								native,
								userPreferred,
							},
							type,
							format,
							status,
							description,
							startDate {
								year,
								month,
								day,
							},
							endDate {
								year,
								month,
								day,
							},
							season,
							seasonYear,
							seasonInt,
							episodes,
							duration,
							chapters,
							volumes,
							countryOfOrigin,
							isLicensed,
							source,
							hashtag,
							trailer {
								id,
							},
							updatedAt,
							coverImage {
								extraLarge,
								large,
								medium,
								color,
							},
							bannerImage,
							genres,
							synonyms,
							averageScore,
							meanScore,
							popularity,
							isLocked,
							trending,
							favourites,
							}`

// NewMediaTrend Create new NewMediaTrend Object
func NewMediaTrend() *MediaTrend {
	m := MediaTrend{}

	return &m
}

func (m *MediaTrend) searchByID(id int) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			MediaTrend(id: %v) {
				%s
			  }
		}
	`, id, mediaTrendQueries),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	cleanData := CleanPageJSON(PostRequest(jsonValue))

	if err := json.Unmarshal(cleanData, &m); err != nil {
		panic(err)
	}
}
