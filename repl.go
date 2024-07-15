package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *Config) {
	cmds := getCommands(config)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("=================")
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		fmt.Println("=================")

		commandName := words[0]
		command, exists := cmds[commandName]
		parameter := ""
		if len(words) > 1 {
			parameter = words[1]
		}
		if exists {
			err := command.callback(command.config, parameter)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	return strings.Fields(output)
}
