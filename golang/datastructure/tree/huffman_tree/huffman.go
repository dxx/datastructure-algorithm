package main

import (
    "fmt"
    "sort"
)

// 哈夫曼树
// 有 N 个权值作为 N 个叶子结点，构造一棵二叉树，如果该树的带权路径长度达到最小
// 称这样的二叉树为最优二叉树，也称为哈夫曼树(Huffman Tree)。哈夫曼树是带权路
// 径长度最短的树，权值较大的结点离根较近。哈夫曼树又称为最优树。

// 构建步骤
// 1.将 w1、w2、…、wn 看成一个序列, 每个数据可以看做一个权值
// 2.将序列从小到大排序
// 3.选出两个根节点的权值最小的树合并，作为一棵新树的左、右子节点，且新树的根节点权值为其左、右子树根节点权值之和
// 4.从序列中删除选出的两个节点，并将新树加入序列
// 5.重复 2、3、4 步，直到序列中只剩一棵树为止，该树即为所求得的哈夫曼树

type Node struct {
    value int   // 权值
    left  *Node // 左子节点
    right *Node // 右子节点
}

type Nodes []*Node

func (nodes Nodes) Len() int {
    return len(nodes)
}

func (nodes Nodes) Less(i, j int) bool {
    return nodes[i].value < nodes[j].value
}

func (nodes Nodes) Swap(i, j int) {
    nodes[i], nodes[j] = nodes[j], nodes[i]
}

func (nodes Nodes) String() string {
    str := "["
    for _, node := range nodes {
        str += fmt.Sprintf("value:%d, ", node.value)
    }
    str += "]"
    return str
}

// 构建哈夫曼树
func createHuffmanTree(nums []int) *Node {
    if nums == nil {
        return nil
    }
    var nodes = make(Nodes, 0)
    for _, num := range nums {
        nodes = append(nodes, &Node{value: num})
    }
    for len(nodes) > 1 {
        // 排序
        sort.Sort(nodes)
        fmt.Println(nodes)

        left := nodes[0]  // 权值最小的元素
        right := nodes[1] // 权值第二小的元素
        // 创建新的根节点
        root := &Node{value: left.value + right.value}
        // 构建二叉树
        root.left = left
        root.right = right
        // 删除处理过的节点
        nodes = deleteNode(nodes, left)
        nodes = deleteNode(nodes, right)
        // 将二叉树加入到 nodes
        nodes = append(nodes, root)
    }
    return nodes[0]
}

func preOrder(node *Node) {
    if node == nil {
        return
    }
    fmt.Println(node.value)
    preOrder(node.left)
    preOrder(node.right)
}

func deleteNode(nodes Nodes, node *Node) Nodes {
    for i := 0; i < len(nodes); i++ {
        if nodes[i].value == node.value {
            nodes = append(nodes[:i], nodes[i+1:]...)
            i-- // 删除了一个元素，修正下标
        }
    }
    return nodes
}

func main() {
    nums := []int{1, 7, 3, 8, 16}
    root := createHuffmanTree(nums)
    preOrder(root)
}
