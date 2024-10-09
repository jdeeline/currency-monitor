package grabber

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"
)

const url = "https://www.cbr-xml-daily.ru/daily_json.js"

type Currency struct {
	ID           string  `json:"ID"`
	Code         string  `json:"CharCode"`
	Nominal      int     `json:"Nominal"`
	Name         string  `json:"Name"`
	Rate         float64 `json:"Value"`
	PreviousRate float64 `json:"Previous"`
}

// Executes a request to the API, parses the response and returns an array of currencies.
func Grab() ([]Currency, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("can't get response: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read body: %w", err)
	}

	var result response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("can't unmarshal response body: %w", err)
	}

	return result.Valute.sorted(), nil
}

type response struct {
	Date         time.Time `json:"Date"`
	PreviousDate time.Time `json:"PreviousDate"`
	PreviousURL  string    `json:"PreviousURL"`
	Timestamp    time.Time `json:"Timestamp"`
	Valute       valute    `json:"Valute"`
}

type valute map[string]Currency

func (v *valute) sorted() []Currency {
	keys := make([]string, 0, len(*v))
	for k := range *v {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sorted := make([]Currency, len(*v))
	for i, k := range keys {
		sorted[i] = (*v)[k]
	}

	return sorted
}
