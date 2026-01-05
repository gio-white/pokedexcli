package main

import (
	"github.com/gio-white/pokedexcli/internal/pokeapi"
	"github.com/gio-white/pokedexcli/repl"
	"time"
)	

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &repl.Config{
		PokeapiClient: 	&pokeClient,
		Pokedex:		make(map[string]pokeapi.Pokemon),
	}
	repl.StartRepl(cfg)
}
