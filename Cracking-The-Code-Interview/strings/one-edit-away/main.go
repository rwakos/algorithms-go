package main

func main() {

}

func oneEditAway(original string, replacebale string) bool {
	if len(original) == len(replacebale) {
		return oneEditReplace(original, replacebale)
	} else if len(original)+1 == len(replacebale) {
		return oneEditInsert(original, replacebale)
	} else if len(original)-1 == len(replacebale) {
		return oneEditInsert(replacebale, original)
	}

	return false
}

func oneEditReplace(s1 string, s2 string) bool {
	flag := false
	for i, v := range s1 {
		if string(v) != string(s2[i]) {
			if flag {
				return false
			}
			flag = true
		}
	}
	return true
}

func oneEditInsert(s1 string, s2 string) bool {
	index1 := 0
	index2 := 0

	for index2 < len(s2) && index1 < len(s1) {
		if string(s1[index1]) != string(s2[index2]) {
			if index1 != index2 {
				return false
			}
			index2++
		} else {
			index1++
			index2++
		}
	}

	return true
}
