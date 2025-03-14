package main

import (
	"fmt"
)

type locationArea struct {
	Location []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
}

func commandMap(conf *config, args ...string) error {
	err := fetchLocations(conf)
	if err != nil {
		return err
	}

	return nil
}

func commandMapB(conf *config, args ...string) error {
	if conf.mapPager <= 40 {
		conf.mapPager = 0
	} else {
		conf.mapPager -= 40
	}

	err := fetchLocations(conf)
	if err != nil {
		return err
	}

	return nil
}

func fetchLocations(conf *config) error {
	locations, err := conf.apiClient.FetchLocationsList(conf.mapPager)
	if err != nil {
		return err
	}

	for _, v := range locations.Locations {
		fmt.Printf("%v\n", v.Name)
	}
	conf.mapPager += 20
	return nil
}
