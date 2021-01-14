package anilist

import (
	anilist "anilistGo/anilist/query"
	"encoding/json"
	"fmt"
)

//SaveMediaListEntry create or update a media list SaveMediaListEntry
/*
	mediaId: Int
	The id of the media the entry is of

	status: MediaListStatus
	The watching/reading status

	CURRENT
	Currently watching/reading

	PLANNING
	Planning to watch/read

	COMPLETED
	Finished watching/reading

	DROPPED
	Stopped watching/reading before completing

	PAUSED
	Paused watching/reading

	REPEATING
	Re-watching/reading

	progress: Int
	The amount of episodes/chapters consumed by the user




*/
func SaveMediaListEntry(mediaID int, status string, progress int, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        SaveMediaListEntry(mediaId:%v, status:%s, progress:%d){
				id,
				mediaId,
				status,
				progress,
			},
		}`, mediaID, status, progress),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}
