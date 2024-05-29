package utils

import (
	"bytes"
	"fmt"
	"sort"
)

// "S" and "L" types represented
// by 0 and 1 ints
const (
	S_TYPE = 0
	L_TYPE = 1
)

// Naively build suffix array for the given string
func NaiveSuffixArray(s string) []int {
	runes := []rune(s)

	suffixes := make([]int, len(runes))
	for i := range suffixes {
		suffixes[i] = i
	}

	sort.Slice(suffixes, func(i, j int) bool {
		return compareSuffixes(runes, suffixes[i], suffixes[j])
	})

	return suffixes
}

func compareSuffixes(runes []rune, i, j int) bool {
	for i < len(runes) && j < len(runes) {
		if runes[i] != runes[j] {
			return runes[i] < runes[j]
		}
		i++
		j++
	}
	return i > j
}

// Return last column of BW-matrix using suffix array
func BWTLastColumn(s string, suffixArray []int) string {
	runes := []rune(s)
	var lastColumn bytes.Buffer

	for _, idx := range suffixArray {
		if idx == 0 {
			lastColumn.WriteRune(runes[len(runes)-1])
		} else {
			lastColumn.WriteRune(runes[idx-1])
		}
	}

	return lastColumn.String()
}

// Return suffixes type array for the given string
func SuffixTypeArray(data string) []byte {
	res := make([]byte, len(data)+1)

	res[len(data)] = S_TYPE
	if len(data) == 0 {
		return res
	}
	res[len(data)-1] = L_TYPE

	for i := len(data) - 2; i >= 0; i-- {
		if data[i] > data[i+1] {
			res[i] = L_TYPE
		} else if data[i] == data[i+1] && res[i+1] == L_TYPE {
			res[i] = L_TYPE
		} else {
			res[i] = S_TYPE
		}
	}

	return res
}

// Check if suffix is Left Most "Smaller"
func IsLMS(offset int, typeArray []byte) bool {
	if offset == 0 {
		return false
	}

	if typeArray[offset] == S_TYPE && typeArray[offset-1] == L_TYPE {
		return true
	}

	return false
}

// Print type for each suffix in the given string
func PrintSuffixTypes(s string) {
	st := SuffixTypeArray(s)
	for i := range st {
		if st[i] == 0 {
			fmt.Printf("S")
		} else {
			fmt.Printf("L")
		}
	}
	fmt.Println()
	for i := range st {
		if IsLMS(i, st) {
			fmt.Printf("^")
		} else {
			fmt.Printf(" ")
		}
	}
}

//! Some parts of code are not working so I've fully commented SA-IS algorithm

// func BuildLCPArray(s string, sa []int) []int {
// 	n := len(s)
// 	rank := make([]int, n)
// 	lcp := make([]int, n)

// 	for i := 0; i < n; i++ {
// 		rank[sa[i]] = i
// 	}

// 	k := 0
// 	for i := 0; i < n; i++ {
// 		if rank[i] == n-1 {
// 			k = 0
// 			continue
// 		}
// 		j := sa[rank[i]+1]
// 		for i+k < n && j+k < n && s[i+k] == s[j+k] {
// 			k++
// 		}
// 		lcp[rank[i]] = k
// 		if k > 0 {
// 			k--
// 		}
// 	}

// 	return lcp
// }

// func lmsSubstringsAreEqual(str string, typeArray []byte, offsetA int, offsetB int) bool {
// 	if offsetA == len(str) || offsetB == len(str) {
// 		return false
// 	}

// 	i := 0
// 	for {
// 		bIsLMS := IsLMS(i+offsetB, typeArray)
// 		aIsLMS := IsLMS(i+offsetA, typeArray)

// 		if i > 0 && aIsLMS && bIsLMS {
// 			return true
// 		}

// 		if aIsLMS != bIsLMS {
// 			return false
// 		}

// 		if str[i+offsetA] != str[i+offsetB] {
// 			return false
// 		}

// 		i++
// 	}
// }

// func guessLMSSort(str string, bucketSizes []int, typeArray []byte) []int {
// 	guessedSuffixArray := make([]int, len(str)+1)

// 	bucketTails := FindBucketTails(bucketSizes)

// 	for i := 0; i < len(str); i++ {
// 		if !IsLMS(i, typeArray) {
// 			continue
// 		}

// 		bucketIndex := int(str[i])
// 		guessedSuffixArray[bucketTails[bucketIndex]] = i
// 		bucketTails[bucketIndex]--
// 	}

// 	guessedSuffixArray[0] = len(str)
// 	return guessedSuffixArray
// }

// func summariseSuffixArray(str string, guessedSuffixArray []int, typeArray []byte) ([]byte, int, []int) {
// 	lmsNames := make([]int, len(str)+1)
// 	currentName := 0
// 	var lastLMSSuffixOffset int

// 	lmsNames[guessedSuffixArray[0]] = currentName
// 	lastLMSSuffixOffset = guessedSuffixArray[0]

// 	for i := 1; i < len(guessedSuffixArray); i++ {
// 		suffixOffset := guessedSuffixArray[i]

// 		if !IsLMS(suffixOffset, typeArray) {
// 			continue
// 		}

// 		if !lmsSubstringsAreEqual(str, typeArray, lastLMSSuffixOffset, suffixOffset) {
// 			currentName++
// 		}

// 		lastLMSSuffixOffset = suffixOffset

// 		lmsNames[suffixOffset] = currentName
// 	}

// 	summaryString := make([]byte, 0)
// 	summarySuffixOffsets := make([]int, 0)
// 	for index, name := range lmsNames {
// 		if name == -1 {
// 			continue
// 		}
// 		summarySuffixOffsets = append(summarySuffixOffsets, index)
// 		summaryString = append(summaryString, byte(name))
// 	}

// 	summaryAlphabetSize := currentName + 1

// 	return summaryString, summaryAlphabetSize, summarySuffixOffsets
// }

// func makeSummarySuffixArray(summaryString []byte, summaryAlphabetSize int) []int {
// 	var summarySuffixArray []int

// 	if summaryAlphabetSize == len(summaryString) {
// 		summarySuffixArray = make([]int, len(summaryString)+1)

// 		summarySuffixArray[0] = len(summaryString)

// 		for x, y := range summaryString {
// 			summarySuffixArray[y+1] = x
// 		}
// 	} else {
// 		summarySuffixArray = SAbyIS(string(summaryString), summaryAlphabetSize)
// 	}

// 	return summarySuffixArray
// }

// func accurateLMSSort(str string, bucketSizes []int, summarySuffixArray []int, summarySuffixOffsets []int) []int {
// 	suffixOffsets := make([]int, len(str)+1)

// 	bucketTails := FindBucketTails(bucketSizes)
// 	for i := len(summarySuffixArray) - 1; i >= 1; i-- {
// 		stringIndex := summarySuffixOffsets[summarySuffixArray[i]]

// 		bucketIndex := str[stringIndex]
// 		suffixOffsets[bucketTails[bucketIndex]] = stringIndex
// 		bucketTails[bucketIndex]--
// 	}

// 	suffixOffsets[0] = len(str)

// 	return suffixOffsets
// }

// func SAbyIS(str string, alphabetSize int) []int {
// 	typeArray := SuffixTypeArray(str)
// 	bucketSizes := FindBucketSizes(str, alphabetSize)

// 	guessedSuffixArray := guessLMSSort(str, bucketSizes, typeArray)

// 	induceSortL(str, guessedSuffixArray, bucketSizes, typeArray)
// 	induceSortS(str, guessedSuffixArray, bucketSizes, typeArray)

// 	summaryString, summaryAlphabetSize, summarySuffixOffsets :=
// 		summariseSuffixArray(str, guessedSuffixArray, typeArray)

// 	summarySuffixArray := makeSummarySuffixArray(
// 		summaryString,
// 		summaryAlphabetSize,
// 	)

// 	result := accurateLMSSort(
// 		str, bucketSizes,
// 		summarySuffixArray, summarySuffixOffsets,
// 	)

// 	induceSortL(str, result, bucketSizes, typeArray)
// 	induceSortS(str, result, bucketSizes, typeArray)

// 	return result
// }
