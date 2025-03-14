package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheFetch(t *testing.T) {
	type testCases struct {
		cacheKey string
		data     []byte
	}

	cases := []testCases{
		{"favorite", []byte("pikachu")},
	}

	client := NewCache(5 * time.Second)

	for _, test := range cases {
		fmt.Printf("\nTest Data %s", test.data)
		client.Add(test.cacheKey, test.data)
	}

	for _, test := range cases {
		data, _ := client.Get(test.cacheKey)
		if string(data) != string(test.data) {
			t.Errorf("Cache data does not match. Expected %s and received %s", string(data), string(data))
		}
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://zietlow.io", []byte("This is the tale of captain Jack Sparrow"))

	_, ok := cache.Get("https://zietlow.io")
	if !ok {
		t.Errorf("Failed to find cache item with key")
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://zietlow.io")
	if ok {
		t.Errorf("Found key, when expected the cache to be empty")
	}
}
