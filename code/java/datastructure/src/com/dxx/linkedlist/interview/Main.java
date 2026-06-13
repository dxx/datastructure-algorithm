package com.dxx.linkedlist.interview;

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
        Node node = headNode.next;
        while (node != null) {
            length++;
            node = node.next;
        }
        return length;
    }

    /**
     * 获取倒数第 n 个结点
     *
     * 快慢指针
     * 1.定义快指针 fast 和 慢指针 slow
     * 2.快慢指针的初始值指向头结点
     * 3.快指针先走 index 步
     * 4.慢指针开始走直到快指针指向了末尾结点
     * 5.此时慢指针就是倒数第 n 个结点
     */
    public static Node getLastIndexNode(Node headNode, int index) {
        // 头结点为空，index 小于等于 0 返回空
        if (headNode == null || index <= 0) {
            return null;
        }
        Node fast = headNode;
        Node slow = headNode;
		int i = index;
		int length = 0;
        while (fast != null) {
			length++;
            if (i > 0) {
                fast = fast.next;
                i--;
                continue;
            }
            // 快慢指针同时走
            fast = fast.next;
            slow = slow.next;
        }
        // index 超过了链表的长度
        if (index > length) {
            return null;
        }
        return slow;
    }

    /**
     * 获取倒数第 n 个结点
     *
     * 遍历
     * 1.获取链表结点数 length
     * 2.遍历到 length - n 个结点
     * 3.然后返回
     */
    public static Node getLastIndexNode2(Node headNode, int index) {
        // 头结点为空，返回空
        if (headNode == null) {
            return null;
        }
        int length = getNodeLength(headNode);
        if (index <= 0 || index > length) {
            return null;
        }
        Node lastNode = headNode.next;
        for (int i = 0; i < length - index; i++) {
            lastNode = lastNode.next;
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
        if (headNode == null || headNode.next == null) {
            return;
        }
        Node reverseHead = new Node();
        Node current = headNode.next;
        Node next;
        while (current != null) {
            // 保存当前结点的下一个结点
            next = current.next;
            // 将 reverseHead 结点的下一个结点放在当前结点的下一个结点
            current.next = reverseHead.next;
            // 当前结点放在 reverseHead 后面
            reverseHead.next = current;
            // 移动当前结点
            current = next;
        }
        // 将头结点的 next 结点指向 reverseHead 的 next 结点
        headNode.next = reverseHead.next;
    }

    public static void testGetLength() {
        Node headNode = new Node();
        Node node1 = new Node("node1");
        Node node2 = new Node("node2");
        Node node3 = new Node("node3");
        headNode.next = node1;
        node1.next = node2;
        node2.next = node3;

        int length = getNodeLength(headNode);
        System.out.printf("单链表结点个数为: %d\n", length);
    }

    public static void testGetLastIndexNode() {
        Node headNode = new Node();
        Node node1 = new Node("node1");
        Node node2 = new Node("node2");
        Node node3 = new Node("node3");
        headNode.next = node1;
        node1.next = node2;
        node2.next = node3;
        int index = 2;
        Node lastNode = getLastIndexNode(headNode, index);
        System.out.printf("单链表结点中倒数第 %d 个结点为: %s\n", index, lastNode.name);
    }

    public static void testReverseNode() {
        Node headNode = new Node();
        Node node1 = new Node("node1");
        Node node2 = new Node("node2");
        Node node3 = new Node("node3");
        headNode.next = node1;
        node1.next = node2;
        node2.next = node3;

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
