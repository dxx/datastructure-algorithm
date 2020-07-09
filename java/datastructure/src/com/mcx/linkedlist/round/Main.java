package com.mcx.linkedlist.round;

/**
 * 循环链表
 * 循环链表的特点是表中最后一个结点的指针域指向头结点，整个链表形成一个环
 *
 * 双向循环链表
 */
public class Main {

    public static class PersonNode {
        private int no;
        private String name;
        private PersonNode prev;
        private PersonNode next;

        public int getNo() {
            return no;
        }

        public void setNo(int no) {
            this.no = no;
        }

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public PersonNode getPrev() {
            return prev;
        }

        public void setPrev(PersonNode prev) {
            this.prev = prev;
        }

        public PersonNode getNext() {
            return next;
        }

        public void setNext(PersonNode next) {
            this.next = next;
        }

        public PersonNode() {}
        public PersonNode(int no, String name) {
            this.no = no;
            this.name = name;
        }
    }

    /**
     * 插入结点
     */
    public static void insertNode(PersonNode headNode, PersonNode newNode) {
        // 判断是否是第一次插入
        if (headNode.getNext() == null) {
            headNode.setNo(newNode.getNo());
            headNode.setName(newNode.getName());
            headNode.setPrev(headNode);
            headNode.setNext(headNode);
            return;
        }
        PersonNode lastNode = headNode;
        // 下一个结点不等于头结点继续循环
        while (lastNode.getNext() != headNode) {
            lastNode = lastNode.getNext();
        }
        // 将新结点添加到链表末尾
        lastNode.setNext(newNode);
        newNode.setPrev(lastNode);
        // 将新结点下一个结点指针指向头结点
        newNode.setNext(headNode);
        headNode.setPrev(newNode);
    }

    /**
     * 删除指定结点，返回头结点
     */
    public static PersonNode deleteNode(PersonNode headNode, PersonNode node) {
        // 没有结点 或者 只有一个头结点
        if (headNode.getNext() == null || headNode.getNext() == headNode) {
            // 头结点就是要删除的结点
            if (headNode.getNo() == node.getNo()) {
                headNode.setNext(null);
                headNode.setPrev(null);
            }
            return headNode;
        }

        PersonNode tempNode = headNode.getNext();
        boolean isExist = false;
        while (true) {
            if (tempNode == headNode) { // 最后一个结点
                if (tempNode.getNo() == node.getNo()) {
                    isExist = true;
                    // 头结点删除了，将头结点的下一个结点作为头结点
                    headNode = tempNode.getNext();
                }
                break;
            } else if (tempNode.getNo() == node.getNo()) {
                isExist = true;
                break;
            }
            tempNode = tempNode.getNext();
        }
        // 存在需要删除的结点
        if (isExist) {
            // 将查找到的结点的上一个结点的下一个结点指针指向当前结点的下一个结点
            tempNode.getPrev().setNext(tempNode.getNext());
            // 将查找到的结点的下一个结点的上一个结点指针指向当前指针的上一个结点
            tempNode.getNext().setPrev(tempNode.getPrev());
        }
        return headNode;
    }

    /**
     * 打印循环链表的信息
     */
    public static void printRoundNodeInfo(PersonNode headNode) {
        if (headNode.getNext() == null) {
            System.out.println("该循环链表没有节点");
            return;
        }
        StringBuilder sb = new StringBuilder("[");
        PersonNode tempNode = headNode;
        while (true) {
            sb.append(String.format("{no:%d, name:%s}",
                    tempNode.getNo(), tempNode.getName()));
            // 表示最后一个结点
            if (tempNode.getNext() == headNode) {
                break;
            }
            tempNode = tempNode.getNext();
        }
        sb.append("]");
        System.out.println(sb.toString());
    }

    public static void testInsertNode() {
        // 创建 head 结点，head 结点不初始化数据，等到添加了第一个结点后才初始化数据
        PersonNode headNode = new PersonNode();
        // 创建第一个结点
        PersonNode personNode1 = new PersonNode(1, "张三");
        // 创建第二个结点
        PersonNode personNode2 = new PersonNode(2, "李四");
        // 创建第三个结点
        PersonNode personNode3 = new PersonNode(3, "王五");

        // 插入结点
        insertNode(headNode, personNode1);
        insertNode(headNode, personNode2);
        insertNode(headNode, personNode3);

        printRoundNodeInfo(headNode);
    }

    public static void testDeleteNode() {
        // 创建结点
        PersonNode headNode = new PersonNode();
        PersonNode personNode1 = new PersonNode(1, "张三");
        PersonNode personNode2 = new PersonNode(2, "李四");
        PersonNode personNode3 = new PersonNode(3, "王五");
        PersonNode personNode4 = new PersonNode(4, "赵六");
        PersonNode personNode5 = new PersonNode(5, "孙七");

        // 插入结点
        insertNode(headNode, personNode1);
        insertNode(headNode, personNode2);
        insertNode(headNode, personNode3);
        insertNode(headNode, personNode4);
        insertNode(headNode, personNode5);

        System.out.println("删除前:");
        printRoundNodeInfo(headNode);

        // 删除 no 为 2 的结点
        headNode = deleteNode(headNode, personNode2);
        System.out.println("删除 no 为 2 的结点后:");
        printRoundNodeInfo(headNode);

        PersonNode newNode = new PersonNode(6, "周八");

        insertNode(headNode, newNode);
        System.out.println("插入新结点:");
        printRoundNodeInfo(headNode);

        // 删除 no 为 1，3 的结点
        headNode = deleteNode(headNode, personNode1);
        headNode = deleteNode(headNode, personNode3);
        System.out.println("删除 no 为 1,3 的结点后:");
        printRoundNodeInfo(headNode);
    }

    public static void main(String[] args) {
        // testInsertNode();
        // testDeleteNode();
    }

}
