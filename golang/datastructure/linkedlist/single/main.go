package main

import "fmt"

// 单向链表
// 单向链表是链表的一种，其特点是链表的链接方向是单向的，对链表的访问要从头部开始
// 链表是由结点构成，head 指针指向第一个称为表头的结点，而最后一个结点的指针指向 NULL

type HeroNode struct {
    no       int       // 编号
    name     string    // 姓名
    nickname string    // 昵称
    next     *HeroNode // 下一个结点
}

func NewHeroNode(no int, name, nickname string) *HeroNode {
    return &HeroNode {
        no: no,
        name: name,
        nickname: nickname,
    }
}

// 在链表尾部插入，通过 head 找到链表的尾部
func insertAtTail(headNode *HeroNode, newNode *HeroNode) {
    lastNode := headNode
    // 下一个结点不为空继续循环
    for lastNode.next != nil {
        // 将下一个结点赋值给当前结点
        lastNode = lastNode.next
    }
    // 将当前结点插入到链表的最后一个结点
    lastNode.next = newNode
}

// 按照 no 升序插入，通过 head 找到合适的插入位置
func sortInsertByNo(headNode *HeroNode, newNode *HeroNode) {
    tempNode := headNode
    for {
        if tempNode.next == nil { // 最后一个结点，跳出循环
            break
        } else if tempNode.next.no > newNode.no { // newNode 结点应该插入到 tempNode 后面
            break
        } else if tempNode.next.no == newNode.no {
            panic("no 相等不能插入") // no 相等不能插入
        }
        tempNode = tempNode.next
    }
    // tempNode 的下一个结点插入到 newNode 的下一个结点
    newNode.next = tempNode.next
    // newNode 结点插入到 tempNode 的下一个结点
    tempNode.next = newNode
}

// 删除指定结点
func deleteNode(headNode *HeroNode, node *HeroNode) {
    tempNode := headNode
    for tempNode.next != nil {
        if tempNode.next.no == node.no {
            // 将下一个结点的下一个结点，链接到被删除结点的上一个结点
            tempNode.next = tempNode.next.next
            return
        }
        tempNode = tempNode.next
    }
}

// 打印单链表结点内容
func printHeadNodeInfo(headNode *HeroNode) {
    if headNode.next == nil {
        fmt.Println("该链表没有结点")
        return
    }
    tempNode := headNode.next
    info := "["
    for tempNode != nil {
        info += fmt.Sprintf("{no:%v, name:%s, nickname:%s}",
            tempNode.no, tempNode.name, tempNode.nickname)
        tempNode = tempNode.next
    }
    info += "]"
    fmt.Println(info)
}

func testInsertAtTail() {
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

func testSortInsertByNo() {
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

func testDeleteNode() {
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

func main() {
    // testInsertAtTail()
    // testSortInsertByNo()
    // testDeleteNode()
}
