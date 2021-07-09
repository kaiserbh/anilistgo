package anilistgo

type MediaTrendConnection struct {
	Edges    []MediaTrendEdge `json:"edges"`
	Nodes    []MediaTrend     `json:"nodes"`
	PageInfo PageInfo         `json:"pageInfo"`
}

type MediaTrendEdge struct {
	Node MediaTrend `json:"node"`
}
