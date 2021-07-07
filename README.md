# AniList Go Wrapper
[![Go Reference](https://pkg.go.dev/badge/github.com/kaiserbh/anilistgo.svg)](https://pkg.go.dev/github.com/kaiserbh/anilistgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/kaiserbh/anilistgo)](https://goreportcard.com/report/github.com/kaiserbh/anilistgo)


Most of the functionality is good to *go* for querying media from Anilist and Mutation such as adding media to the list or removing it. However there are still some left such as Threads, Reviews.

Anilist is an online manga/anime database and social networking service
More information on Anilist can be found here : [anilist](https://anilist.co/home)

# INSTALLATION

```go
go get github.com/kaiserbh/anilistgo
```

# Examples

## Query

### Character Query

```go
import "github.com/kaiserbh/anilistgo"
c := anilistgo.NewCharacterQuery()

ok, err := c.FilterCharacterByName("Nezuko")
if err != nil{
	panic(err)
}

// ok is optional basically a boolean for further use cases.
if ok{ 
	fmt.Printf("%+v\n", c)
}
	
// can do this as well without ok variable.
import "github.com/kaiserbh/anilistgo"
c := anilistgo.NewCharacterQuery()

_, err := c.FilterCharacterByName("Nezuko-channnnnn")
if err != nil{
	panic(err)
}

fmt.Printf("%+v\n", c)
```

### User Query

```go
import "github.com/kaiserbh/anilistgo"
u := anilistgo.NewUserQuery()

ok, err : = u.FilterUserByID(197690)
if err != nil{
	panic(err)
}

// ok is optional basically a boolean for further use cases.
if ok{
fmt.Printf("%+v\n", u)
}

// can do this as well without ok variable.
import "github.com/kaiserbh/anilistgo"
u := NewUserQuery

_, err : = u.FilterUserByID(197690)
if err != nil{
	panic(err)
}

fmt.Printf("%+v\n", u)
```

### Media Query

```go
import "github.com/kaiserbh/anilistgo"
m := NewMediaQuery()

ok, err := m.FilterByID(1)
if err != nil{
	panic(err)
}

// ok is optional basically a boolean for further use cases.
if ok{
fmt.Printf("%+v\n", m)
}

// can do this as well without ok variable.
import "github.com/kaiserbh/anilistgo"
m := NewMediaQuery()

_, err := m.FilterByID(1)
if err != nil{
	panic(err)
}

fmt.Printf("%+v\n", u)
  
```

### Page Query

```go
import "github.com/kaiserbh/anilistgo"
p := NewPageQuery()

ok, err := p.PaginationByTitle("Attack on Titan", 1, 10) // takes title, page number, and per page amount
if err != nil{
	panic(err)
}

// ok is optional basically a boolean for further use cases.
if ok{
fmt.Printf("%+v\n", p)
}
```

### Mediatrend

```go
import "github.com/kaiserbh/anilistgo"
mt := NewMediaTrendQuery()
ok,err := mt.SearchByMediaID(1)
if err != nil{
	panic(err)
}
if ok{
    fmt.Printf("%+v\n", mt)
}
  
```

### Staff query

```go
import "github.com/kaiserbh/anilistgo"
s := NewStaffQuery()

ok, err := s.FilterStaffByName("Hideaki Sorachi")
if err  != nil {
	panic(err)
}

if ok{
	fmt.Printf("%+v\n", s)
}
```

## Mutation

### UserMediaList requires Auth/Token

```go
import  "github.com/kaiserbh/anilistgo"
authKey := "long ass string"
ok, err := anilistgo.SaveMediaListEntry(1, "CURRENT", 0, authKey)
if err != nil{
	panic(err)
}
if ok{
	log.Info("succesfully updated uwu")
}
```

### updateUser requires Auth/Token

```go
import "github.com/kaiserbh/anilistgo"
authKey := "long ass string"

ok, err := anilist.UpdateUserAbout("uwu I love anime", authKey)
if err != nil{
    panic(err)
}
if ok{
    log.Info("succesfully updated uwu")
}
```

# Notes

You will need auth token for some of the methods to work to get your own token feel free to use https://anilist.co/api/v2/oauth/authorize?client_id=4684&response_type=token

# Docs
[Documentation](https://github.com/KaiserBh/AniListGo/blob/master/docs/index.md)
