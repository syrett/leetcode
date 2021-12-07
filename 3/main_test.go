package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	type Case struct {
		str    string
		maxLen int
	}

	caseList := []Case{
		Case{"abcabcbb", 3},
		Case{"bbbb", 1},
		Case{"", 0},
		Case{"abcbcdde", 3},
		Case{"aabaab!bb", 3},
	}

	for _, c := range caseList {
		result := LengthOfLongestSubstring2(c.str)
		if result != c.maxLen {
			t.Errorf("Failed:%s, %d!=%d", c.str, c.maxLen, result)
			t.Fail()
		} else {
			t.Logf("SUCCESS:%s, %d", c.str, result)
		}

	}
}
