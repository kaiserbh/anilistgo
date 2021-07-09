package anilistgo

type ReviewConnection struct {
	Edges    []ReviewEdge `json:"edges"`
	Nodes    []Review     `json:"nodes"`
	PageInfo PageInfo     `json:"pageInfo"`
}

type ReviewEdge struct {
	Node []Review `json:"node"`
}
