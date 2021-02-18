package round

import (
    "fmt"
    "testing"
)

func TestInsertNode(t *testing.T) {
    // 创建 head 结点，head 结点不初始化数据，等到添加了第一个结点后才初始化数据
    headNode := &PersonNode{}
    // 创建第一个结点
    personNode1 := &PersonNode{no: 1, name: "张三"}
    // 创建第二个结点
    personNode2 := &PersonNode{no: 2, name: "李四"}
    // 创建第三个结点
    personNode3 := &PersonNode{no: 3, name: "王五"}

    // 插入结点
    insertNode(headNode, personNode1)
    insertNode(headNode, personNode2)
    insertNode(headNode, personNode3)

    printRoundNodeInfo(headNode)
}

func TestDeleteNode(t *testing.T) {
    // 创建结点
    headNode := &PersonNode{}
    personNode1 := &PersonNode{no: 1, name: "张三"}
    personNode2 := &PersonNode{no: 2, name: "李四"}
    personNode3 := &PersonNode{no: 3, name: "王五"}
    personNode4 := &PersonNode{no: 4, name: "赵六"}
    personNode5 := &PersonNode{no: 5, name: "孙七"}

    // 插入结点
    insertNode(headNode, personNode1)
    insertNode(headNode, personNode2)
    insertNode(headNode, personNode3)
    insertNode(headNode, personNode4)
    insertNode(headNode, personNode5)

    fmt.Println("删除前:")
    printRoundNodeInfo(headNode)

    // 删除 no 为 2 的结点
    headNode = deleteNode(headNode, personNode2)
    fmt.Println("删除 no 为 2 的结点后:")
    printRoundNodeInfo(headNode)

    newNode := &PersonNode{no: 6, name: "周八"}
    insertNode(headNode, newNode)
    fmt.Println("插入新结点:")
    printRoundNodeInfo(headNode)

    // 删除 no 为 1，3 的结点
    headNode = deleteNode(headNode, personNode1)
    headNode = deleteNode(headNode, personNode3)
    fmt.Println("删除 no 为 1,3 的结点后:")
    printRoundNodeInfo(headNode)
}
