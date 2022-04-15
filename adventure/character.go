package gomud

type Character struct {
	Name		string
	Class		string
	Stats		Stats
}

func NewCharacter(name string) Character {
	min := 8
	max := 18
	var c Character
	c.SetName(name)
	c.ClassPrompt()
	c.RollStats(min, max)
	return c
}

func (c *Character) SetName(name string) {
	c.Name = name
}
