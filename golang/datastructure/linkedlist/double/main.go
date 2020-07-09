package main

import "fmt"

// 双向链表
// 双向链表也叫双链表，是链表的一种，它的每个数据结点中都有两个指针，分别指向前一个和后一个结点
// 所以，从双向链表中的任意一个结点开始，都可以很方便地访问它的前一个结点和后一个结点。

type HeroNode struct {
    no       int       // 编号
    name     string    // 姓名
    nickname string    // 昵称
    previous *HeroNode // 上一个结点
    next     *HeroNode // 下一个结点
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
    // 将新结点的上一个结点指向当前结点
    newNode.previous = lastNode
}

// 删除指定结点
func deleteNode(headNode *HeroNode, node *HeroNode) {
    tempNode := headNode.next
    for tempNode != nil {
        if tempNode.no == node.no {
            // 将查找到的结点的上一个结点的下一个结点指针指向当前结点的下一个结点
            tempNode.previous.next = tempNode.next
            // 最后一个结点的 next 指向空
            if tempNode.next != nil {
                // 将查找到的结点的下一个结点的上一个结点指针指向当前指针的上一个结点
                tempNode.next.previous = tempNode.previous
            }
            return
        }
        tempNode = tempNode.next
    }
}

// 打印链表结点内容
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
    heroNode1 := &HeroNode{no: 3, name: "吴用", nickname: "智多星"}
    // 创建第二个结点
    heroNode2 := &HeroNode{no: 6, name: "林冲", nickname: "豹子头"}
    // 创建第三个结点
    heroNode3 := &HeroNode{no: 7, name: "秦明", nickname: "霹雳火"}

    // 将结点添加到链表尾部
    insertAtTail(headNode, heroNode1)
    insertAtTail(headNode, heroNode2)
    insertAtTail(headNode, heroNode3)

    printHeadNodeInfo(headNode)
}

func testDeleteNode() {
    // 创建结点
    headNode := new(HeroNode)
    heroNode1 := &HeroNode{no: 1, name: "宋江", nickname: "呼保义"}
    heroNode2 := &HeroNode{no: 2, name: "卢俊义", nickname: "玉麒麟"}
    heroNode3 := &HeroNode{no: 3, name: "吴用", nickname: "智多星"}
    heroNode4 := &HeroNode{no: 4, name: "公孙胜", nickname: "入云龙"}
    heroNode5 := &HeroNode{no: 5, name: "关胜", nickname: "大刀"}

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
    // testDeleteNode()
}
