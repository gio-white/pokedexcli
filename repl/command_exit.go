package repl

import (
	"fmt"
	"os"
)

func commandExit(c *Config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}