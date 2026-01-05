package repl

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gio-white/pokedexcli/internal/pokeapi"
)


func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(c *Config) error {
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

func commandMapb(c *Config) error {
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