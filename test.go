package main

import (
	"fmt"
	"sort"
)

func main() {
	res := rollUp("aaabbbccccc")
	res1 := rollUp("aaabbbcccccaaaaa")
	res2 := rollUp("zzzzcccUUUuu")
	res3 := rollUp("ЯЯЯБББддд")

	fmt.Println(res)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}

func rollUp(s string) string {
	var result string

	m := make(map[rune]int)
	str := []rune(s)

	for _, v := range str {
		m[v] += 1
	}

	keys := make([]rune, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, v := range keys {
		res += fmt.Sprintf("%c%d", v, m[v])
	}

	return res
}
