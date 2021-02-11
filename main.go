package main

import (
	"fmt"

	"github.com/kaiserbh/anilistgo/anilist/query"
)

func main() {
	newq := query.NewMedia()
	newq.FilterAnimeByID(21)

	fmt.Println(newq.Title.Romaji)
}
