package anilistgo

type StaffConnection struct {
	Edges    []StaffEdge `json:"edges"`
	Nodes    []Staff     `json:"nodes"`
	PageInfo PageInfo    `json:"pageInfo"`
}

type StaffEdge struct {
	Node Staff `json:"node"`
	// ID The id of the connection
	ID int `json:"id"`
	// Role The role of the staff member in the production of the media
	Role string `json:"role"`
	// FavouriteOrder The order the staff should be displayed from the users favourites
	FavouriteOrder int `json:"favourite_order"`
}
