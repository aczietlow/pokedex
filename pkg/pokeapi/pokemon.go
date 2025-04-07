package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokemon struct {
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

func (c *Client) FetchPokemon(pokemonName string) (pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if body, exists := c.cache.Get(url); exists {
		p := pokemon{}
		if err := json.Unmarshal(body, &p); err != nil {
			return pokemon{}, err
		}
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return pokemon{}, fmt.Errorf("Received a %v response from api", resp.StatusCode)
	}

	if resp.StatusCode != http.StatusOK {
		return pokemon{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemon{}, err
	}

	p := pokemon{}
	if err = json.Unmarshal(body, &p); err != nil {
		return pokemon{}, err
	}

	c.cache.Add(url, body)
	return p, nil

}
