package com.mcx.tree.binarytree;

/**
 * 二叉树
 * 二叉树是每个结点最多有两个子树的树结构
 * 通常子树被称作“左子树”（left subtree）和“右子树”（right subtree）
 * 二叉树常被用于实现二叉查找树和二叉堆
 */
public class BinaryTree {

    public static class BinaryTreeNode {
        private int no;
        private BinaryTreeNode left;
        private BinaryTreeNode right;

        public BinaryTreeNode(int no) {
            this.no = no;
        }

        @Override
        public String toString() {
            return "no:" + no;
        }
    }

    /**
     * 前序遍历
     */
    public static void preOrder(BinaryTreeNode node) {
        if (node == null) {
            return;
        }
        // 当前节点
        System.out.println(node);
        // 遍历左子树
        preOrder(node.left);
        // 遍历右子树
        preOrder(node.right);
    }

    /**
     * 中序遍历
     */
    public static void infixOrder(BinaryTreeNode node) {
        if (node == null) {
            return;
        }
        // 遍历左子树
        infixOrder(node.left);
        // 当前节点
        System.out.println(node);
        // 遍历右子树
        infixOrder(node.right);
    }

    /**
     * 后序遍历
     */
    public static void postOrder(BinaryTreeNode node) {
        if (node == null) {
            return;
        }
        // 遍历左子树
        postOrder(node.left);
        // 遍历右子树
        postOrder(node.right);
        // 当前节点
        System.out.println(node);
    }

    /**
     * 前序查找
     */
    public static BinaryTreeNode preOrderSearch(BinaryTreeNode node, int no) {
        if (node == null) {
            return null;
        }
        System.out.println("进入查找");
        if (node.no == no) {
            return node;
        }
        // 左边查找
        BinaryTreeNode returnNode = preOrderSearch(node.left, no);
        if (returnNode != null) {
            // 左边找到了节点，返回
            return returnNode;
        }
        // 右边查找
        return preOrderSearch(node.right, no);
    }

    /**
     * 中序查找
     */
    public static BinaryTreeNode infixOrderSearch(BinaryTreeNode node, int no) {
        if (node == null) {
            return null;
        }
        // 左边查找
        BinaryTreeNode returnNode = infixOrderSearch(node.left, no);
        if (returnNode != null) {
            // 左边找到了节点，返回
            return returnNode;
        }
        System.out.println("进入查找");
        if (node.no == no) {
            return node;
        }
        // 右边查找
        return infixOrderSearch(node.right, no);
    }

    /**
     * 后序查找
     */
    public static BinaryTreeNode postOrderSearch(BinaryTreeNode node, int no) {
        if (node == null) {
            return null;
        }
        // 左边查找
        BinaryTreeNode returnNode = postOrderSearch(node.left, no);
        if (returnNode != null) {
            // 左边找到了节点，返回
            return returnNode;
        }
        // 右边查找
        returnNode = postOrderSearch(node.right, no);
        if (returnNode != null) {
            // 右边找到了节点，返回
            return returnNode;
        }
        System.out.println("进入查找");
        if (node.no == no) {
            returnNode = node;
        }
        return returnNode;
    }

    public static BinaryTreeNode initNode() {
        BinaryTreeNode root = new BinaryTreeNode(1);
        BinaryTreeNode node2 = new BinaryTreeNode(2);
        BinaryTreeNode node3 = new BinaryTreeNode(3);
        BinaryTreeNode node4 = new BinaryTreeNode(4);
        BinaryTreeNode node5 = new BinaryTreeNode(5);

        // 手动建立树的关系
        root.left = node2;
        root.right = node5;
        node2.left = node3;
        node2.right = node4;

        return root;
    }

    public static void testOrder() {
        BinaryTreeNode root = initNode();

        System.out.println("======前序遍历======");
        preOrder(root);

        System.out.println("======中续遍历======");
        infixOrder(root);

        System.out.println("======后续遍历======");
        postOrder(root);
    }

    public static void testSearch() {
        BinaryTreeNode root = initNode();

        int no = 4;

        System.out.println("======前序查找======");
        System.out.printf("查找no=%d\n", no);
        BinaryTreeNode node = preOrderSearch(root, no);
        System.out.printf("查找结果: no=%d\n", node.no);

        no = 4;
        System.out.println("======中序查找======");
        System.out.printf("查找no=%d\n", no);
        node = infixOrderSearch(root, no);
        System.out.printf("查找结果: no=%d\n", node.no);

        no = 4;
        System.out.println("======后序查找======");
        System.out.printf("查找no=%d\n", no);
        node = postOrderSearch(root, no);
        System.out.printf("查找结果: no=%d\n", node.no);
    }

    public static void main(String[] args) {
        // testOrder();
        // testSearch();
    }
}
