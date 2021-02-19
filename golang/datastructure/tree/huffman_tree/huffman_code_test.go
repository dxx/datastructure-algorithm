package huffman_tree

import (
    "fmt"
    "testing"
)

func TestHuffmanCode(t *testing.T) {
    msg := "i love go"
    // 获取到要压缩的数据对应的字节数组
    bytes := []byte(msg)
    // 构建哈夫曼树
    root := buildHuffmanTree(bytes)
    fmt.Println("======前序遍历======")
    dataNodePreOrder(root)

    fmt.Println("======创建码表======")
    // 获取码表
    codeMap := getCodes(root)
    fmt.Println(codeMap)
    fmt.Printf("未编码时的长度:%d\n", len(bytes))

    fmt.Println("======编码======")
    // 编码
    encodeBytes := encodeBytes(bytes, codeMap)
    fmt.Println(encodeBytes)
    fmt.Printf("编码后的长度:%d\n", len(encodeBytes))

    fmt.Println("======解码======")
    // 解码
    sourceBytes := decodeBytes(encodeBytes, codeMap)
    fmt.Printf("解码后的内容:%s\n", string(sourceBytes))
}
