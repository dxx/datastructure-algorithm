package com.dxx.linkedlist.ovonic;

/**
 * 双向链表
 * 双向链表也叫双链表，是链表的一种，它的每个数据结点中都有两个指针，分别指向前一个和后一个结点
 * 所以，从双向链表中的任意一个结点开始，都可以很方便地访问它的前一个结点和后一个结点。
 */
public class Main {

    public static class HeroNode {
        private int no; // 编号
        private String name; // 姓名
        private String nickname; // 昵称
        private HeroNode prev; // 上一个节点
        private HeroNode next; // 下一个节点

        public HeroNode() {}

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
        // 将新结点的上一个结点指向当前结点
        newNode.prev = lastNode;
    }

    /**
     * 删除指定结点
     */
    public static void deleteNode(HeroNode headNode, HeroNode node) {
        HeroNode tempNode = headNode.next;
        while (tempNode != null) {
            if (tempNode.no == node.no) {
                // 将查找到的结点的上一个结点的下一个结点指针指向当前结点的下一个结点
                tempNode.prev.next = tempNode.next;
                // 最后一个结点的 next 指向空
                if (tempNode.next != null) {
                    // 将查找到的结点的下一个结点的上一个结点指针指向当前指针的上一个结点
                    tempNode.next.prev =tempNode.prev;
                }
                return;
            }
            tempNode = tempNode.next;
        }
    }

    /**
     * 打印链表结点内容
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
        HeroNode heroNode1 = new HeroNode(3, "吴用", "智多星");
        // 创建第二个结点
        HeroNode heroNode2 = new HeroNode(6, "林冲", "豹子头");
        // 创建第三个结点
        HeroNode heroNode3 = new HeroNode(7, "秦明", "霹雳火");

        // 将结点添加到链表尾部
        insertAtTail(headNode, heroNode1);
        insertAtTail(headNode, heroNode2);
        insertAtTail(headNode, heroNode3);

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
        // testDeleteNode();
    }
}
