package repl

import (
	"fmt"
)

func commandInspect(c *Config, params []string) error {
	if len(params) < 1 {
		return fmt.Errorf("usage: catch <pokemon> - Add the pokemon that you want try to catch")
	}

	pokemonName := params[0]
	pokemon, exist := c.Pokedex[pokemonName]
	if !exist {
		fmt.Println("You have not caught that Pokémon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}