package gomud

import (
	"strings"
	"testing"
)

func TestLoginWithReader(t *testing.T) {
	defer quiet()()
	input := "david"
	c := LoginWithReader(strings.NewReader(input))
	if c.Name != input {
		t.Errorf("Login mismatch")
	}
}

func TestClassPromptWithReader(t *testing.T) {
	defer quiet()()
	var c Character
	c.Name = "martin"
	input := "cleric"
	c.ClassPromptWithReader(strings.NewReader(input))
	if c.Class != input {
		t.Errorf("ClassPrompt mismatch")
	}
}

func TestClassHandler(t *testing.T) {
	defer quiet()()
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
