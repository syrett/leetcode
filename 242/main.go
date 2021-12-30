package main

func isAnagram(s string, t string) bool {
	keyMap := make(map[rune]int)
	for _, v := range s {
		if i, ok := keyMap[v]; ok {
			keyMap[v] = i + 1
		} else {
			keyMap[v] = 1
		}
	}

	for _, v := range t {
		if i, ok := keyMap[v]; ok {
			keyMap[v] = i - 1
		} else {
			return false
		}
	}

	for _, count := range keyMap {
		if count != 0 {
			return false
		}
	}

	return true
}
