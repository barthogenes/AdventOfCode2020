package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// AdventOfCodeAPI Used for interacting with the AdventOfCode web API. This is a public struct (aka class) because the name is in Pascal-casing.
type AdventOfCodeAPI struct {
	Year   int
	Cookie string
}

const adventOfCodeSessionCookieName = "session"

// GetInputInBytes Get the input for the given advent day. Note: Public (aka Pascal-cased) functions must always have a documentation comment above them in Go.
func (api AdventOfCodeAPI) GetInputInBytes(day int) []byte {
	adventOfCodeURL := getAdventOfCodeURL(api.Year, day)

	req, err := http.NewRequest("GET", adventOfCodeURL, nil)
	if err != nil {
		panic(err)
	}

	req.AddCookie(&http.Cookie{Name: adventOfCodeSessionCookieName, Value: api.Cookie})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	htmlBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return htmlBytes
}

// Construct the url of the adventofcode web API. Note: This is a private function because the name is camelCased.
func getAdventOfCodeURL(year int, day int) string {
	return fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
}
