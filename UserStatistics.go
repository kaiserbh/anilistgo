package anilistgo

// UserStatistics holds information about user stats what they watched or read and scores etc.
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

// UserFormatStatistic holds the stats of type format such as : TV, MOVIE, OVA, ONA etc.
type UserFormatStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Format         string  `json:"format"`
}

// UserStatusStatistic holds the stats of type Status such as : COMPLETED, PLANNING, CURRENT, and PAUSED
type UserStatusStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Status         string  `json:"status"`
}

// UserScoreStatistic holds the stats of type Score such as : 70, 80, 75 basically any scores made and it will return the count meanScore and mediaIds.
type UserScoreStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Score          int     `json:"score"`
}

// UserLengthStatistic holds the stats of type Length such as : 7-16 minutes, 101+, 1, 29-55. minutes or chapter read for manga.
type UserLengthStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Length         string  `json:"length"`
}

// UserReleaseYearStatistic holds the stats of type ReleaseYear such as : 2020, 2019, 2018 returns the score and MediaIds.
type UserReleaseYearStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	ReleaseYear    int     `json:"releaseYear"`
}

// UserStartYearStatistic holds the stats of StartYear  such as : 2020, 2019, 2018 returns the score and MediaIds.
type UserStartYearStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	StartYear      int     `json:"startYear"`
}

// UserGenreStatistic holds the stats of Genre  such as : Action, Comedy, Fantasy, Adventure etc.
type UserGenreStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Genre          string  `json:"genre"`
}

// UserTagStatistic holds the stats of Tag users tags such as; Male Protagonist, Shounen, Super Power etc
type UserTagStatistic struct {
	Count          int      `json:"count"`
	MeanScore      float64  `json:"meanScore"`
	MinutesWatched int      `json:"minutesWatched"`
	ChaptersRead   int      `json:"chaptersRead"`
	MediaIds       []int    `json:"mediaIds"`
	Tag            MediaTag `json:"tag"`
}

// UserCountryStatistic  holds the stats of Country which country the anime/manga user watched from contains: "JP" == Japan, "CN" == China and etc
type UserCountryStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Country        string  `json:"country"`
}

// UserVoiceActorStatistic holds the stats of VoiceActor such as the count or how many anime they have voiced in the user anime stats.
type UserVoiceActorStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	VoiceActor     Staff   `json:"voiceActor"`
	CharacterIds   []int   `json:"characterIds"`
}

// UserStaffStatistic holds the stats of Staff such as the count or how many anime they have worked on etc.
type UserStaffStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Staff          Staff   `json:"staff"`
}

// UserStudioStatistic holds the stats of Studio such as the count or how many anime the user watched of this particular studio and what are their ratings for the anime etc.
type UserStudioStatistic struct {
	Count          int     `json:"count"`
	MeanScore      float64 `json:"meanScore"`
	MinutesWatched int     `json:"minutesWatched"`
	ChaptersRead   int     `json:"chaptersRead"`
	MediaIds       []int   `json:"mediaIds"`
	Studio         Studio  `json:"studio"`
}
