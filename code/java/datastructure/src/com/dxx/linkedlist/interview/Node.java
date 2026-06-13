package com.dxx.linkedlist.interview;

/**
 * 单链表节点
 */
public class Node {

    public String name;
    public Node next;

    public Node() {}

    public Node(String name) {
        this.name = name;
    }

    public void printNodeInfo() {
        if (this.next == null) {
            System.out.println("该链表没有节点");
            return;
        }
        StringBuilder sb = new StringBuilder("[");
        Node tempNode = this.next;
        while (tempNode != null) {
            sb.append(String.format("{name:%s}", tempNode.name));
            tempNode = tempNode.next;
        }
        sb.append("]");
        System.out.println(sb.toString());
    }

}
