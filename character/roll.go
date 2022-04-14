package gomud

import (
	"time"
	"math/rand"
)

type Stats struct {
	Raw	[]int
}

func RollStats(min, max int) Stats {
	var s Stats

	for i := 0; i < 6; i++ {
		roll := getRandomStat(min, max)
		s.Raw = append(s.Raw, roll)
	}
	return s
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
