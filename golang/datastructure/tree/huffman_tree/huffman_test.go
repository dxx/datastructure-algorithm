package huffman_tree

import "testing"

func TestHuffmanTree(t *testing.T) {
    nums := []int{1, 7, 3, 8, 16}
    root := createHuffmanTree(nums)
    preOrder(root)
}
