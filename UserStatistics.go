package anilistgo

type UserStatistics struct {
	Count             int                        `json:"count"`
	MeanScore         float64                    `json:"meanScore"`
	StandardDeviation float64                    `json:"standardDeviation"`
	MinutesWatched    int                        `json:"minutesWatched"`
	EpisodeWatched    int                        `json:"episodeWatched"`
	ChaptersRead      int                        `json:"chaptersRead"`
	VolumesRead       int                        `json:"volumesRead"`
	Formats           []UserFormatStatistic      `json:"formats"`
	Statuses          []UserStatusStatistic      `json:"statuses"`
	Scores            []UserScoreStatistic       `json:"scores"`
	Lengths           []UserLengthStatistic      `json:"lengths"`
	ReleaseYears      []UserReleaseYearStatistic `json:"releaseYears"`
	StartYears        []UserStartYearStatistic   `json:"startYears"`
	Genres            []UserGenreStatistic       `json:"genres"`
	Tags              []UserTagStatistic         `json:"tags"`
	Countries         []UserCountryStatistic     `json:"countries"`
	VoiceActors       []UserVoiceActorStatistic  `json:"voiceActors"`
	Staff             []UserStaffStatistic       `json:"staff"`
	Studios           []UserStudioStatistic      `json:"studios"`
}

type UserFormatStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Format         string  `json:"format"`
}

type UserStatusStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Status         string  `json:"status"`
}

type UserScoreStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Score          int     `json:"score"`
}

type UserLengthStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Length         string  `json:"length"`
}

type UserReleaseYearStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	ReleaseYear    int     `json:"releaseYear"`
}

type UserStartYearStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	StartYear      int     `json:"startYear"`
}

type UserGenreStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Genre          string  `json:"genre"`
}

type UserTagStatistic struct {
	Count          int      `json:"count"`
	MeanScore      float64  `json:"meanScore"`
	MinutesWatched int      `json:"minutesWatched"`
	ChaptersRead   int      `json:"chaptersRead"`
	MediaIds       []int    `json:"mediaIds"`
	Tag            MediaTag `json:"tag"`
}

type UserCountryStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Country        string  `json:"country"`
}

type UserVoiceActorStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	VoiceActor     Staff   `json:"voiceActor"`
	CharacterIds   []int   `json:"characterIds"`
}

type UserStaffStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Staff          Staff   `json:"staff"`
}

type UserStudioStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Studio         Studio  `json:"studio"`
}
