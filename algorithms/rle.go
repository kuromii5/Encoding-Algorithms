package algorithms

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

// Optimized version of RLE which
// encodes UTF-8 strings properly
// and doesn't exceed size of singly
// positioned characters
func RLEncodeUTF8(input []byte) []byte {
	var encoded bytes.Buffer
	runes := []rune(string(input))
	length := len(runes)

	for i := 0; i < length; {
		char := runes[i]
		size := utf8.RuneLen(char)

		count := 1

		for i+1 < length && runes[i+1] == char {
			count++
			i++
		}
		i++

		if count > 1 {
			charBytes := make([]byte, size)
			utf8.EncodeRune(charBytes, char)
			if unicode.IsDigit(char) {
				encoded.WriteString(fmt.Sprintf("\x00%d\x00%s", count, string(charBytes)))
			} else {
				encoded.WriteString(fmt.Sprintf("%d%s", count, string(charBytes)))
			}
		} else {
			if unicode.IsDigit(char) {
				encoded.WriteString("\x01")
				encoded.WriteRune(char)
			} else {
				encoded.WriteRune(char)
			}
		}
	}

	return encoded.Bytes()
}

// Decoding algorithm
func RLDecodeUTF8(input []byte) []byte {
	var decoded bytes.Buffer

	for i := 0; i < len(input); {
		count := 0

		if input[i] == '\x00' {
			i++
			for isDigit(input[i]) {
				count = count*10 + int(input[i]-'0')
				i++
			}
			i++
		} else if input[i] == '\x01' {
			i++
		} else {
			for isDigit(input[i]) {
				count = count*10 + int(input[i]-'0')
				i++
			}
		}

		if count == 0 {
			decoded.WriteByte(input[i])
			i++
		} else {
			char, size := utf8.DecodeRune(input[i:])
			i += size

			for j := 0; j < count; j++ {
				decoded.WriteRune(char)
			}
		}
	}

	return decoded.Bytes()
}

// util function to check
// if we are in digit pos
func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
