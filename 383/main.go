func canConstruct(ransomNote string, magazine string) bool {
	magazineCharMap := make(map[byte]int, 0)
	for i := 0; i < len(magazine); i++ {
		magazineCharMap[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		if v, ok := magazineCharMap[ransomNote[i]]; !ok || v < 1 {
			return false
		}
		magazineCharMap[ransomNote[i]]--		
	}

	return true
}
