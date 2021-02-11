package mutation

import (
	"encoding/json"
	"fmt"

	anilist "github.com/kaiserbh/anilistgo/anilist/query"
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
	        SaveMediaListEntry( mediaId:%v, status:%s, progress:%d){
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

// SaveMediaListEntryWithScore Create or update a media list entry
/*
score: Float
The score of the media in the user's chosen scoring method
*/
func SaveMediaListEntryWithScore(mediaID int, status string, progress int, score float64, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        SaveMediaListEntry( mediaId:%v, status:%s, progress:%d, score:%f){
				id,
				mediaId,
				status,
				progress,
				score,
			},
		}`, mediaID, status, progress, score),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

// SaveMediaListEntryWithScoreRaw Create or update a media list entry
/*
scoreRaw: Int
The score of the media in 100 point
*/
func SaveMediaListEntryWithScoreRaw(mediaID int, status string, progress int, score int, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        SaveMediaListEntry( mediaId:%v, status:%s, progress:%d, score:%d){
				id,
				mediaId,
				status,
				progress,
				score,
			},
		}`, mediaID, status, progress, score),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

// SaveMediaListEntryProgressVolume  Create or update a media list entry
/*
progressVolumes: Int
The amount of volumes read by the user
*/
func SaveMediaListEntryProgressVolume(mediaID int, status string, progressVolumes int, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        SaveMediaListEntry( mediaId:%v, status:%s, progressVolumes:%d){
				id,
				mediaId,
				status,
				progress,
				progressVolumes,
			},
		}`, mediaID, status, progressVolumes),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

// SaveMediaListEntryWithRepeat Create or update a media list entry
/*
repeat: Int
The amount of times the user has rewatched/read the media
*/
func SaveMediaListEntryWithRepeat(mediaID int, repeat int, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        SaveMediaListEntry( mediaId:%v, repeat:%d){
				id,
				mediaId,
				status,
				progress,
				repeat,
			},
		}`, mediaID, repeat),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

// SaveMediaListEntryPrivate Create or update a media list entry
/*
private: Boolean
If the entry should only be visible to authenticated user

*/
func SaveMediaListEntryPrivate(mediaID int, private bool, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        SaveMediaListEntry( mediaId:%v, private:%v){
				id,
				mediaId,
				status,
				progress,
				repeat,
				private,
			},
		}`, mediaID, private),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

// SaveMediaListEntryAddNotes Create or update a media list entry
/*
notes: String
Text notes add notes to the media
*/
func SaveMediaListEntryAddNotes(mediaID int, notes string, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        SaveMediaListEntry( mediaId:%v, notes:"%v"){
				id,
				mediaId,
				status,
				progress,
				repeat,
				private,
				notes,
			},
		}`, mediaID, notes),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

// DeleteMediaListEntry Delete a media list entry
// just provide media ID it will then be converted to the mediaList ID
/*
id: Int
The id of the media list entry to delete
*/
func DeleteMediaListEntry(ID int, authToken string) {
	u := anilist.MediaListEntry{}
	u.GetUserMediaList(ID, authToken)

	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        DeleteMediaListEntry( id:%v){
				deleted,
			},
		}`, u.ID),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}
