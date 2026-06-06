from single_node import Node, get_last_index_node, get_node_length, reverse_node


def test_get_length():
    head_node = Node()
    node1 = Node("node1")
    node2 = Node("node2")
    node3 = Node("node3")

    head_node.next = node1
    node1.next = node2
    node2.next = node3

    length = get_node_length(head_node)
    print("单链表结点个数为: " + str(length))


def test_get_last_index_node():
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


def test_reverse_node():
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


def main():
    # test_get_length()
    # test_get_last_index_node()
    # test_reverse_node()
    pass


if __name__ == "__main__":
    main()
