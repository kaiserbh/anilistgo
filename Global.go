package anilistgo

// FuzzyDate date type that's being used widely.
type FuzzyDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

// StaffRoleType Voice actor role for a character
type StaffRoleType struct {
	// VoiceActor The voice actors of the character
	VoiceActor Staff `json:"voiceActor"`
	// RoleNotes Notes regarding the VA's role for the character
	RoleNotes string `json:"roleNotes"`
	// DubGroup Used for grouping roles where multiple dubs exist for the same language. Either dubbing company name or language variant.
	DubGroup string `json:"dubGroup"`
}

// NotificationOption Notification option
type NotificationOption struct {
	// Type The type of notification
	Type string `json:"type"`
	// Enabled Whether this type of notification is enabled
	Enabled bool `json:"enabled"`
}

// UserStatisticTypes  holds user statistics anime|manga
type UserStatisticTypes struct {
	Anime UserStatistics `json:"anime"`
	Manga UserStatistics `json:"manga"`
}

// ScoreDistribution A user's list score distribution. It's an array so [0] to get the item.
type ScoreDistribution struct {
	Score  int `json:"score"`
	Amount int `json:"amount"`
}

// StatusDistribution The distribution of the watching/reading status of media or a user's list. It's an array so [0] to get the item.
type StatusDistribution struct {
	Status string `json:"status"`
	Amount int    `json:"amount"`
}
