package algorithms

import (
	"bytes"
	"sort"
	"strconv"
	"strings"
)

// function to build alphabet
func buildAlphabet(input []rune) []rune {
	alphabetMap := make(map[rune]bool)

	// add all unique characters
	for _, char := range input {
		alphabetMap[char] = true
	}

	// add characters to the slice
	alphabet := make([]rune, 0, len(alphabetMap))
	for char := range alphabetMap {
		alphabet = append(alphabet, char)
	}

	// sort alphabet
	sort.Slice(alphabet, func(i, j int) bool {
		return alphabet[i] < alphabet[j]
	})

	return alphabet
}

// linear search character in alphabet
func search(c rune, alphabet []rune) int {
	for i, char := range alphabet {
		if char == c {
			return i
		}
	}

	return -1
}

// Move to Front transform
func MTF(input string) ([]int, []rune) {
	runes := []rune(input)
	alphabet := buildAlphabet(runes)

	result := make([]int, len(runes))

	for i, char := range runes {
		// find unicode character in alphabet
		result[i] = search(char, alphabet)

		// move this character to the front
		for j := result[i]; j > 0; j-- {
			alphabet[j], alphabet[j-1] = alphabet[j-1], alphabet[j]
		}
	}

	// Sorting alphabet after transformation
	// to prevent errors in inverse transformation
	sort.Slice(alphabet, func(i, j int) bool {
		return alphabet[i] < alphabet[j]
	})

	return result, alphabet
}

// inverse Move to Front transform
func InverseMTF(indexes []int, alphabet []rune) string {
	var result bytes.Buffer

	// take char by index and
	for _, idx := range indexes {
		char := alphabet[idx]
		result.WriteRune(char)
		copy(alphabet[1:], alphabet[:idx])
		alphabet[0] = char
	}

	return result.String()
}

func ConvertToString(transformed []int) string {
	strArr := make([]string, len(transformed))

	for i, num := range transformed {
		strArr[i] = strconv.Itoa(num)
	}

	return strings.Join(strArr, "")
}
