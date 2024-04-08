package main

import (
	"EncodingAlgorithms/algorithms"
	"EncodingAlgorithms/compressors/bwt_rle"
	"EncodingAlgorithms/compressors/huffman"
	"EncodingAlgorithms/compressors/rle"
	"fmt"
)

const randomPath = "test_data/texts/random.txt"
const sequencesPath = "test_data/texts/sequences.txt"
const enwik8Path = "test_data/texts/enwik8"

func testRLE() {
	data := "22856 и 55555555555555555555555555 777 ааааааааааааа хууууууй залупа"

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
	bwt_rle.CompressFile(sequencesPath)
	fmt.Println("Compression completed")
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
	data := "A3982rvn29d8jdjdkl2dpkl;sdfl,bwp[1qzawsxedcrftvgvbyvhbunbuijnmniomkmpl,.']/][9=68-50AAA€AAAABBB烷烷🙂烷烷烃BBCDGG烃烃烃あ烃烃GGGN"
	fmt.Println("Initial string:", data)
	result, freqs := algorithms.HuffmanEncode(data)
	fmt.Println("Encoded data:", result)
	decoded := algorithms.HuffmanDecode(result, freqs)
	fmt.Println("decoded string:", decoded)
}

func testFileHC() {
	huffman.CompressFile(enwik8Path)
	fmt.Println("Compression completed")
	huffman.DecompressFile("compressors/huffman/compressed.huf")
	fmt.Println("Decompression completed")
}

func testBWT() {
	a := "abracadabra"
	i, t := algorithms.NaiveBWT(a)
	inversed := algorithms.InverseBWT(i, t)
	fmt.Println(inversed)
}

func testLZ77() {
	a := "abracadabra"
	encoded := algorithms.LZ77Encode(a, 5)
	fmt.Println(encoded)
	decoded := algorithms.LZ77Decode(encoded)
	fmt.Println(decoded)
}

func main() {
	// //data := "AAAA€AAAABBB烷烷🙂烷烷烃BBCDGG烃烃烃あ烃烃GGGN"
	// //data := "абракадабра"
	// data := "7-8 так называемых пидоров в ванной 1488 пасхалко ооооо 727"
	// fmt.Println("initial string:", data)

	// // BWT
	// pos, transformed := algorithms.NaiveBWT(data)
	// fmt.Println("transformed string:", transformed)

	// // MTF
	// mtf, alphabet := algorithms.MTF(transformed)
	// mtfString := algorithms.ConvertToString(mtf)
	// fmt.Println("transformed MTF:", mtfString)

	// // RLE encode
	// byteArray := []byte(mtfString)
	// result := algorithms.RLEncodeUTF8(byteArray)
	// fmt.Println("encoded string:", string(result))

	// // RLE decode
	// decoded := string(algorithms.RLDecodeUTF8(result))
	// fmt.Println("decoded after RLE:", decoded)

	// // MTF inverse
	// inversed := algorithms.InverseMTF(mtf, alphabet)
	// fmt.Println("inversed MTF", inversed)

	// // BWT inverse
	// original := algorithms.NaiveBWTInverse(pos, decoded)
	// fmt.Println("original string:", original)

	//a := "jkfoevwbjnf owjndvi wjiwessdljcjn kljnsdjaskdvkwvj d;klwfhj;welfvhjw"
	//testHC()
	//testFileHC()
	//testBWT()

	// k := "GTCCCGATGTCATGTCAGGA$"
	// sa := utils.NaiveSuffixArray(k)
	// fmt.Println(utils.MakeLCPArray(k, sa))
	// fmt.Println(utils.SuffixTypes(k, sa))
	//testLZ77()
	//testFileRLE()
	testFileBWT_RLE()
}
