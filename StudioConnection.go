package anilistgo

type StudioConnection struct {
	Edges    []StudioEdge `json:"edges"`
	Nodes    []Studio     `json:"nodes"`
	PageInfo PageInfo     `json:"pageInfo"`
}

type StudioEdge struct {
	Node Studio `json:"node"`
	// ID The id of the connection
	ID int `json:"id"`
	// If the studio is the main animation studio of the anime
	IsMain bool `json:"isMain"`
	// FavouriteOrder The order the character should be displayed from the users favourites
	FavouriteOrder int `json:"favouriteOrder"`
}
