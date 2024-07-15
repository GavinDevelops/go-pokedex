module github.com/GavinDevelops/pokedexcli

replace github.com/GavinDevelops/pokedexcli/pokeapi v0.0.0 => ./pokeapi/

replace github.com/GavinDevelops/pokedex/cli/pokecache v0.0.0 => ./pokecache/

require (
	github.com/GavinDevelops/pokedexcli/pokeapi v0.0.0
	github.com/GavinDevelops/pokedex/cli/pokecache v0.0.0
)

go 1.22.4
