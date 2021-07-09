package anilistgo

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
