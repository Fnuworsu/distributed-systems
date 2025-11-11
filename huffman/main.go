package main

import (
	"fmt"
	"huffman_encoding/encode"
	"huffman_encoding/decode"
)

func main() {
	s := "geeksforgeeks"

	pq := encode.BuildHeap(s)
	ans, root := encode.BuildHuffmanTree(pq)
	res := decode.DecodeString(ans, root)

	fmt.Println("Encoded", ans)
	fmt.Println("Decoded", res)
}