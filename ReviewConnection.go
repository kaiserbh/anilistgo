package anilistgo

// ReviewConnection holds Edges, Nodes, and PageInfo of ReviewConnection
type ReviewConnection struct {
	Edges    []ReviewEdge `json:"edges"`
	Nodes    []Review     `json:"nodes"`
	PageInfo PageInfo     `json:"pageInfo"`
}

// ReviewEdge Review connection edge
type ReviewEdge struct {
	Node []Review `json:"node"`
}
