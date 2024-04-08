package utils

func FindBucketSizes(str string, alphabetSize int) []int {
	res := make([]int, alphabetSize)

	for _, char := range str {
		res[char]++
	}

	return res
}

func FindBucketHeads(bucketSizes []int) []int {
	offset := 1
	res := make([]int, len(bucketSizes))
	for i, size := range bucketSizes {
		res[i] = offset
		offset += size
	}
	return res
}

func FindBucketTails(bucketSizes []int) []int {
	offset := 1
	res := make([]int, len(bucketSizes))
	for i, size := range bucketSizes {
		offset += size
		res[i] = offset - 1
	}
	return res
}
