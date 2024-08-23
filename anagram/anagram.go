package anagram

type empty struct{}

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// ss := strings.Split(s, "")
	// sort.Strings(ss)
	// sc := strings.Join(ss, "")
	//
	// ts := strings.Split(t, "")
	// sort.Strings(ts)
	// tc := strings.Join(ts, "")

	freq := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		freq[rune(s[i])]++
		freq[rune(t[i])]--
	}

	for _, v := range freq {
		if v != 0 {
			return false
		}
	}

	return true
}
