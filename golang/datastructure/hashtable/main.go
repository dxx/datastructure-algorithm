package main

import (
    "errors"
    "fmt"
)

// 哈希表
// 哈希表也叫散列表，是根据关键值 key 直接进行访问的数据结构
// 它通过把 key 映射到表中一个位置来记录，以加快查找的速度

// 员工结构体
type Employee struct {
    id   int
    name string
}

// 单链表结构体
type LinkNode struct {
    employee *Employee
    next     *LinkNode
}

// 哈希表结构体，包含一个 LinkNode 数组
type Hashtable struct {
    linkArray [5]*LinkNode
}

// 创建哈希表
func NewHashtable() *Hashtable {
    return &Hashtable{[5]*LinkNode{}}
}

// 添加员工的方法，按照 id 升序插入
func (h *Hashtable) Add(employee *Employee) error {
    linkNode := &LinkNode{employee: employee}
    // 计算下标
    index := employee.id % 5
    headNode := h.linkArray[index]
    // 数组没有元素
    if headNode == nil {
        // 添加到链表数组中，作为头结点
        h.linkArray[index] = linkNode
        return nil
    }
    // 判断头结点
    if headNode.employee.id > employee.id {
        // 新结点作为头结点
        linkNode.next = headNode
        h.linkArray[index] = linkNode
        return nil
    } else if headNode.employee.id == employee.id {
        return errors.New("员工 id 重复")
    }

    // 查找后续结点中合适的位置插入
    tempNode := headNode
    for {
        if tempNode.next == nil { // 最后一个结点
            break
        } else if tempNode.next.employee.id > employee.id {
            break
        } else if tempNode.next.employee.id == employee.id {
            return errors.New("员工 id 重复")
        }
        tempNode = tempNode.next
    }
    // tempNode 的下一个结点插入到 linkNode 的下一个结点
    linkNode.next = tempNode.next
    // 将 linkNode 插入到 tempNode 后面
    tempNode.next = linkNode
    return nil
}

// 修改员工的方法
func (h *Hashtable) Update(employee *Employee) {
    emp := h.getEmployeeById(employee.id)
    if emp != nil {
        emp.name = employee.name
    }
}

// 删除员工的方法
func (h *Hashtable) Delete(id int) {
    // 计算下标
    index := id % 5
    headNode := h.linkArray[index]
    // 数组没有元素
    if headNode == nil {
        return
    }
    // 判断头结点
    if headNode.employee.id == id {
        h.linkArray[index] = headNode.next
        return
    }

    // 查找后续结点中需要删除的员工
    tempNode := headNode
    for tempNode.next != nil {
        if tempNode.next.employee.id == id {
            // 将要删除结点的下一个结点链接到该结点的上一个结点
            tempNode.next = tempNode.next.next
            return
        }
        tempNode = tempNode.next
    }
}

// 通过 id 查找员工
func (h *Hashtable) getEmployeeById(id int) *Employee {
    // 计算下标
    index := id % 5
    headNode := h.linkArray[index]
    // 数组没有元素
    if headNode == nil {
        return nil
    }
    // 查找链表中是否存在相同 id 的员工
    tempNode := headNode
    for tempNode != nil {
        if tempNode.employee.id == id {
            return tempNode.employee
        }
        tempNode = tempNode.next
    }
    return nil
}

// 显示哈希表内容的方法
func (h *Hashtable) List() {
    for i := 0; i < len(h.linkArray); i++ {
        employeeInfo := ""
        headNode := h.linkArray[i]
        if headNode != nil {
            employeeInfo += "["
            tempNode := headNode
            // 循环所有结点
            for tempNode != nil {
                employeeInfo += fmt.Sprintf("{id=%v, name=%s}",
                    tempNode.employee.id, tempNode.employee.name)
                tempNode = tempNode.next
            }
            employeeInfo += "]"
        }
        fmt.Printf("linkArray[%v]=%s\n", i, employeeInfo)
    }
}

func testAddEmployee() {
    // 创建一个哈希表
    hashtable := NewHashtable()
    // 创建员工
    employee1 := &Employee{1, "张三"}
    employee2 := &Employee{2, "李四"}
    employee3 := &Employee{5, "孙七"}
    // 添加员工
    _ = hashtable.Add(employee1)
    _ = hashtable.Add(employee2)
    _ = hashtable.Add(employee3)

    fmt.Println("添加员工后:")
    // 显示哈希表内容
    hashtable.List()
}

func testUpdateEmployee() {
    hashtable := NewHashtable()
    employee1 := &Employee{1, "张三"}
    employee2 := &Employee{2, "李四"}
    employee3 := &Employee{6, "周八"}
    _ = hashtable.Add(employee1)
    _ = hashtable.Add(employee2)
    _ = hashtable.Add(employee3)

    fmt.Println("修改员工前:")
    // 显示哈希表内容
    hashtable.List()

    // 修改员工
    employee := &Employee{6, "菜菜"}
    hashtable.Update(employee)

    fmt.Println("修改员工后:")
    hashtable.List()
}

func testDeleteEmployee() {
    hashtable := NewHashtable()
    employee1 := &Employee{2, "李四"}
    employee2 := &Employee{5, "孙七"}
    _ = hashtable.Add(employee1)
    _ = hashtable.Add(employee2)

    fmt.Println("删除员工前:")
    hashtable.List()

    // 删除员工
    hashtable.Delete(2)

    fmt.Println("删除员工后:")
    hashtable.List()
}

func main() {
    // testAddEmployee()
    // testUpdateEmployee()
    // testDeleteEmployee()
}
