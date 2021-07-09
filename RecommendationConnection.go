package anilistgo

type RecommendationConnection struct {
	Edges    []RecommendationEdge `json:"edges"`
	Nodes    []Recommendation     `json:"nodes"`
	PageInfo PageInfo             `json:"pageInfo"`
}

type RecommendationEdge struct {
	Node Recommendation `json:"node"`
}
