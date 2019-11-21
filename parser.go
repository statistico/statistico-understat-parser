package understat

import (
	"fmt"
)

func GetLeagueFixtures(league, season string) ([]Fixture, error) {
	url := fmt.Sprintf(LeagueURL+"/%s/%s", league, season)

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
