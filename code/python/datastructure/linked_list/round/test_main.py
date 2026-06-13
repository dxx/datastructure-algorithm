import unittest
from main import PersonNode, delete_node, insert_node, print_round_node_info


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
