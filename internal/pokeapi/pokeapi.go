package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) GetLocationArea(url string) (Location, error) {
	fullURL := url


	if val, ok := client.cache.Get(url); ok {
		locationsResp := Location{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return Location{}, err
		}

		return locationsResp, nil
	}


	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return Location{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return Location{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("HTTP error:", resp.Status)
		return Location{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return Location{}, err
	}

	data := Location{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return Location{}, err
	}

	client.cache.Add(url, body)
	return data, nil
}