## 哈希表(散列)

> 各种语言实现代码：[Go](./golang/datastructure/hashtable)   [Java](./java/datastructure/src/com/dxx/hashtable)   [JavaScript](./javascript/datastructure/hashtable)   [TypeScript](./typescript/datastructure/hashtable)   [Python](./python/datastructure/hashtable)   [Rust](./rust/datastructure/src/hashtable)
>
> 默认使用 **Python** 语言实现。

### 简介

哈希表也叫散列表，是根据关键值 key 直接进行访问的数据结构。它通过把 key 映射到表中一个位置来记录，以加快查找的速度。这个映射函数叫做散列函数，存放记录的数组叫做散列表。

![data_structure_hashtable_01](https://dxx.github.io/static-resource/datastructure-algorithm/images/data_structure_hashtable_01.png)

### 实现

使用哈希表来管理员工信息，实现添加，修改，删除，根据员工编号查找员工信息等功能。

思路如下：

![data_structure_hashtable_02](https://dxx.github.io/static-resource/datastructure-algorithm/images/data_structure_hashtable_02.png)

* 定义一个类，声明一个存放记录的数组字段，数组大小定义为 5（可根据需要增加）
* 定义一个链表类，声明员工对象和指向下一个结点的指针两个字段
* 定义员工类，声明编号 id 和名字 name 字段

先定义链表和员工类：

```python
class Employee:
    """员工"""

    def __init__(self, emp_id: int, name: str) -> None:
        self.id = emp_id
        self.name = name


class LinkNode:
    """单链表"""

    def __init__(self, employee: Employee, next_node: "LinkNode | None" = None) -> None:
        self.employee = employee
        self.next = next_node
```

定义哈希表类和一个用来输出哈希表内容的 `list` 方法：

```python
class Hashtable:
    """哈希表结，包含一个 LinkNode 数组"""

    def __init__(self, length: int) -> None:
        self.link_array: list[LinkNode | None] = [None] * length

    def list(self) -> None:
        """显示哈希表内容的方法"""
        for i, head_node in enumerate(self.link_array):
            employee_info = ""
            if head_node is not None:
                employee_info += "["
                temp_node = head_node
                while temp_node is not None:
                    employee_info += "{id=" + str(temp_node.employee.id) + ", name=" + temp_node.employee.name + "}"
                    temp_node = temp_node.next
                employee_info += "]"
            print("linkArray[" + str(i) + "]=" + employee_info)
```

添加员工方法实现如下：

```python
def add(self, employee: Employee) -> None:
    """添加员工的方法，按照 id 升序插入"""
    link_node = LinkNode(employee)
    # 计算下标
    index = employee.id % len(self.link_array)
    head_node = self.link_array[index]
    # 数组没有元素
    if not head_node:
        # 添加到链表数组中，作为头结点
        self.link_array[index] = link_node
        return
    # 判断头结点
    if employee.id < head_node.employee.id:
        # 新结点作为头结点
        link_node.next = head_node
        self.link_array[index] = link_node
        return
    if employee.id == head_node.employee.id:
        print("员工 id 重复")
        return
    # 查找后续结点中合适的位置插入
    temp_node = head_node
    while True:
        if temp_node.next is None:
            break
        if temp_node.next.employee.id > employee.id:
            break
        if temp_node.next.employee.id == employee.id:
            print("员工 id 重复")
            return
        temp_node = temp_node.next
    # temp_node 的下一个结点插入到 link_node 的下一个结点
    link_node.next = temp_node.next
    # 将 link_node 插入到 temp_node 后面
    temp_node.next = link_node
```

编写测试代码：

```python
class Test(unittest.TestCase):
    
    def test_add_employee(self):
        # 创建一个哈希表
        hashtable = Hashtable(5)
        # 创建员工
        employee1 = Employee(1, "张三")
        employee2 = Employee(2, "李四")
        employee3 = Employee(5, "孙七")
        # 添加员工
        hashtable.add(employee1)
        hashtable.add(employee2)
        hashtable.add(employee3)

        print("添加员工后:")
        # 显示哈希表内容
        hashtable.list()
```

运行：

```shell
❯ python -m unittest test_main.Test.test_add_employee
添加员工后:
linkArray[0]=[{id=5, name=孙七}]
linkArray[1]=[{id=1, name=张三}]
linkArray[2]=[{id=2, name=李四}]
linkArray[3]=
linkArray[4]=
```

要修改员工，就要先找到需要修改的员工，先实现通过 id 查找员工的方法：

```python
def get_employee_by_id(self, emp_id: int) -> Employee | None:
    """通过 id 查找员工"""
    # 计算下标
    index = emp_id % len(self.link_array)
    head_node = self.link_array[index]
    # 数组没有元素
    if not head_node:
        return None
    # 查找链表中是否存在相同 id 的员工
    temp_node = head_node
    while temp_node is not None:
        if temp_node.employee.id == emp_id:
            return temp_node.employee
        temp_node = temp_node.next
    return None
```

然后调用通过 id 查找员工后进行修改：

```python
def update(self, employee: Employee) -> None:
    """修改员工的方法"""
    emp = self.get_employee_by_id(employee.id)
    if emp is not None:
        emp.name = employee.name
```

编写测试修改的函数：

```python
class Test(unittest.TestCase):
    
    def test_update_employee(self):
        hashtable = Hashtable(5)
        employee1 = Employee(1, "张三")
        employee2 = Employee(2, "李四")
        employee3 = Employee(6, "周八")
        hashtable.add(employee1)
        hashtable.add(employee2)
        hashtable.add(employee3)

        print("修改员工前:")
        # 显示哈希表内容
        hashtable.list()

        # 修改员工
        employee = Employee(6, "菜菜")
        hashtable.update(employee)

        print("修改员工后:")
        hashtable.list()
```

运行：

```shell
❯ python -m unittest test_main.Test.test_update_employee
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

```python
def delete(self, emp_id: int) -> None:
    """删除员工的方法"""
    # 计算下标
    index = emp_id % len(self.link_array)
    head_node = self.link_array[index]
    # 数组没有元素
    if not head_node:
        return
    # 判断头结点
    if head_node.employee.id == emp_id:
        self.link_array[index] = head_node.next
        return

    # 查找后续结点中需要删除的员工
    temp_node = head_node
    while temp_node.next is not None:
        if temp_node.next.employee.id == emp_id:
            # 将要删除结点的下一个结点链接到该结点的上一个结点
            temp_node.next = temp_node.next.next
            return
        temp_node = temp_node.next
```

编写测试删除的函数：

```python
class Test(unittest.TestCase):
    
    def test_delete_employee(self):
        hashtable = Hashtable(5)
        employee1 = Employee(2, "李四")
        employee2 = Employee(5, "孙七")
        hashtable.add(employee1)
        hashtable.add(employee2)

        print("删除员工前:")
        hashtable.list()

        # 删除员工
        hashtable.delete(2)

        print("删除员工后:")
        hashtable.list()
```

运行：

```shell
❯ python -m unittest test_main.Test.test_delete_employee
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
