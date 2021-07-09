package anilistgo

type Review struct {
	Id           int    `json:"id"`
	UserId       int    `json:"userId"`
	MediaId      int    `json:"mediaId"`
	MediaType    string `json:"mediaType"`
	Summary      string `json:"summary"`
	Body         string `json:"body"`
	Rating       int    `json:"rating"`
	RatingAmount int    `json:"ratingAmount"`
	UserRating   string `json:"userRating"`
	Score        int    `json:"score"`
	Private      bool   `json:"private"`
	SiteUrl      string `json:"siteUrl"`
	CreatedAt    int    `json:"createdAt"`
	UpdatedAt    int    `json:"updatedAt"`
	User         User   `json:"user"`
}
