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
	c, err := Login(strings.NewReader(r), w)
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

func TestClassHandler(t *testing.T) {
	testTrim := "  mage  "
	testClassPatterns := []string{"rogue","mage","cleric","fighter"}
	testAntiPatterns := []string{"blaster","vogue","","12345","two words"}

	expectedTrim := "mage"
	gotTrim := classHandler(testTrim)
	if gotTrim != expectedTrim {
		t.Errorf("unexpected trim: %s", gotTrim)
	}

	for _, c := range testClassPatterns {
		expected := c
		got := classHandler(c)
		if expected != got {
			t.Errorf("mismatch: expected: %s got: %s", expected, got)
		}
	}

	for _, c := range testAntiPatterns {
		expected := "fighter"
		got := classHandler(c)
		if expected != got {
			t.Errorf("unexpected result: gave: %s got: %s expected: %s", c, got, expected)
		}
	}
}

func TestCommandHandler(t *testing.T) {
	expectedCmd := "look"
	expectedArgs := []string{"at","me"}
	action := "look at me"

	cmd, args := commandHandler(action)
	if expectedCmd != cmd {
		t.Errorf("commandHandler mismatch: expectedCmd: %s cmd: %s", expectedCmd, cmd)
	}
	if len(expectedArgs) != len(args) {
		t.Errorf("args size mismatch")
	}
	for i := 0; i < len(args); i++ {
		if args[i] != expectedArgs[i] {
			t.Errorf("args mismatch: expected: %s got: %s", expectedArgs[i], args[i])
		}
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
