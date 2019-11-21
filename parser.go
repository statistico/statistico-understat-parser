package understat

import (
	"fmt"
)
type Parser struct {
	BaseURL string
}

func (p Parser) LeagueFixtures(league, season string) ([]Fixture, error) {
	url := fmt.Sprintf(p.BaseURL + "/league/%s/%s", league, season)

	var response []Fixture

	body, err := sendRequest(url)

	if err != nil {
		return response, err
	}

	if err := parseStringMatch(body, "datesData", &response); err != nil {
		return response, err
	}

	return response, nil
}
