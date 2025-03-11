package pokeapi

import (
	"net/http"
	"time"

	"github.com/aczietlow/pokedex/pkg/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{},
		cache:      pokecache.NewCache(cacheInterval),
	}
}
