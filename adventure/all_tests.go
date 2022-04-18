package gomud

func initializeTestCharacter() Character {
	return Character{
		Name:		"tester",
		Class:		"mage",
		Stats:		Stats{
			Raw:	[]int{15,14,13,12,10,10},
			Str:	10,
			Dex:	12,
			Wis:	13,
			Int:	15,
			Con:	14,
			Cha:	10,
		},
		Cursor:		"tester the mage-\u003e ",
		Level:		2,
		CanSave:	true,
	}
}
