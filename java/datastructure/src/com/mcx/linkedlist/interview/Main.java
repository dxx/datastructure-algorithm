package com.mcx.linkedlist.interview;

/**
 * 单链表相关面试题
 */
public class Main {

    /**
     * 获取单链表有效的结点数
     * 1.遍历结点数
     * 2.定义一个长度遍历，每遍历一次长度 +1
     */
    public static int getNodeLength(Node headNode) {
        // 头结点为空，返回 0
        if (headNode == null) {
            return 0;
        }
        int length = 0;
        Node node = headNode.getNext();
        while (node != null) {
            length++;
            node = node.getNext();
        }
        return length;
    }

    /**
     * 获取倒数第 n 个节点
     * 1.获取链表节点数 length
     * 2.遍历到 length - n 个节点
     * 3.然后返回
     */
    public static Node getLastIndexNode(Node headNode, int index) {
        // 头结点为空，返回空
        if (headNode == null) {
            return null;
        }
        int length = getNodeLength(headNode);
        if (index <= 0 || index > length) {
            return null;
        }
        Node lastNode = headNode.getNext();
        for (int i = 0; i < length - index; i++) {
            lastNode = lastNode.getNext();
        }
        return lastNode;
    }

    /**
     * 单链表反转
     * 1.定义一个新的头结点 reverseHead
     * 2.遍历链表，每遍历一个结点，将其取出，放在新的头结点 reverseHead 的后面
     * 3.最后将头结点的 next 结点指向 reverseHead 的 next 结点
     */
    public static void reverseNode(Node headNode) {
        if (headNode == null || headNode.getNext() == null) {
            return;
        }
        Node reverseHead = new Node();
        Node current = headNode.getNext();
        Node next;
        while (current != null) {
            // 保存当前结点的下一个结点
            next = current.getNext();
            // 将 reverseHead 结点的下一个结点放在当前结点的下一个结点
            current.setNext(reverseHead.getNext());
            // 当前结点放在 reverseHead 后面
            reverseHead.setNext(current);
            // 移动当前结点
            current = next;
        }
        // 将头结点的 next 结点指向 reverseHead 的 next 结点
        headNode.setNext(reverseHead.getNext());
    }

    public static void testGetLength() {
        Node headNode = new Node();
        Node node1 = new Node("node1");
        Node node2 = new Node("node2");
        Node node3 = new Node("node3");
        headNode.setNext(node1);
        node1.setNext(node2);
        node2.setNext(node3);

        int length = getNodeLength(headNode);
        System.out.printf("单链表结点个数为: %d\n", length);
    }

    public static void testGetLastIndexNode() {
        Node headNode = new Node();
        Node node1 = new Node("node1");
        Node node2 = new Node("node2");
        Node node3 = new Node("node3");
        headNode.setNext(node1);
        node1.setNext(node2);
        node2.setNext(node3);
        int index = 2;
        Node lastNode = getLastIndexNode(headNode, index);
        System.out.printf("单链表结点中倒数第%d个结点为: %s\n", index, lastNode.getName());
    }

    public static void testReverseNode() {
        Node headNode = new Node();
        Node node1 = new Node("node1");
        Node node2 = new Node("node2");
        Node node3 = new Node("node3");
        headNode.setNext(node1);
        node1.setNext(node2);
        node2.setNext(node3);

        System.out.println("反转前:");
        headNode.printNodeInfo();

        reverseNode(headNode);

        System.out.println("反转后:");
        headNode.printNodeInfo();
    }

    public static void main(String[] args) {
        // testGetLength();
        // testGetLastIndexNode();
        // testReverseNode();
    }
}
