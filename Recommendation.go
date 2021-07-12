package anilistgo

// Recommendation Media recommendation
type Recommendation struct {
	ID         int    `json:"id"`
	Rating     int    `json:"rating"`
	UserRating string `json:"userRating"`
	Media      Media  `json:"media"`
	User       User   `json:"user"`
}
