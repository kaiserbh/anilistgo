package query

import (
	"encoding/json"
	"fmt"
)

//MediaListEntry The authenticated user's media list entry for the media
type MediaListEntry struct {
	ID              int64  `json:"id"`
	UserID          int64  `json:"userId"`
	MediaID         int64  `json:"mediaId"`
	Status          string `json:"status"`
	Score           int64  `json:"score"`
	Progress        int64  `json:"progress"`
	ProgressVolumes int64  `json:"progressVolumes"`
	Media           Media  `json:"media"`
	User            User   `json:"user"`
}

// NewMediaListQuery Create new MediaListEntry Object
func NewMediaListQuery() *MediaListEntry {
	m := MediaListEntry{}
	return &m
}

// GetUserMediaList will return the authenticated user media list
func (m *MediaListEntry) GetUserMediaList(ID int, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`
		{
			Media(id:%d){
	        	mediaListEntry{
					id,
					userId,
					mediaId,
					status,
					score,
					progress,
					progressVolumes,
					media {
						id,
					}
					user {
						id,
					},
				},
			},
		}`, ID),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	cleanData := CleanUserListJSON(PostRequestAuth(jsonValue, authToken))

	if err := json.Unmarshal(cleanData, &m); err != nil {
		panic(err)
	}
}
