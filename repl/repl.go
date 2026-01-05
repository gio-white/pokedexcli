package repl

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/gio-white/pokedexcli/internal/pokeapi"
)

func StartRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		params := words[1:]

		command, exists := GetCommands()[commandName]
		if exists {
			err := command.callback(cfg, params)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}


func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config, params []string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:		"map",
			description: "Listing the following 20 Location Areas",
			callback: 	commandMap,
		},
		"mapb": {
			name:		"mapb",
			description: "Listing the previous 20 Location Areas",
			callback: 	commandMapb,
		},
		"explore": {
			name:		"explore",
			description: "Listing Pokemon in a specified Area",
			callback:	commandExplore,
		},
		"catch": {
			name:		"catch",
			description: "Try to catch a Pokemon",
			callback:	commandCatch,
		},
		"inspect": {
			name:		"inspect",
			description: "Describe the attributes of a Pokemon",
			callback:	commandInspect,
		},
		"pokedex": {
			name:		"pokedex",
			description: "List the captured Pokemon",
			callback:	commandPokedex,
		},
	}
}

type Config struct{
	PokeapiClient 	*pokeapi.Client
	Next 			url.URL
	Previous 		url.URL
	Pokedex			map[string]pokeapi.Pokemon
}
