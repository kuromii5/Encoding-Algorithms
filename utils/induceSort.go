package utils

func induceSortL(str string, guessedSuffixArray []int, bucketSizes []int, typeArray []byte) {
	bucketHeads := FindBucketHeads(bucketSizes)

	for i := range guessedSuffixArray {
		if guessedSuffixArray[i] == -1 {
			continue
		}

		j := guessedSuffixArray[i] - 1
		if j < 0 || typeArray[j] != L_TYPE {
			continue
		}

		bucketIndex := int(str[j])
		guessedSuffixArray[bucketHeads[bucketIndex]] = j
		bucketHeads[bucketIndex]++
	}
}

func induceSortS(str string, guessedSuffixArray []int, bucketSizes []int, typeArray []byte) {
	bucketTails := FindBucketTails(bucketSizes)

	for i := len(guessedSuffixArray) - 1; i >= 0; i-- {
		j := guessedSuffixArray[i] - 1
		if j < 0 || typeArray[j] != S_TYPE {
			continue
		}

		bucketIndex := int(str[j])
		guessedSuffixArray[bucketTails[bucketIndex]] = j
		bucketTails[bucketIndex]--
	}
}
