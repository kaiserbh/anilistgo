package anilistgo

type Recommendation struct {
	ID         int   `json:"id"`
	Rating     int   `json:"rating"`
	UserRating int   `json:"userRating"`
	Media      Media `json:"media"`
	User       User  `json:"user"`
}
