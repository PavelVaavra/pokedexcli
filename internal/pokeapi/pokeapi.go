package pokeapi

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Urls struct {
	Next string
	Previous string
}
	
type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func CommandMap(urls *Urls) error {
	res, err := http.Get(urls.Next)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var locationArea LocationArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationArea)
	if err != nil {
		return err
	}

	var locations []string
	for _, location := range locationArea.Results {
		locations = append(locations, location.Name)
	}

	for _, location := range locations {
		fmt.Println(location)
	}
	
	urls.Next, urls.Previous = locationArea.Next, locationArea.Previous

	return nil
}