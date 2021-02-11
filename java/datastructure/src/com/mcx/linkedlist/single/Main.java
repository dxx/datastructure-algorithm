package com.mcx.linkedlist.single;

/**
 * 单向链表
 * 单向链表是链表的一种，其特点是链表的链接方向是单向的，对链表的访问要从头部开始
 * 链表是由结点构成，head 指针指向第一个称为表头的结点，而最后一个结点的指针指向 NULL
 */
public class Main {
    public static class HeroNode {
        private int no; // 编号
        private String name; // 姓名
        private String nickname; // 昵称
        private HeroNode next; // 下一个节点

        public HeroNode(){}

        public HeroNode(int no, String name, String nickname) {
            this.no = no;
            this.name = name;
            this.nickname = nickname;
        }

    }

    /**
     * 在链表尾部插入，通过 head 找到链表的尾部
     */
    public static void insertAtTail(HeroNode headNode, HeroNode newNode) {
        HeroNode lastNode = headNode;
        // 下一个结点不为空继续循环
        while (lastNode.next != null) {
            // 将下一个结点赋值给当前结点
            lastNode = lastNode.next;
        }
        // 将当前结点插入到链表的最后一个结点
        lastNode.next = newNode;
    }

    /**
     * 按照 no 升序插入，通过 head 找到合适的插入位置
     */
    public static void sortInsertByNo(HeroNode headNode, HeroNode newNode) {
        HeroNode tempNode = headNode;
        while (true) {
            if (tempNode.next == null) {
                break;
            } else if (tempNode.next.no > newNode.no) {
                break;
            } else if (tempNode.next.no == newNode.no) {
                throw new IllegalStateException("no 相等不能插入");
            }
            tempNode = tempNode.next;
        }
        // tempNode 的下一个结点插入到 newNode 的下一个结点
        newNode.next = tempNode.next;
        // newNode 结点插入到 tempNode 的下一个结点
        tempNode.next = newNode;
    }

    /**
     * 删除指定结点
     */
    public static void deleteNode(HeroNode headNode, HeroNode node) {
        HeroNode tempNode = headNode;
        while (tempNode.next != null) {
            if (tempNode.next.no == node.no) {
                // 将下一个结点的下一个结点，链接到被删除结点的上一个结点
                tempNode.next = tempNode.next.next;
                return;
            }
            tempNode = tempNode.next;
        }
    }

    /**
     * 打印单链表结点内容
     */
    public static void printHeadNodeInfo(HeroNode headNode) {
        if (headNode.next == null) {
            System.out.println("该链表没有节点");
            return;
        }
        StringBuilder sb = new StringBuilder("[");
        HeroNode tempNode = headNode.next;
        while (tempNode != null) {
            sb.append(String.format("{no:%d, name:%s, nickname:%s}",
                    tempNode.no, tempNode.name, tempNode.nickname));
            tempNode = tempNode.next;
        }
        sb.append("]");
        System.out.println(sb.toString());
    }

    public static void testInsertAtTail() {
        // 创建 head 结点，head 结点不包含数据
        HeroNode headNode = new HeroNode();
        // 创建第一个结点
        HeroNode heroNode1 = new HeroNode(1, "宋江", "呼保义");
        // 创建第二个结点
        HeroNode heroNode2 = new HeroNode(2, "卢俊义", "玉麒麟");
        // 创建第三个结点
        HeroNode heroNode3 = new HeroNode(3, "吴用", "智多星");

        // 将结点添加到链表尾部
        insertAtTail(headNode, heroNode1);
        insertAtTail(headNode, heroNode2);
        insertAtTail(headNode, heroNode3);

        printHeadNodeInfo(headNode);
    }

    public static void testSortInsertByNo() {
        // 创建结点，用来做尾部插入
        HeroNode head = new HeroNode();
        HeroNode node1 = new HeroNode(1, "宋江", "呼保义");
        HeroNode node2 = new HeroNode(2, "卢俊义", "玉麒麟");
        HeroNode node3 = new HeroNode(3, "吴用", "智多星");

        insertAtTail(head, node1);
        insertAtTail(head, node3); // 将第三个结点插入到第二个位置
        insertAtTail(head, node2);

        System.out.println("尾部插入的结果:");
        printHeadNodeInfo(head);

        // 创建 head 结点
        HeroNode headNode = new HeroNode();
        // 创建第一个结点
        HeroNode heroNode1 = new HeroNode(1, "宋江", "呼保义");
        // 创建第二个结点
        HeroNode heroNode2 = new HeroNode(2, "卢俊义", "玉麒麟");
        // 创建第三个结点
        HeroNode heroNode3 = new HeroNode(3, "吴用", "智多星");

        // 将结点按照 no 升序插入
        sortInsertByNo(headNode, heroNode1);
        sortInsertByNo(headNode, heroNode3);
        sortInsertByNo(headNode, heroNode2);

        System.out.println("按照 no 升序插入的结果:");
        printHeadNodeInfo(headNode);
    }

    public static void testDeleteNode() {
        // 创建结点
        HeroNode headNode = new HeroNode();
        HeroNode heroNode1 = new HeroNode(1, "宋江", "呼保义");
        HeroNode heroNode2 = new HeroNode(2, "卢俊义", "玉麒麟");
        HeroNode heroNode3 = new HeroNode(3, "吴用", "智多星");
        HeroNode heroNode4 = new HeroNode(4, "公孙胜", "入云龙");
        HeroNode heroNode5 = new HeroNode(5, "关胜", "大刀");

        // 插入结点
        insertAtTail(headNode, heroNode1);
        insertAtTail(headNode, heroNode2);
        insertAtTail(headNode, heroNode3);
        insertAtTail(headNode, heroNode4);
        insertAtTail(headNode, heroNode5);

        System.out.println("删除前:");
        printHeadNodeInfo(headNode);

        // 删除 no 为 2 的结点
        deleteNode(headNode, heroNode2);
        System.out.println("删除 no 为 2 的结点后:");
        printHeadNodeInfo(headNode);

        // 删除 no 为 3, 4 的结点
        deleteNode(headNode, heroNode3);
        deleteNode(headNode, heroNode4);
        System.out.println("删除 no 为 3,4 的结点后:");
        printHeadNodeInfo(headNode);
    }

    public static void main(String[] args) {
        // testInsertAtTail();
        // testSortInsertByNo();
        // testDeleteNode();
    }
}
