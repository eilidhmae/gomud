package gomud

import (
	"reflect"
)

const AREAS_PATH string = `../areas/`

var TestRealm Realm

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

func initializeTestRealm() {
	if reflect.DeepEqual(TestRealm, Realm{}) {
		r, err := BuildRealm(AREAS_PATH)
		if err == nil {
			TestRealm = r
		}
	}
}
