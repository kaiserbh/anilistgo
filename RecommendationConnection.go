package anilistgo

// RecommendationConnection holds Edges, Nodes, and PageInfo
type RecommendationConnection struct {
	Edges    []RecommendationEdge `json:"edges"`
	Nodes    []Recommendation     `json:"nodes"`
	PageInfo PageInfo             `json:"pageInfo"`
}

// RecommendationEdge Recommendation connection edge
type RecommendationEdge struct {
	Node Recommendation `json:"node"`
}
