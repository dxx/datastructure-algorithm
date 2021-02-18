package joseph

import "fmt"

// 约瑟夫问题
// 设编号为 1,2,...n 的 n 个人围坐一圈约定编号为 k(1<=k<=n) 的人
// 从 1 开始报数，数到 m 的那个人出列它的下一位又从 1 开始报数数到 m
// 的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列

type Person struct {
    no       int     // 编号
    prev     *Person // 上一个人
    next     *Person // 下一个人
}

type PersonLinkedList struct {
    first  *Person // 第一个小孩
    length int     // 小孩的数量
}

func NewBoyLinkedList(count int) *PersonLinkedList {
    if count < 1 {
        fmt.Printf("链表至少需要一个元素")
        return nil
    }

    personLinkedList := PersonLinkedList{first: nil, length: count}

    // 初始化小孩，构成双向环形链表
    var prev = personLinkedList.first
    for i := 1; i <= count; i++ {
        person := &Person{no: i}

        if i == 1 {
            // 初始化 First 节点
            personLinkedList.first = person
            personLinkedList.first.next = personLinkedList.first
            personLinkedList.first.prev = personLinkedList.first

            // 将 prev 指向 First 节点，继续下一次循环
            prev = personLinkedList.first
            continue
        }
        prev.next = person
        person.prev = prev

        // 新增加的节点的下一个节点指向第一节点
        person.next = personLinkedList.first
        // 第一个节点的上一个节点指向新增加的节点
        personLinkedList.first.prev = person

        prev = person
    }
    return &personLinkedList
}

func (personLinkedList *PersonLinkedList) ShowPersons() {
    if personLinkedList.first == nil {
        return
    }
    if personLinkedList.first == personLinkedList.first.next {
        fmt.Printf("num:%d\n", personLinkedList.first.no)
        return
    }
    current := personLinkedList.first
    for {
        fmt.Printf("num:%d\n", current.no)
        current = current.next

        if current == personLinkedList.first {
            break
        }
    }
}

func (personLinkedList *PersonLinkedList) Count(start, num int) {
    if start < 1 || start > personLinkedList.length {
        fmt.Printf("start 不能小于 1 或者不能大于 %d\n", personLinkedList.length)
        return
    }
    if num > personLinkedList.length {
        fmt.Printf("num 不能大于元素个数: %d\n", personLinkedList.length)
        return
    }

    current := personLinkedList.first

    // 循环 start - 1 次
    for i := 1; i <= start-1; i++ {
        current = current.next
    }

    for {
        // 表示只有一个节点
        if current.prev == current && current.next == current {
            break
        }

        // 循环 num - 1 次
        for i := 1; i <= num-1; i++ {
            current = current.next
        }
        // 删除元素
        current.prev.next = current.next
        current.next.prev = current.prev

        fmt.Printf("出队人的编号: %d\n", current.no)
        current = current.next
    }
    fmt.Printf("最后留下人的编号: %d\n", current.no)
}
