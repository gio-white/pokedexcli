package repl

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gio-white/pokedexcli/pokeapi"
)


func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
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

func commandMap(c *config) error {
	var data map[string]any
	var err error

	if c.Previous.String() == "" {
		data, err = pokeapi.GetLocationArea("https://pokeapi.co/api/v2/location-area/")
	} else {
		data, err = pokeapi.GetLocationArea(c.Next.String())
	}
	if err != nil {
		fmt.Println("Error with PokeAPI:", err)
		return err
	}
	
	results, ok := data["results"].([]interface{})
	if !ok {
		return fmt.Errorf("unexpected type for results")
	}

	for _, r := range results {
		resultMap, ok := r.(map[string]interface{})
		if !ok {
			continue
		}
		fmt.Println(resultMap["name"])
	}
	nextStr, _ := data["next"].(string)
	if nextStr != "" {
		nextPage, err := url.Parse(nextStr)
		c.Next = *nextPage
		if err != nil {
			return err
		}
	} else {
		c.Next = url.URL{}
	}
	return nil
}

func commandMapb(c *config) error {
	var data map[string]any
	var err error

	if c.Previous.String() == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		data, err = pokeapi.GetLocationArea(c.Previous.String())
	}
	if err != nil {
		fmt.Println("Error with PokeAPI:", err)
		return err
	}
	
	results, ok := data["results"].([]interface{})
	if !ok {
		return fmt.Errorf("unexpected type for results")
	}

	for _, r := range results {
		resultMap, ok := r.(map[string]interface{})
		if !ok {
			continue
		}
		fmt.Println(resultMap["name"])
	}
	previousStr, _ := data["previous"].(string)
	if previousStr != "" {
		previousPage, err := url.Parse(previousStr)
		c.Previous = *previousPage
		if err != nil {
			return err
		}
	} else {
		c.Next = url.URL{}
	}
	return nil
}