package huffman_tree

import (
    "fmt"
    "sort"
    "strconv"
    "strings"
)

// 哈夫曼编码
// 哈夫曼编码(Huffman Coding)，又称霍夫曼编码，是一种编码方式，哈夫曼编码
// 是可变字长编码(VLC)的一种。Huffman于1952年提出一种编码方法，该方法完全
// 依据字符出现概率来构造异字头的平均长度最短的码字，有时称之为最佳编码

// 定义树的存储结构，使用字节作为域(data), 出现的次数作为权值(weight)
type DataNode struct {
    data   byte // 保存字节
    weight int  // 保存出现的次数
    left   *DataNode
    right  *DataNode
}

type DataNodes []*DataNode

func (nodes DataNodes) Len() int {
    return len(nodes)
}

func (nodes DataNodes) Less(i, j int) bool {
    return nodes[i].weight < nodes[j].weight
}

func (nodes DataNodes) Swap(i, j int) {
    nodes[i], nodes[j] = nodes[j], nodes[i]
}

func (nodes DataNodes) String() string {
    str := "["
    for _, node := range nodes {
        str += fmt.Sprintf("{data:%v, value:%d}, ", node.data, node.weight)
    }
    str += "]"
    return str
}

// 转哈夫曼树
// 1.获取到要压缩的数据对应的字节数组
// 2.统计每个字节出现的次数, 将字节作为 key, 次数作为 value
// 3.定义树的存储结构，使用字节作为域(data), 出现的次数作为权值(weight)
// 4.每一个 key 对应一个节点, 构建哈夫曼树

// 构建哈夫曼树
func buildHuffmanTree(bytes []byte) *DataNode {
    // 统计每个字节出现的次数, 将字节作为 key, 次数作为 value
    byteMap := count(bytes)

    dataNodes := initDataNode(byteMap)

    // 每一个 key 对应一个节点, 构建哈夫曼树
    for len(dataNodes) > 1 {
        // 稳定排序，保证每次生成的哈夫曼树保持不变
        // 也可以使用不稳定排序，如果排序的顺序不一样，会导致生成的哈夫曼树不一样，但最后编码的长度是一样的
        sort.Stable(dataNodes)

        left := dataNodes[0]  // 权值最小的元素
        right := dataNodes[1] // 权值第二小的元素
        // 创建新的根节点
        root := &DataNode{weight: left.weight + right.weight}
        // 构建二叉树
        root.left = left
        root.right = right
        // 删除处理过的节点
        dataNodes = deleteDataNode(dataNodes, left)
        dataNodes = deleteDataNode(dataNodes, right)
        // 将二叉树加入到 nodes
        dataNodes = append(dataNodes, root)
    }
    return dataNodes[0]
}

// 获取哈夫曼码表
// 生成字符和编码对应关系
// node: 哈夫曼码树
// 返回 map[32:00 101:010 103:011 105:100 108:101 111:111 118:110]
func getCodes(node *DataNode) map[byte]string {
    codeMap := make(map[byte]string, 0)
    var codes []string
    if node != nil {
        getLeafCodes(node.left, "0", codes, codeMap)
        getLeafCodes(node.right, "1", codes, codeMap)
    }
    return codeMap
}

// 递归拼接 0 或 1, 形成 101:010
func getLeafCodes(node *DataNode, flag string, codes []string, codeMap map[byte]string) {
    codes = append(codes, flag)
    if node != nil {
        if node.data == 0 { // 表示非叶子节点
            getLeafCodes(node.left, "0", codes, codeMap)
            getLeafCodes(node.right, "1", codes, codeMap)
        } else { // 叶子节点
            codeMap[node.data] = strings.Join(codes, "")
        }
    }
}

// 编码成字节切片
// 1.根据原始数据，和哈夫曼码表，将编码顺序拼接
// 2.以每八位作为一个字节的二进制数，转成十进制，并顺序放入切片中

// 将原始数据和码表编码成字节切片
// source: 原始字节切片
// codeMap: 哈夫曼码表
// 返回: 编码后的字节切片
func encodeBytes(source []byte, codeMap map[byte]string) []byte {
    var codes []string
    for _, b := range source {
        codeStr, ok := codeMap[b]
        if ok {
            codes = append(codes, codeStr)
        }
    }
    // 转成编码序列
    codeString := strings.Join(codes, "")

    fmt.Printf("编码序列:%s\n", codeString) // 1000010111111001000011111

    // 将编码序列每八位看成一个字节，放入字节切片中
    // 八位看成一个字节
    // 10000101 11111001 00001111 1
    targetByteLength := (len(codeString) + 7) / 8
    // 增加 1 个字节，用来表示最后一个字节的有效比特位
    targetBytes := make([]byte, targetByteLength+1)
    start, end := 0, 0
    lastBitLength := 8
    for i := 0; i < targetByteLength; i++ {
        start = i * 8
        end = (i + 1) * 8
        if end > len(codeString) {
            end = len(codeString)
            // 记录不足八位的有效长度
            lastBitLength = len(codeString) % 8
        }
        // 每八位截取一次作为一个字节
        bitStr := codeString[start:end]

        n, _ := strconv.ParseInt(bitStr, 2, 64)
        targetBytes[i] = byte(n)
    }
    // 最后一个字节记录原始序列中最后一个字节真实有效的比特长度
    targetBytes[targetByteLength] = byte(lastBitLength)

    return targetBytes
}

// 解码成原始字节切片
// 1.将编码后的字节切片，还原成原始字符串序列，每个字节对应八个字符长度(除最后一个字节外)
// 2.反转哈夫曼编码表，将 key 作为新编码表的 value，将 value 作为新编码表的 value
// 3.从字符串序列起始位置开始，每向右增加一个位置就从反转后的编码中搜索是否存在 key，将
//   存在的 key 对应的 value 添加到切片中，知道所有的字符串序列搜索完成

// 将编码后的字节数据按照哈夫曼编码表解码成字节切片
// target: 编码后的字节切片
// codeMap: 哈夫曼码表
// 返回解码后的字节切片
func decodeBytes(target []byte, codeMap map[byte]string) []byte {
    var codes []string
    codeLength := len(target)
    // 实际长度
    codeRealLength := codeLength - 1
    for i := 0; i < codeRealLength; i++ {
        n := int64(target[i])
        bitStr := ""
        if i != codeRealLength-1 {
            // 除了最后一个字节，不足 8 位，高位补 0
            // 如果不补 0，会导致长度不正确
            n = n | 256
            bitStr = strconv.FormatInt(n, 2)
            // 截取最后八位
            bitStr = bitStr[len(bitStr)-8:]
        } else { // 最后一个字节
            lastBitLength := int(target[codeLength-1])
            bitStr = strconv.FormatInt(n, 2)
            if len(bitStr) < lastBitLength {
                // 前面补 0
                bitStr = strings.Repeat("0", lastBitLength-len(bitStr)) + bitStr
            }
        }

        codes = append(codes, bitStr)
    }

    sourceString := strings.Join(codes, "") // 拼接成原始序列

    fmt.Printf("解码序列:%s\n", sourceString) // 1000010111111001000011111

    // 反转哈夫曼编码表
    // 32:00 => 00:32
    reverseCodeMap := make(map[string]byte, len(codeMap))
    for key, value := range codeMap {
        reverseCodeMap[value] = key
    }

    var sourceBytes []byte
    var str string
    // 根据反转后的哈夫曼编码表，查找对应的字节数值
    for i := 0; i < len(sourceString); {
        count := 1
        for {
            if i+count > len(sourceString) {
                break
            }
            str = sourceString[i : i+count]
            b, ok := reverseCodeMap[str]
            if ok {
                // 搜索到就添加到切片中
                sourceBytes = append(sourceBytes, b)
                break
            } else {
                count++
            }
        }
        i += count
    }

    return sourceBytes
}

func count(bytes []byte) map[byte]int {
    count := make(map[byte]int, len(bytes))
    for _, key := range bytes {
        value, ok := count[key]
        if ok {
            count[key] = value + 1
        } else {
            count[key] = 1
        }
    }
    return count
}

func initDataNode(byteMap map[byte]int) DataNodes {
    dataNodes := make(DataNodes, 0)
    var slice []string
    // 保存 key
    for key := range byteMap {
        slice = append(slice, strconv.Itoa(int(key)))
    }
    // 排序，保证每次遍历 key 的顺序一致
    sort.Strings(slice)
    for _, data := range slice {
        n, _ := strconv.Atoi(data)
        key := byte(n)
        dataNodes = append(dataNodes, &DataNode{data: key, weight: byteMap[key]})
    }

    return dataNodes
}

func dataNodePreOrder(node *DataNode) {
    if node == nil {
       return
    }

    fmt.Printf("{data:%v, value:%d}\n", node.data, node.weight)
    dataNodePreOrder(node.left)
    dataNodePreOrder(node.right)
}

func deleteDataNode(nodes DataNodes, node *DataNode) DataNodes {
    for i := 0; i < len(nodes); i++ {
        if nodes[i].data == node.data {
            nodes = append(nodes[:i], nodes[i+1:]...)
            // 避免删除重复的数据
            return nodes
        }
    }
    return nodes
}
