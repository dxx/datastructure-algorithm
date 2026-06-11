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


def test_insert_at_tail() -> None:
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


def test_sort_insert_by_no() -> None:
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
    # 创建第一个结点
    hero_node1 = HeroNode(1, "宋江", "呼保义")
    # 创建第二个结点
    hero_node2 = HeroNode(2, "卢俊义", "玉麒麟")
    # 创建第三个结点
    hero_node3 = HeroNode(3, "吴用", "智多星")

    # 将结点按照 no 升序插入
    sort_insert_by_no(head_node, hero_node1)
    sort_insert_by_no(head_node, hero_node3)
    sort_insert_by_no(head_node, hero_node2)

    print("按照 no 升序插入的结果:")
    print_head_node_info(head_node)


def test_delete_node() -> None:
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
