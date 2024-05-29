package main

import (
	"EncodingAlgorithms/algorithms"
	bmh "EncodingAlgorithms/compressors/bwt_mtf_ha"
	"EncodingAlgorithms/compressors/bwt_rle"
	"EncodingAlgorithms/compressors/huffman"
	"EncodingAlgorithms/compressors/rle"
	"fmt"
)

const randomPath = "test_data/texts/random.txt"
const sequencesPath = "test_data/texts/sequences.txt"
const enwik8Path = "test_data/texts/enwik8"
const sentencesPath = "test_data/texts/sentences.txt"

func testRLE() {
	data := "22856 и 55555555555555555555555555 777 аааааааааааа"

	byteArray := []byte(data)
	result := algorithms.RLEncodeUTF8(byteArray)
	fmt.Println("encoded string:", string(result))

	decoded := string(algorithms.RLDecodeUTF8(result))
	fmt.Println("decoded after RLE:", decoded)
}

func testFileRLE() {
	rle.CompressFile(enwik8Path)
	fmt.Println("Compression completed")
	rle.DecompressFile("compressors/rle/compressed.rle")
	fmt.Println("Decompression completed")
}

func testFileBWT_RLE() {
	bwt_rle.CompressFile(sentencesPath)
	fmt.Println("Compression completed")
	bwt_rle.DecompressFile("compressors/bwt_rle/compressed.brl")
	fmt.Println("Decompression completed")
}

func testFileHC() {
	huffman.CompressFile(enwik8Path)
	fmt.Println("Compression completed")
	huffman.DecompressFile("compressors/huffman/compressed.huf")
	fmt.Println("Decompression completed")
}

func testFileBMH() {
	bmh.CompressFile(sequencesPath)
	fmt.Println("Compression completed")
	bmh.DecompressFile("compressors/bwt_mtf_ha/compressed.bmh")
	fmt.Println("Decompression completed")
}

func testAC() {
	a := "aaaaaaaaaaaaaaaaaaa ddddddddddddddd ggggggggggggggggg cccccccc"
	fmt.Println("Initial string:", a)
	array, reversedBorders := algorithms.ArithmeticEncodeChunk(string(a))
	fmt.Println("Encoded data:", array)
	initial := algorithms.ArithmeticDecodeChunk(array, reversedBorders)
	fmt.Println("Decoded data:", initial)
}

func testHC() {
	data := "\x00A3982rvn29d8\x00jdjdkl2dpkl;sdfl,bw\x02\x02p[1qzawsxedcrftvgvbyvhbun\x01buijnmniomkmpl,.']/][9=68-50AAA€AAAABBB烷烷🙂"
	fmt.Println("Initial string:", data)
	result, freqs := algorithms.HuffmanEncode([]byte(data))
	fmt.Println("Encoded data:", string(result))
	decoded := algorithms.HuffmanDecode(result, freqs)
	fmt.Println("decoded string:", string(decoded))
}

func testBWT() {
	a := "абракадабраабракадабраабракадабра"
	fmt.Println("Initial string:", a)
	// fmt.Println("\nNaive BWT:")
	// i, t := algorithms.NaiveBWT(a)
	// fmt.Println("index:", i, "string:", t)
	// inversed := algorithms.NaiveInverseBWT(i, t)
	// fmt.Println("reverse algorithm:", inversed)

	fmt.Println("\nSuffix Array BWT:")
	t1 := algorithms.SuffixArrayBWT(a)
	fmt.Println("string:", t1)
	inversed1 := algorithms.InverseBWT(t1)
	fmt.Println("reverse algorithm:", inversed1)
}

func testMTF() {
	//a := "abracadabra"
	a := "\x00A3982rvn29d8\x00jdjdkl2dpkl;sdfl,bw\x02\x02p[1qzawsxedcrftvgvbyvhbun\x01buijnmniomkmpl,.']/][9=68-50AAA€AAAABBB烷烷🙂"
	t := algorithms.MtFTransform([]byte(a))
	fmt.Println(t)
	inv := algorithms.MtFInverse(t)
	fmt.Println(string(inv))
}

func testLZ77() {
	a := "abracadabra"
	encoded := algorithms.LZ77Encode(a, 5)
	fmt.Println(encoded)
	decoded := algorithms.LZ77Decode(encoded)
	fmt.Println(decoded)
}

func test() {
	a := []byte("abracadabra")
	bwt := algorithms.MakeBWTString(a)
	fmt.Println("BWT:", string(bwt))
	mtf := algorithms.MtFTransform(bwt)
	fmt.Println(string(mtf), mtf)
	hc, freqs := algorithms.HuffmanEncode(mtf)
	fmt.Println(string(hc))

	decodedHC := algorithms.HuffmanDecode(hc, freqs)
	fmt.Println(string(decodedHC))
	decodedMTF := algorithms.MtFInverse(decodedHC)
	fmt.Println(string(decodedMTF))
	decodedBWT := algorithms.InverseTextBWT(decodedMTF)
	fmt.Println(decodedBWT)
}

func main() {
	test()
	//testLZ77()
	//testFileRLE()
	//testBWT()
	//testFileBWT_RLE()
	//testHueta()
	//testMTF()
	//testHC()
	//testFileHC()
	//testMTF()
	//testFileBMH()
	//test()
}
