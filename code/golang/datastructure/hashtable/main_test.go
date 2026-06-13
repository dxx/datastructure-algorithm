package hashtable

import (
    "fmt"
    "testing"
)

func TestAddEmployee(t *testing.T) {
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

func TestUpdateEmployee(t *testing.T) {
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

func TestDeleteEmployee(t *testing.T) {
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
