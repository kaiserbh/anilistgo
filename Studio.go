package anilistgo

import (
	"encoding/json"
	"fmt"
)

type Studio struct {
	// ID The id of the studio
	ID int `json:"id"`
	// Name The name of the studio
	Name string `json:"name"`
	// IsAnimationStudio If the studio is an animation studio or a different kind of company
	IsAnimationStudio bool `json:"isAnimationStudio"`
	// Media The media the studio has worked on
	Media MediaConnection `json:"media"`
	// SiteURL The url for the studio page on the AniList website
	SiteUrl string `json:"siteUrl"`
	// IsFavourite If the studio is marked as favourite by the currently authenticated user
	IsFavourite bool `json:"isFavourite"`
	// Favourites The amount of user's who have favourited the studio
	Favourites int `json:"favourites"`
}

const StudioQuery = `id
    name
    isAnimationStudio
    media {
      edges {
        id
        isMainStudio
        node {
          id
          idMal
          title {
            romaji
            english
            native
            userPreferred
          }
          type
          format
          status
          description
          startDate {
            year
            month
            day
          }
          endDate {
            year
            month
            day
          }
          season
          seasonYear
          seasonInt
          episodes
          duration
          chapters
          volumes
          countryOfOrigin
          isLicensed
          source
          hashtag
          trailer {
            id
            site
            thumbnail
          }
          updatedAt
          coverImage {
            extraLarge
            large
            medium
            color
          }
          bannerImage
          genres
          synonyms
          averageScore
          meanScore
          popularity
          isLocked
          trending
          favourites
          tags {
            id
            name
            description
            category
            rank
            isGeneralSpoiler
            isMediaSpoiler
            isAdult
          }
          isFavourite
          isAdult
          nextAiringEpisode {
            id
            airingAt
            timeUntilAiring
            episode
            mediaId
          }
          externalLinks {
            id
            url
            site
          }
          streamingEpisodes {
            title
            thumbnail
            url
            site
          }
          rankings {
            id
            rank
            type
            format
            year
            season
            allTime
            context
          }
          stats {
            scoreDistribution {
              score
              amount
            }
            statusDistribution {
              status
              amount
            }
          }
          siteUrl
          autoCreateForumThread
          isRecommendationBlocked
          modNotes
        }
      }
      nodes {
        id
        idMal
        title {
          romaji
          english
          native
          userPreferred
        }
        type
        format
        status
        description
        startDate {
          year
          month
          day
        }
        endDate {
          year
          month
          day
        }
        season
        seasonYear
        seasonInt
        episodes
        duration
        chapters
        volumes
        countryOfOrigin
        isLicensed
        source
        hashtag
        trailer {
          id
          site
          thumbnail
        }
        updatedAt
        coverImage {
          extraLarge
          large
          medium
          color
        }
        bannerImage
        genres
        synonyms
        averageScore
        meanScore
        popularity
        isLocked
        trending
        favourites
        tags {
          id
          name
          description
          category
          rank
          isGeneralSpoiler
          isMediaSpoiler
          isAdult
        }
        isFavourite
        isAdult
        nextAiringEpisode {
          id
          airingAt
          timeUntilAiring
          episode
          mediaId
        }
        externalLinks {
          id
          url
          site
        }
        streamingEpisodes {
          title
          thumbnail
          url
          site
        }
        rankings {
          id
          rank
          type
          format
          year
          season
          allTime
          context
        }
        stats {
          scoreDistribution {
            score
            amount
          }
          statusDistribution {
            status
            amount
          }
        }
        siteUrl
        autoCreateForumThread
        isRecommendationBlocked
        modNotes
      }
      pageInfo {
        total
        perPage
        currentPage
        lastPage
        hasNextPage
      }
    }
    siteUrl
    isFavourite
    favourites`

// NewStudioQuery Create new Studio Object
func NewStudioQuery() *Studio {
	s := Studio{}
	return &s
}

// FilterStudioByID Search Anilist Studio by it's ID
func (s *Studio) FilterStudioByID(id int) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Studio(id: %v) {
				%s
			  }
		}
	`, id, StudioQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := CleanStudioJSON(request)
	if err := json.Unmarshal(cleanData, &s); err != nil {
		return false, err
	}

	return true, nil
}

// FilterStudioByName Search Anilist Studio by it's name
func (s *Studio) FilterStudioByName(name string) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Studio(search: "%v") {
				%s
			  }
		}
	`, name, StudioQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}
	cleanData := CleanStudioJSON(request)
	if err := json.Unmarshal(cleanData, &s); err != nil {
		return false, err
	}

	return true, nil
}
