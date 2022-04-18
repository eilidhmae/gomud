package gomud

import (
	"os"
	"log"
)

func quiet() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr

	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)

	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}

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
