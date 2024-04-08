package huffman

import (
	"EncodingAlgorithms/algorithms"
	"bytes"
	"encoding/binary"
	"os"
	"unicode/utf8"

	"github.com/emirpasic/gods/maps/linkedhashmap"
)

func CompressFile(path string) {
	data, _ := os.ReadFile(path)
	s := string(data)
	encoded, freqs := algorithms.HuffmanEncode(s)
	file, _ := os.Create("compressors/huffman/compressed.huf")

	// frequencies length
	sizeFreqs := uint32(freqs.Size())
	binary.Write(file, binary.LittleEndian, sizeFreqs)

	// frequencies data
	for _, char := range freqs.Keys() {
		charCodeBytes := make([]byte, utf8.RuneLen(char.(rune)))
		utf8.EncodeRune(charCodeBytes, char.(rune))
		binary.Write(file, binary.LittleEndian, charCodeBytes)

		value, _ := freqs.Get(char)
		binary.Write(file, binary.LittleEndian, uint32(value.(int)))
	}

	// bits length
	bitsLength := len(encoded)
	lengthBuffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(lengthBuffer, uint32(bitsLength))
	file.Write(lengthBuffer)

	// huffman codes
	byteLength := (bitsLength + 7) / 8
	dataBuffer := make([]byte, byteLength)
	for i := 0; i < bitsLength; i++ {
		byteIndex := i / 8
		bitIndex := uint(7 - (i % 8))
		if encoded[i] == '1' {
			dataBuffer[byteIndex] |= 1 << bitIndex
		}
	}
	file.Write(dataBuffer)
	file.Close()
}

func DecompressFile(path string) {
	data, _ := os.ReadFile(path)

	// frequencies size
	sizeFreqs := binary.LittleEndian.Uint32(data[:4])
	data = data[4:]

	freqs := linkedhashmap.New()
	for i := 0; i < int(sizeFreqs); i++ {
		// read unicode character
		char, size := utf8.DecodeRune(data)
		data = data[size:]

		// read frequency
		freq := binary.LittleEndian.Uint32(data)
		data = data[4:]

		freqs.Put(char, int(freq))
	}

	// data size
	bitsLength := binary.LittleEndian.Uint32(data)
	data = data[4:]

	// read data
	var encoded bytes.Buffer
	byteLength := int(bitsLength / 8)
	bitsLeft := int(bitsLength) % 8
	buffer := data[:byteLength]
	for _, b := range buffer {
		for i := 7; i >= 0; i-- {
			if b&(1<<uint(i)) != 0 {
				encoded.WriteString("1")
			} else {
				encoded.WriteString("0")
			}
		}
	}

	if bitsLeft > 0 {
		lastByte := data[byteLength]
		for i := 7; i >= 8-bitsLeft; i-- {
			if lastByte&(1<<uint(i)) != 0 {
				encoded.WriteString("1")
			} else {
				encoded.WriteString("0")
			}
		}
	}

	decoded := algorithms.HuffmanDecode(encoded.String(), freqs)

	file, _ := os.Create("compressors/huffman/decompressed.txt")
	file.WriteString(decoded)
}
