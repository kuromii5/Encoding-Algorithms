package huffman

import (
	"EncodingAlgorithms/algorithms"
	"os"
)

func CompressFile(path string) {
	file, _ := os.Create("compressors/huffman/compressed.huf")
	data, _ := os.ReadFile(path)

	encoded, freqs := algorithms.HuffmanEncode(data)

	algorithms.WriteDataToFile(file, freqs, encoded)

	file.Close()
}

func DecompressFile(path string) {
	file, _ := os.Create("compressors/huffman/decompressed.txt")
	data, _ := os.ReadFile(path)

	encoded, freqs := algorithms.ReadDataFromFile(data)

	decoded := algorithms.HuffmanDecode(encoded, freqs)
	file.Write(decoded)

	file.Close()
}
