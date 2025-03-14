package pokeapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type locationAreas struct {
	// Location []location `json:"results"`
	Locations []struct {
		Name string `json:"name"`
	} `json:"results"`
}

// type location struct {
// 	Name string `json:"name"`
// }

func (c *Client) FetchLocationsList(pager int) (locationAreas, error) {
	offset := strconv.Itoa(pager)
	url := baseURL + "/location-area?limit=20&offset=" + offset

	// _, ok := c.cache.Get(url)
	// fmt.Printf("debug\nurl:%v\ncache:%v\n", url, ok)
	if body, exists := c.cache.Get(url); exists {
		// fmt.Println("cache hit")
		// fmt.Printf("data:\n%s\n", body)
		var locations locationAreas
		decoder := json.NewDecoder(bytes.NewReader(body))
		if err := decoder.Decode(&locations); err != nil {
			return locationAreas{}, err
		}

		return locations, nil
	}

	response, err := http.Get(url)
	if err != nil {
		return locationAreas{}, err
	}
	defer response.Body.Close()

	// TODO: Reading from the same io.Reader twice is dirty, should fix this.
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return locationAreas{}, err
	}

	var locations locationAreas
	decoder := json.NewDecoder(bytes.NewReader(body))
	if err = decoder.Decode(&locations); err != nil {
		return locationAreas{}, err
	}

	c.cache.Add(url, body)

	return locations, nil
}
