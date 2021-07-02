package anilistgo

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// regex Helper

// PostRequest sends POST request that takes []byte as parameter then returns back html Body as []byte
func PostRequest(jsonValue []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	// Check if the response is 200 if not then media is not found or server might be down.
	if resp.StatusCode != 200 {
		log.Error("Media Not Found got Error with statusCode: ", resp.StatusCode)
		log.Warn(resp.Body)
		return nil, nil
	}

	if resp.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(resp.Body)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// PostRequestAuth Request with token and auth token can be retrieved from anilist
func PostRequestAuth(jsonValue []byte, authKey string) ([]byte, error) {
	var request *http.Request
	var err error

	timeout := 5 * time.Second
	client := &http.Client{
		Timeout: timeout,
	}

	if jsonValue != nil {
		body := bytes.NewBuffer(jsonValue)
		request, err = http.NewRequest("POST", url, body)
		if err != nil {
			return nil, err
		}
	} else {
		request, err = http.NewRequest("application/json", url, nil)
		if err != nil {
			return nil, err
		}
	}
	request.Header.Set("Authorization", "Bearer "+authKey)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Error("StatusCode 200: ", resp.Body)
		return nil, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
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

// CleanPageJSON cleans the html body
func CleanPageJSON(str []byte) []byte {
	re := regexp.MustCompile(`(?m){"data":{"Page":|}}$`)
	substitution := ""
	firstItr := re.ReplaceAll(str, []byte(substitution))

	return firstItr
}

// CleanMediaTrendPageJSON cleans the html
func CleanMediaTrendPageJSON(str []byte) []byte {
	re := regexp.MustCompile(`(?m){"data":{"MediaTrend":|}}$`)
	substitution := ""
	firstItr := re.ReplaceAll(str, []byte(substitution))

	return firstItr
}

// CleanStaffJSON cleans the html body
func CleanStaffJSON(str []byte) []byte {
	re := regexp.MustCompile(`(?m){"data":{"Staff":|}}$`)
	substitution := ""
	firstItr := re.ReplaceAll(str, []byte(substitution))

	return firstItr
}

// CleanCharacterJSON cleans the html body
func CleanCharacterJSON(str []byte) []byte {
	re := regexp.MustCompile(`(?m){"data":{"Character":|}}$`)
	substitution := ""
	firstItr := re.ReplaceAll(str, []byte(substitution))

	return firstItr
}

// CleanUserJSON cleans the html body
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
