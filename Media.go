package anilistgo

import (
	"encoding/json"
	"fmt"
)

// Media Anime or Manga
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
	StartDate FuzzyDate `json:"startDate"`
	// EndDate The last official release date of the media
	EndDate FuzzyDate `json:"endDate"`
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
	Tags []MediaTag `json:"tags"`
	// Other media in the same or connecting franchise
	Relations MediaConnection `json:"relations"`
	// The characters in the media
	Characters CharacterConnection `json:"characters"`
	// The staff who produced the media
	Staff StaffConnection `json:"staff"`
	// The companies who produced the media
	Studios StudioConnection `json:"studios"`
	// If the media is marked as favourite by the current authenticated user
	IsFavourite bool `json:"isFavourite"`
	// If the media is intended only for 18+ adult audiences
	IsAdult bool `json:"isAdult"`
	// The media's next episode airing schedule
	NextAiringEpisode AiringSchedule `json:"nextAiringEpisode"`
	// The media's entire airing schedule
	AiringSchedule AiringScheduleConnection `json:"airingSchedule"`
	// Trends The media's daily trend stats
	Trends MediaTrendConnection `json:"trends"`
	// External links to another site related to the media
	ExternalLinks []MediaExternalLinks `json:"externalLinks"`
	// Data and links to legal streaming episodes on external sites
	StreamingEpisodes []MediaStreamingEpisode `json:"streamingEpisodes"`
	// The ranking of the media in a particular time span and format compared to other media
	Rankings []MediaRank `json:"rankings"`
	// User reviews of the media
	Reviews Review `json:"reviews"`
	// User recommendations for similar media
	Recommendations RecommendationConnection `json:"recommendations"`
	Stats           MediaStats               `json:"stats"`
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

// Trailer contains YouTube Unique code to the video.
type Trailer struct {
	ID        string `json:"id"`
	Site      string `json:"site"`
	Thumbnail string `json:"thumbnail"`
}

// CoverImage of the anime ExtraLarge, Large, Medium, and Color for some reason.
type CoverImage struct {
	ExtraLarge string `json:"extraLarge"`
	Large      string `json:"large"`
	Medium     string `json:"medium"`
	Color      string `json:"color"`
}

// MediaTag A tag that describes a theme or element of the media
type MediaTag struct {
	// ID The id of the tag
	ID int `json:"id"`
	// Name The name of the tag
	Name string `json:"name"`
	// Description A general description of the tag
	Description string `json:"description"`
	// Category The categories of tags this tag belongs to
	Category string `json:"category"`
	// Rank The relevance ranking of the tag out of the 100 for this media
	Rank int `json:"rank"`
	// IsGeneralSpoiler If the tag could be a spoiler for any media
	IsGeneralSpoiler bool `json:"isGeneralSpoiler"`
	// IsMediaSpoiler If the tag is a spoiler for this media
	IsMediaSpoiler bool `json:"isMediaSpoiler"`
	// isAdult If the tag is only for adult 18+ media
	IsAdult bool `json:"isAdult"`
}

// MediaExternalLinks An external link to another site related to the media
type MediaExternalLinks struct {
	ID   int    `json:"id"`
	Url  string `json:"url"`
	Site string `json:"site"`
}

// MediaStreamingEpisode Data and links to legal streaming episodes on external sites
type MediaStreamingEpisode struct {
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	URL       string `json:"url"`
	Site      string `json:"site"`
}

// MediaRank The ranking of a media in a particular time span and format compared to other media
type MediaRank struct {
	//ID The id of the rank
	ID int `json:"id"`
	// Rank The numerical rank of the media
	Rank int `json:"rank"`
	// Type The type of ranking [RATED, POPULAR]
	Type string `json:"type"`
	// Format The format the media is ranked within [TV, TV_SHORT, MOVIE, SPECIAL, OVA, ONA, MUSIC, MANGA, NOVEL, ONE_SHOT]
	Format string `json:"format"`
	// Year The year the media is ranked within
	Year int `json:"year"`
	//Season The season the media is ranked within
	Season string `json:"season"`
	// AllTime If the ranking is based on all time instead of a season/year
	AllTime bool `json:"allTime"`
	// Context String that gives context to the ranking type and time span
	Context string `json:"context"`
}

// MediaStats of the current anime array of Score such as 10, 20, 30, 100 tells you how many people scored it such and Status such as PLANNING, WATCHING, DROPPED
type MediaStats struct {
	ScoreDistribution  []ScoreDistribution  `json:"scoreDistribution"`
	StatusDistribution []StatusDistribution `json:"statusDistribution"`
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
      site,
      thumbnail,
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
      name,
      description,
      category,
      rank,
      isGeneralSpoiler,
      isMediaSpoiler,
      isAdult,
    },
    relations {
      edges {
        id,
        relationType,
        node {
          id,
          idMal,
          title {
            romaji,
            english,
            native,
            userPreferred,
          }
        }
        node {
          id,
          idMal,
          title {
            romaji,
            english,
            native,
            userPreferred,
          },
        },
      },
      pageInfo {
        total,
        perPage,
        currentPage,
        lastPage,
        hasNextPage,
      },
    },
    characters {
      edges {
        id,
        role,
        name,
        voiceActors {
          id,
          name {
            first,
            middle,
            last,
            full,
            native,
            userPreferred,
          },
          languageV2,
          image {
            large,
            medium,
          },
          description,
          primaryOccupations,
          gender,
          dateOfBirth {
            year,
            month,
            day,
          },
          dateOfDeath {
            year,
            month,
            day,
          },
          age,
          yearsActive,
          homeTown,
          isFavourite,
          isFavouriteBlocked,
          siteUrl,
          submitter {
            id,
            name,
          },
          submissionStatus,
          submissionNotes,
          favourites,
          modNotes,
        },
        voiceActorRoles {
          voiceActor {
            id,
            name {
              first,
              middle,
              last,
              full,
              native,
              userPreferred,
            },
          },
          roleNotes,
          dubGroup,
        },
        node {
          id,
          name {
            first,
            middle,
            last,
            full,
            native,
            userPreferred,
          },
          image {
            large,
            medium,
          },
          description,
          gender,
          dateOfBirth {
            year,
            month,
            day,
          },
          age,
          isFavourite,
          isFavouriteBlocked,
          siteUrl,
          favourites,
          modNotes,
        },
      },
      nodes {
        id,
        name {
          first,
          middle,
          last,
          full,
          native,
          userPreferred,
        },
        image {
          large,
          medium,
        },
        description,
        gender,
        dateOfBirth {
          year,
          month,
          day,
        },
        age,
        isFavourite,
        isFavouriteBlocked,
        siteUrl,
        favourites,
        modNotes,
      },
      pageInfo {
        total,
        perPage,
        currentPage,
        lastPage,
        hasNextPage,
      },
    },
    staff {
      edges {
        id,
        role,
        favouriteOrder,
        node {
          id,
          name {
            first,
            middle,
            last,
            full,
            native,
            userPreferred,
          },
          languageV2,
          image {
            large,
            medium,
          },
          description,
          primaryOccupations,
          gender,
          dateOfBirth {
            year,
            month,
            day,
          }
          dateOfDeath {
            year,
            month,
            day,
          },
          age,
          yearsActive,
          homeTown,
          isFavourite,
          isFavouriteBlocked,
          siteUrl,
          submitter {
            id,
          }
          submissionNotes,
          favourites,
          modNotes,
        },
      },
      nodes {
        id,
        name {
          first,
          middle,
          last,
          full,
          native,
          userPreferred,
        },
        languageV2,
        image {
          large,
          medium,
        },
        description,
        primaryOccupations,
        gender,
        dateOfBirth {
          year,
          month,
          day,
        },
        dateOfDeath {
          year,
          month,
          day,
        },
        age,
        yearsActive,
        homeTown,
        isFavourite,
        isFavouriteBlocked,
        siteUrl,
        submitter {
          id,
        }
        submissionNotes,
        favourites,
        modNotes,
      },
      pageInfo {
        total,
        perPage,
        currentPage,
        lastPage,
        hasNextPage,
      },
    },
    studios {
      edges {
        id,
        isMain,
        favouriteOrder,
        node {
          id,
          name,
          isAnimationStudio,
          siteUrl,
          isFavourite,
          favourites,
        },
      },
      nodes {
        id,
        name,
        isAnimationStudio,
        siteUrl,
        isFavourite,
        favourites,
      },
      pageInfo {
        total,
        perPage,
        currentPage,
        lastPage,
        hasNextPage,
      },
    },
    isFavourite,
    isAdult,
    nextAiringEpisode {
      id,
      airingAt,
      timeUntilAiring,
      episode,
      media {
        id,
      },
    },
    airingSchedule {
      edges {
        id,
        node {
          id,
          airingAt,
          timeUntilAiring,
          episode,
          mediaId,
        },
      },
      nodes {
        id,
        airingAt,
        timeUntilAiring,
        episode,
        mediaId,
      },
      pageInfo {
        total,
        perPage,
        currentPage,
        lastPage,
        hasNextPage,
      },
    },
    trends {
      edges {
        node {
          mediaId,
          date,
          trending,
          averageScore,
          popularity,
          inProgress,
          releasing,
          episode,
        },
      },
      nodes {
        mediaId,
        date,
        trending,
        averageScore,
        popularity,
        inProgress,
        releasing,
        episode,
      },
      pageInfo {
        total,
        perPage,
        currentPage,
        lastPage,
        hasNextPage,
      },
    },
    externalLinks {
      id,
      url,
      site,
    },
    streamingEpisodes {
      title,
      thumbnail,
      url,
      site,
    },
    rankings {
      id,
      rank,
      type,
      format,
      year,
      season,
      allTime,
      context,
    },
    reviews {
      edges {
        node {
          id,
          userId,
          mediaId,
          mediaType,
          summary,
          body,
          rating,
          ratingAmount,
          userRating,
          score,
          private,
          siteUrl,
          createdAt,
          updatedAt,
          user {
            id,
          },
        },
      },
      nodes {
        user {
          id,
          name,
          about,
          avatar {
            large,
            medium,
          },
          bannerImage,
          siteUrl,
        },
        id,
        userId,
        mediaId,
        mediaType,
        summary,
        body,
        rating,
        ratingAmount,
        userRating,
        score,
        private,
        siteUrl,
        createdAt,
        updatedAt,
        user {
          id,
          name,
          about,
          avatar {
            large,
            medium,
          },
          bannerImage,
          siteUrl,
        },
      },
      pageInfo {
        total,
        perPage,
        currentPage,
        lastPage,
        hasNextPage,
      },
    },
    recommendations {
      edges {
        node {
          id,
          rating,
          userRating,
          media {
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
          },
          user {
            id,
            name,
            about,
            avatar {
              large,
              medium,
            },
            bannerImage,
          },
        },
      },
      nodes {
        id,
        rating,
        userRating,
        media {
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
        },
        user {
          id,
          name,
          about,
          avatar {
            large,
            medium,
          },
          bannerImage,
        },
      },
      pageInfo {
        total,
        perPage,
        currentPage,
        lastPage,
        hasNextPage,
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
    },
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

// FilterMediaByTitle search Anilist Media by title of the anime or manga
func (m *Media) FilterMediaByTitle(title string) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Media(search: "%v") {
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

// FilterAnimeByTitle search Anilist Media by title of the anime or manga
func (m *Media) FilterAnimeByTitle(title string) (bool, error) {
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

// FilterMangaByTitle search Anilist Media by title of the manga
func (m *Media) FilterMangaByTitle(title string) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Media(search: "%v" type: MANGA ) {
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
