## 链表

> 各种语言实现代码：[Go](../code/golang/datastructure/linkedlist)   [Java](../code/java/datastructure/src/com/dxx/linkedlist)   [JavaScript](../code/javascript/datastructure/linkedlist)   [TypeScript](../code/typescript/datastructure/linkedlist)   [Python](../code/python/datastructure/linked_list)   [Rust](../code/rust/datastructure/src/linked_list)
>
> 默认使用 **Python** 语言实现。

### 简介

链表是一种物理存储单元上非连续、非顺序，但是在逻辑上是顺序的存储结构。链表的每个结点包括两个部分，一个是存储数据元素的数据域，另一个是存储下一个结点地址的指针域，链表使用一个 head 指针来指向第一个表头的结点，而最后一个结点的指针指向 NULL。由于链表不必须按顺序存储，所以链表在插入的时候比线性顺序表快得多，但是查找或者访问某个节点则比线性表和顺序表要慢。链表有很多种不同的类型：单向链表，双向链表以及循环链表。

![data_structure_linkedlist_01](https://dxx.github.io/static-resource/datastructure-algorithm/images/data_structure_linkedlist_01.png)

### 单向链表

单向链表是链表的一种，其特点是链表的链接方向是单向的，对链表的访问要从头部开始。

![data_structure_linkedlist_02](https://dxx.github.io/static-resource/datastructure-algorithm/images/data_structure_linkedlist_02.png)

我们都知道水浒传里面有梁山英雄排名，以梁山英雄排名为例，用单向链表来写一个示例。

![shuihuzhuan](https://dxx.github.io/static-resource/datastructure-algorithm/images/shuihuzhuan.png)

首先定义一个类用来存储数据和下一个指针：

```python
class HeroNode:
    def __init__(self, no: int | None = None, name: str | None = None, nickname: str | None = None) -> None:
        self.no = no  # 编号
        self.name = name  # 姓名
        self.nickname = nickname  # 昵称
        self.next: HeroNode | None = None  # 下一个节点
```

为了方便查看内容，定义函数用来输出：

```python
def print_head_node_info(head_node: HeroNode) -> None:
    """打印单链表结点内容"""
    if not head_node.next:
        print("该链表没有节点")
        return
    result = "["
    temp_node = head_node.next
    while temp_node is not None:
        result += "{no:" + str(temp_node.no) + ", name:" + str(temp_node.name) + ", nickname:" + str(temp_node.nickname) + "}"
        temp_node = temp_node.next
    result += "]"
    print(result)
```

**尾部插入**

单向链表是从 head 结点开始，形成一条链状，最后一个结点的指针是 NULL。在尾部插入，通过 head 结点找到最后一个结点，然后将新的结点插入到最后一个结点即可。

实现如下：

```python
def insert_at_tail(head_node: HeroNode, new_node: HeroNode) -> None:
    """在链表尾部插入，通过 head 找到链表的尾部"""
    last_node = head_node
    # 下一个结点不为空继续循环
    while last_node.next is not None:
        # 将下一个结点赋值给当前结点
        last_node = last_node.next
    # 将当前结点插入到链表的最后一个结点
    last_node.next = new_node
```

编写测试方法，然后运行：

```python
class Test(unittest.TestCase):
    
    def test_insert_at_tail(self):
        # 创建 head 结点，head 结点不包含数据
        head_node = HeroNode()
        # 创建第一个结点
        hero_node1 = HeroNode(1, "宋江", "呼保义")
        # 创建第二个结点
        hero_node2 = HeroNode(2, "卢俊义", "玉麒麟")
        # 创建第三个结点
        hero_node3 = HeroNode(3, "吴用", "智多星")

        # 将结点添加到链表尾部
        insert_at_tail(head_node, hero_node1)
        insert_at_tail(head_node, hero_node2)
        insert_at_tail(head_node, hero_node3)

        print_head_node_info(head_node)
```

运行：

```shell
❯ python -m unittest test_main.Test.test_insert_at_tail
[{no:1, name:宋江, nickname:呼保义}{no:2, name:卢俊义, nickname:玉麒麟}{no:3, name:吴用, nickname:智多星}]
```

**顺序插入**

上例中实现了尾部插入，如果要按照某种顺序插入，在插入结点的时候找到合适的结点，然后将该结点的下一个结点指向要插入的结点。下面是一个按照 no 升序的顺序插入实现：

```python
def sort_insert_by_no(head_node: HeroNode, new_node: HeroNode) -> None:
    """按照 no 升序插入，通过 head 找到合适的插入位置"""
    temp_node = head_node
    while True:
        if temp_node.next is None:
            break
        if temp_node.next.no > new_node.no:
            break
        if temp_node.next.no == new_node.no:
            print("no 相等不能插入")
            return
        temp_node = temp_node.next
    # temp_node 的下一个结点插入到 new_node 的下一个结点
    new_node.next = temp_node.next
    # new_node 结点插入到 temp_node 的下一个结点
    temp_node.next = new_node
```

对比尾部插入，测试代码如下：

```python
class Test(unittest.TestCase):
    
    def test_sort_insert_by_no(self):
        # 创建结点，用来做尾部插入
        head = HeroNode()
        node1 = HeroNode(1, "宋江", "呼保义")
        node2 = HeroNode(2, "卢俊义", "玉麒麟")
        node3 = HeroNode(3, "吴用", "智多星")

        insert_at_tail(head, node1)
        insert_at_tail(head, node3)  # 将第三个结点插入到第二个位置
        insert_at_tail(head, node2)

        print("尾部插入的结果:")
        print_head_node_info(head)

        # 创建 head 结点
        head_node = HeroNode()
        hero_node1 = HeroNode(1, "宋江", "呼保义")
        hero_node2 = HeroNode(2, "卢俊义", "玉麒麟")
        hero_node3 = HeroNode(3, "吴用", "智多星")

        # 将结点按照 no 升序插入
        sort_insert_by_no(head_node, hero_node1)
        sort_insert_by_no(head_node, hero_node3)
        sort_insert_by_no(head_node, hero_node2)

        print("按照 no 升序插入的结果:")
        print_head_node_info(head_node)
```

运行：

```shell
❯ python -m unittest test_main.Test.test_sort_insert_by_no
尾部插入的结果:
[{no:1, name:宋江, nickname:呼保义}{no:3, name:吴用, nickname:智多星}{no:2, name:卢俊义, nickname:玉麒麟}]
按照 no 升序插入的结果:
[{no:1, name:宋江, nickname:呼保义}{no:2, name:卢俊义, nickname:玉麒麟}{no:3, name:吴用, nickname:智多星}]
```

可以看出使用顺序插入，插入节点时**对代码调用顺序无关**，其链表元素的顺序总是有序的。

**单向链表删除**

对于删除，先找到需要删除的结点，然后将需要删除的结点的上一个结点的下一个结点指针指向该删除结点的下一个结点。

实现如下：

```python
def delete_node(head_node: HeroNode, node: HeroNode) -> None:
    """删除指定结点"""
    temp_node = head_node
    while temp_node.next is not None:
        if temp_node.next.no == node.no:
            # 将下一个结点的下一个结点，链接到被删除结点的上一个结点
            temp_node.next = temp_node.next.next
            return
        temp_node = temp_node.next
```

编写测试方法：

```python
class Test(unittest.TestCase):
    
    def test_delete_node(self):
        # 创建结点
        head_node = HeroNode()
        hero_node1 = HeroNode(1, "宋江", "呼保义")
        hero_node2 = HeroNode(2, "卢俊义", "玉麒麟")
        hero_node3 = HeroNode(3, "吴用", "智多星")
        hero_node4 = HeroNode(4, "公孙胜", "入云龙")
        hero_node5 = HeroNode(5, "关胜", "大刀")

        # 插入结点
        insert_at_tail(head_node, hero_node1)
        insert_at_tail(head_node, hero_node2)
        insert_at_tail(head_node, hero_node3)
        insert_at_tail(head_node, hero_node4)
        insert_at_tail(head_node, hero_node5)

        print("删除前:")
        print_head_node_info(head_node)

        # 删除 no 为 2 的结点
        delete_node(head_node, hero_node2)
        print("删除 no 为 2 的结点后:")
        print_head_node_info(head_node)

        # 删除 no 为 3, 4 的结点
        delete_node(head_node, hero_node3)
        delete_node(head_node, hero_node4)
        print("删除 no 为 3,4 的结点后:")
        print_head_node_info(head_node)
```

运行：

```shell
❯ python -m unittest test_main.Test.test_delete_node
删除前:
[{no:1, name:宋江, nickname:呼保义}{no:2, name:卢俊义, nickname:玉麒麟}{no:3, name:吴用, nickname:智多星}{no:4, name:公孙胜, nickname:入云龙}{no:5, name:关胜, nickname:大刀}]
删除 no 为 2 的结点后:
[{no:1, name:宋江, nickname:呼保义}{no:3, name:吴用, nickname:智多星}{no:4, name:公孙胜, nickname:入云龙}{no:5, name:关胜, nickname:大刀}]
删除 no 为 3,4 的结点后:
[{no:1, name:宋江, nickname:呼保义}{no:5, name:关胜, nickname:大刀}]
```

### 单链表面试题

#### 获取单链表结点个数

遍历链表，统计出了头结点以外的结点，头结点不存储数据。

```python
class Node:
    def __init__(self, name: str | None = None) -> None:
        self.name = name
        self.next: Node | None = None

    def print_node_info(self) -> None:
        if not self.next:
            print("该链表没有节点")
            return
        result = "["
        temp_node = self.next
        while temp_node is not None:
            result += "{name:" + str(temp_node.name) + "}"
            temp_node = temp_node.next
        result += "]"
        print(result)


def get_node_length(head_node: Node | None) -> int:
    """获取单链表有效的结点数"""
    # 头结点为空，返回 0
    if not head_node:
        return 0
    length = 0
    node = head_node.next
    while node is not None:
        length += 1
        node = node.next
    return length
```

测试代码：

```python
class Test(unittest.TestCase):

    def test_get_length(self) -> None:
        head_node = Node()
        node1 = Node("node1")
        node2 = Node("node2")
        node3 = Node("node3")

        head_node.next = node1
        node1.next = node2
        node2.next = node3

        length = get_node_length(head_node)
        print("单链表结点个数为: " + str(length))
```

运行：

```shell
❯ python -m unittest test_main.Test.test_get_length
单链表结点个数为: 3
```

#### 获取倒数第 N 个结点

快慢指针：快指针先走 index 步，然后快慢指针一起移动，快指针到达末尾时，慢指针就是倒数第 N 个结点。

```python
def get_last_index_node(head_node: Node | None, index: int) -> Node | None:
    """获取倒数第 n 个结点"""
    # 头结点为空，index 小于等于 0 返回空
    if head_node is None or index <= 0:
        return None
    fast = head_node
    slow = head_node
    i = index
    length = 0
    while fast is not None:
        length += 1
        if i > 0:
            fast = fast.next
            i -= 1
            continue
        # 快慢指针同时走
        fast = fast.next
        if slow is None:
            return None
        slow = slow.next
    # index 超过了链表的长度
    if index > length:
        return None
    return slow
```

遍历方式：先获取链表结点数 length，再遍历到 length - n 个结点。

```python
def get_last_index_node2(head_node: Node | None, index: int) -> Node | None:
    """获取倒数第 n 个结点"""
    # 头结点为空，返回空
    if not head_node:
        return None
    length = get_node_length(head_node)
    if index <= 0 or index > length:
        return None
    last_node = head_node.next
    for _ in range(length - index):
        if last_node is None:
            return None
        last_node = last_node.next
    return last_node
```

测试代码：

```python
class Test(unittest.TestCase):

    def test_get_last_index_node(self) -> None:
        head_node = Node()
        node1 = Node("node1")
        node2 = Node("node2")
        node3 = Node("node3")
        head_node.next = node1
        node1.next = node2
        node2.next = node3
        index = 2
        last_node = get_last_index_node(head_node, index)
        if last_node is not None:
            print("单链表结点中倒数第 " + str(index) + " 个结点为: " + str(last_node.name))
```

运行：

```shell
❯ python -m unittest test_main.Test.test_get_last_index_node
单链表结点中倒数第 2 个结点为: node2
```

#### 单链表反转

1. 定义一个新的头结点 reverseHead。
2. 遍历链表，每遍历一个结点，将其取出，放在新的头结点 reverseHead 的后面。
3. 最后将头结点的 next 结点指向 reverseHead 的 next 结点。

```python
def reverse_node(head_node: Node | None) -> None:
    """单链表反转"""
    if head_node is None or head_node.next is None:
        return
    reverse_head = Node()
    current = head_node.next
    while current is not None:
        # 保存当前结点的下一个结点
        next_node = current.next
        # 将 reverseHead 结点的下一个结点放在当前结点的下一个结点
        current.next = reverse_head.next
        # 当前结点放在 reverseHead 后面
        reverse_head.next = current
        # 移动当前结点
        current = next_node
    # 将头结点的 next 结点指向 reverseHead 的 next 结点
    head_node.next = reverse_head.next
```

测试代码：

```python
class Test(unittest.TestCase):

    def test_reverse_node(self) -> None:
        head_node = Node()
        node1 = Node("node1")
        node2 = Node("node2")
        node3 = Node("node3")
        head_node.next = node1
        node1.next = node2
        node2.next = node3

        print("反转前:")
        head_node.print_node_info()

        reverse_node(head_node)

        print("反转后:")
        head_node.print_node_info()
```

运行：

```shell
❯ python -m unittest test_main.Test.test_reverse_node
反转前:
[{name:node1}{name:node2}{name:node3}]
反转后:
[{name:node3}{name:node2}{name:node1}]
```

#### 从尾到头打印单链表

反向遍历：先将链表反转，再遍历打印。问题是会破坏链表原来的结构。

使用栈：遍历链表，将每一个结点压入栈中，再遍历栈元素，将其打印输出。

### 双向链表

双向链表也叫双链表，是链表的一种，它的每个数据结点中都有两个指针，分别指向前一个和后一个结点。所以，从双向链表中的任意一个结点开始，都可以很方便地访问它的前一个结点和后一个结点。

![data_structure_linkedlist_03](https://dxx.github.io/static-resource/datastructure-algorithm/images/data_structure_linkedlist_03.png)

双向链表相比于单向链表有以下优点：

* 单向链表只能向后查找，双向链表可以向前或向后查找
* 单向链表是通过查找到被删除结点的上一个结点来删除，双向链表则是自我删除

接下来使用双向链表来实现上面梁山英雄排名的示例。

定义一个类用来存储数据和两个结点指针：

```python
class HeroNode:
    def __init__(self, no: int | None = None, name: str | None = None, nickname: str | None = None) -> None:
        self.no = no  # 编号
        self.name = name  # 姓名
        self.nickname = nickname  # 昵称
        self.prev: HeroNode | None = None  # 上一个节点
        self.next: HeroNode | None = None  # 下一个节点
```

为了方便查看内容，定义函数用来输出：

```python
def print_head_node_info(head_node: HeroNode) -> None:
    """打印链表结点内容"""
    if not head_node.next:
        print("该链表没有节点")
        return
    result = "["
    temp_node = head_node.next
    while temp_node is not None:
        result += "{no:" + str(temp_node.no) + ", name:" + str(temp_node.name) + ", nickname:" + str(temp_node.nickname) + "}"
        temp_node = temp_node.next
    result += "]"
    print(result)
```

**尾部插入**

和单向链表一样，通过 head 结点找到最后一个结点，然后将尾部结点的下一个指针结点指向新的结点，新结点的上一个结点指针指向尾部结点。

实现如下：

```python
def insert_at_tail(head_node: HeroNode, new_node: HeroNode) -> None:
    """在链表尾部插入，通过 head 找到链表的尾部"""
    last_node = head_node
    # 下一个结点不为空继续循环
    while last_node.next is not None:
        # 将下一个结点赋值给当前结点
        last_node = last_node.next
    # 将当前结点插入到链表的最后一个结点
    last_node.next = new_node
    # 将新结点的上一个结点指向当前结点
    new_node.prev = last_node
```

编写测试函数：

```python
class Test(unittest.TestCase):
    
    def test_insert_at_tail(self):
        # 创建 head 结点，head 结点不包含数据
        head_node = HeroNode()
        # 创建第一个结点
        hero_node1 = HeroNode(3, "吴用", "智多星")
        # 创建第二个结点
        hero_node2 = HeroNode(6, "林冲", "豹子头")
        # 创建第三个结点
        hero_node3 = HeroNode(7, "秦明", "霹雳火")

        # 将结点添加到链表尾部
        insert_at_tail(head_node, hero_node1)
        insert_at_tail(head_node, hero_node2)
        insert_at_tail(head_node, hero_node3)

        print_head_node_info(head_node)
```

运行：

```shell
❯ python -m unittest test_main.Test.test_insert_at_tail
[{no:3, name:吴用, nickname:智多星}{no:6, name:林冲, nickname:豹子头}{no:7, name:秦明, nickname:霹雳火}]
```

**双向链表删除**

双链表删除可以和单链表一样，只是在去掉被删除的节点后，需要将该删除结点的上下两个结点链接起来，实现如下：

```python
def delete_node(head_node: HeroNode, node: HeroNode) -> None:
    """删除指定结点"""
    temp_node = head_node.next
    while temp_node is not None:
        if temp_node.no == node.no:
            prev_node = temp_node.prev
            if prev_node is None:
                return
            # 将查找到的结点的上一个结点的下一个结点指针指向当前结点的下一个结点
            prev_node.next = temp_node.next
            # 最后一个结点的 next 指向空
            if temp_node.next is not None:
                # 将查找到的结点的下一个结点的上一个结点指针指向当前指针的上一个结点
                temp_node.next.prev = temp_node.prev
        temp_node = temp_node.next
```

编写测试代码：

```python
class Test(unittest.TestCase):
    
    def test_delete_node(self):
        # 创建结点
        head_node = HeroNode()
        hero_node1 = HeroNode(1, "宋江", "呼保义")
        hero_node2 = HeroNode(2, "卢俊义", "玉麒麟")
        hero_node3 = HeroNode(3, "吴用", "智多星")
        hero_node4 = HeroNode(4, "公孙胜", "入云龙")
        hero_node5 = HeroNode(5, "关胜", "大刀")

        # 插入结点
        insert_at_tail(head_node, hero_node1)
        insert_at_tail(head_node, hero_node2)
        insert_at_tail(head_node, hero_node3)
        insert_at_tail(head_node, hero_node4)
        insert_at_tail(head_node, hero_node5)

        print("删除前:")
        print_head_node_info(head_node)

        # 删除 no 为 2 的结点
        delete_node(head_node, hero_node2)
        print("删除 no 为 2 的结点后:")
        print_head_node_info(head_node)

        # 删除 no 为 3, 4 的结点
        delete_node(head_node, hero_node3)
        delete_node(head_node, hero_node4)
        print("删除 no 为 3,4 的结点后:")
        print_head_node_info(head_node)
```

运行：

```shell
❯ python -m unittest test_main.Test.test_delete_node
删除前:
[{no:1, name:宋江, nickname:呼保义}{no:2, name:卢俊义, nickname:玉麒麟}{no:3, name:吴用, nickname:智多星}{no:4, name:公孙胜, nickname:入云龙}{no:5, name:关胜, nickname:大刀}]
删除 no 为 2 的结点后:
[{no:1, name:宋江, nickname:呼保义}{no:3, name:吴用, nickname:智多星}{no:4, name:公孙胜, nickname:入云龙}{no:5, name:关胜, nickname:大刀}]
删除 no 为 3,4 的结点后:
[{no:1, name:宋江, nickname:呼保义}{no:5, name:关胜, nickname:大刀}]
```

### 循环链表(双向)

循环链表的特点是表中最后一个结点的指针域指向头结点，整个链表形成一个环。

![data_structure_linkedlist_04](https://dxx.github.io/static-resource/datastructure-algorithm/images/data_structure_linkedlist_04.png)

下面创建一个双向的循环链表类：

```python
class PersonNode:
    def __init__(self, no: int | None = None, name: str | None = None) -> None:
        self.no = no
        self.name = name
        self.prev: PersonNode | None = None
        self.next: PersonNode | None = None
```

定义函数用来输出：

```python
def print_round_node_info(head_node: PersonNode) -> None:
    """打印循环链表的信息"""
    if not head_node.next:
        print("该链表没有节点")
        return
    result = "["
    temp_node = head_node
    while True:
        result += "{no:" + str(temp_node.no) + ", name:" + str(temp_node.name) + "}"
        # 表示最后一个结点
        if temp_node.next == head_node:
            break
        temp_node = temp_node.next
    result += "]"
    print(result)
```

**尾部插入**

循环链表和非循环链表的头结点不一样，循环链表的头结点会存储第一个结点信息并且本身作为第一个结点，插入结点方法如下：

```python
def insert_node(head_node: PersonNode, new_node: PersonNode) -> None:
    """插入结点"""
    # 判断是否第一次插入
    if head_node.next is None:
        head_node.no = new_node.no
        head_node.name = new_node.name
        head_node.prev = head_node
        head_node.next = head_node
        return
    last_node = head_node
    # 下一个结点不等于头结点继续循环
    while last_node.next != head_node:
        last_node = last_node.next
    # 将新结点添加到链表末尾
    last_node.next = new_node
    new_node.prev = last_node
    # 将新结点下一个结点指针指向头结点
    new_node.next = head_node
    head_node.prev = new_node
```

编写测试函数代码：

```python
class Test(unittest.TestCase):
    
    def test_insert_node(self):
        # 创建 head 结点，head 结点不初始化数据，等到添加了第一个结点后才初始化数据
        head_node = PersonNode()
        # 创建第一个结点
        person_node1 = PersonNode(1, "张三")
        # 创建第二个结点
        person_node2 = PersonNode(2, "李四")
        # 创建第三个结点
        person_node3 = PersonNode(3, "王五")

        # 插入结点
        insert_node(head_node, person_node1)
        insert_node(head_node, person_node2)
        insert_node(head_node, person_node3)

        print_round_node_info(head_node)
```

运行：

```shell
❯ python -m unittest test_main.Test.test_insert_node
[{no:1, name:张三}{no:2, name:李四}{no:3, name:王五}]
```

**删除结点**

双向循环链表的头结点存储了第一个结点信息，也要作为判断是否要被删除的结点。遍历结点，当结点的下一个结点指针等于头结点就说明遍历完成，循环链表的删除有可能会删除头结点，当头结点被删除后，头结点就不在链表中，需要返回一个新的头结点，实现如下：

```python
def delete_node(head_node: PersonNode, node: PersonNode) -> PersonNode:
    """删除指定结点，返回头结点"""
    # 没有结点 或者 只有一个头结点
    if head_node.next is None or head_node.next == head_node:
        # 头结点就是要删除的结点
        if head_node.no == node.no:
            head_node.prev = None
            head_node.next = None
        return head_node

    temp_node = head_node.next
    is_exist = False
    while True:
        if temp_node == head_node:  # 最后一个结点
            if temp_node.no == node.no:
                is_exist = True
                # 头结点删除了，将头结点的下一个结点作为头结点
                head_node = temp_node.next
            break
        if temp_node.no == node.no:
            is_exist = True
            break
        temp_node = temp_node.next
    # 存在需要删除的结点
    if is_exist:
        prev_node = temp_node.prev
        next_node = temp_node.next
        # 将查找到的结点的上一个结点的下一个结点指针指向当前结点的下一个结点
        prev_node.next = next_node
        # 将查找到的结点的下一个结点的上一个结点指针指向当前指针的上一个结点
        next_node.prev = prev_node
    return head_node
```

编写测试删除结点的函数：

```python
class Test(unittest.TestCase):
    
    def test_delete_node(self):
        # 创建结点
        head_node = PersonNode()
        person_node1 = PersonNode(1, "张三")
        person_node2 = PersonNode(2, "李四")
        person_node3 = PersonNode(3, "王五")
        person_node4 = PersonNode(4, "赵六")
        person_node5 = PersonNode(5, "孙七")

        # 插入结点
        insert_node(head_node, person_node1)
        insert_node(head_node, person_node2)
        insert_node(head_node, person_node3)
        insert_node(head_node, person_node4)
        insert_node(head_node, person_node5)

        print("删除前:")
        print_round_node_info(head_node)

        # 删除 no 为 2 的结点
        head_node = delete_node(head_node, person_node2)
        print("删除 no 为 2 的结点后:")
        print_round_node_info(head_node)

        new_node = PersonNode(6, "周八")

        insert_node(head_node, new_node)
        print("插入新结点:")
        print_round_node_info(head_node)

        # 删除 no 为 1，3 的结点
        head_node = delete_node(head_node, person_node1)
        head_node = delete_node(head_node, person_node3)
        print("删除 no 为 1,3 的结点后:")
        print_round_node_info(head_node)
```

运行：

```shell
❯ python -m unittest test_main.Test.test_delete_node
删除前:
[{no:1, name:张三}{no:2, name:李四}{no:3, name:王五}{no:4, name:赵六}{no:5, name:孙七}]
删除 no 为 2 的结点后:
[{no:1, name:张三}{no:3, name:王五}{no:4, name:赵六}{no:5, name:孙七}]
插入新结点:
[{no:1, name:张三}{no:3, name:王五}{no:4, name:赵六}{no:5, name:孙七}{no:6, name:周八}]
删除 no 为 1,3 的结点后:
[{no:4, name:赵六}{no:5, name:孙七}{no:6, name:周八}]
```

### Joseph 问题

设编号为 1,2,...n 的 n 个人围坐一圈，约定编号为 k(1<=k<=n) 的人从 1 开始报数，数到 m 的那个人出列，它的下一位又从 1 开始报数，数到 m 的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列。

提示：用一个不带头节点的循环链表来处理 Josephu 问题：先构成一个有 n 个结点的单循环链表，然后由 k 节点起从 1 开始计数，计到 m 时，对应节点的人从链表中删除，然后再从被删除节点的下一个节点又从 1 开始计数，直到链表中留下最后一个节点。

假设有 5 个人，从第 1 个人开始，每次数到第 3 个人就出列，画图分析如下：

![josephu_01](https://dxx.github.io/static-resource/datastructure-algorithm/images/josephu_01.png)

![josephu_02](https://dxx.github.io/static-resource/datastructure-algorithm/images/josephu_02.png)

![josephu_03](https://dxx.github.io/static-resource/datastructure-algorithm/images/josephu_03.png)

![josephu_04](https://dxx.github.io/static-resource/datastructure-algorithm/images/josephu_04.png)

**代码实现**

创建类：

```python
class Person:
    def __init__(self, no: int) -> None:
        self.no = no
        self.prev: Person | None = None
        self.next: Person | None = None


class PersonLinkedList:
    def __init__(self, count: int) -> None:
        if count <= 0:
            print("链表至少需要一个元素")
            self.first = None
            self.length = 0
            return
        self.first = None
        self.length = count

        prev = self.first
        # 初始化小孩，构成双向循环链表
        for i in range(1, count + 1):
            person = Person(i)
            if i == 1:
                # 初始化 First 节点
                self.first = person
                self.first.next = self.first
                self.first.prev = self.first

                # 将 prev 指向 First 节点，继续下一次循环
                prev = self.first
                continue

            prev.next = person
            person.prev = prev

            # 新增加的节点的下一个节点指向第一节点
            person.next = self.first
            # 第一个节点的上一个节点指向新增加的节点
            self.first.prev = person

            prev = person
```

显示所有人的编号方法：

```python
def show_persons(self) -> None:
    if self.first is None:
        return
    current = self.first
    while True:
        print("num:" + str(current.no))
        if current.next is None:
            break
        current = current.next
        if current == self.first:
            break
```

报数方法：

```python
def count(self, start: int, num: int) -> None:
    if start < 1 or start > self.length:
        print("start 不能小于 1 或者不能大于 " + str(self.length))
        return
    if num > self.length:
        print("num 不能大于元素个数: " + str(self.length))
        return

    current = self.first
    if current is None:
        return

    for _ in range(1, start):
        if current.next is None:
            return
        current = current.next

    while True:
        if current.prev is None or current.next is None:
            return
        if current.prev == current and current.next == current:
            break

        for _ in range(1, num):
            if current.next is None:
                return
            current = current.next

        prev_node = current.prev
        next_node = current.next
        if prev_node is None or next_node is None:
            return
        prev_node.next = next_node
        next_node.prev = prev_node

        print("出队人的编号: " + str(current.no))
        current = next_node
    print("最后留下人的编号: " + str(current.no))
```

测试代码如下：

```python
class Test(unittest.TestCase):
    
    def test_person_linked_list(self):
        person_linked_list = PersonLinkedList(5)
        person_linked_list.show_persons()

        person_linked_list.count(1, 3)
```

运行：

```shell
❯ python -m unittest test_main.Test.test_person_linked_list
num:1
num:2
num:3
num:4
num:5
出队人的编号: 3
出队人的编号: 1
出队人的编号: 5
出队人的编号: 2
最后留下人的编号: 4
```
