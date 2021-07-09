package anilistgo

// FuzzyDate date type that's being used widely.
type FuzzyDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type StaffRoleType struct {
	// VoiceActor The voice actors of the character
	VoiceActor Staff `json:"voiceActor"`
	// RoleNotes Notes regarding the VA's role for the character
	RoleNotes string `json:"roleNotes"`
	// DubGroup Used for grouping roles where multiple dubs exist for the same language. Either dubbing company name or language variant.
	DubGroup string `json:"dubGroup"`
}

type NotificationOption struct {
	// Type The type of notification
	Type string `json:"type"`
	// Enabled Whether this type of notification is enabled
	Enabled bool `json:"enabled"`
}

type UserStatisticTypes struct {
	Anime UserStatistics `json:"anime"`
	Manga UserStatistics `json:"manga"`
}

// ScoreDistribution individual item to access it use array such as ScoreDistribution[0] to get the first array.
type ScoreDistribution struct {
	Score  int `json:"score"`
	Amount int `json:"amount"`
}

// StatusDistribution individual item to access it use array such as StatusDistribution[0] to get the first array.
type StatusDistribution struct {
	Status string `json:"status"`
	Amount int    `json:"amount"`
}
