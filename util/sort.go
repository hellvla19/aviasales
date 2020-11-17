package util

import "sort"

// SortString сортируем строку по символам с помощью рун. Если используется только ASCII, можно работать через []byte.
func SortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
