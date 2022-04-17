package gomud

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func WriteToPlayer(w io.Writer, m string) (int, error) {
	return w.Write([]byte(m))
}

func LoginWithOS(areasPath string) (Character, error) {
	return Login(os.Stdin, os.Stdout, areasPath)
}

func Login(r io.Reader, w io.Writer, areasPath string) (Character, error) {
	_, err := WriteToPlayer(w, "Hello adventurer. What is your name?\n")
	if err != nil {
		return Character{}, err
	}
	s := bufio.NewScanner(r)
	s.Scan()
	name := s.Text()
	c, err := NewCharacter(name, r, w)
	if err != nil {
		return Character{}, err
	}
	al, err := BuildAreaList(areasPath)
	if err != nil {
		return Character{}, err
	}
	c.Arealist = al
	return c, nil
}

func (c *Character) ClassPrompt(r io.Reader, w io.Writer) error {
	msg := fmt.Sprintf("Please select a class %s: [fighter], mage, cleric, rogue\n", c.Name)
	c.Mutex.Lock()
	count, err := WriteToPlayer(w, msg)
	if err != nil {
		return err
	}
	if count != len(msg) {
		return fmt.Errorf("ClassPrompt count mismatch: sent: %d recvd: %d", len(msg), count)
	}
	s := bufio.NewScanner(r)
	s.Scan()
	class := s.Text()
	c.Mutex.Unlock()
	c.Class = classHandler(class)
	return nil
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
	errorHandler <- fmt.Errorf("%s the %s has arrived.", c.Name, c.Class)
	msg := "\nWhat would you like to do?\n"
	for {
		c.Mutex.Lock()
		count, err := WriteToPlayer(w, msg)
		if err != nil {
			errorHandler <- err
		}
		if count != len(msg) {
			errorHandler <- fmt.Errorf("Write mismatch Prompt->WriteToPlayer: sent: %d recvd: %d", len(msg), count)
		}
		s := bufio.NewScanner(r)
		s.Scan()
		c.Action = s.Text()
		c.Mutex.Unlock()
		if q := c.Do(r, w, errorHandler); q == true {
			quit <- true
			break
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
		args = []string{"none"}
	}
	return cmd, args
}

// Do returns true for quit, and false for any other command
func (c *Character) Do(r io.Reader, w io.Writer, errorHandler chan error) bool {
	command, args := commandHandler(c.Action)
	c.Mutex.Lock()
	switch command {
	case "quit":
		msg := fmt.Sprintf("%s once again fades away into the mists of time.\n", c.Name)
		_, err := WriteToPlayer(w, msg)
		if err != nil {
			errorHandler <- err
		}
		c.Mutex.Unlock()
		return true
	case "help":
		_, err := WriteToPlayer(w, "you can: areas,get,help,inventory,look,quit,stats\n")
		if err != nil {
			errorHandler <- err
		}
	case "stats":
		msg := fmt.Sprintf("%s the %s\n%s\n", c.Name, c.Class, c.Stats.Text())
		_, err := WriteToPlayer(w, msg)
		if err != nil {
			errorHandler <- err
		}
	case "inventory":
		_, err := WriteToPlayer(w, "you ain't got shit.\n")
		if err != nil {
			errorHandler <- err
		}
	case "look":
		switch args[0] {
		case "none":
			_, err := WriteToPlayer(w,
			  "There's a tree. it doesn't move much. There's a wooden crate under the tree, and a pot of coffee with clean mugs.\n")
			if err != nil {
				errorHandler <- err
			}
		case "tree":
			_, err := WriteToPlayer(w,
			  "An old and sturdy tree stands here. Its knotty roots sink deeply into the ground and its broad, leafy branches tame the wind.\n")
			if err != nil {
				errorHandler <- err
			}
		case "coffee":
			_, err := WriteToPlayer(w, "A shiny pot seems to have an endless supply of coffee. Clean, teal mugs wait to be filled.\n")
			if err != nil {
				errorHandler <- err
			}
		case "crate":
			_, err := WriteToPlayer(w, "An stained crate decorated with a tablecloth serves as a surface for serving coffee.\n")
			if err != nil {
				errorHandler <- err
			}
		}
	case "areas":
		cur := c.Arealist.Head
		for cur.Next != nil {
			msg := fmt.Sprintf("%2d %s", cur.Id, cur.Title)
			_, err := WriteToPlayer(w, msg)
			if err != nil {
				errorHandler <- err
			}
			cur = cur.Next
		}
		msg := fmt.Sprintf("%2d %s", cur.Id, cur.Title)
		_, err := WriteToPlayer(w, msg)
		if err != nil {
			errorHandler <- err
		}
	case "dance":
		_, err := WriteToPlayer(w, "shake your booty.\n")
		if err != nil {
			errorHandler <- err
		}
	case "get":
		switch args[0] {
		case "coffee":
			_, err := WriteToPlayer(w, "you get warm coffee in a fresh mug.\n")
			if err != nil {
				errorHandler <- err
			}
		case "mug":
			_, err := WriteToPlayer(w, "you grab an empty mug, but you don't have any pockets.\n")
			if err != nil {
				errorHandler <- err
			}
		case "crate":
			_, err := WriteToPlayer(w, "it's too heavy.\n")
			if err != nil {
				errorHandler <- err
			}
		default:
			_, err := WriteToPlayer(w, "get what?\n")
			if err != nil {
				errorHandler <- err
			}
		}
	default:
		_, err := WriteToPlayer(w, "not possible.\n")
		if err != nil {
			errorHandler <- err
		}
	}
	c.Mutex.Unlock()
	return false
}
