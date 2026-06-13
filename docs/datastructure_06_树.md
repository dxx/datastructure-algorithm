## 树

> 各种语言实现代码：[Go](./golang/datastructure/tree)   [Java](java/datastructure/src/com/dxx/tree)   [JavaScript](javascript/datastructure/tree)   [TypeScript](./typescript/datastructure/tree)   [Python](./python/datastructure/tree)   [Rust](./rust/datastructure/src/tree)
>
> 默认使用 **Python** 语言实现。

[二叉树](./datastructure_06_树2.md#二叉树)

[哈夫曼树](./datastructure_06_树2.md#哈夫曼树)

[二叉排序树](./datastructure_06_树2.md#二叉排序树)

[AVL树](./datastructure_06_树2.md#AVL树)

### 简介

树是一种数据结构，它是由 n（n>=0）个结点组成一个具有层次关系的集合。因为它看起来像一棵倒挂的树，所以把它叫做树，也就是说它是根朝上，而叶朝下的。它具有以下的特点：

* 每个结点有零个或多个子结点。
* 没有父结点的结点称为根结点。
* 每一个非根结点有且只有一个父结点。
* 除了根结点外，每个子结点可以分为多个不相交的子树。

### 二叉树

二叉树是每个结点最多有两个子树的树结构。通常子树被称作“左子树”（left subtree）和“右子树”（right subtree）。二叉树常被用于实现二叉查找树和二叉堆。

#### 遍历

遍历二叉树就是按一定的规则和顺序走遍二叉树的所有结点，使每一个结点都被访问一次，而且只被访问一次。

![data_structure_tree_01](https://dxx.github.io/static-resource/datastructure-algorithm/images/data_structure_tree_01.png)

创建二叉树节点类：

```python
class BinaryTreeNode:
    def __init__(self, no: int) -> None:
        self.no = no
        self.left: BinaryTreeNode | None = None
        self.right: BinaryTreeNode | None = None
```

前序、中序、后序遍历：

```python
def pre_order(node: BinaryTreeNode | None) -> None:
    """前序遍历"""
    if not node:
        return
    print(node.no)
    pre_order(node.left)
    pre_order(node.right)


def infix_order(node: BinaryTreeNode | None) -> None:
    """中序遍历"""
    if not node:
        return
    infix_order(node.left)
    print(node.no)
    infix_order(node.right)


def post_order(node: BinaryTreeNode | None) -> None:
    """后序遍历"""
    if not node:
        return
    post_order(node.left)
    post_order(node.right)
    print(node.no)
```

手动初始化节点函数：

```python
def init_node() -> BinaryTreeNode:
    root = BinaryTreeNode(1)
    node2 = BinaryTreeNode(2)
    node3 = BinaryTreeNode(3)
    node4 = BinaryTreeNode(4)
    node5 = BinaryTreeNode(5)

    root.left = node2
    root.right = node5
    node2.left = node3
    node2.right = node4

    return root
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_order(self):
        root = init_node()

        print("======前序遍历======")
        pre_order(root)

        print("======中续遍历======")
        infix_order(root)

        print("======后续遍历======")
        post_order(root)
```

运行：

```shell
❯ python -m unittest test_binary_tree.Test.test_order
======前序遍历======
1
2
3
4
5
======中续遍历======
3
2
4
1
5
======后续遍历======
3
4
2
5
1
```

#### 查找节点

```python
def pre_order_search(node: BinaryTreeNode | None, no: int) -> BinaryTreeNode | None:
    """前序查找"""
    if not node:
        return None
    print("进入查找")
    if node.no == no:
        return node
    return_node = pre_order_search(node.left, no)
    if return_node:
        return return_node
    return pre_order_search(node.right, no)


def infix_order_search(node: BinaryTreeNode | None, no: int) -> BinaryTreeNode | None:
    """中序查找"""
    if not node:
        return None
    return_node = infix_order_search(node.left, no)
    if return_node:
        return return_node
    print("进入查找")
    if node.no == no:
        return node
    return infix_order_search(node.right, no)


def post_order_search(node: BinaryTreeNode | None, no: int) -> BinaryTreeNode | None:
    """后序查找"""
    if not node:
        return None
    return_node = post_order_search(node.left, no)
    if return_node:
        return return_node
    return_node = post_order_search(node.right, no)
    if return_node is not None:
        return return_node
    print("进入查找")
    if node.no == no:
        return_node = node
    return return_node
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_search(self):
        root = init_node()
        no = 4

        print("======前序查找======")
        print("查找no=" + str(no))
        node = pre_order_search(root, no)
        if node is not None:
            print("查找结果: no=" + str(node.no))

        print("======中序查找======")
        print("查找no=" + str(no))
        node = infix_order_search(root, no)
        if node is not None:
            print("查找结果: no=" + str(node.no))

        print("======后序查找======")
        print("查找no=" + str(no))
        node = post_order_search(root, no)
        if node is not None:
            print("查找结果: no=" + str(node.no))
```

运行：

```shell
❯ python -m unittest test_binary_tree.Test.test_search
======前序查找======
查找no=4
进入查找
进入查找
进入查找
进入查找
查找结果: no=4
======中序查找======
查找no=4
进入查找
进入查找
进入查找
查找结果: no=4
======后序查找======
查找no=4
进入查找
进入查找
查找结果: no=4
```

#### 二叉树顺序存储

将二叉树存储在一个数组中，通过存储元素的下标反映元素之间的父子关系。

```python
class SeqBinaryTree:
    def __init__(self, array: list[int]) -> None:
        self.array = array

    def pre_order(self) -> None:
        self._pre_order_from_index(0)

    def _pre_order_from_index(self, index: int) -> None:
        if not self.array:
            return
        length = len(self.array)
        if length == 0 or index >= length:
            return
        print(self.array[index])
        left_index = 2 * index + 1
        right_index = 2 * index + 2
        self._pre_order_from_index(left_index)
        self._pre_order_from_index(right_index)
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_pre_order(self):
        nos = [1, 2, 3, 4, 5, 6, 7]
        seq_binary_tree = SeqBinaryTree(nos)

        print("======前序遍历======")
        seq_binary_tree.pre_order()
```

运行：

```shell
❯ python -m unittest test_binary_tree_seq_storage.Test.test_pre_order
======前序遍历======
1
2
4
5
3
6
7
```

#### 二叉树线索化

对二叉树以某种遍历方式进行遍历，使其变为线索二叉树的过程称为对二叉树进行线索化。

```python
class ThreadedBinaryTreeNode:
    def __init__(self, no: int) -> None:
        self.no = no
        self.left: ThreadedBinaryTreeNode | None = None
        self.right: ThreadedBinaryTreeNode | None = None
        self.left_tag = 0
        self.right_tag = 0

    def search(self, no: int) -> "ThreadedBinaryTreeNode | None":
        left_child_node = self
        while left_child_node.left is not None:
            left_child_node = left_child_node.left
        while left_child_node is not None:
            if left_child_node.no == no:
                break
            left_child_node = left_child_node.right
        return left_child_node


previous: ThreadedBinaryTreeNode | None = None


def infix_thread_tree(node: ThreadedBinaryTreeNode | None) -> None:
    """中序线索化二叉树"""
    global previous
    if not node:
        return
    infix_thread_tree(node.left)
    if node.left is None:
        node.left = previous
        node.left_tag = 1
    if previous is not None and previous.right is None:
        previous.right = node
        previous.right_tag = 1
    previous = node
    infix_thread_tree(node.right)
```

中序遍历线索化二叉树：

```python
def infix_order_threaded_tree(node: ThreadedBinaryTreeNode | None) -> None:
    current_node = node
    while current_node is not None:
        while current_node.left_tag == 0:
            if current_node.left is None:
                return
            current_node = current_node.left
        print("id:" + str(current_node.no))
        while current_node.right_tag == 1:
            if current_node.right is None:
                return
            current_node = current_node.right
            print("id:" + str(current_node.no))
        current_node = current_node.right
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_infix_threaded_tree(self):
        root = init_threaded_node()
        infix_thread_tree(root)
        node = root.search(10)
        if node is None or node.left is None or node.right is None:
            return
        print("no=" + str(node.no) + "的前驱节点为" + str(node.left.no))
        print("no=" + str(node.no) + "的后继节点为" + str(node.right.no))

    def test_infix_order_threaded_tree(self):
        root = init_threaded_node()
        infix_thread_tree(root)
        infix_order_threaded_tree(root)
```

运行：

```shell
❯ python -m unittest test_binary_tree_threaded.Test.test_infix_threaded_tree
no=10的前驱节点为2
no=10的后继节点为1

❯ python -m unittest test_binary_tree_threaded.Test.test_infix_order_threaded_tree
id:8
id:2
id:10
id:1
id:16
id:6
```

#### 堆排序

```python
def heap_sort(nums: list[int]) -> None:
    if not nums:
        return
    for i in range(len(nums) // 2 - 1, -1, -1):
        adjust_heap(nums, i, len(nums))
    for i in range(len(nums) - 1, 0, -1):
        nums[0], nums[i] = nums[i], nums[0]
        adjust_heap(nums, 0, i)


def adjust_heap(nums: list[int], i: int, count: int) -> None:
    temp = nums[i]
    j = 2 * i + 1
    while j < count:
        if j + 1 < count and nums[j] < nums[j + 1]:
            j += 1
        if nums[j] > temp:
            nums[i] = nums[j]
            i = j
        else:
            break
        j = 2 * j + 1
    nums[i] = temp
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_heap_sort(self):
        nums = [1, 7, 5, 2, 8]
        print("排序前: " + str(nums))
        heap_sort(nums)
        print("排序后: " + str(nums))
```

运行：

```shell
❯ python -m unittest test_heap_sort.Test.test_heap_sort
排序前: [1, 7, 5, 2, 8]
排序后: [1, 2, 5, 7, 8]
```

### 哈夫曼树

有 N 个权值作为 N 个叶子结点，构造一棵二叉树，如果该树的带权路径长度达到最小，称这样的二叉树为最优二叉树，也称为哈夫曼树(Huffman Tree)。

```python
class Node:
    def __init__(self, value: int) -> None:
        self.value = value
        self.left: Node | None = None
        self.right: Node | None = None


def create_huffman_tree(nums: list[int]) -> Node | None:
    if not nums:
        return None
    nodes = []
    for num in nums:
        nodes.append(Node(num))
    while len(nodes) > 1:
        nodes.sort(key=lambda node: node.value)
        left = nodes[0]
        right = nodes[1]
        root = Node(left.value + right.value)
        root.left = left
        root.right = right
        nodes.pop(0)
        nodes.pop(0)
        nodes.append(root)
    return nodes[0]


def pre_order(node: Node | None) -> None:
    if node is None:
        return
    print(node.value)
    pre_order(node.left)
    pre_order(node.right)
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_huffman(self):
        nums = [1, 7, 3, 8, 16]
        root = create_huffman_tree(nums)
        pre_order(root)
```

运行：

```shell
❯ python -m unittest test_huffman.Test.test_huffman
35
16
19
8
11
4
1
3
7
```

### 二叉排序树

二叉排序树又称二叉查找树，也叫二叉搜索树。对于任何一个非叶子节点，它的左子节点的值小于自身的值，它的右子节点的值大于自身的值。

```python
class BinaryTreeNode:
    def __init__(self, no: int) -> None:
        self.no = no
        self.left: BinaryTreeNode | None = None
        self.right: BinaryTreeNode | None = None


class BinarySortTree:
    def __init__(self) -> None:
        self.root: BinaryTreeNode | None = None

    def add(self, node: BinaryTreeNode | None) -> None:
        if node is None:
            return
        if self.root is None:
            self.root = node
            return
        self._add(self.root, node)

    def _add(self, root: BinaryTreeNode, node: BinaryTreeNode) -> None:
        if node.no < root.no:
            if root.left is None:
                root.left = node
                return
            self._add(root.left, node)
        else:
            if root.right is None:
                root.right = node
                return
            self._add(root.right, node)

    def infix_order(self) -> None:
        self._infix_order(self.root)

    def _infix_order(self, node: BinaryTreeNode | None) -> None:
        if node is None:
            return
        self._infix_order(node.left)
        print("no:" + str(node.no))
        self._infix_order(node.right)
```

删除节点：

```python
def delete(self, no: int) -> None:
    binary_tree_nodes = self.search(no)
    if not binary_tree_nodes:
        return
    parent_node = binary_tree_nodes[0]
    node = binary_tree_nodes[1]
    if node.left is None and node.right is None:
        if parent_node is None:
            self.root = None
            return
        if parent_node.left is not None and parent_node.left.no == no:
            parent_node.left = None
        if parent_node.right is not None and parent_node.right.no == no:
            parent_node.right = None
        return
    if node.left is not None and node.right is not None:
        left_child_node = node.right
        while left_child_node.left is not None:
            left_child_node = left_child_node.left
        self.delete(left_child_node.no)
        node.no = left_child_node.no
    else:
        replace_node = node.left if node.left is not None else node.right
        if parent_node is None:
            self.root = replace_node
            return
        if parent_node.left is not None and parent_node.left.no == no:
            parent_node.left = replace_node
        if parent_node.right is not None and parent_node.right.no == no:
            parent_node.right = replace_node
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_binary_sort_tree(self):
        nos = [8, 5, 10, 3, 6, 9, 12, 2]
        binary_sort_tree = BinarySortTree()
        for no in nos:
            binary_sort_tree.add(BinaryTreeNode(no))
        print("======中序遍历======")
        binary_sort_tree.infix_order()

        binary_sort_tree.delete(6)
        print("======删除叶子节点 6======")
        binary_sort_tree.infix_order()

        binary_sort_tree.delete(5)
        print("======删除只有一颗子树的节点 5======")
        binary_sort_tree.infix_order()

        binary_sort_tree.delete(10)
        print("======删除有两颗子树的节点 10======")
        binary_sort_tree.infix_order()
```

运行：

```shell
❯ python -m unittest test_binary_sort_tree.Test.test_binary_sort_tree
======中序遍历======
no:2
no:3
no:5
no:6
no:8
no:9
no:10
no:12
======删除叶子节点 6======
no:2
no:3
no:5
no:8
no:9
no:10
no:12
======删除只有一颗子树的节点 5======
no:2
no:3
no:8
no:9
no:10
no:12
======删除有两颗子树的节点 10======
no:2
no:3
no:8
no:9
no:12
```

### AVL树

AVL 树本质上是带了平衡功能的二叉排序树。任何节点的两个子树的高度最大差别为 1，所以它也被称为高度平衡树。

```python
class AVLTree:
    def __init__(self) -> None:
        self.root: BinaryTreeNode | None = None

    def _height(self, node: BinaryTreeNode | None) -> int:
        if not node:
            return 0
        return max(self._height(node.left), self._height(node.right)) + 1

    def _left_height(self) -> int:
        if self.root is None:
            return 0
        return self._height(self.root.left)

    def _right_height(self) -> int:
        if self.root is None:
            return 0
        return self._height(self.root.right)
```

左旋转和右旋转：

```python
def _left_rotate(self, node: BinaryTreeNode | None) -> None:
    if not node:
        return
    new_node = BinaryTreeNode(node.no)
    if node.right is None:
        return
    new_node.left = node.left
    new_node.right = node.right.left
    node.no = node.right.no
    node.right = node.right.right
    node.left = new_node


def _right_rotate(self, node: BinaryTreeNode | None) -> None:
    if not node:
        return
    new_node = BinaryTreeNode(node.no)
    if node.left is None:
        return
    new_node.right = node.right
    new_node.left = node.left.right
    node.no = node.left.no
    node.left = node.left.left
    node.right = new_node
```

添加节点时判断是否旋转：

```python
def add(self, node: BinaryTreeNode | None) -> None:
    if node is None:
        return
    if self.root is None:
        self.root = node
        return
    self._add(self.root, node)

    if self._right_height() - self._left_height() > 1:
        right_node = self.root.right
        if right_node is not None and self._height(right_node.left) > self._height(right_node.right):
            self._right_rotate(right_node)
        self._left_rotate(self.root)
        return

    if self._left_height() - self._right_height() > 1:
        left_node = self.root.left
        if left_node is not None and self._height(left_node.right) > self._height(left_node.left):
            self._left_rotate(left_node)
        self._right_rotate(self.root)
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_left_rotate(self):
        nos = [3, 2, 5, 4, 6, 7]
        avl_tree = AVLTree()
        for no in nos:
            avl_tree.add(BinaryTreeNode(no))
        print("左旋转后")
        avl_tree.infix_order()
        if avl_tree.root is not None:
            print("根节点 = ", avl_tree.root.no)
        print("左子树的高度为: ", avl_tree._left_height())
        print("右子树的高度为: ", avl_tree._right_height())

    def test_right_rotate(self):
        nos = [6, 4, 7, 3, 5, 2]
        avl_tree = AVLTree()
        for no in nos:
            avl_tree.add(BinaryTreeNode(no))
        print("右旋转后")
        avl_tree.infix_order()
        if avl_tree.root is not None:
            print("根节点 = ", avl_tree.root.no)
        print("左子树的高度为: ", avl_tree._left_height())
        print("右子树的高度为: ", avl_tree._right_height())
```

运行：

```shell
❯ python -m unittest test_avl_tree.Test.test_left_rotate
左旋转后
no:2
no:3
no:4
no:5
no:6
no:7
根节点 =  5
左子树的高度为:  2
右子树的高度为:  2

❯ python -m unittest test_avl_tree.Test.test_right_rotate
右旋转后
no:2
no:3
no:4
no:5
no:6
no:7
根节点 =  4
左子树的高度为:  2
右子树的高度为:  2

❯ python -m unittest test_avl_tree.Test.test_double_rotate
双旋转后
no:2
no:3
no:4
no:5
no:6
no:7
根节点 =  4
左子树的高度为:  2
右子树的高度为:  2
```
