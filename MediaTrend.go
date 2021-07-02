package anilistgo

import (
	"encoding/json"
	"fmt"
)

// MediaTrend Daily Media Statistics
type MediaTrend struct {
	MediaID      int64 `json:"mediaID"`
	Date         int64 `json:"date"`
	Trending     int   `json:"trending"`
	AverageScore int   `json:"averageScore"`
	Popularity   int   `json:"popularity"`
	InProgress   int   `json:"inProgress"`
	Releasing    bool  `json:"releasing"`
	Episode      int   `json:"episode"`
	Media        Media `json:"media"`
}

const MediaTrendQueries = `mediaId,
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

// NewMediaTrendQuery Create new NewMediaTrend Object
func NewMediaTrendQuery() *MediaTrend {
	m := MediaTrend{}

	return &m
}

// SearchByMediaID searches mediaTrend by anilist ID
func (m *MediaTrend) SearchByMediaID(id int) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			MediaTrend(mediaId: %v) {
				%s
			  }
		}
	`, id, MediaTrendQueries),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := CleanMediaTrendPageJSON(request)
	if err := json.Unmarshal(cleanData, &m); err != nil {
		return false, err
	}

	return true, nil
}

// FilterByTrendingAmount filters by trending amount
func (m *MediaTrend) FilterByTrendingAmount(trending int) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			MediaTrend(trending: %v) {
				%s
			  }
		}
	`, trending, MediaTrendQueries),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := CleanMediaTrendPageJSON(request)
	if err := json.Unmarshal(cleanData, &m); err != nil {
		return false, err
	}

	return true, nil
}

// FilterByTrendingAverageScore filters by average score
func (m *MediaTrend) FilterByTrendingAverageScore(averageScore int) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			MediaTrend(averageScore: %v) {
				%s
			  }
		}
	`, averageScore, MediaTrendQueries),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := CleanMediaTrendPageJSON(request)
	if err := json.Unmarshal(cleanData, &m); err != nil {
		return false, err
	}

	return true, nil
}

// FilterByPopularity filters by popularity
func (m *MediaTrend) FilterByPopularity(popularity int) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			MediaTrend(popularity: %v) {
				%s
			  }
		}
	`, popularity, MediaTrendQueries),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := CleanMediaTrendPageJSON(request)
	if err := json.Unmarshal(cleanData, &m); err != nil {
		return false, err
	}

	return true, nil
}

// FilterByEpisode filters by episode number
func (m *MediaTrend) FilterByEpisode(episode int) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			MediaTrend(episode: %v) {
				%s
			  }
		}
	`, episode, MediaTrendQueries),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := CleanMediaTrendPageJSON(request)
	if err := json.Unmarshal(cleanData, &m); err != nil {
		return false, err
	}

	return true, nil
}
