package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Species                struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (c *Client) FetchPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if body, exists := c.cache.Get(url); exists {
		p := Pokemon{}
		if err := json.Unmarshal(body, &p); err != nil {
			return Pokemon{}, err
		}
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Received a %v response from api", resp.StatusCode)
	}

	if resp.StatusCode == http.StatusNotFound {
		return Pokemon{}, fmt.Errorf("Pokemon with name or id of %s not found\n", pokemonName)
	}

	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("Recieved a %v response from api\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	p := Pokemon{}
	if err = json.Unmarshal(body, &p); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)
	return p, nil
}
