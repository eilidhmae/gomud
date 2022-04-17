package gomud

import (
	"io"
	"sync"
)

type Character struct {
	Name		string
	Class		string
	Stats		Stats
	Action		string
	Mutex		sync.Mutex
	Arealist	Arealist
}

func NewCharacter(name string, r io.Reader, w io.Writer) (Character, error) {
	min := 8
	max := 18
	var c Character
	c.SetName(name)
	if err := c.ClassPrompt(r, w); err != nil {
		return Character{}, err
	}
	c.RollStats(min, max)
	return c, nil
}

func (c *Character) SetName(name string) {
	c.Name = name
}
