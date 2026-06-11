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
    """
    获取单链表有效的结点数
    1.遍历结点数
    2.定义一个长度遍历，每遍历一次长度 +1
    """
    # 头结点为空，返回 0
    if not head_node:
        return 0
    length = 0
    node = head_node.next
    while node is not None:
        length += 1
        node = node.next
    return length


def get_last_index_node(head_node: Node | None, index: int) -> Node | None:
    """
    获取倒数第 n 个结点

    快慢指针
    1.定义快指针 fast 和 慢指针 slow
    2.快慢指针的初始值指向头结点
    3.快指针先走 index 步
    4.慢指针开始走直到快指针指向了末尾结点
    5.此时慢指针就是倒数第 n 个结点
    """
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


def get_last_index_node2(head_node: Node | None, index: int) -> Node | None:
    """
    获取倒数第 n 个结点

    遍历
    1.获取链表结点数 length
    2.遍历到 length - n 个结点
    3.然后返回
    """
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


def reverse_node(head_node: Node | None) -> None:
    """
    单链表反转
    1.定义一个新的头结点 reverseHead
    2.遍历链表，每遍历一个结点，将其取出，放在新的头结点 reverseHead 的后面
    3.最后将头结点的 next 结点指向 reverseHead 的 next 结点
    """
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
