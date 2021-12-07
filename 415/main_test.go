package main

import ("testing")

func TestAddStrings(t *testing.T) {
	type testCase struct {
		num1 string
		num2 string
		result string
	}

	caseList := []testCase{
    	testCase{"11","123", "134"},
	}

	for _, item := range caseList{
		result := addStrings(item.num1, item.num2)
		if result != item.result {
			t.Errorf("Error:%s!=%s", result, item.result)
		}
	}
}