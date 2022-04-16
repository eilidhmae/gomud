package gomud

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func LoginWithOS() Character {
	return Login(os.Stdin, os.Stdout)
}

func Login(r io.Reader, w io.Writer) Character {
	w.Write([]byte("Hello adventurer. What is your name?\n"))
	s := bufio.NewScanner(r)
	s.Scan()
	name := s.Text()
	c := NewCharacter(name, r, w)
	return c
}

func (c *Character) ClassPrompt(r io.Reader, w io.Writer) {
	msg := fmt.Sprintf("Please select a class %s: [fighter], mage, cleric, rogue\n", c.Name)
	c.Mutex.Lock()
	w.Write([]byte(msg))
	s := bufio.NewScanner(r)
	s.Scan()
	class := s.Text()
	c.Mutex.Unlock()
	c.Class = classHandler(class)
}

func classHandler(class string) string {
	t := strings.TrimSpace(class)
	switch t {
	case "mage":
		return t
	case "cleric":
		return t
	case "rogue":
		return t
	default:
		return "fighter"
	}
}

func (c *Character) PromptWithOS(quit chan bool, errorHandler chan error) {
	go c.Prompt(os.Stdin, os.Stdout, quit, errorHandler)
}

func (c *Character) Prompt(r io.Reader, w io.Writer, quit chan bool, errorHandler chan error) {
	for {
		c.Mutex.Lock()
		w.Write([]byte("\nWhat would you like to do?\n"))
		s := bufio.NewScanner(r)
		s.Scan()
		c.Action = s.Text()
		c.Mutex.Unlock()
		if q := c.Do(r, w, errorHandler); q == true {
			quit <- true
		}
	}
}

func commandHandler(action string) (string, []string) {
	var args []string
	w := strings.Split(action, " ")
	cmd := w[0]
	if len(w) > 1 {
		args = w[1:]
	} else {
		args = []string{"unknown"}
	}
	return cmd, args
}

// Do returns true for quit, and false for any other command
func (c *Character) Do(r io.Reader, w io.Writer, errorHandler chan error) bool {
	command, args := commandHandler(c.Action)
	c.Mutex.Lock()
	switch command {
	case "quit":
		msg := fmt.Sprintf("Until next time %s!\n", c.Name)
		w.Write([]byte(msg))
		return true
	case "help":
		w.Write([]byte("you can: areas,get,help,inventory,look,quit,stats\n"))
	case "stats":
		msg := fmt.Sprintf("%s the %s\n%s\n", c.Name, c.Class, c.Stats.Text())
		c.Writer.Write([]byte(msg))
	case "inventory":
		w.Write([]byte("you ain't got shit.\n"))
	case "look":
		w.Write([]byte("There's a tree. it doesn't move much. There's a wooden crate under the tree, and a pot of coffee with clean mugs.\n"))
	case "areas":
		// future support for area files
		// https://github.com/alexmchale/merc-mud/blob/master/doc/area.txt
		w.Write([]byte("#AREA	{ 5 35} Eilidh    The Coffeehouse~\n"))
	case "dance":
		w.Write([]byte("shake your booty.\n"))
	case "get":
		switch args[0] {
		case "coffee":
			w.Write([]byte("you get warm coffee in a fresh mug.\n"))
		case "mug":
			w.Write([]byte("you grab an empty mug, but you don't have any pockets.\n"))
		case "crate":
			w.Write([]byte("it's too heavy.\n"))
		default:
			w.Write([]byte("get what?\n"))
		}
	default:
		w.Write([]byte("not possible.\n"))
	}
	c.Mutex.Unlock()
	return false
}
