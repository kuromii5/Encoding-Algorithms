package bwt_rle

import (
	"EncodingAlgorithms/algorithms"
	"os"
)

func CompressFile(path string) {
	file, _ := os.Create("compressors/bwt_rle/compressed.brl")
	data, _ := os.ReadFile(path)

	BWT := algorithms.MakeBWTString(data)

	compressed := algorithms.RLEncodeUTF8(BWT)
	file.Write(compressed)

	file.Close()
}

func DecompressFile(path string) {
	file, _ := os.Create("compressors/bwt_rle/decompressed.txt")
	data, _ := os.ReadFile(path)

	decodedRLE := algorithms.RLDecodeUTF8(data)

	decodedBWT := algorithms.InverseTextBWT(decodedRLE)
	file.WriteString(decodedBWT)

	file.Close()
}
