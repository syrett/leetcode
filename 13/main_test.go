package main

import "testing"

func TestRomanToInt(t *testing.T) {
	type Case struct {
		roman  string
		result int
	}

	caseList := []Case{
		Case{"III", 3},
		Case{"IV", 4},
		Case{"IX", 9},
		Case{"LVIII", 58},
		Case{"MCMXCIV", 1994},
	}

	for _, c := range caseList {
		result := romanToInt(c.roman)
		if result != c.result {
			t.Errorf("FAIL:%s %d!=%d", c.roman, c.result, result)
		}
	}
}
