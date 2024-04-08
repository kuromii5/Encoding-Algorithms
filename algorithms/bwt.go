package algorithms

import (
	"EncodingAlgorithms/utils"
	"sort"
)

type Pair struct {
	Rune  rune
	Index int
}

// Naive BWT implementation using string matrix
// Time Complexity O(N^2) ; Space Complexity O(N^2)
func NaiveBWT(data string) (int, string) {
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

	return slices[dataPos]
}

func bwt(sa []int, data string) (int, string) {
	var index int
	for i, v := range sa {
		if v == 0 {
			index = i
			break
		}
	}

	res := utils.BWTLastColumn(data, sa)
	return index, res
}

// function using naive building of suffix array
// Time Complexity O(N) ; Space Complexity O(N)
func SuffixArrayBWT(data string) (int, string) {
	sa := utils.NaiveSuffixArray(data)
	return bwt(sa, data)
}

//! Not working
// Effective function using suffix array with induce sorting
// Time Complexity O(N) ; Space Complexity O(N)
// func FastSuffixArrayBWT(data string) (int, string) {
// 	sa := utils.SAbyIS(data, 256)
// 	return bwt(sa, data)
// }

// Effective inverse function using cycling permutations
// Time Complexity O(N*logN) ; Space Complexity O(N)
func InverseBWT(pos int, data string) string {
	dataPairs := make([]Pair, len(data))
	for i, char := range data {
		dataPairs[i] = Pair{Rune: char, Index: i}
	}

	sortedPairs := make([]Pair, len(data))
	copy(sortedPairs, dataPairs)
	sort.Slice(sortedPairs, func(i, j int) bool {
		return sortedPairs[i].Rune < sortedPairs[j].Rune
	})

	var result []rune
	for i := 0; i < len(data); i++ {
		result = append(result, sortedPairs[pos].Rune)
		pos = sortedPairs[pos].Index
	}

	return string(result)
}
