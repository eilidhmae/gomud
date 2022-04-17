package gomud

import (
	"math/rand"
	"strings"
	"time"
)

const SPLIT_CHAR string = ` `

func packageBytes(lines []string) *[]byte {
	b := []byte(strings.Join(lines, "\n"))
	return &b
}

func joinArgs(args []string) string {
	return strings.Join(args, SPLIT_CHAR)
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

func commandHandler(action string) (string, []string) {
	var args []string
	w := strings.Split(action, SPLIT_CHAR)
	cmd := w[0]
	if len(w) > 1 {
		args = w[1:]
	} else {
		args = []string{"none"}
	}
	return cmd, args
}

func getRandomStat(min, max int) int {
	if min > max {
		t := max
		max = min
		min = t
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min + 1) + min
}
