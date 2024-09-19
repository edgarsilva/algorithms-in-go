package groupanagrams

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	groupMap := make(map[string][]string, len(strs))

	for i := 0; i < len(strs); i++ {
		s := strs[i]
		ss := strings.Split(s, "")
		sort.Strings(ss)
		key := strings.Join(ss, "")
		groupMap[key] = append(groupMap[key], s)
	}

	result := make([][]string, 0, len(groupMap))
	for _, v := range groupMap {
		result = append(result, v)
	}

	return result
}
