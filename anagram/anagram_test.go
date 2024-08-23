package anagram

import "testing"

func TestAnagram(t *testing.T) {
	testcases := []struct {
		s    string
		t    string
		want bool
	}{
		{"racecar", "carrace", true},
		{"jar", "jam", false},
		{"heart", "earth", true},
	}

	for _, tc := range testcases {
		got := IsAnagram(tc.s, tc.t)

		if got != tc.want {
			t.Errorf("got %#v want %#v", got, tc.want)
		}
	}
}
