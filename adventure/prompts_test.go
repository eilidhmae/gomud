package gomud

import (
	"os"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	r := "david"
	w, err := os.OpenFile(os.DevNull, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()
	c, err := Login(strings.NewReader(r), w, AREAS_PATH)
	if err != nil {
		t.Error(err)
	}
	if c.Name != r {
		t.Errorf("Login mismatch")
	}
}

func TestClassPrompt(t *testing.T) {
	var c Character
	c.Name = "martin"
	r := "cleric"
	w, err := os.OpenFile(os.DevNull, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()
	c.ClassPrompt(strings.NewReader(r), w)
	if c.Class != r {
		t.Errorf("ClassPrompt mismatch")
	}
}

func TestDo(t *testing.T) {
	w, err := os.OpenFile(os.DevNull, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()
	ec := make(chan error)
	var c Character
	c.Action = "quit"
	if c.Do(strings.NewReader(""), w, ec) != true {
		t.Errorf("quit returned false")
	}
}

func TestWriteToPlayer(t *testing.T) {
	m := "Hello!\n"
	w, err := os.OpenFile(os.DevNull, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	if err !=nil {
		t.Fatal(err)
	}
	defer w.Close()
	c, err := WriteToPlayer(w, m)
	if err != nil {
		t.Error(err)
	}
	if c != len(m) {
		t.Errorf("TestWriteToPlayer count mismatch: sent: %d received %d", len(m), c)
	}
}
