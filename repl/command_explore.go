package repl

import "fmt"

func commandExplore (c *Config, params []string) error {
	if len(params) < 1 {
		return fmt.Errorf("usage: explore <location-area> - Add the location that you want to explore")
	}
	location := params[0]

	data, err := c.PokeapiClient.GetLocationPokemon(location)
	if err != nil {
		return fmt.Errorf("ERROR - explore %s: %w", location, err)
	}
	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokémon:")

	for _, encounter := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}