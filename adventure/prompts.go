package gomud

import (
	"fmt"
	"strings"
	"sync"
	"os"
	"bufio"
	"io"
)

var UserPrompt sync.Mutex

func Login() Character {
	return LoginWithReader(os.Stdin)
}

func LoginWithReader(input io.Reader) Character {
	UserPrompt.Lock()
	fmt.Println("Hello adventurer. What is your name?")
	s := bufio.NewScanner(input)
	s.Scan()
	name := s.Text()
	UserPrompt.Unlock()
	c := NewCharacter(name)
	return c
}

func (c *Character) ClassPrompt() {
	c.ClassPromptWithReader(os.Stdin)
}

func (c *Character) ClassPromptWithReader(input io.Reader) {
	UserPrompt.Lock()
	fmt.Printf("Please select a class %s: [fighter], mage, cleric, rogue\n", c.Name)
	s := bufio.NewScanner(input)
	s.Scan()
	class := s.Text()
	UserPrompt.Unlock()
	c.Class = classHandler(class)
}

func classHandler(class string) string {
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

func (c *Character) Prompt(quit chan bool, errorHandler chan string) {
	c.PromptWithReader(os.Stdin, quit, errorHandler)
}

func (c *Character) PromptWithReader(input io.Reader, quit chan bool, errorHandler chan string) {
	done := make(chan bool)
	for {
		UserPrompt.Lock()
		fmt.Println("\nWhat would you like to do?")
		s := bufio.NewScanner(input)
		s.Scan()
		c.Action = s.Text()
		UserPrompt.Unlock()
		go c.Do(errorHandler, quit, done)
		<-done
	}
}

func commandHandler(action string) (string, []string) {
	w := strings.Split(action, " ")
	cmd := w[0]
	args := w[1:]
	return cmd, args
}

func (c *Character) Do(errorHandler chan string, quit,done chan bool) {
	command, args := commandHandler(c.Action)
	UserPrompt.Lock()
	switch command {
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
	case "get":
		getHandler(args[0])
	default:
		fmt.Println("not possible.")
	}
	UserPrompt.Unlock()
	done <- true
}

func getHandler(arg string) {
	switch arg {
	case "coffee":
		fmt.Println("you get warm coffee in a fresh mug.")
	case "mug":
		fmt.Println("you grab an empty mug, but you don't have any pockets.")
	case "crate":
		fmt.Println("it's too heavy.")
	default:
		fmt.Println("get what?")
	}
}

func statsHandler(c *Character) {
	fmt.Printf("%s the %s\n%s\n", c.Name, c.Class, c.Stats.Text())
}

func helpHandler() {
	fmt.Println("you can: help,look,inventory,areas,stats,quit")
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
