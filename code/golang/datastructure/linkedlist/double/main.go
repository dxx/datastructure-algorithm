package double

import "fmt"

// 双向链表
// 双向链表也叫双链表，是链表的一种，它的每个数据结点中都有两个指针，分别指向前一个和后一个结点
// 所以，从双向链表中的任意一个结点开始，都可以很方便地访问它的前一个结点和后一个结点。

type HeroNode struct {
    no       int       // 编号
    name     string    // 姓名
    nickname string    // 昵称
    prev *HeroNode // 上一个结点
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
    // 将新结点的上一个结点指向当前结点
    newNode.prev = lastNode
}

// 删除指定结点
func deleteNode(headNode *HeroNode, node *HeroNode) {
    tempNode := headNode.next
    for tempNode != nil {
        if tempNode.no == node.no {
            // 将查找到的结点的上一个结点的下一个结点指针指向当前结点的下一个结点
            tempNode.prev.next = tempNode.next
            // 最后一个结点的 next 指向空
            if tempNode.next != nil {
                // 将查找到的结点的下一个结点的上一个结点指针指向当前指针的上一个结点
                tempNode.next.prev = tempNode.prev
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
