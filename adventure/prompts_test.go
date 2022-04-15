package gomud

import "testing"

func TestParseClassPrompt(t *testing.T) {
	testTrim := "  mage  "
	testClassPatterns := []string{"rogue","mage","cleric","fighter"}
	testAntiPatterns := []string{"blaster","vogue","","12345","two words"}

	expectedTrim := "mage"
	gotTrim := parseClassPrompt(testTrim)
	if gotTrim != expectedTrim {
		t.Errorf("unexpected trim: %s", gotTrim)
	}

	for _, c := range testClassPatterns {
		expected := c
		got := parseClassPrompt(c)
		if expected != got {
			t.Errorf("mismatch: expected: %s got: %s", expected, got)
		}
	}

	for _, c := range testAntiPatterns {
		expected := "fighter"
		got := parseClassPrompt(c)
		if expected != got {
			t.Errorf("unexpected result: gave: %s got: %s expected: %s", c, got, expected)
		}
	}
}
