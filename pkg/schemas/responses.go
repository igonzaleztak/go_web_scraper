package schemas

import (
	"encoding/json"
	"fmt"
	"time"
)

type LiveDataHackerNews []int

type StoryHackerNews struct {
	ID          int       `json:"id"`
	Type        string    `json:"type"`
	By          string    `json:"by"`
	Time        Timestamp `json:"time"`
	Text        string    `json:"text"`
	Dead        bool      `json:"dead"`
	Parent      string    `json:"parent"`
	Poll        string    `json:"poll"`
	Kids        []int     `json:"kids,omitempty"`
	URL         string    `json:"url"`
	Score       int       `json:"score"`
	Title       string    `json:"title"`
	Parts       []string  `json:"parts"`
	Descendants int       `json:"descendants"`
}

// Timestamp is a custom type to handle Unix timestamps. By default, the Hacker news API returns Unix timestamps as int64.
type Timestamp struct {
	time.Time
}

// UnmarshalJSON decodes an int64 timestamp into a time.Time object
func (p *Timestamp) UnmarshalJSON(bytes []byte) error {
	var raw int64
	err := json.Unmarshal(bytes, &raw)

	if err != nil {
		fmt.Printf("error decoding timestamp: %s\n", err)
		return err
	}

	p.Time = time.Unix(raw, 0)
	return nil
}

// SpaceFlightResponse is the response from the Space Flight News API.
type SpaceFlightResponse struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []SpaceFlightNews `json:"results"`
}

// SpaceFlightNews is the structure of a news item from the Space Flight News API.
type SpaceFlightNews struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	ImageURL    string    `json:"image_url"`
	NewsSite    string    `json:"news_site"`
	Summary     string    `json:"summary"`
	PublishedAt time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Featured    bool      `json:"featured"`
	Launches    any       `json:"launches"`
	Events      any       `json:"events"`
}
