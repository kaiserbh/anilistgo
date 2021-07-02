package anilistgo

import (
	"encoding/json"
	"fmt"
)

// Media object to store the json from Anilist
type Media struct {
	// Media ID from anilist
	ID int `json:"id"`
	// Media ID from mal (MyAnimeList not always accurate with the new animes or upcoming ones)
	IDMAL int `json:"idMal"`
	// Title of the Media
	Title Title `json:"title"`
	// MediaType such as MANGA, ANIME
	MediaType string `json:"type"`
	// MediaFormat such as TV, TV_SHORT (under 15 min), MOVIE, SPECIAL, OVA, ONA, MUSIC, MANGA, NOVEL, ONE_SHOT
	MediaFormat string `json:"format"`
	// Media Status such as FINISHED, RELEASING, NOT_YET_RELEASED, CANCELLED and HIATUS
	Status string `json:"status"`
	// Description Short description of the media's story and characters
	Description string `json:"description"`
	// StartDate The first official release date of the media
	StartDate StartDate `json:"startDate"`
	// EndDate The last official release date of the media
	EndDate EndDate `json:"endDate"`
	// Season The season the media was initially released in (Winter Months December to February, SPRING Months March to May, SUMMER Months June to August and FALL Months September to November)
	Season string `json:"season"`
	// The season year the media was initially released in
	SeasonYear int `json:"seasonYear"`
	// The year & season the media was initially released in
	SeasonInt int `json:"seasonInt"`
	// The amount of episodes the anime has when complete
	Episodes int `json:"episodes"`
	// The general length of each anime episode in minutes
	Duration int `json:"duration"`
	// The amount of chapters the manga has when complete
	Chapters int `json:"chapters"`
	// The amount of volumes the manga has when complete
	Volumes int `json:"volumes"`
	// Where the media was created. (ISO 3166-1 alpha-2)
	CountryOfOrigin string `json:"countryOfOrigin"`
	// If the media is officially licensed or a self-published doujin release
	IsLicensed bool `json:"isLicensed"`
	// Source type the media was adapted from. (ORIGINAL, MANGA, LIGHT_NOVEL, VISUAL_NOVEL, VIDEO_GAME, OTHER, NOVEL, DOUJINSHI, ANIME )
	Source string `json:"source"`
	// Official Twitter hashtags for the media
	HashTag string `json:"hashtag"`
	// Media trailer or advertisement unique YouTube string
	Trailer Trailer `json:"trailer"`
	// When the media's data was last updated
	UpdatedAt int `json:"updatedAt"`
	// The cover images of the media
	CoverImage CoverImage `json:"coverImage"`
	// The banner image of the media
	BannerImage string `json:"bannerImage"`
	// The genres of the media
	Genres []string `json:"genres"`
	// Alternative titles of the media
	Synonyms []string `json:"synonyms"`
	// media's average score
	AverageScore int `json:"averageScore"`
	// Mean score of all the user's scores of the media
	MeanScore int `json:"meanScore"`
	// Popularity of the media
	Popularity int `json:"popularity"`
	// Locked media may not be added to lists our favorited. This may be due to the entry pending for deletion or other reasons.
	IsLocked bool `json:"isLocked"`
	// The amount of related activity in the past hour
	Trending int `json:"trending"`
	// The amount of user's who have favourited the media
	Favourites int `json:"favourites"`
	// List of tags that describes elements and themes of the media
	Tags []Tags `json:"tags"`
	// Other media in the same or connecting franchise
	Relations Relations `json:"relations"`
	// The characters in the media
	Characters Characters `json:"characters"`
	// The staff who produced the media
	Staff Staff `json:"staff"`
	// The companies who produced the media
	Studio Studio `json:"studio"`
	// If the media is marked as favourite by the current authenticated user
	IsFavourite bool `json:"isFavourite"`
	// If the media is intended only for 18+ adult audiences
	IsAdult bool `json:"isAdult"`
	// The media's next episode airing schedule
	NextAiringEpisode NextAiringEpisode `json:"nextAiringEpisode"`
	// The media's entire airing schedule
	AiringSchedule AiringSchedule `json:"airingSchedule"`
	// External links to another site related to the media
	ExternalLinks []ExternalLinks `json:"externalLinks"`
	// Data and links to legal streaming episodes on external sites
	StreamingEpisodes []StreamingEpisodes `json:"streamingEpisodes"`
	// The ranking of the media in a particular time span and format compared to other media
	Rankings []Rankings `json:"rankings"`
	// User reviews of the media
	Reviews Reviews `json:"reviews"`
	// User recommendations for similar media
	Recommendations Recommendations `json:"recommendations"`
	Stats           Stats           `json:"stats"`
	// The url for the media page on the AniList website
	SiteURL string `json:"siteUrl"`
	// If the media should have forum thread automatically created for it on airing episode release
	AutoCreateForumThread bool `json:"autoCreateForumThread"`
	// If the media is blocked from being recommended to/from
	IsRecommendationBlocked bool `json:"isRecommendationBlocked"`
	// Notes for site moderators
	ModNotes string `json:"modNotes"`
}

// Title object that has Romaji, English, Native and UserPreffered if they are given.
type Title struct {
	Romaji        string `json:"romaji"`
	English       string `json:"english"`
	Native        string `json:"native"`
	UserPreferred string `json:"userPreferred"`
}

// StartDate object that have year month and day of the anime
type StartDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

// EndDate object to store json and endDate of the anime
type EndDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

// Trailer contains YouTube Unique code to the video.
type Trailer struct {
	ID string `json:"id"`
}

// CoverImage of the anime ExtraLarge, Large, Medium, and Color for some reason.
type CoverImage struct {
	ExtraLarge string `json:"extraLarge"`
	Large      string `json:"large"`
	Medium     string `json:"medium"`
	Color      string `json:"color"`
}

// Tags object contais the ID for tags
type Tags struct {
	ID int `json:"id"`
}

// Relations Object I believe relations to the anime/manga
type Relations struct {
	Edges []Edges `json:"edges"`
	Nodes []Node  `json:"nodes"`
}

// Edges some data have edges so need a different one.
type Edges struct {
	ID int `json:"id"`
}

// Characters Characters ID must use Edges to call it and it's an array.
type Characters struct {
	Edges []Edges `json:"edges"`
}

// Studio studios that worked on the anime/manga (anime mostly)
type Studio struct {
	Edges []Edges `json:"edges"`
}

// NextAiringEpisode contains the ID of the anime episode?
type NextAiringEpisode struct {
	ID int `json:"id"`
}

// AiringSchedule Anime airing schedule if it's provided uses Edges and ID
type AiringSchedule struct {
	Edges []Edges `json:"edges"`
}

// ExternalLinks for the anime, uses ID
type ExternalLinks struct {
	ID int `json:"id"`
}

// StreamingEpisodes returns the Episode title, Thumbnail, URL to the streaming site, and the Site
type StreamingEpisodes struct {
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	URL       string `json:"url"`
	Site      string `json:"site"`
}

// Rankings of the Anime ID
type Rankings struct {
	ID int `json:"id"`
}

// Reviews If the authenticated user have made review for it I believe.
type Reviews struct {
	Edges []RecroRevEdges `json:"edges"`
}

// RecroRevEdges (temporaty name type find a better name for different edges :( )) Node Object nested from Media
type RecroRevEdges struct {
	Node Node `json:"node"`
}

// Recommendations of the similar anime.
type Recommendations struct {
	Edges []RecroRevEdges `json:"edges"`
	Node  Node            `json:"node"`
}

// Node that contains the ID
type Node struct {
	ID int `json:"id"`
}

// Stats of the current anime array of Score such as 10, 20, 30, 100 tells you how many people scored it such and Status such as PLANNING, WATCHING, DROPPED
type Stats struct {
	ScoreDistribution  []ScoreDistribution  `json:"scoreDistribution"`
	StatusDistribution []StatusDistribution `json:"statusDistribution"`
}

// ScoreDistribution individual item to access it use array such as ScoreDistribution[0] to get the first array.
type ScoreDistribution struct {
	Score  int `json:"score"`
	Amount int `json:"amount"`
}

// StatusDistribution individual item to access it use array such as StatusDistribution[0] to get the first array.
type StatusDistribution struct {
	Status string `json:"status"`
	Amount int    `json:"amount"`
}

// url constant of Anilist Api
const url = "https://graphql.anilist.co"

// MediaQuery Query Constant what will be posted to the api and send as post request
const MediaQuery = `id,
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
				tags {
				  id,
				},
				relations {
				  edges {
					id,
				  },
				  nodes {
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
				isFavourite,
				isAdult,
				nextAiringEpisode {
				  id,
				},
				airingSchedule {
				  edges {
					id,
				  },
				},
				externalLinks {
				  id,
				},
				streamingEpisodes {
				  title,
				  thumbnail,
				  url,
				  site,
				},
				rankings {
				  id,
				},
				reviews {
				  edges {
					node {
					  id,
					},
				  },
				},
				recommendations {
				  edges {
					node {
					  id,
					},
				  },
				},
				stats {
					scoreDistribution {
					  score,
					  amount,
					},
					statusDistribution {
					  status,
					  amount,
					},
				  }
				  siteUrl,
				  autoCreateForumThread,
				  isRecommendationBlocked,
				  modNotes,`

// NewMediaQuery creates Media objects
func NewMediaQuery() *Media {
	m := Media{}

	return &m
}

// FilterByID Search Anilist Media by it's ID
func (m *Media) FilterByID(id int) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Media(id: %v) {
				%s
			  }
		}
	`, id, MediaQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := cleanMediaJSON(request)
	if err := json.Unmarshal(cleanData, &m); err != nil {
		return false, err
	}

	return true, nil
}

// FilterByMalID search Anilist Media by it's MAL(MyAnimeList) ID
func (m *Media) FilterByMalID(malID int) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Media(idMal: %v) {
				%s
			  }
		}
	`, malID, MediaQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := cleanMediaJSON(request)
	if err := json.Unmarshal(cleanData, &m); err != nil {
		return false, err
	}

	return true, nil
}

// FilterByStartDate search Anilist Media by start Date of the show 8 digit 2013-04-08 == 20130408
func (m *Media) FilterByStartDate(date int32) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Media(startDate: %v) {
				%s
			  }
		}
	`, date, MediaQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := cleanMediaJSON(request)
	if err := json.Unmarshal(cleanData, &m); err != nil {
		return false, err
	}

	return true, nil
}

// FilterByEndDate search Anilist Media by start Date of the show 8 digit 2013-04-08 == 20130408
func (m *Media) FilterByEndDate(date int32) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Media(endDate: %v) {
				%s
			  }
		}
	`, date, MediaQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := cleanMediaJSON(request)
	if err := json.Unmarshal(cleanData, &m); err != nil {
		return false, err
	}

	return true, nil
}

// FilterBySeason search Anilist Media by Season (WINTER, SPRING, SUMMER, FALL)
func (m *Media) FilterBySeason(season string) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Media(season: %v) {
				%s
			  }
		}
	`, season, MediaQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := cleanMediaJSON(request)
	if err := json.Unmarshal(cleanData, &m); err != nil {
		return false, err
	}

	return true, nil
}

// FilterByTitle search Anilist Media by title of the anime or manga
func (m *Media) FilterByTitle(title string) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Media(search: "%v" type: ANIME ) {
				%s
			  }
		}
	`, title, MediaQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := cleanMediaJSON(request)
	if err := json.Unmarshal(cleanData, &m); err != nil {
		return false, err
	}

	return true, nil
}

// FilterAnimeByID search Anilist Anime only type: ANIME is hard-coded in the query
func (m *Media) FilterAnimeByID(id int) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Media(id: %v type: ANIME) {
				%s
			  }
		}
	`, id, MediaQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := cleanMediaJSON(request)
	if err := json.Unmarshal(cleanData, &m); err != nil {
		return false, err
	}

	return true, nil
}

// FilterMangaByID search Anilist Manga type: MANGA is hard-coded in the query
func (m *Media) FilterMangaByID(id int) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Media(id: %v type: MANGA) {
				%s
			  }
		}
	`, id, MediaQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := cleanMediaJSON(request)
	if err := json.Unmarshal(cleanData, &m); err != nil {
		return false, err
	}

	return true, nil
}
