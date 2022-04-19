package gomud

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

const SPLIT_CHAR string = ` `
const PLAYER_SAVE_EXT string = `.save`

func getCharacterData(name string) ([][]byte, error) {
	var d [][]byte
	loc, err := findPlayerFiles()			// loc is absolute filepath ending with slash
	if err != nil {
		return d, err
	}
	f := loc + name + PLAYER_SAVE_EXT
	fh, err := os.OpenFile(f, os.O_RDONLY, 0)
	if err != nil {
		return d, err
	}
	defer fh.Close()
	s := bufio.NewScanner(fh)
	for s.Scan() {
		b := s.Bytes()
		d = append(d, b)
	}
	return d, nil
}

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

func matches(reg, text string) bool {
	m, err := regexp.Match(reg, []byte(text))
	if err != nil {
		return false
	}
	return m
}

func findPlayerFiles() (string, error) {
	var playersPath string
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	ss := strings.Split(cwd, "/")
	switch {
	case ss[len(ss)-1] == "adventure":
		playersPath = strings.Join(ss[:len(ss)-1], "/") + "/players/"
		return playersPath, nil
	case ss[len(ss)-1] == "gomud":
		playersPath = cwd + "/players/"
		return playersPath, nil
	default:
		return cwd, fmt.Errorf("%s: players directory not found.", cwd)
	}
}

func ParseObjects(data []byte) Tree {
	mugName := "a mug of coffee"
	mugData := packageBytes([]string{"#1",mugName,"A mug of coffee lies here."})
	objs := NewTree(NewNodeByName(mugName, *mugData))
	leafName := "a leaf"
	leafData := packageBytes([]string{"#3",leafName,"A dewy leaf lies here."})
	objs.Add(NewNodeByName(leafName, *leafData))
	newObjectPattern := `^#[0-9]*$` // object index
	newObjectRegex := regexp.MustCompile(newObjectPattern)
	lines := strings.Split(string(data), "\n")
	for index, l := range lines {
		if newObjectRegex.MatchString(l) {
			// from location get line 0, 2, 3
			// create object Node{Data: []string{#<objnum>,<short desc>,<long desc>}, Name: <short desc>}
			id := lines[index]
			short := lines[index+2]
			long := lines[index+3]
			data := packageBytes([]string{id,short,long})
			objs.Add(NewNodeByName(short, *data))
		}
	}
	return objs
}
