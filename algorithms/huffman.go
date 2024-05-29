package algorithms

import (
	"EncodingAlgorithms/utils"
	"bytes"
	"encoding/binary"
	"os"
	"sort"
	"unicode/utf8"

	"github.com/emirpasic/gods/maps/linkedhashmap"
)

type indexedNode struct {
	index int
	node  HuffmanNode
}

// HuffmanNode represents a node in the Huffman Tree
type HuffmanNode struct {
	Content *rune
	Weight  int
	Left    *HuffmanNode
	Right   *HuffmanNode
}

func sortNodes(nodes []HuffmanNode) []HuffmanNode {
	indexedNodes := make([]indexedNode, len(nodes))
	for i, node := range nodes {
		indexedNodes[i] = indexedNode{i, node}
	}

	sort.Slice(indexedNodes, func(i, j int) bool {
		return indexedNodes[i].node.Weight > indexedNodes[j].node.Weight
	})

	sortedNodes := make([]HuffmanNode, len(nodes))
	for i, indexedNode := range indexedNodes {
		sortedNodes[i] = indexedNode.node
	}

	return sortedNodes
}

// GetCodeForCharacter traverses the Huffman Tree to find the code for a character
func (c *HuffmanNode) GetCodeForCharacter(char rune, parentPath string) string {
	if c.Content != nil && *c.Content == char {
		return parentPath
	}

	if c.Left != nil {
		path := c.Left.GetCodeForCharacter(char, parentPath+"0")
		if path != "" {
			return path
		}
	}

	if c.Right != nil {
		path := c.Right.GetCodeForCharacter(char, parentPath+"1")
		if path != "" {
			return path
		}
	}

	return ""
}

func buildHuffmanTree(freqs *linkedhashmap.Map) HuffmanNode {
	var nodeArray []HuffmanNode

	for _, k := range freqs.Keys() {
		key := k.(rune)
		value, _ := freqs.Get(key)

		nodeArray = append(nodeArray, HuffmanNode{Content: &key, Weight: value.(int), Left: nil, Right: nil})
	}

	// for _, node := range nodeArray {
	// 	if node.Content != nil {
	// 		fmt.Printf("Content: %q\n", *node.Content)
	// 	} else {
	// 		fmt.Println("Content: nil")
	// 	}
	// }

	for len(nodeArray) > 1 {
		nodeArray = sortNodes(nodeArray)

		left := nodeArray[len(nodeArray)-1]
		right := nodeArray[len(nodeArray)-2]
		parent := HuffmanNode{
			Content: nil,
			Weight:  left.Weight + right.Weight,
			Left:    &left,
			Right:   &right,
		}

		nodeArray = append(nodeArray[:len(nodeArray)-2], parent)
	}

	return nodeArray[0]
}

func buildHuffmanMap(freqs *linkedhashmap.Map) map[rune]string {
	huffmanTree := buildHuffmanTree(freqs)

	huffmanCodesMap := make(map[rune]string)
	for _, char := range freqs.Keys() {
		huffmanCodesMap[char.(rune)] = huffmanTree.GetCodeForCharacter(char.(rune), "")
	}

	return huffmanCodesMap
}

func WriteDataToFile(file *os.File, freqs *linkedhashmap.Map, encoded []byte) {
	encodedString := string(encoded)

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
	bitsLength := len(encodedString)
	lengthBuffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(lengthBuffer, uint32(bitsLength))
	file.Write(lengthBuffer)

	// huffman codes
	byteLength := (bitsLength + 7) / 8
	dataBuffer := make([]byte, byteLength)
	for i := 0; i < bitsLength; i++ {
		byteIndex := i / 8
		bitIndex := uint(7 - (i % 8))
		if encodedString[i] == '1' {
			dataBuffer[byteIndex] |= 1 << bitIndex
		}
	}

	// write data to the file
	file.Write(dataBuffer)
}

func ReadDataFromFile(data []byte) ([]byte, *linkedhashmap.Map) {
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

	return encoded.Bytes(), freqs
}

func HuffmanEncode(input []byte) ([]byte, *linkedhashmap.Map) {
	data := string(input)
	var buffer bytes.Buffer

	freqs := utils.CountFrequenciesSorted(data)
	huffmanMap := buildHuffmanMap(freqs)

	for _, char := range data {
		buffer.WriteString(huffmanMap[rune(char)])
	}

	return buffer.Bytes(), freqs
}

func HuffmanDecode(encoded []byte, freqs *linkedhashmap.Map) []byte {
	var decoded bytes.Buffer
	huffmanTree := buildHuffmanTree(freqs)

	// copy huffmanTree
	root := huffmanTree
	for _, bit := range string(encoded) {
		if bit == '0' {
			root = *root.Left
		} else {
			root = *root.Right
		}

		if root.Content != nil {
			decoded.WriteRune(*root.Content)
			root = huffmanTree
		}
	}

	return decoded.Bytes()
}
