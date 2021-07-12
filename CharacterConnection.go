package anilistgo

// CharacterConnection connection for Character
type CharacterConnection struct {
	Edges []CharacterEdge `json:"edges"`
	Nodes []Character     `json:"nodes"`
	// PageInfo The pagination information
	PageInfo PageInfo `json:"pageInfo"`
}

// CharacterEdge Character connection edge
type CharacterEdge struct {
	Node Character `json:"node"`
	// ID The id of the connection
	ID int `json:"id"`
	// Name Media specific character name
	Name string `json:"name"`
	// VoiceActors The voice actors of the character
	VoiceActors []Staff `json:"voiceActors"`
	// VoiceActorsRoles The voice actors of the character with role date
	VoiceActorRoles []StaffRoleType `json:"voiceActorRoles"`
	// Media The media the character is in
	Media []Media `json:"media"`
	// favouriteOrder The order the character should be displayed from the users favourites
	FavouriteOrder int `json:"favouriteOrder"`
}
