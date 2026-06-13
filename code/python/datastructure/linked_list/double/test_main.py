import unittest
from main import HeroNode, delete_node, insert_at_tail, print_head_node_info


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
