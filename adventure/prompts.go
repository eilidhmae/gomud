package gomud

import (
	"fmt"
	"strings"
)

func ClassPrompt() string {
	var class string
	fmt.Println("Please select a class: [fighter], mage, cleric, rogue")
	fmt.Scanln(&class)
	return parseClassPrompt(class)
}

func parseClassPrompt(class string) string {
	trimmed := strings.TrimSpace(class)
	switch trimmed {
	case "mage":
		return trimmed
	case "cleric":
		return trimmed
	case "rogue":
		return trimmed
	default:
		return "fighter"
	}
}

func Prompt(quit chan bool) {
	var action string
	done := make(chan bool)
	for {
		fmt.Println("What would you like to do?")
		fmt.Scanln(&action)
		go parsePrompt(quit, done, action)
		<-done
	}
}

func parsePrompt(quit,done chan bool, action string) {
	switch strings.TrimSpace(action) {
	case "quit":
		fmt.Println("goodbye adventurer.")
		quit <- true
	case "help":
		fmt.Printf("you can: help,look,inventory,quit\n\n")
	case "inventory":
		fmt.Printf("you ain't got shit.\n\n")
	case "look":
		fmt.Printf("there's a tree. it doesn't move much.\n\n")
	default:
		fmt.Printf("not possible.\n\n")
	}
	done <- true
}
