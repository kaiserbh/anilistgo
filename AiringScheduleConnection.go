package anilistgo

// AiringScheduleConnection connection for AiringSchedule
type AiringScheduleConnection struct {
	Edges    []AiringScheduleEdge `json:"edges"`
	Nodes    []AiringSchedule     `json:"nodes"`
	PageInfo PageInfo             `json:"pageInfo"`
}

// AiringScheduleEdge AiringSchedule connection edge
type AiringScheduleEdge struct {
	Node AiringSchedule `json:"node"`
	ID   int            `json:"id"`
}

// AiringSchedule Media Airing Schedule. NOTE: We only aim to guarantee that FUTURE airing data is present and accurate.
type AiringSchedule struct {
	// ID The id of the airing schedule item
	ID int `json:"id"`
	// AiringAt The time the episode airs at
	AiringAt int `json:"airingAt"`
	// TimeUntilAiring Seconds until episode starts airing
	TimeUntilAiring int `json:"timeUntilAiring"`
	// Episode The airing episode number
	Episode int `json:"episode"`
	// MediaId The associate media id of the airing episode
	MediaId int `json:"mediaId"`
}
