package anilistgo

// MediaTrendConnection holds Edges, Nodes, and PageInfo
type MediaTrendConnection struct {
	Edges    []MediaTrendEdge `json:"edges"`
	Nodes    []MediaTrend     `json:"nodes"`
	PageInfo PageInfo         `json:"pageInfo"`
}

// MediaTrendEdge trend connection edge
type MediaTrendEdge struct {
	Node MediaTrend `json:"node"`
}
