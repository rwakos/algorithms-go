package main

/*
	Create a function that returns true if the strings contains unique Ascii letters
*/
func hasUniquesCharacters(s string) bool {
	if len(s) > 128 {
		return false //This means that it contains duplicates
	}

	m := map[string]bool{}

	for _, v := range s {
		_, ok := m[string(v)]

		if ok {
			return false
		}
		m[string(v)] = true
	}

	return true
}
