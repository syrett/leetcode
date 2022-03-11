package main

func replaceSpace(s string) string {
	t := []byte(s)
	r := []byte{}

	for i := 0; i < len(t); i++ {
		if t[i] == ' ' {
			r = append(r, []byte("%20")...)
		} else {
			r = append(r, t[i])
		}
	}

	return string(r)
}
