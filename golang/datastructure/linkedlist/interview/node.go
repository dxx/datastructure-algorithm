package main

import "fmt"

type Node struct {
    name string
    next *Node
}

func insertAtTail(headNode *Node, newNode *Node) {
    lastNode := headNode
    for lastNode.next != nil {
        lastNode = lastNode.next
    }
    lastNode.next = newNode
}

func printNodeInfo(headNode *Node) {
    if headNode.next == nil {
        fmt.Println("该链表没有结点")
        return
    }
    tempNode := headNode.next
    info := "["
    for tempNode != nil {
        info += fmt.Sprintf("{name:%s}", tempNode.name)
        tempNode = tempNode.next
    }
    info += "]"
    fmt.Println(info)
}

// 获取单链表有效的结点数
// 1.遍历结点数
// 2.定义一个长度遍历，每遍历一次长度 +1
func getNodeLength(headNode *Node) int {
    // 头结点为空，返回 0
    if headNode == nil {
        return 0
    }
    var length = 0
    node := headNode.next
    for node != nil {
        length++
        node = node.next
    }
    return length
}

// 获取倒数第 n 个节点
// 1.获取链表节点数 length
// 2.遍历到 length - n 个节点
// 3.然后返回
func getLastIndexNode(headNode *Node, index int) *Node {
    // 头结点为空，返回空
    if headNode == nil {
        return nil
    }
    length := getNodeLength(headNode)
    if index <= 0 || index > length {
        return nil
    }
    lastNode := headNode.next
    for i := 0; i < length - index; i++ {
        lastNode = lastNode.next
    }
    return lastNode
}

// 单链表反转
// 1.定义一个新的头结点 reverseHead
// 2.遍历链表，每遍历一个结点，将其取出，放在新的头结点 reverseHead 的后面
// 3.最后将头结点的 next 结点指向 reverseHead 的 next 结点
func reverseNode(headNode *Node) {
    if headNode == nil || headNode.next == nil {
        return
    }
    // 定义新的头结点
    reverseHead := &Node{}
    current := headNode.next
    var next *Node
    for current != nil {
        // 保存当前结点的下一个结点
        next = current.next
        // 将 reverseHead 结点的下一个结点放在当前结点的下一个结点
        current.next = reverseHead.next
        // 当前结点放在 reverseHead 后面
        reverseHead.next = current
        // 移动当前结点
        current = next
    }
    // 将头结点的 next 结点指向 reverseHead 的 next 结点
    headNode.next = reverseHead.next
}
