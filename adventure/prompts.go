package gomud

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func WriteToPlayer(w io.Writer, m string) (int, error) {
	return w.Write([]byte(m))
}

func LoginWithOS(areasPath string) (Character, error) {
	return Login(os.Stdin, os.Stdout, areasPath)
}

func Login(r io.Reader, w io.Writer, areasPath string) (Character, error) {
	var char Character
	_, err := WriteToPlayer(w, "Hello adventurer. What is your name?\n")
	if err != nil {
		return char, err
	}
	s := bufio.NewScanner(r)
	s.Scan()
	name := s.Text()
	if CharacterExists(name) {
		c, err := LoadCharacter(name)
		if err != nil {
			return c, err
		}
		char = c
		_, err = WriteToPlayer(w, fmt.Sprintf("Welcome back %s.\n", c.Name))
		if err != nil {
			return char, err
		}
	} else {
		c, err := NewCharacter(name, r, w)
		if err != nil {
			return c, err
		}
		char = c
	}
	l, err := BuildAreaList(areasPath)
	if err != nil {
		return char, err
	}
	char.Arealist = l
	return char, nil
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

func (c *Character) PromptWithOS(quit chan bool, errorHandler chan error) {
	go c.Prompt(os.Stdin, os.Stdout, quit, errorHandler)
}

func (c *Character) Prompt(r io.Reader, w io.Writer, quit chan bool, errorHandler chan error) {
	errorHandler <- fmt.Errorf("%s the %s has arrived.", c.Name, c.Class)
	for {
		c.Mutex.Lock()
		count, err := WriteToPlayer(w, "\n" + c.Cursor)
		if err != nil {
			errorHandler <- err
			err = nil
		}
		if count != len(c.Cursor) + 1 {
			errorHandler <- fmt.Errorf("custom Cursor len mismatch: sent: %d recvd: %d", len(c.Cursor), count)
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

/* Do returns true for quit, and false for any other command.
 * Currently, many player actions are hard-coded here which means
 * the system is not extensible without a new deployment and restart.
 * Eventually, action logic should be moved to commandHandler().
 * Then commandHandler() should determine types of actions like
 * emotes, skills, look, get, etc. Some very basic items like 'help'
 * and 'commands' could be kept here if they are considered core commands.
 * There's a strong case to even factor out emotes into a separate
 * handler because anything factored out can be driven by a file
 * on the host filesystem. Emotes, actions and skills could be added
 * by modifying a file and using a command within the realm to read
 * those config files again without a restart.
 */
func (c *Character) Do(r io.Reader, w io.Writer, errorHandler chan error) bool {
	command, args := commandHandler(c.Action)
	c.Mutex.Lock()
	switch command {
	case "quit", "q":
		msg := fmt.Sprintf("%s once again fades away into the mists of time.\n", c.Name)
		_, err := WriteToPlayer(w, msg)
		if err != nil {
			errorHandler <- err
			err = nil
		}
		c.Mutex.Unlock()
		return true
	case "help":
		_, err := WriteToPlayer(w, "you can: areas,get,help,inventory,look,prompt,quit,stats\n")
		if err != nil {
			errorHandler <- err
			err = nil
		}
	case "stats":
		msg := fmt.Sprintf("%s the %s. (Level %d)\n%s\n", c.Name, c.Class, c.Level, c.Stats.Text())
		_, err := WriteToPlayer(w, msg)
		if err != nil {
			errorHandler <- err
			err = nil
		}
	case "inventory", "inv", "i":
		if c.Inventory.Head == nil {
			_, err := WriteToPlayer(w, "you ain't got shit.\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
		} else {
			cur := c.Inventory.Head
			_, err := WriteToPlayer(w, "inventory:\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
			for cur != nil {
				_, err := WriteToPlayer(w, fmt.Sprintf("%s\n", cur.Desc))
				if err != nil {
					errorHandler <- err
					err = nil
				}
				if cur.Next != nil {
					cur = cur.Next
				} else {
					cur = nil
				}
			}
		}
	case "drink":
		switch args[0] {
		case "none":
			_, err := WriteToPlayer(w, "drink what?\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
		case "coffee", "mug":
			if c.Inventory.HasObject(args[0]) {
				_, err := WriteToPlayer(w, "you drink coffee from a mug.\n")
				if err != nil {
					errorHandler <- err
					err = nil
				}
				if c.CanSave == false {
					c.CanSave = true
					_, err := WriteToPlayer(w,
						"Congratulations " + c.Name + "\n" +
						"You have gained a level!\n" +
						"You may now save your character with 'save'.\n" +
						"You also get a fancy new prompt!\n")
					c.FancyCursor()
					c.Level = c.Level + 1
					if err != nil {
						errorHandler <- err
						err = nil
					}
				}
			}
		default:
			_, err := WriteToPlayer(w, "you don't seem to have that.\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
		}
	case "look", "l":
		switch args[0] {
		case "none":
			_, err := WriteToPlayer(w,
			  "The Shady Coffeehouse\n" +
			  "There's a tree. it doesn't move much. There's a wooden crate under the tree, and a pot of coffee with clean mugs.\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
		case "tree":
			_, err := WriteToPlayer(w,
			  "An old and sturdy tree stands here. Its knotty roots sink deeply into the ground.\n" +
			  "A leaf flutters to the ground as the large branches attempt to tame the wind.\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
		case "leaf":
			_, err := WriteToPlayer(w, "A dewy leaf lies upon the ground.\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
		case "coffee", "mugs", "mug", "pot":
			_, err := WriteToPlayer(w, "A shiny pot seems to have an endless supply of coffee. Clean, teal mugs wait to be filled.\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
		case "crate":
			_, err := WriteToPlayer(w, "An stained crate decorated with a tablecloth serves as a surface for serving coffee.\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
		}
	case "areas":
		cur := c.Arealist.Head
		for cur.Next != nil {
			msg := fmt.Sprintf("%2d %s", cur.Id, cur.Title)
			_, err := WriteToPlayer(w, msg)
			if err != nil {
				errorHandler <- err
				err = nil
			}
			cur = cur.Next
		}
		msg := fmt.Sprintf("%2d %s", cur.Id, cur.Title)
		_, err := WriteToPlayer(w, msg)
		if err != nil {
			errorHandler <- err
			err = nil
		}
	case "dance":
		_, err := WriteToPlayer(w, "shake your booty.\n")
		if err != nil {
			errorHandler <- err
			err = nil
		}
	case "get":
		switch args[0] {
		case "coffee", "mug":
			_, err := WriteToPlayer(w, "you get warm coffee in a fresh mug.\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
			o := Object{Desc: "a mug of warm coffee"}
			if c.Inventory.Head == nil {
				c.Inventory = NewObjlist(&o)
			} else {
				c.Inventory.Add(&o)
			}
		case "leaf":
			_, err := WriteToPlayer(w, "you pick up a leaf.\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
			o := Object{Desc: "a leaf"}
			if c.Inventory.Head == nil {
				c.Inventory = NewObjlist(&o)
			} else {
				c.Inventory.Add(&o)
			}
		case "crate":
			_, err := WriteToPlayer(w, "it's too heavy.\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
		default:
			_, err := WriteToPlayer(w, "get what?\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
		}
	case "goto":
		index, err := strconv.Atoi(string(args[0]))
		if err != nil {
			errorHandler <- err
			err = nil
		}
		cur := c.Arealist.Lookup(index)
		c.Arealist.Current = cur
		if err := cur.Build(); err != nil {
			errorHandler <- err
			err = nil
		}
		if cur.Rooms != nil {
			msg := fmt.Sprintf("%s\n", cur.Rooms.Data)
			_, err = WriteToPlayer(w, msg)
			if err != nil {
				errorHandler <- err
				err = nil
			}
		}
	case "prompt":
		cursor := joinArgs(args)
		c.SetCursor(cursor)
		if cursor == "fancy" {
			c.FancyCursor()
		}
	case "drop":
		if args[0] == "none" {
			_, err := WriteToPlayer(w, "drop what?\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
			break
		}
		if args[0] == "all" {
			_, err := WriteToPlayer(w, "You drop everything and it all disappears in smoke.\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
			c.Inventory = Objlist{}
			break
		}
		cur := c.Inventory.Head
		for cur != nil {
			m, err := regexp.Match(args[0], []byte(cur.Desc))
			if err != nil {
				errorHandler <- err
				err = nil
			}
			if m {
				if c.Inventory.Count == 1 {
					c.Inventory = Objlist{}
					break
				}
				_ = c.Inventory.Drop(cur.Id)
				_, err := WriteToPlayer(w, fmt.Sprintf("you drop %s on the ground and it dissolves.\n", cur.Desc))
				if err != nil {
					errorHandler <- err
					err = nil
				}
				break
			}
			cur = cur.Next
		}
	case "save":
		if c.CanSave == true {
			err := c.Save()
			if err != nil {
				errorHandler <- err
				err = nil
			}
			errorHandler <- fmt.Errorf("%s has saved.", c.Name)
		} else {
			_, err := WriteToPlayer(w, "not possible.\n")
			if err != nil {
				errorHandler <- err
				err = nil
			}
		}
	default:
		_, err := WriteToPlayer(w, "not possible.\n")
		if err != nil {
			errorHandler <- err
			err = nil
		}
	}
	c.Mutex.Unlock()
	return false
}
