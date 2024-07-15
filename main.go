package main

func main() {
	config := Config{
		next:     "https://pokeapi.co/api/v2/location/",
		previous: "",
	}
	startRepl(&config)
}
