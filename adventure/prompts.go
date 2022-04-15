package gomud

import (
	"fmt"
	"strings"
)

func Login() Character {
	var name string
	fmt.Println("Hello adventurer. What is your name?")
	fmt.Scanln(&name)
	c := NewCharacter(name)
	return c
}

func (c *Character) ClassPrompt() {
	var class string
	fmt.Printf("Please select a class %s: [fighter], mage, cleric, rogue\n", c.Name)
	fmt.Scanln(&class)
	c.Class = parseClassPrompt(class)
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

func (c *Character) Prompt(quit chan bool) {
	var action string
	done := make(chan bool)
	for {
		fmt.Println("\nWhat would you like to do?")
		fmt.Scanln(&action)
		go c.Do(quit, done, action)
		<-done
	}
}

func (c *Character) Do(quit,done chan bool, action string) {
	switch strings.TrimSpace(action) {
	case "quit":
		fmt.Println("goodbye adventurer.")
		quit <- true
	case "help":
		fmt.Println("you can: help,look,inventory,quit,stats")
	case "stats":
		statsHandler(c)
	case "inventory":
		fmt.Println("you ain't got shit.")
	case "look":
		fmt.Println("there's a tree. it doesn't move much.")
	case "dance":
		fmt.Println("shake your booty.")
	default:
		fmt.Println("not possible.")
	}
	done <- true
}

func statsHandler(c *Character) {
	fmt.Printf("%s - %s\n%s\n", c.Name, c.Class, c.Stats.Text())
}
