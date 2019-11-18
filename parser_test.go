package understat

import (
	"fmt"
	"testing"
)

func TestGetLeagueFixturesNow(t *testing.T) {
	fix, _ := GetLeagueFixtures("EPL", "2019")

	fmt.Printf("Fixtures: %+v", fix)
}
