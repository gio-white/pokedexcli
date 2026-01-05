package repl

import (
	"fmt"
	"github.com/gio-white/pokedexcli/internal/pokeapi"
	"net/url"
)
func commandMap(c *Config, params []string) error {
	var data pokeapi.Location
	var err error

	if c.Previous.String() == "" {
		data, err = c.PokeapiClient.GetLocationArea("https://pokeapi.co/api/v2/location-area/")
	} else {
		data, err = c.PokeapiClient.GetLocationArea(c.Next.String())
	}
	if err != nil {
		fmt.Println("Error with PokeAPI:", err)
		return err
	}
	
	for _, result := range data.Results {
		fmt.Println(result.Name)
	}

	if data.Next != nil {
		nextURL, err := url.Parse(*data.Next)
		if err != nil {
			return err
		}
		c.Next = *nextURL
	} else {
		c.Next = url.URL{}
	}

	return nil
}

func commandMapb(c *Config, params []string) error {
	var data pokeapi.Location
	var err error

	if c.Previous.String() == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	data, err = c.PokeapiClient.GetLocationArea(c.Previous.String())
	if err != nil {
		fmt.Println("Error with PokeAPI:", err)
		return err
	}

	for _, result := range data.Results {
		fmt.Println(result.Name)
	}

	if data.Previous != nil {
		prevURL, err := url.Parse(*data.Previous)
		if err != nil {
			return err
		}
		c.Previous = *prevURL
	} else {
		c.Previous = url.URL{}
	}

	return nil
}