package utils

func AverageRepeatLength(s string) float64 {
	totalLength := 0
	repeatedSequences := 0
	sequenceLength := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			sequenceLength++
		} else {
			if sequenceLength > 1 {
				totalLength += sequenceLength
				repeatedSequences++
			}
			sequenceLength = 1
		}
	}

	// Check if the last sequence was repeated
	if sequenceLength > 1 {
		totalLength += sequenceLength
		repeatedSequences++
	}

	// Calculate the average length of repeated sequences
	average := 0.0
	if repeatedSequences > 0 {
		average = float64(totalLength-repeatedSequences*2) / float64(len(s))
	}

	return average
}
