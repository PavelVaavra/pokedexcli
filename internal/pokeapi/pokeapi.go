package pokeapi

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"io"
	"github.com/PavelVaavra/pokedexcli/internal/pokecache"
)

var cache = pokecache.NewCache(time.Duration(20) * time.Second)

type Urls struct {
	Next string
	Previous string
	ExploreBasis string
	CatchBasis string
	ExploreArea string
	CatchPokemon string
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
	url := urls.Next

	bytes, err := getBytes(url)
	if err != nil {
		return err
	}

	next, previous, err := parseLocationArea(bytes)
	
	urls.Next, urls.Previous = next, previous
	return err
}

func CommandMapb(urls *Urls) error {
	if urls.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	url := urls.Previous

	bytes, err := getBytes(url)
	if err != nil {
		return err
	}

	next, previous, err := parseLocationArea(bytes)
	
	urls.Next, urls.Previous = next, previous
	return err
}

func parseLocationArea(bytes []byte) (next, previous string, err error) {
	var locationArea LocationArea

	err = json.Unmarshal(bytes, &locationArea)
	if err != nil {
		return "", "", err
	}

	var names []string
	for _, results := range locationArea.Results {
		names = append(names, results.Name)
	}

	for i, name := range names {
		if i == 0 || i == 1 {
			fmt.Println(name)
		}
	}

	return locationArea.Next, locationArea.Previous, nil
}

func getBytes(url string) (bytes []byte, err error) {
	bytes, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return nil, fmt.Errorf("Response failed with status code: %d", res.StatusCode)
		}

		bytes, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		cache.Add(url, bytes)
	}
	
	return bytes, nil
}