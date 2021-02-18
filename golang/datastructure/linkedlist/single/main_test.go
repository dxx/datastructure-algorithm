package single

import (
    "fmt"
    "testing"
)

func TestInsertAtTail(t *testing.T) {
    // 创建 head 结点，head 结点不包含数据
    headNode := new(HeroNode)
    // 创建第一个结点
    heroNode1 := NewHeroNode(1, "宋江", "呼保义")
    // 创建第二个结点
    heroNode2 := NewHeroNode(2, "卢俊义", "玉麒麟")
    // 创建第三个结点
    heroNode3 := NewHeroNode(3, "吴用", "智多星")

    // 将结点添加到链表尾部
    insertAtTail(headNode, heroNode1)
    insertAtTail(headNode, heroNode2)
    insertAtTail(headNode, heroNode3)

    printHeadNodeInfo(headNode)
}

func TestSortInsertByNo(t *testing.T) {
    // 创建结点，用来做尾部插入
    head := new(HeroNode)
    node1 := NewHeroNode(1, "宋江", "呼保义")
    node2 := NewHeroNode(2, "卢俊义", "玉麒麟")
    node3 := NewHeroNode(3, "吴用", "智多星")

    insertAtTail(head, node1)
    insertAtTail(head, node3) // 将第三个结点插入到第二个位置
    insertAtTail(head, node2)

    fmt.Println("尾部插入的结果:")
    printHeadNodeInfo(head)

    // 创建 head 结点
    headNode := new(HeroNode)
    // 创建第一个结点
    heroNode1 := NewHeroNode(1, "宋江", "呼保义")
    // 创建第二个结点
    heroNode2 := NewHeroNode(2, "卢俊义", "玉麒麟")
    // 创建第三个结点
    heroNode3 := NewHeroNode(3, "吴用", "智多星")

    // 将结点按照 no 升序插入
    sortInsertByNo(headNode, heroNode1)
    sortInsertByNo(headNode, heroNode3)
    sortInsertByNo(headNode, heroNode2)

    fmt.Println("按照 no 升序插入的结果:")
    printHeadNodeInfo(headNode)
}

func TestDeleteNode(t *testing.T) {
    // 创建结点
    headNode := new(HeroNode)
    heroNode1 := NewHeroNode(1, "宋江", "呼保义")
    heroNode2 := NewHeroNode(2, "卢俊义", "玉麒麟")
    heroNode3 := NewHeroNode(3, "吴用", "智多星")
    heroNode4 := NewHeroNode(4, "公孙胜", "入云龙")
    heroNode5 := NewHeroNode(5, "关胜", "大刀")

    // 插入结点
    insertAtTail(headNode, heroNode1)
    insertAtTail(headNode, heroNode2)
    insertAtTail(headNode, heroNode3)
    insertAtTail(headNode, heroNode4)
    insertAtTail(headNode, heroNode5)

    fmt.Println("删除前:")
    printHeadNodeInfo(headNode)

    // 删除 no 为 2 的结点
    deleteNode(headNode, heroNode2)
    fmt.Println("删除 no 为 2 的结点后:")
    printHeadNodeInfo(headNode)

    // 删除 no 为 3, 4 的结点
    deleteNode(headNode, heroNode3)
    deleteNode(headNode, heroNode4)
    fmt.Println("删除 no 为 3,4 的结点后:")
    printHeadNodeInfo(headNode)
}
