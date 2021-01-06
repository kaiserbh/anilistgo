package main

import (
	"anilist/anilist"
	"fmt"
)

func main() {
	p := anilist.Page{}
	p.PaginationByTitle("To Love Ru", 1, 5)
	fmt.Printf("%+v\n", p.Media[0])
	fmt.Printf("%+v\n", p.Media[1])
}
