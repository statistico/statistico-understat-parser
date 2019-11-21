package understat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var pattern = `\s+=\s+JSON.parse\(\'(.*?)\'\)`

func newStringReplace(str string) (string, error) {
	s := strings.Replace(str, " ", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = `"` + s + `"`

	s, err := strconv.Unquote(s)

	if err != nil {
		return "", err
	}

	return s, nil
}

func sendRequest(url string) (string, error) {
	resp, err := http.Get(url)

	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func parseStringMatch(body, key string, response interface{}) error {
	str, err := findMatch(body, key)

	if err != nil {
		return err
	}

	if err := marshalJson(str, response); err != nil {
		return err
	}

	return nil
}

func findMatch(body, key string) (string, error) {
	re := regexp.MustCompile(key + pattern)

	matches := re.FindAllStringSubmatch(body, -1)

	if len(matches) < 1 {
		return "", fmt.Errorf("no matches found")
	}

	sub := matches[0][1]

	return newStringReplace(sub)
}

func marshalJson(str string, response interface{}) error {
	if err := json.Unmarshal([]byte(str), &response); err != nil {
		return err
	}

	return nil
}
