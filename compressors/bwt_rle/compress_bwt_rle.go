package bwt_rle

import (
	"EncodingAlgorithms/algorithms"
	"encoding/binary"
	"fmt"
	"os"
)

type BWTstring struct {
	Index       uint16
	Transformed []byte
}

// 1000 symbols
const CHUNK_SIZE = 1000

func CompressFile(path string) {
	data, _ := os.ReadFile(path)
	s := string(data)
	//runes := []rune(s)

	var transformedStrings []BWTstring
	for i := 0; i < len(s); i += CHUNK_SIZE {
		end := i + CHUNK_SIZE
		if end > len(s) {
			end = len(s)
		}
		chunk := s[i:end]
		s = s[i:]

		index, transformed := algorithms.SuffixArrayBWT(chunk)
		transformedStrings = append(transformedStrings, BWTstring{uint16(index), []byte(transformed)})
	}

	var transformedStringsRLE []BWTstring
	for _, bwtstr := range transformedStrings {
		compressed := algorithms.RLEncodeUTF8(bwtstr.Transformed)
		bwtstr.Transformed = compressed
		fmt.Println(string(bwtstr.Transformed))
		transformedStringsRLE = append(transformedStringsRLE, bwtstr)
	}

	file, _ := os.Create("compressors/bwt_rle/compressed.brl")
	for _, bwtstr := range transformedStringsRLE {
		// Write index
		binary.Write(file, binary.LittleEndian, bwtstr.Index)

		// Write data
		file.Write(bwtstr.Transformed)
	}
}

func DecompressFile(path string) {
	file, _ := os.ReadFile(path)
	fmt.Println(string(file)[:100])
}
