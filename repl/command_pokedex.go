package repl

import "fmt"

func commandPokedex(c *Config, params []string) error {
	if len(params) > 0 {
		return fmt.Errorf("usage: pokedex (no arguments)")
	}

	if len(c.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty. Catch some Pokémon first!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range c.Pokedex {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
