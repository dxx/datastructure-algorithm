"""
单向链表
单向链表是链表的一种，其特点是链表的链接方向是单向的，对链表的访问要从头部开始
链表是由结点构成，head 指针指向第一个称为表头的结点，而最后一个结点的指针指向 NULL
"""

from typing import cast


class HeroNode:
    def __init__(self, no: int | None = None, name: str | None = None, nickname: str | None = None) -> None:
        self.no = no  # 编号
        self.name = name  # 姓名
        self.nickname = nickname  # 昵称
        self.next: HeroNode | None = None  # 下一个节点


def insert_at_tail(head_node: HeroNode, new_node: HeroNode) -> None:
    """在链表尾部插入，通过 head 找到链表的尾部"""
    last_node = head_node
    # 下一个结点不为空继续循环
    while last_node.next is not None:
        # 将下一个结点赋值给当前结点
        last_node = last_node.next
    # 将当前结点插入到链表的最后一个结点
    last_node.next = new_node


def sort_insert_by_no(head_node: HeroNode, new_node: HeroNode) -> None:
    """按照 no 升序插入，通过 head 找到合适的插入位置"""
    temp_node = head_node
    while True:
        if temp_node.next is None:
            break
        if cast(int, temp_node.next.no) > cast(int, new_node.no):
            break
        if temp_node.next.no == new_node.no:
            print("no 相等不能插入")
            return
        temp_node = temp_node.next
    # tempNode 的下一个结点插入到 newNode 的下一个结点
    new_node.next = temp_node.next
    # newNode 结点插入到 tempNode 的下一个结点
    temp_node.next = new_node


def delete_node(head_node: HeroNode, node: HeroNode) -> None:
    """删除指定结点"""
    temp_node = head_node
    while temp_node.next is not None:
        if temp_node.next.no == node.no:
            # 将下一个结点的下一个结点，链接到被删除结点的上一个结点
            temp_node.next = temp_node.next.next
            return
        temp_node = temp_node.next


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
