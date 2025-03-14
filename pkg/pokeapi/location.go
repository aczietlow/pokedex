package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationArea struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) FetchLocationArea(locationName string) (locationArea, error) {
	url := baseURL + "/location-area/" + locationName

	if body, exists := c.cache.Get(url); exists {
		location := locationArea{}
		if err := json.Unmarshal(body, &location); err != nil {
			return locationArea{}, err
		}
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return locationArea{}, fmt.Errorf("Received a %s response from api", resp.StatusCode)
	}

	if resp.StatusCode != http.StatusOK {
		return locationArea{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationArea{}, err
	}

	location := locationArea{}
	if err = json.Unmarshal(body, &location); err != nil {
		return locationArea{}, err
	}

	c.cache.Add(url, body)
	return location, nil

}
