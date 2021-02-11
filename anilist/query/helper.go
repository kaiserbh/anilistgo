package query

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// regex Helper

// PostRequest sends POST request that takes []byte as parameter then returns back httml Body as []byte
func PostRequest(jsonValue []byte) []byte {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Check if the respone is 200 if not then media is not found or server might be down.
	if resp.StatusCode != 200 {
		fmt.Println("Media Not Found got Error with statusCode: ", resp.StatusCode)
		fmt.Println(resp.Body)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return data
}

// PostRequestAuth Request with token and auth token can be retrived from anilist
func PostRequestAuth(jsonValue []byte, authKey string) []byte {
	var request *http.Request
	var err error

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	if jsonValue != nil {
		body := bytes.NewBuffer(jsonValue)
		request, err = http.NewRequest("POST", url, body)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		request, err = http.NewRequest("application/json", url, nil)
	}
	request.Header.Set("Authorization", "Bearer "+authKey)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("Error: ", resp.Body)
	} else {
		fmt.Println(" ------------ SUCCESS ------------")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return body
}

// CleanJSON cleans the html body by removing unnecessary html tags using regex.
// That then returns []byte data and takes []byte as parameter
func cleanMediaJSON(str []byte) []byte {
	re := regexp.MustCompile(`(?m){"data":{"Media":`)
	substitution := ""
	firstItr := re.ReplaceAll(str, []byte(substitution))
	// clean the last two ]] from the string
	reCleanReminders := regexp.MustCompile(`(?m)}}$`)
	cleanJSON := reCleanReminders.ReplaceAll(firstItr, []byte(substitution))

	// Clean the html tags in description
	reCleanHTMLDescription := regexp.MustCompile(`(?m)<br>\\n`)
	cleanDescription := reCleanHTMLDescription.ReplaceAll(cleanJSON, []byte(substitution))

	// Clean the i tags
	cleanTags := regexp.MustCompile(`(?m)<i>`)
	cleanTagsI := regexp.MustCompile(`(?m)<\\/i>`)

	secondItr := cleanTags.ReplaceAll(cleanDescription, []byte(substitution))
	lastItr := cleanTagsI.ReplaceAll(secondItr, []byte(substitution))

	return lastItr
}

// CleanPageJSON cleans the hmtl body
func CleanPageJSON(str []byte) []byte {
	re := regexp.MustCompile(`(?m){"data":{"Page":|}}$`)
	substitution := ""
	firstItr := re.ReplaceAll(str, []byte(substitution))

	return firstItr
}

// CleanMediaTrendPageJSON cleans the hmtl
func CleanMediaTrendPageJSON(str []byte) []byte {
	re := regexp.MustCompile(`(?m){"data":{"MediaTrend":|}}$`)
	substitution := ""
	firstItr := re.ReplaceAll(str, []byte(substitution))

	return firstItr
}

// CleanStaffJSON cleans the hmtl body
func CleanStaffJSON(str []byte) []byte {
	re := regexp.MustCompile(`(?m){"data":{"Staff":|}}$`)
	substitution := ""
	firstItr := re.ReplaceAll(str, []byte(substitution))

	return firstItr
}

// CleanCharacterJSON cleans the hmtl body
func CleanCharacterJSON(str []byte) []byte {
	re := regexp.MustCompile(`(?m){"data":{"Character":|}}$`)
	substitution := ""
	firstItr := re.ReplaceAll(str, []byte(substitution))

	return firstItr
}

// CleanUserJSON cleans the hmtl body
func CleanUserJSON(str []byte) []byte {
	re := regexp.MustCompile(`(?m){"data":{"User":|}}$`)
	substitution := ""
	firstItr := re.ReplaceAll(str, []byte(substitution))

	return firstItr
}

// CleanUserListJSON clean the body for the usermedialist
func CleanUserListJSON(str []byte) []byte {
	re := regexp.MustCompile(`(?m){"data":{"Media":{"mediaListEntry":|}}}$`)
	substitution := ""
	firstItr := re.ReplaceAll(str, []byte(substitution))
	return firstItr
}
