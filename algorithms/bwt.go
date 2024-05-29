package algorithms

import (
	"EncodingAlgorithms/utils"
	"bytes"
	"sort"
	"strings"
)

type Pair struct {
	Rune  rune
	Index int
}

// amount of symbols for 1 bwt string
const BWT_CHUNK_SIZE = 1000

// Naive BWT implementation using string matrix
// Time Complexity O(N^2) ; Space Complexity O(N^2)
func NaiveBWT(data string) (int, string) {
	data += "\x02"
	runes := []rune(data)
	length := len(runes)

	slices := make([]string, length)
	for i := 0; i < length; i++ {
		slices[i] = string(runes[i:]) + string(runes[:i])
	}

	sort.Strings(slices)

	dataPos := 0
	for i, str := range slices {
		if str == data {
			dataPos = i
			break
		}
	}

	result := ""
	for _, str := range slices {
		r := []rune(str)
		result += string(r[length-1])
	}

	return dataPos, result
}

// Naive BWT inverse building initial string matrix
// Time Complexity O(N^2*logN) ; Space Complexity O(N^2)
func NaiveInverseBWT(dataPos int, transformed string) string {
	runes := []rune(transformed)
	length := len(runes)

	slices := make([]string, length)

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			slices[j] = string(runes[j]) + slices[j]
		}
		sort.Strings(slices)
	}

	return slices[dataPos][:len(slices)-1]
}

// function using naive building of suffix array
// Time Complexity O(N^2) ; Space Complexity O(N)
const EOF = -1

func SuffixArrayBWT(data string) string {
	data += "\x02"
	sa := utils.NaiveSuffixArray(data)
	res := utils.BWTLastColumn(data, sa)
	return res
}

// Effective inverse function using cycling permutations
// Time Complexity O(N*logN) ; Space Complexity O(N)
func InverseBWT(data string) string {
	runes := []rune(data)
	dataPairs := make([]Pair, len(runes))
	for i, char := range runes {
		dataPairs[i] = Pair{Rune: char, Index: i}
	}

	sortedPairs := make([]Pair, len(runes))
	copy(sortedPairs, dataPairs)
	sort.Slice(sortedPairs, func(i, j int) bool {
		if sortedPairs[i].Rune == sortedPairs[j].Rune {
			return sortedPairs[i].Index < sortedPairs[j].Index
		}
		return sortedPairs[i].Rune < sortedPairs[j].Rune
	})

	var result []rune
	pos := sortedPairs[0].Index
	for i := 1; i < len(runes); i++ {
		result = append(result, sortedPairs[pos].Rune)
		pos = sortedPairs[pos].Index
	}

	return string(result)
}

func MakeBWTString(data []byte) []byte {
	s := string(data)
	runes := []rune(s)

	var buffer bytes.Buffer
	for i := 0; i < len(runes); i += BWT_CHUNK_SIZE {
		end := i + BWT_CHUNK_SIZE
		if end > len(runes) {
			end = len(runes)
		}
		chunk := runes[i:end]
		transformed := SuffixArrayBWT(string(chunk))
		buffer.WriteString(transformed)
	}

	return buffer.Bytes()
}

func InverseTextBWT(data []byte) string {
	s := string(data)
	runes := []rune(s)

	var decodedStringsArray []string
	for i := 0; i < len(runes); i += BWT_CHUNK_SIZE + 1 {
		end := i + BWT_CHUNK_SIZE + 1
		if end > len(runes) {
			end = len(runes)
		}
		chunk := runes[i:end]

		inversed := InverseBWT(string(chunk))
		decodedStringsArray = append(decodedStringsArray, inversed)
	}

	return strings.Join(decodedStringsArray, "")
}

//! Not working
// Effective function using suffix array with induce sorting
// Time Complexity O(N) ; Space Complexity O(N)
// func FastSuffixArrayBWT(data string) (int, string) {
// 	sa := utils.SAbyIS(data, 256)
// 	return bwt(sa, data)
// }
