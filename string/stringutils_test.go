package stringutils

import (
	"testing"
)

func TestMatches(t *testing.T) {
	testPattern := "[0-9]{2}_\\w+"
	testString := "12_abcd"

	matched, err := Matches(testPattern, testString)
	if err != nil {
		t.Error("Error...", err)
	}
	if !matched {
		t.Errorf("Expected matched to be true, but was false."+
			"testPattern=%v, testString=%v", testPattern, testString)
	}
}

func TestDoesntMatch(t *testing.T) {
	testPattern := "[0-9]{2}_\\w+"
	testString := "abcd_12"

	matched, err := Matches(testPattern, testString)
	if err != nil {
		t.Error("Error...", err)
	}
	if matched {
		t.Errorf("Expected matched to be false, but was true."+
			"testPattern=%v, testString=%v", testPattern, testString)
	}
}
