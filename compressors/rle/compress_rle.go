package rle

import (
	"EncodingAlgorithms/algorithms"
	"os"
)

// RLE compression
func CompressFile(path string) {
	data, _ := os.ReadFile(path)
	encoded := algorithms.RLEncodeUTF8(data)
	os.WriteFile("compressors/rle/compressed.rle", encoded, 0644)
}

// RLE decompression
func DecompressFile(path string) {
	data, _ := os.ReadFile(path)
	decompressed := algorithms.RLDecodeUTF8(data)
	os.WriteFile("compressors/rle/decompressed.txt", decompressed, 0644)
}
