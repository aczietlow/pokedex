package pokeapi

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type locationAreas struct {
	// Location []location `json:"results"`
	Locations []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type location struct {
	Name string `json:"name"`
}

func (c *Client) FetchLocations(pager int) (locationAreas, error) {
	offset := strconv.Itoa(pager)
	url := baseURL + "/location-area?limit=20&offset=" + offset
	response, err := http.Get(url)
	if err != nil {
		return locationAreas{}, err
	}
	defer response.Body.Close()

	var locations locationAreas
	decoder := json.NewDecoder(response.Body)
	if err = decoder.Decode(&locations); err != nil {
		return locationAreas{}, err
	}

	return locations, nil
}
