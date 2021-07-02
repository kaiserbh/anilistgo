package mutation

import (
	"encoding/json"
	"fmt"

	anilist "github.com/kaiserbh/anilistgo/anilist/query"
)

//UpdateUserAbout User's about/bio text.
func UpdateUserAbout(about string, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(about:"%v"){
				about,
			},
		}`, about),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

//UpdateUserTitleLanguage User's title language; The language the user wants to see media titles in such as
//ROMAJI, ENGLISH, NATIVE, ROMAJI_STYLISED, ENGLISH_STYLISED, NATIVE_STYLISED
func UpdateUserTitleLanguage(titleLanguage string, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(titleLanguage:%v){
				options {
      				titleLanguage,
    			},
			},
		}`, titleLanguage),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

//UpdateUserAdultContent If the user should see media marked as adult-only boolean
func UpdateUserAdultContent(displayAdultContent bool, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(displayAdultContent:%v){
				options {
      				displayAdultContent,
    			},
			},
		}`, displayAdultContent),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

//UpdateUserAiringNotification If the user should get notifications when a show they are watching aires
func UpdateUserAiringNotification(airingNotifications bool, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(airingNotifications:%v){
				options {
					  displayAdultContent,
					  airingNotifications,
    			},
			},
		}`, airingNotifications),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

//UpdateUserScoreFormat The user's list scoring system/Media list scoring type such as POINT_100, POINT_10_DECIMAL, POINT_10, POINT_5, POINT_3
func UpdateUserScoreFormat(scoreFormat string, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(scoreFormat:%v){
				mediaListOptions {
					scoreFormat,
    			},
			},
		}`, scoreFormat),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

//UpdateUserRowOrder The user's default list order such as Score, title, last updated, last added
func UpdateUserRowOrder(rowOrder string, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(rowOrder:"%v"){
				mediaListOptions {
					rowOrder,
    			},
			},
		}`, rowOrder),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

//UpdateUserProfileColor The user's Profile highlight color: blue, purple, green, orange, red, pink
func UpdateUserProfileColor(profileColor string, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(profileColor:"%v"){
				options {
					profileColor,
    			},
			},
		}`, profileColor),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateUserNotificationOption The user's Notification options

//ACTIVITY_MESSAGE
//A user has sent you message
//
//ACTIVITY_REPLY
//A user has replied to your activity
//
//FOLLOWING
//A user has followed you
//
//ACTIVITY_MENTION
//A user has mentioned you in their activity
//
//THREAD_COMMENT_MENTION
//A user has mentioned you in a forum comment
//
//THREAD_SUBSCRIBED
//A user has commented in one of your subscribed forum threads
//
//THREAD_COMMENT_REPLY
//A user has replied to your forum comment
//
//AIRING
//An anime you are currently watching has aired
//
//ACTIVITY_LIKE
//A user has liked your activity
//
//ACTIVITY_REPLY_LIKE
//A user has liked your activity reply
//
//THREAD_LIKE
//A user has liked your forum thread
//
//THREAD_COMMENT_LIKE
//A user has liked your forum comment
//
//ACTIVITY_REPLY_SUBSCRIBED
//A user has replied to activity you have also replied to
//
//RELATED_MEDIA_ADDITION
//A new anime or manga has been added to the site where its related media is on the user's list

func UpdateUserNotificationOption(notificationType string, enabled bool, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(notificationOptions:{type:%v,enabled:%v}) {
				options {
					notificationOptions {
						type,
						enabled,
					},
				},
			},
		}`, notificationType, enabled),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateUserAnimeListOptionSectionOrder A user's list options for anime or manga lists
func UpdateUserAnimeListOptionSectionOrder(sectionOrder string, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(animeListOptions:{sectionOrder:"%v"}) {
				mediaListOptions {
					animeList {
						sectionOrder
						splitCompletedSectionByFormat
						customLists
						advancedScoring
						advancedScoringEnabled
					}
				},
			},
		}`, sectionOrder),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateUserAnimeListOptionsplitCompletedSectionByFormat If the completed sections of the list should be separated by format
func UpdateUserAnimeListOptionsplitCompletedSectionByFormat(format bool, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(animeListOptions:{splitCompletedSectionByFormat:%v}) {
				mediaListOptions {
					animeList {
						sectionOrder
						splitCompletedSectionByFormat
						customLists
						advancedScoring
						advancedScoringEnabled
					}
				},
			},
		}`, format),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateUserAnimeListOptionsCustomLists The names of the user's custom lists
func UpdateUserAnimeListOptionsCustomLists(customList string, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(animeListOptions:{customLists:"%v"}) {
				mediaListOptions {
					animeList {
						sectionOrder
						splitCompletedSectionByFormat
						customLists
						advancedScoring
						advancedScoringEnabled
					}
				},
			},
		}`, customList),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateUserAdvancedScoringEnabled If advanced scoring is enabled
func UpdateUserAdvancedScoringEnabled(scoring bool, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(animeListOptions:{advancedScoringEnabled:%v}) {
				mediaListOptions {
					animeList {
						sectionOrder
						splitCompletedSectionByFormat
						customLists
						advancedScoring
						advancedScoringEnabled
					}
				},
			},
		}`, scoring),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateUserAdvancedScoring The names of the user's advanced scoring sections
func UpdateUserAdvancedScoring(name string, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(animeListOptions:{advancedScoring:"%v"}) {
				mediaListOptions {
					animeList {
						sectionOrder
						splitCompletedSectionByFormat
						customLists
						advancedScoring
						advancedScoringEnabled
					}
				},
			},
		}`, name),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Manga

// UpdateUserMangaListOptionsSectionOrder A user's list options for anime or manga lists
func UpdateUserMangaListOptionsSectionOrder(sectionOrder string, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(mangaListOptions:{sectionOrder:"%v"}) {
				mediaListOptions {
					animeList {
						sectionOrder
						splitCompletedSectionByFormat
						customLists
						advancedScoring
						advancedScoringEnabled
					}
				},
			},
		}`, sectionOrder),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateUserMangaListOptionSplitCompletedSectionByFormat If the completed sections of the list should be separated by format
func UpdateUserMangaListOptionSplitCompletedSectionByFormat(format bool, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(mangaListOptions:{splitCompletedSectionByFormat:%v}) {
				mediaListOptions {
					animeList {
						sectionOrder
						splitCompletedSectionByFormat
						customLists
						advancedScoring
						advancedScoringEnabled
					}
				},
			},
		}`, format),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateUserMangaListOptionsCustomLists The names of the user's custom lists
func UpdateUserMangaListOptionsCustomLists(customList string, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(mangaListOptions:{customLists:"%v"}) {
				mediaListOptions {
					animeList {
						sectionOrder
						splitCompletedSectionByFormat
						customLists
						advancedScoring
						advancedScoringEnabled
					}
				},
			},
		}`, customList),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

// DeleteCustomListManga Delete a custom list and remove the list entries from it
func DeleteCustomListManga(customList string, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        DeleteCustomList(type:MANGA, customList:"%v") {
    			deleted,
  			},
		}`, customList),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

// DeleteCustomLisAnime Delete a custom list and remove the list entries from it
func DeleteCustomLisAnime(customList string, authToken string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        DeleteCustomList(type:ANIME, customList:"%v") {
    			deleted,
  			},
		}`, customList),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

// ResetTheme will reset theme to default values which is "list" if encounter anime list that can't be shown even when the list is not big try using this
func ResetTheme(authToken string) (bool, error) {
	query := map[string]string{
		"query": `mutation{
	          UpdateUser(animeListOptions: {theme:"list"}){
				mediaListOptions {
					animeList {
						theme,
					},
				},
			},
		}`,
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	_, err = anilist.PostRequestAuth(jsonValue, authToken)
	if err != nil {
		return false, err
	}

	return true, nil
}
