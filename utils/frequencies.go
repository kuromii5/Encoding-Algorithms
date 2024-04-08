package utils

import "github.com/emirpasic/gods/maps/linkedhashmap"

func CountFrequencies(s string) map[rune]int {
	freqs := make(map[rune]int)

	for _, char := range s {
		freqs[char]++
	}

	return freqs
}

func CountFrequenciesSorted(s string) *linkedhashmap.Map {
	m := linkedhashmap.New()

	for _, char := range s {
		value, found := m.Get(char)
		if found {
			count := value.(int) + 1
			m.Put(char, count)
		} else {
			m.Put(char, 1)
		}
	}

	return m
}
