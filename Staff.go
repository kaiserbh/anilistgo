package anilistgo

import (
	"encoding/json"
	"fmt"
)

// Staff Voice actors or production staff
type Staff struct {
	// ID The id of the staff member
	ID int `json:"id"`
	// Name The names of the staff member
	Name StaffName `json:"name"`
	// LanguageV2 The primary language of the staff member. Current values: Japanese, English,
	// Korean, Italian, Spanish, Portuguese, French, German, Hebrew, Hungarian, Chinese, Arabic, Filipino, Catalan
	LanguageV2 string `json:"languageV2"`
	// Image The staff images
	Image StaffImage `json:"image"`
	// Description A general description of the staff member
	Description string `json:"description"`
	// PrimaryOccupations The person's primary occupations
	PrimaryOccupations []string `json:"primaryOccupations:"`
	// Gender The staff's gender. Usually Male, Female, or Non-binary but can be any string.
	Gender      string    `json:"gender"`
	DateOfBirth FuzzyDate `json:"dateOfBirth"`
	DateOfDeath FuzzyDate `json:"dateOfDeath"`
	// Age The person's age in years
	Age int `json:"age"`
	// YearsActive [startYear, endYear] (If the 2nd value is not present staff is still active)
	YearsActive []int `json:"yearsActive"`
	// HomeTown The persons birthplace or hometown
	HomeTown string `json:"homeTown"`
	// IsFavourite If the staff member is marked as favourite by the currently authenticated user
	IsFavourite bool `json:"isFavourite"`
	// IsFavouriteBlocked If the staff member is blocked from being added to favourites
	IsFavouriteBlocked bool `json:"isFavouriteBlocked"`
	// SiteURL The url for the staff page on the AniList website
	SiteURL string `json:"siteUrl"`
	// StaffMedia Media where the staff member has a production role
	StaffMedia MediaConnection `json:"staffMedia"`
	// Characters Characters voiced by the actor
	Characters CharacterConnection `json:"characters"`
	// CharacterMedia Media the actor voiced characters in. (Same data as characters with media as node instead of characters)
	CharacterMedia MediaConnection `json:"characterMedia"`
	// Staff Staff member that the submission is referencing
	Staff string `json:"staff"`
	// Submitter Submitter for the submission
	Submitter User `json:"submitter"`
	// SubmissionStatus
	SubmissionStatus int `json:"submissionStatus"`
	// Favourites
	Favourites int `json:"favourites"`
	// ModNotes
	ModNotes string `json:"modNotes"`
}

// StaffName The names of the staff member
type StaffName struct {
	// First The person's given name
	First string `json:"first"`
	// Middle The person's middle name
	Middle string `json:"middle"`
	// Last The person's surname
	Last string `json:"last"`
	// Full The person's first and last name
	Full string `json:"full"`
	// Native The person's full name in their native language
	Native string `json:"native"`
	// Alternative Other names the staff member might be referred to as (pen names)
	Alternative []string `json:"alternative"`
	// UserPreferred The currently authenticated users preferred name language. Default romaji for non-authenticated
	UserPreferred string `json:"userPreferred"`
}

// StaffImage of the staff Large or Medium
type StaffImage struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

// StaffQuery graphql constant for StaffQuery
const StaffQuery = `id
    name {
      first
      middle
      last
      full
      native
      userPreferred
    }
    languageV2
    image {
      large
      medium
    }
    description
    gender
    dateOfBirth {
      year
      month
      day
    }
    dateOfDeath {
      year
      month
      day
    }
    age
    yearsActive
    homeTown
    isFavourite
    isFavouriteBlocked
    siteUrl
    staffMedia {
      edges {
        id
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
    characters {
      edges {
        id
        role
        node {
          id
          name {
            first
            middle
            last
            full
            native
            userPreferred
          }
          image {
            large
            medium
          }
          description
          gender
          dateOfBirth {
            year
            month
            day
          }
          age
          isFavourite
          isFavouriteBlocked
          siteUrl
        }
      }
      nodes {
        id
        name {
          first
          middle
          last
          full
          native
          userPreferred
        }
        image {
          large
          medium
        }
        description
        gender
        dateOfBirth {
          year
          month
          day
        }
        age
        isFavourite
        isFavouriteBlocked
        siteUrl
      }
      pageInfo {
        total
        perPage
        currentPage
        lastPage
        hasNextPage
      }
    }
    submitter {
      id
      name
    }
    submissionStatus
    submissionNotes
    favourites`

// NewStaffQuery Create new Staff Object
func NewStaffQuery() *Staff {
	s := Staff{}
	return &s
}

// FilterStaffByID Search Anilist Staff by it's ID
func (s *Staff) FilterStaffByID(id int) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Staff(id: %v) {
				%s
			  }
		}
	`, id, StaffQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := CleanStaffJSON(request)
	if err := json.Unmarshal(cleanData, &s); err != nil {
		return false, err
	}

	return true, nil
}

// FilterStaffByName Search Anilist Staff by it's ID
func (s *Staff) FilterStaffByName(name string) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			Staff(search: "%v") {
				%s
			  }
		}
	`, name, StaffQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := CleanStaffJSON(request)
	if err := json.Unmarshal(cleanData, &s); err != nil {
		return false, err
	}

	return true, nil
}
