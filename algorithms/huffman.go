package algorithms

import (
	"EncodingAlgorithms/utils"
	"bytes"
	"sort"

	"github.com/emirpasic/gods/maps/linkedhashmap"
)

type indexedNode struct {
	index int
	node  HuffmanNode
}

// HuffmanNode represents a node in the Huffman Tree
type HuffmanNode struct {
	Content rune
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
	if c.Content == char {
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
		nodeArray = append(nodeArray, HuffmanNode{key, value.(int), nil, nil})
	}

	for len(nodeArray) > 1 {
		nodeArray = sortNodes(nodeArray)

		left := nodeArray[len(nodeArray)-1]
		right := nodeArray[len(nodeArray)-2]
		parent := HuffmanNode{
			Weight: left.Weight + right.Weight,
			Left:   &left,
			Right:  &right,
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

func HuffmanEncode(input string) (string, *linkedhashmap.Map) {
	var buffer bytes.Buffer

	freqs := utils.CountFrequenciesSorted(input)
	huffmanMap := buildHuffmanMap(freqs)

	for _, char := range input {
		buffer.WriteString(huffmanMap[char])
	}

	return buffer.String(), freqs
}

func HuffmanDecode(encoded string, freqs *linkedhashmap.Map) string {
	var decoded bytes.Buffer
	huffmanTree := buildHuffmanTree(freqs)

	// copy huffmanTree
	root := huffmanTree
	for _, bit := range encoded {
		if bit == '0' {
			root = *root.Left
		} else {
			root = *root.Right
		}

		if root.Content != 0 {
			decoded.WriteRune(root.Content)
			root = huffmanTree
		}
	}

	return decoded.String()
}
