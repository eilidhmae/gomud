package gomud

import "testing"

func TestPackageBytes(t *testing.T) {
	msg := []string{"this is the song that never ends.",
			"it just goes on and on my friends.",
			"some people started singing it not knowing what it was",
			"and they'll continue singing it forever just because"}
	expected := []byte("this is the song that never ends.\nit just goes on and on my friends.\nsome people started singing it not knowing what it was\nand they'll continue singing it forever just because")

	p := packageBytes(msg)
	if p == nil {
		t.Errorf("packageBytes returned nil")
	}
	if len(*p) != len(expected) {
		t.Errorf("packageBytes len mismatch: got: %d expected: %d", len(*p), len(expected))
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

func TestGetRandomStat(t *testing.T) {
	low := 1
	high := 20
	for i := 1; i <= 100; i++ {
		result := getRandomStat(low, high)
		if result < low {
			t.Errorf("result too low")
		}
		if result > high {
			t.Errorf("result too high")
		}
	}

	flipped := getRandomStat(high, low)
	if flipped < low {
		t.Errorf("flipped too low")
	}
	if flipped > high {
		t.Errorf("flipped too high")
	}
}

func TestJoinArgs(t *testing.T) {
	args := []string{"one","two","three"}
	expected := "one two three"
	got := joinArgs(args)
	if got != expected {
		t.Errorf("joinArgs mismatch: got: %s expected: %s", got, expected)
	}
}

func TestMatches(t *testing.T) {
	reg := "free"
	text := "all free people"
	if matches(reg, text) != true {
		t.Errorf("matches did not match.")
	}
}
