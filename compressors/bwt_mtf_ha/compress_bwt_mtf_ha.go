package bmh

import (
	"EncodingAlgorithms/algorithms"
	"os"
)

func CompressFile(path string) {
	file, _ := os.Create("compressors/bwt_mtf_ha/compressed.bmh")
	data, _ := os.ReadFile(path)

	// Make BWT
	//BWT := algorithms.MakeBWTString(data)

	// Make MTF
	MTF := algorithms.MtFTransform(data)

	// Huffman encode
	encoded, freqs := algorithms.HuffmanEncode(MTF)
	algorithms.WriteDataToFile(file, freqs, encoded)

	file.Close()
}

func DecompressFile(path string) {
	file, _ := os.Create("compressors/bwt_mtf_ha/decompressed.txt")
	data, _ := os.ReadFile(path)

	encoded, freqs := algorithms.ReadDataFromFile(data)
	decodedHuffman := algorithms.HuffmanDecode(encoded, freqs)

	inversedMTF := algorithms.MtFInverse(decodedHuffman)

	//decodedBWT := algorithms.InverseTextBWT(inversedMTF)

	file.WriteString(string(inversedMTF))
	file.Close()
}
