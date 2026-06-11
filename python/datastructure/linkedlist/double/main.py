"""
双向链表
双向链表也叫双链表，是链表的一种，它的每个数据结点中都有两个指针，分别指向前一个和后一个结点
所以，从双向链表中的任意一个结点开始，都可以很方便地访问它的前一个结点和后一个结点。
"""

class HeroNode:
    def __init__(self, no: int | None = None, name: str | None = None, nickname: str | None = None) -> None:
        self.no = no  # 编号
        self.name = name  # 姓名
        self.nickname = nickname  # 昵称
        self.prev: HeroNode | None = None  # 上一个节点
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
    # 将新结点的上一个结点指向当前结点
    new_node.prev = last_node


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


def test_insert_at_tail() -> None:
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
