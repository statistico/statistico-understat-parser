package understat

type (
	Fixture struct {
		ID       string `json:"id"`
		Home     Team   `json:"h"`
		Away     Team   `json:"a"`
		IsResult bool   `json:"isresult"`
		XG       XG     `json:"xg"`
		Goals    Goals  `json:"goals"`
		DateTime string `json:"datetime"`
	}

	Forecast struct {
		Win  string `json:"w"`
		Lose string `json:"l"`
		Draw string `json:"d"`
	}

	Goals struct {
		Home *string `json:"h"`
		Away *string `json:"a"`
	}

	Team struct {
		ID         string `json:"id"`
		Title      string `json:"title"`
		ShortTitle string `json:"short_title"`
	}

	XG struct {
		Home *string `json:"h"`
		Away *string `json:"a"`
	}
)
