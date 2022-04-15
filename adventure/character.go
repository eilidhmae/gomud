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
	Reader		io.Reader
	Writer		io.Writer
	Mutex		sync.Mutex
}

func NewCharacter(name string, r io.Reader, w io.Writer) Character {
	min := 8
	max := 18
	var c Character
	c.SetName(name)
	c.Reader = r
	c.Writer = w
	c.ClassPrompt(r, w)
	c.RollStats(min, max)
	return c
}

func (c *Character) SetName(name string) {
	c.Name = name
}
