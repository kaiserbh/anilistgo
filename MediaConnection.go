package anilistgo

type MediaConnection struct {
	Edges    []MediaEdge `json:"edges"`
	Nodes    []Media     `json:"nodes"`
	PageInfo PageInfo    `json:"pageInfo"`
}

type MediaEdge struct {
	Node Media `json:"node"`
	// Id The id of the connection
	Id int `json:"id"`
	// RelationType The type of relation to the parent model
	RelationType string `json:"relationType"`
	// IsMainStudio If the studio is the main animation studio of the media (For Studio->MediaConnection field only)
	IsMainStudio bool `json:"isMainStudio"`
	// Characters The characters in the media voiced by the parent actor
	Characters []Character `json:"characters"`
	// CharacterRole The characters role in the media
	CharacterRole string `json:"characterRole"`
	// CharacterName Media specific character name
	CharacterName string `json:"characterName"`
	// RoleNotes Notes regarding the VA's role for the character
	RoleNotes string `json:"roleNotes"`
	// DubGroup Used for grouping roles where multiple dubs exist for the same language. Either dubbing company name or language variant.
	DubGroup string `json:"dubGroup"`
	// StaffRole The role of the staff member in the production of the media
	StaffRole string `json:"staffRole"`

	VoiceActors []Staff `json:"voiceActors"`
}
