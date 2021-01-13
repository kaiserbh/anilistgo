package anilist

import (
	anilist "AniListGoWrapper/anilist/query"
	"encoding/json"
	"fmt"
)

//UpdateUserAbout User's about/bio text.
func UpdateUserAbout(about string, authToken string) {
	query := map[string]string{
		"query": fmt.Sprintf(`mutation{
	        UpdateUser(about:"%v"){
				about,
			},
		}`, about),
	}

	jsonValue, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

//UpdateUserTitleLanguage User's title language/ The language the user wants to see media titles in such as ROMAJI, ENGLISH, NATIVE, ROMAJI_STYLISED, ENGLISH_STYLISED, NATIVE_STYLISED
func UpdateUserTitleLanguage(titleLanguage string, authToken string) {
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
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

//UpdateUserAdultContent If the user should see media marked as adult-only boolean
func UpdateUserAdultContent(displayAdultContent bool, authToken string) {
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
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

//UpdateUserAiringNotification If the user should get notifications when a show they are watching aires
func UpdateUserAiringNotification(airingNotifications bool, authToken string) {
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
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

//UpdateUserScoreFormat The user's list scoring system/Media list scoring type such as POINT_100, POINT_10_DECIMAL, POINT_10, POINT_5, POINT_3
func UpdateUserScoreFormat(scoreFormat string, authToken string) {
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
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

//UpdateUserRowOrder The user's default list order such as Score, title, last updated, last added
func UpdateUserRowOrder(rowOrder string, authToken string) {
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
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}

//UpdateUserProfileColor The user's Profile highlight color: blue, purple, green, orange, red, pink
func UpdateUserProfileColor(profileColor string, authToken string) {
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
		panic(err)
	}

	fmt.Println(string(anilist.PostRequestAuth(jsonValue, authToken)))
}
