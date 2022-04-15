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
		helpHandler()
	case "stats":
		statsHandler(c)
	case "inventory":
		inventoryHandler()
	case "look":
		lookHandler()
	case "areas":
		areasHandler()
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

func helpHandler() {
	fmt.Println("you can: help,look,inventory,areas,quit")
}

func inventoryHandler() {
	fmt.Println("you ain't got shit.")
}

func lookHandler() {
	fmt.Println("There's a tree. it doesn't move much. There's a wooden crate under the tree, and a pot of coffee with clean mugs.")
}

func areasHandler() {
	// future support for area files
	// https://github.com/alexmchale/merc-mud/blob/master/doc/area.txt
	fmt.Println("#AREA	{ 5 35} Eilidh    The Coffeehouse~")
}
