package algorithms

import "bytes"

// Represent the return type
// of LZ77 encoding
type LZ77tuple struct {
	Offset int
	Length int
	Next   rune
}

func LZ77Encode(input string, windowSize int) []LZ77tuple {
	var compressed []LZ77tuple
	runes := []rune(input)

	for i := 0; i < len(runes); {
		bestOffset := 0
		bestLength := 0

		for j := 1; j <= windowSize && i-j >= 0; j++ {
			length := 0
			for length < len(runes)-i && runes[i-length] == runes[i-length-j] {
				length++
			}
			if length > bestLength {
				bestOffset = j
				bestLength = length
			}
		}

		var next rune
		if i+bestLength < len(runes) {
			next = runes[i+bestLength]
		} else {
			next = 0
		}

		token := LZ77tuple{Offset: bestOffset, Length: bestLength, Next: next}
		compressed = append(compressed, token)

		i += bestLength + 1
	}

	return compressed
}

func LZ77Decode(compressed []LZ77tuple) string {
	var decompressed bytes.Buffer

	for _, token := range compressed {
		if token.Length == 0 {
			decompressed.WriteRune(token.Next)
		} else {
			start := decompressed.Len() - token.Offset
			end := start + token.Length
			runes := []rune(decompressed.String())
			decompressed.WriteString(string(runes[start:end]))
			decompressed.WriteRune(token.Next)
		}
	}

	return decompressed.String()
}
