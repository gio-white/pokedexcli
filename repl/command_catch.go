package repl
import (
	"fmt"
	"math/rand"
)
func commandCatch (c *Config, params []string) error {
	if len(params) < 1 {
		return fmt.Errorf("usage: catch <pokemon> - Add the pokemon that you want try to catch")
	}

	pokemonName := params[0]
	pokemonData, err := c.PokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("catch %s: %w", pokemonName, err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	baseExp := pokemonData.BaseExperience
	catchChance := 1.0 - float64(baseExp)/400.0

	if catchChance < 0.05 {
		catchChance = 0.05
	}
	if catchChance > 0.9 {
		catchChance = 0.9
	}

	roll := rand.Float64()
	if roll < catchChance {
		fmt.Printf("%s was caught!\n", pokemonName)
		c.Pokedex[pokemonName] = pokemonData
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}