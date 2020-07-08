package com.mcx.linkedlist.interview;

/**
 * 单链表节点
 */
public class Node {

    private String name;
    private Node next;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Node getNext() {
        return next;
    }

    public void setNext(Node next) {
        this.next = next;
    }

    public Node() {}

    public Node(String name) {
        this.name = name;
    }

    public void printNodeInfo() {
        if (this.getNext() == null) {
            System.out.println("该链表没有节点");
            return;
        }
        StringBuilder sb = new StringBuilder("[");
        Node tempNode = this.getNext();
        while (tempNode != null) {
            sb.append(String.format("{name:%s}", tempNode.getName()));
            tempNode = tempNode.getNext();
        }
        sb.append("]");
        System.out.println(sb.toString());
    }

}
