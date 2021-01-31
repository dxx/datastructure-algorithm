## 哈希表(散列)

> 各种语言实现代码：[Go](./golang/datastructure/hashtable)   [Java](./java/datastructure/src/com/mcx/hashtable)   [JavaScript](./javascript/datastructure/hashtable)   [Rust](./rust/datastructure/src/hashtable)
>
> 默认使用 **Go** 语言实现。

### 简介

哈希表也叫散列表，是根据关键值 key 直接进行访问的数据结构。它通过把 key 映射到表中一个位置来记录，以加快查找的速度。这个映射函数叫做散列函数，存放记录的数组叫做散列表。

![data_structure_hashtable_01](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/data_structure_hashtable_01.png)

### 实现

使用哈希表来管理员工信息，实现添加，修改，删除，根据员工编号查找员工信息等功能。

思路如下：

![data_structure_hashtable_02](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/data_structure_hashtable_02.png)

* 定义一个结构体，声明一个存放记录的数组字段，数组大小定义为 5（可根据需要增加）
* 定义一个链表结构体，声明员工结构体和指向下一个结点的指针两个字段
* 定义员工结构体，声明编号 id 和名字 name 字段

先定义链表和员工结构体：

```go
// 员工结构体
type Employee struct {
    id int
    name string
}

// 单链表结构体
type LinkNode struct {
    employee *Employee
    next *LinkNode
}
```

定义哈希表结构体和一个用来输出哈希表内容的 list 方法：

```go
// 哈希表结构体，包含一个 LinkNode 数组
type Hashtable struct {
    linkArray [5]*LinkNode
}

// 创建哈希表
func NewHashtable() *Hashtable {
    return &Hashtable{[5]*LinkNode{}}
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
```

添加员工方法实现如下：

```go
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
```

编写测试代码：

```go
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

func main() {
    testAddEmployee()
}
```

运行输出：

```
添加员工后:
linkArray[0]=[{id=5, name=孙七}]
linkArray[1]=[{id=1, name=张三}]
linkArray[2]=[{id=2, name=李四}]
linkArray[3]=
linkArray[4]=
```

要修改员工，就要先找到需要修改的员工，先实现通过 id 查找员工的方法：

```go
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
```

然后调用通过 id 查找员工后进行修改：

```go
// 修改员工的方法
func (h *Hashtable) Update(employee *Employee) {
    emp := h.getEmployeeById(employee.id)
    if emp != nil {
        emp.name = employee.name
    }
}
```

编写测试修改的函数：

```go
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

func main() {
    testUpdateEmployee()
}
```

运行后输出：

```
修改员工前:
linkArray[0]=
linkArray[1]=[{id=1, name=张三}{id=6, name=周八}]
linkArray[2]=[{id=2, name=李四}]
linkArray[3]=
linkArray[4]=
修改员工后:
linkArray[0]=
linkArray[1]=[{id=1, name=张三}{id=6, name=菜菜}]
linkArray[2]=[{id=2, name=李四}]
linkArray[3]=
linkArray[4]=
```

删除方法实现如下：

```go
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
```

编写测试删除的函数：

```go
func testDeleteEmployee() {
    hashtable := NewHashtable()
    employee1 := &Employee{2, "李四"}
    employee2 := &Employee{5, "孙七"}
    _ = hashtable.Add(employee1)
    _ = hashtable.Add(employee2)

    fmt.Println("删除员工前:")
    hashtable.List()

    // 修改员工
    hashtable.Delete(2)

    fmt.Println("删除员工后:")
    hashtable.List()
}

func main() {
    testDeleteEmployee()
}
```

运行后输出：

```
删除员工前:
linkArray[0]=[{id=5, name=孙七}]
linkArray[1]=
linkArray[2]=[{id=2, name=李四}]
linkArray[3]=
linkArray[4]=
删除员工后:
linkArray[0]=[{id=5, name=孙七}]
linkArray[1]=
linkArray[2]=
linkArray[3]=
linkArray[4]=
```
