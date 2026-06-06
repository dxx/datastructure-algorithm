"""
循环链表
循环链表的特点是表中最后一个结点的指针域指向头结点，整个链表形成一个环

双向循环链表
"""


class PersonNode:
    def __init__(self, no=None, name=None):
        self.no = no
        self.name = name
        self.prev = None
        self.next = None


def insert_node(head_node, new_node):
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


def delete_node(head_node, node):
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
        # 将查找到的结点的上一个结点的下一个结点指针指向当前结点的下一个结点
        temp_node.prev.next = temp_node.next
        # 将查找到的结点的下一个结点的上一个结点指针指向当前指针的上一个结点
        temp_node.next.prev = temp_node.prev
    return head_node


def print_round_node_info(head_node):
    """打印循环链表的信息"""
    if not head_node.next:
        print("该链表没有节点")
        return
    result = "["
    temp_node = head_node
    while True:
        result += "{no:" + str(temp_node.no) + ", name:" + temp_node.name + "}"
        # 表示最后一个结点
        if temp_node.next == head_node:
            break
        temp_node = temp_node.next
    result += "]"
    print(result)


def test_insert_node():
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


def test_delete_node():
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


def main():
    # test_insert_node()
    # test_delete_node()
    pass


if __name__ == "__main__":
    main()
