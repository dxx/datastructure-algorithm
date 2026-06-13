"""
循环链表
循环链表的特点是表中最后一个结点的指针域指向头结点，整个链表形成一个环

双向循环链表
"""

from typing import cast


class PersonNode:
    def __init__(self, no: int | None = None, name: str | None = None) -> None:
        self.no = no
        self.name = name
        self.prev: PersonNode | None = None
        self.next: PersonNode | None = None


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
        last_node = cast(PersonNode, last_node.next)
    # 将新结点添加到链表末尾
    last_node.next = new_node
    new_node.prev = last_node
    # 将新结点下一个结点指针指向头结点
    new_node.next = head_node
    head_node.prev = new_node


def delete_node(head_node: PersonNode, node: PersonNode) -> PersonNode:
    """删除指定结点，返回头结点"""
    # 没有结点 或者 只有一个头结点
    if head_node.next is None or head_node.next == head_node:
        # 头结点就是要删除的结点
        if head_node.no == node.no:
            head_node.prev = None
            head_node.next = None
        return head_node

    temp_node = cast(PersonNode, head_node.next)
    is_exist = False
    while True:
        if temp_node == head_node:  # 最后一个结点
            if temp_node.no == node.no:
                is_exist = True
                # 头结点删除了，将头结点的下一个结点作为头结点
                head_node = cast(PersonNode, temp_node.next)
            break
        if temp_node.no == node.no:
            is_exist = True
            break
        temp_node = cast(PersonNode, temp_node.next)
    # 存在需要删除的结点
    if is_exist:
        prev_node = cast(PersonNode, temp_node.prev)
        next_node = cast(PersonNode, temp_node.next)
        # 将查找到的结点的上一个结点的下一个结点指针指向当前结点的下一个结点
        prev_node.next = next_node
        # 将查找到的结点的下一个结点的上一个结点指针指向当前指针的上一个结点
        next_node.prev = prev_node
    return head_node


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
        temp_node = cast(PersonNode, temp_node.next)
    result += "]"
    print(result)
