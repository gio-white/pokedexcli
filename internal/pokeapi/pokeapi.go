// package pokeapi

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func (client *Client) GetLocationArea(url string) (Location, error) {
// 	baseURL := "https://pokeapi.co/api/v2/location-area/"
// 	fullURL := baseURL
// 	if url != "" {
// 		fullURL += url
// 	}

// 	if val, ok := client.cache.Get(url); ok {
// 		locationsResp := Location{}
// 		err := json.Unmarshal(val, &locationsResp)
// 		if err != nil {
// 			return Location{}, err
// 		}

// 		return locationsResp, nil
// 	}

// 	req, err := http.NewRequest("GET", fullURL, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return Location{}, err
// 	}

// 	resp, err := client.httpClient.Do(req)
// 	if err != nil {
// 		fmt.Println("Error making request:", err)
// 		return Location{}, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Println("HTTP error:", resp.Status)
// 		return Location{}, err
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading body:", err)
// 		return Location{}, err
// 	}

// 	data := Location{}
// 	err = json.Unmarshal(body, &data)
// 	if err != nil {
// 		fmt.Println("Error parsing JSON:", err)
// 		return Location{}, err
// 	}

// 	client.cache.Add(url, body)
// 	return data, nil
// }

package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) fetchJSON(fullURL string, cacheKey string, target any) error {
	if val, ok := client.cache.Get(cacheKey); ok {
		return json.Unmarshal(val, target)
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, target); err != nil {
		return err
	}

	client.cache.Add(cacheKey, body)

	return nil
}


func (client *Client) GetLocationArea(url string) (Location, error) {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	fullURL := baseURL

	if url != "" {
		fullURL = url
	}

	var location Location
	err := client.fetchJSON(fullURL, fullURL, &location)
	if err != nil {
		return Location{}, err
	}

	return location, nil
}

func (client *Client) GetLocationPokemon(location string) (LocationArea, error) {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	fullURL := baseURL + location + "/"
	if location == "" {
		return LocationArea{}, errors.New("No location provided")
	}

	var locationArea LocationArea
	err := client.fetchJSON(fullURL, fullURL, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}

func (client *Client) GetPokemon(pokemon string) (Pokemon, error) {
	baseURL := "https://pokeapi.co/api/v2/pokemon/"
	fullURL := baseURL + pokemon

	var pokemonInfo Pokemon
	err := client.fetchJSON(fullURL, fullURL, &pokemonInfo)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemonInfo, nil
}