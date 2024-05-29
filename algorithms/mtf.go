package algorithms

import (
	"sort"
)

var alphabet []byte

func setByteAlphabet() {
	alphabet = make([]byte, 256)
	for i := 0; i < len(alphabet); i++ {
		alphabet[i] = byte(i)
	}
}

// function to build alphabet
func buildCustomAlphabet(input []rune) []rune {
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
func search(inputByte byte) byte {
	for i, b := range alphabet {
		if b == inputByte {
			return byte(i)
		}
	}

	return 0
}

func moveToFront(byteIndex byte) {
	toFront := alphabet[byteIndex]
	for j := byteIndex; j > 0; j-- {
		alphabet[j] = alphabet[j-1]
	}
	alphabet[0] = toFront
}

func MtFTransform(input []byte) []byte {
	setByteAlphabet()

	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = search(input[i])
		moveToFront(output[i])
	}

	return output
}

func MtFInverse(input []byte) []byte {
	setByteAlphabet()

	output := make([]byte, len(input))
	for j := 0; j < len(input); j++ {
		output[j] = alphabet[input[j]]
		moveToFront(input[j])
	}

	return output
}

// Move to Front transform
// func MTF(input string) ([]int, []rune) {
// 	runes := []rune(input)
// 	alphabet := buildAlphabet(runes)

// 	result := make([]int, len(runes))

// 	for i, char := range runes {
// 		// find unicode character in alphabet
// 		result[i] = search(char, alphabet)

// 		// move this character to the front
// 		for j := result[i]; j > 0; j-- {
// 			alphabet[j], alphabet[j-1] = alphabet[j-1], alphabet[j]
// 		}
// 	}

// 	// Sorting alphabet after transformation
// 	// to prevent errors in inverse transformation
// 	sort.Slice(alphabet, func(i, j int) bool {
// 		return alphabet[i] < alphabet[j]
// 	})

// 	return result, alphabet
// }

// // inverse Move to Front transform
// func InverseMTF(indexes []int, alphabet []rune) string {
// 	var result bytes.Buffer

// 	// take char by index and
// 	for _, idx := range indexes {
// 		char := alphabet[idx]
// 		result.WriteRune(char)
// 		copy(alphabet[1:], alphabet[:idx])
// 		alphabet[0] = char
// 	}

// 	return result.String()
// }
