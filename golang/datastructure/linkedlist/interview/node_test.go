package interview

import (
    "fmt"
    "testing"
)

func TestGetLength(t *testing.T) {
    headNode := &Node{}
    node1 := &Node{"node1", nil}
    node2 := &Node{"node2", nil}
    node3 := &Node{"node3", nil}
    insertAtTail(headNode, node1)
    insertAtTail(headNode, node2)
    insertAtTail(headNode, node3)

    length := getNodeLength(headNode)
    fmt.Printf("单链表结点个数为: %v\n", length)
}

func TestGetLastIndexNode(t *testing.T) {
    headNode := &Node{}
    node1 := &Node{"node1", nil}
    node2 := &Node{"node2", nil}
    node3 := &Node{"node3", nil}
    insertAtTail(headNode, node1)
    insertAtTail(headNode, node2)
    insertAtTail(headNode, node3)
    var index = 2
    lastNode := getLastIndexNode(headNode, index)
    fmt.Printf("单链表结点中倒数第 %d 个结点为: %s\n", index, lastNode.name)
}

func TestReverseNode(t *testing.T) {
    headNode := &Node{}
    node1 := &Node{"node1", nil}
    node2 := &Node{"node2", nil}
    node3 := &Node{"node3", nil}
    insertAtTail(headNode, node1)
    insertAtTail(headNode, node2)
    insertAtTail(headNode, node3)

    fmt.Println("反转前:")
    printNodeInfo(headNode)

    reverseNode(headNode)

    fmt.Println("反转后:")
    printNodeInfo(headNode)
}
