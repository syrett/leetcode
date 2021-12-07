package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	type Case struct {
		input  []string
		result string
	}

	caseList := []Case{
		Case{[]string{"flower", "flow", "flight"}, "fl"},
		Case{[]string{"", "", ""}, ""},
	}

	for _, c := range caseList {
		result := longestCommonPrefix(c.input)
		if result != c.result {
			t.Errorf("FAIL:%v %s!=%s", c.input, c.result, result)
		}
	}
}
