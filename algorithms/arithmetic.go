package algorithms

import (
	"EncodingAlgorithms/utils"
)

// 10 MB of information
const AC_CHUNK_SIZE = 10_000_000

// 80 symbols is upper bound for encode
const STRING_SIZE = 100

// Struct for encoded string
type Encode struct {
	Value      float64
	CodeLength int
}

// Bound type
type Bound struct {
	Lower float64
	Upper float64
}

// building map of probabilities for every unicode character
func buildProbabilities(frequencies map[rune]int, length int) map[rune]float64 {
	probs := make(map[rune]float64)
	for key, value := range frequencies {
		probs[key] = float64(value) / float64(length)
	}

	return probs
}

// calculating map of borders for every unicode character
func calculateBounds(probs map[rune]float64) (map[rune]Bound, map[Bound]rune) {
	borders := make(map[rune]Bound)
	reversedBounds := make(map[Bound]rune)

	var bound = Bound{0.0, 0.0}
	for char, probability := range probs {
		bound.Upper += probability

		if bound.Upper > 1 {
			bound.Upper = 1
		}

		borders[char] = Bound{bound.Lower, bound.Upper}
		reversedBounds[bound] = char
		bound.Lower = bound.Upper
	}

	return borders, reversedBounds
}

// Arithmetic encode  function
func ArithmeticEncodeString(input string, probs map[rune]float64, borders map[rune]Bound) Encode {
	runes := []rune(input)

	res := borders[runes[0]]
	prevRes := res
	count := 1
	for i := 1; i < len(runes); i++ {
		char := runes[i]
		old := res
		new := borders[char]
		prevRes = res

		calculatedBound := Bound{(old.Lower + (old.Upper-old.Lower)*new.Lower), (old.Lower + (old.Upper-old.Lower)*new.Upper)}
		if calculatedBound == prevRes {
			break
		}

		res = calculatedBound
		count++
	}

	avg := (res.Lower + res.Upper) / 2
	return Encode{avg, count}
}

func ArithmeticEncodeChunk(input string) ([]Encode, map[Bound]rune) {
	runes := []rune(input)
	if len(input) == 0 {
		return []Encode{}, make(map[Bound]rune)
	}

	var encodedPairs []Encode
	probs := buildProbabilities(utils.CountFrequencies(input), len(runes))
	borders, reversedBounds := calculateBounds(probs)
	chunkCounter := 0

	for i := 0; i < len(runes); {

		end := i + STRING_SIZE
		if end > len(runes) {
			end = len(runes)
		}

		chunk := string(runes[i:end])
		pair := ArithmeticEncodeString(chunk, probs, borders)
		i += pair.CodeLength

		encodedPairs = append(encodedPairs, pair)
		chunkCounter++
	}

	return encodedPairs, reversedBounds
}

func ArithmeticDecodeString(pair Encode, reversedBounds map[Bound]rune) string {
	var decoded []rune

	for pair.CodeLength > 0 {
		for borders, value := range reversedBounds {
			bound := borders
			if bound.Lower <= pair.Value && pair.Value < bound.Upper {
				decoded = append(decoded, value)

				pair.Value = (pair.Value - bound.Lower) / (bound.Upper - bound.Lower)
				pair.CodeLength--
				break
			}
		}
	}

	return string(decoded)
}

func ArithmeticDecodeChunk(pairs []Encode, reversedBounds map[Bound]rune) string {
	chunkString := ""

	for _, pair := range pairs {
		chunkString += ArithmeticDecodeString(pair, reversedBounds)
	}

	return chunkString
}
