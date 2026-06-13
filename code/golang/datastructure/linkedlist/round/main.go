package round

import "fmt"

// 循环链表
// 循环链表的特点是表中最后一个结点的指针域指向头结点，整个链表形成一个环

// 双向循环链表

type PersonNode struct {
    no       int
    name     string
    prev *PersonNode
    next     *PersonNode
}

// 插入结点
func insertNode(headNode *PersonNode, newNode *PersonNode) {
    // 判断是否是第一次插入
    if headNode.next == nil {
        headNode.no = newNode.no
        headNode.name = newNode.name
        headNode.prev = headNode
        headNode.next = headNode
        return
    }
    lastNode := headNode
    // 下一个结点不等于头结点继续循环
    for lastNode.next != headNode {
        lastNode = lastNode.next
    }
    // 将新结点添加到链表末尾
    lastNode.next = newNode
    newNode.prev = lastNode
    // 将新结点下一个结点指针指向头结点
    newNode.next = headNode
    headNode.prev = newNode
}

// 删除指定结点，返回头结点
func deleteNode(headNode *PersonNode, node *PersonNode) *PersonNode {
    // 没有结点 或者 只有一个头结点
    if headNode.next == nil || headNode.next == headNode {
        // 头结点就是要删除的结点
        if headNode.no == node.no {
            headNode.next = nil
            headNode.prev = nil
        }
        return headNode
    }
    tempNode := headNode.next
    isExist := false
    for {
        if tempNode == headNode { // 最后一个结点
            if tempNode.no == node.no {
                isExist = true
                // 头结点删除了，将头结点的下一个结点作为头结点
                headNode = headNode.next
            }
            break
        } else if tempNode.no == node.no {
            isExist = true
            break
        }
        tempNode = tempNode.next
    }
    // 存在需要删除的结点
    if isExist {
        // 将查找到的结点的上一个结点的下一个结点指针指向当前结点的下一个结点
        tempNode.prev.next = tempNode.next
        // 将查找到的结点的下一个结点的上一个结点指针指向当前指针的上一个结点
        tempNode.next.prev = tempNode.prev
    }
    return headNode
}

// 打印循环链表的信息
func printRoundNodeInfo(headNode *PersonNode) {
    if headNode.next == nil {
        fmt.Println("该循环链表没有节点")
        return
    }
    tempNode := headNode
    info := "["
    for {
        info += fmt.Sprintf("{no:%v, name:%s}",
            tempNode.no, tempNode.name)
        // 表示最后一个结点
        if tempNode.next == headNode {
            break
        }
        tempNode = tempNode.next
    }
    info += "]"
    fmt.Println(info)
}
