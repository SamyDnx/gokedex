package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("gokedex -> ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		r_input := sc.Text()
		input := cleanInput(r_input)

		if input[0] == "exit" {
			fmt.Println("Exiting...")
			break
		} else {
			fmt.Println("Your command was:", input[0], "")
		}
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}
