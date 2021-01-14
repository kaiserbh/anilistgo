# AniList Go Wrapper

Most of the functionality is completed for querying media from Anilist, while mutation still in work so far we can query user, page, staff, and media.
User authentication and getting user list still work in progress.

Anilist is an online manga/anime database and social networking service
More information on Anilist can be found here : [anilist](https://anilist.co/home)

# Examples

## Query

### Character Query

```go
import query "anilistGo/anilist/query"
c := query.Character{}
	c.FilterCharacterByName("Nezuko")
    fmt.Printf("%+v\n", c)
```

### User Query

```go
import query "anilistGo/anilist/query"
u := query.User{}
	u.FilterUserByID(197690)
	fmt.Printf("%+v\n", u)
```

### Media Query

```go
import query "anilistGo/anilist/query"
m := query.Media{}
	m.FilterByID(1)
    fmt.Printf("%+v\n", m)
```

### Page Query

```go
import query "anilistGo/anilist/query"
p := query.Page{}
	p.PaginationByTitle("Attack on Titan", 1, 10) // takes title, page number, and per page amount
    fmt.Printf("%+v\n", p)
```

### Mediatrend

```go
import query "anilistGo/anilist/query"
mt := query.MediaTrend{}
	mt.SearchByMediaID(1)
    fmt.Printf("%+v\n", mt)
```

### Staff query

```go
import query "anilistGo/anilist/query"
s := query.Staff{}
	s.FilterStaffByName("Hideaki Sorachi")
    fmt.Printf("%+v\n", s)
```

## Mutation

### UserMediaList requires Auth/Token

```go
import mutation "anilistGo/anilist/mutation"
authKey := "long ass string"
mutation.SaveMediaListEntry(1, "CURRENT", 0, authKey)
```

### updateUser requires Auth/Token

```go
import mutation "anilistGo/anilist/mutation"
authKey := "long ass string"
mutation.UpdateUserAbout("uwu I love anime", authKey)
```

# Notes

You will need auth token for some of the methods to work to get your own token feel free to use https://anilist.co/api/v2/oauth/authorize?client_id=4684&response_type=token

# Docs

Full
[docs](https://github.com/KaiserBh/AniListGo/blob/master/docs/index.md)
