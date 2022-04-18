package gomud

import (
	"fmt"
	"io"
	"os"
	"sync"
	"encoding/json"
)

type Character struct {
	Name		string		`json:"name"`
	Class		string		`json:"class"`
	Stats		Stats		`json:"stats"`
	Action		string		`json:"-"`
	Mutex		sync.Mutex	`json:"-"`
	Arealist	Arealist	`json:"-"`
	Inventory	Objlist		`json:"-"`
	Cursor		string		`json:"cursor"`
	Level		int		`json:"level"`
	CanSave		bool		`json:"can_save"`
}

func NewCharacter(name string, r io.Reader, w io.Writer) (Character, error) {
	min := 8
	max := 18
	var c Character
	c.SetName(name)
	c.SetCursor("What would you like to do?\n")
	c.Level = 1
	c.CanSave = false
	if err := c.ClassPrompt(r, w); err != nil {
		return Character{}, err
	}
	c.RollStats(min, max)
	return c, nil
}

func CharacterExists(name string) bool {
	_, err := LoadCharacter(name)
	if err != nil {
		return false
	}
	return true
}

func LoadCharacter(name string) (Character, error) {
	var c Character

	data, err := getCharacterData(name)
	if err != nil {
		return c, err
	}
	if len(data) == 0 {
		return c, fmt.Errorf("no character data found for %s.", name)
	}
	if err := json.Unmarshal(data[0], &c); err != nil {
		return c, err
	}
	if len(data) > 1 {
		var inventory []string
		var objs Objlist
		if err := json.Unmarshal(data[1], &inventory); err != nil {
			return c, err
		}
		for id, desc := range inventory {
			switch {
			case id == 0:
				objs = NewObjlist(&Object{Id: id, Desc: desc})
			default:
				objs.Add(&Object{Id: id, Desc: desc})
			}
		}
		c.Inventory = objs
	}
	return c, nil
}

func (c *Character) SetName(name string) {
	c.Name = name
}

func (c *Character) SetCursor(cursor string) {
	c.Cursor = cursor
}

func (c *Character) FancyCursor() {
	c.Cursor = c.Name + " the " + c.Class + "-> "
}

func (c *Character) Save() error {
	location, err := findPlayerFiles()
	if err != nil {
		return err
	}
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	fh, err := os.Create(location + c.Name + PLAYER_SAVE_EXT)
	if err != nil {
		return err
	}
	defer fh.Close()
	fh.Write(data)
	fh.Write([]byte("\n"))
	if c.Inventory.Head != nil {
		inventory := []string{}
		cur := c.Inventory.Head
		for cur != nil {
			inventory = append(inventory, cur.Desc)
			cur = cur.Next
		}
		binv, err := json.Marshal(inventory)
		if err != nil {
			return err
		}
		fh.Write(binv)
		fh.Write([]byte("\n"))
	}
	return nil
}
