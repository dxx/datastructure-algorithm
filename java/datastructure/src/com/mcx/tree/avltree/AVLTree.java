package com.mcx.tree.avltree;

/**
 * AVL 树
 * 在 AVL 树中任何节点的两个子树的高度最大差别为 1 所以它也被称为高度平衡树
 * 增加和删除可能需要通过一次或多次树旋转来重新平衡这个树。AVL 树本质上是带了
 * 平衡功能的二叉排序树（二叉查找树，二叉搜索树）
 */
public class AVLTree {

    private BinaryTreeNode root; // 树的根节点

    public static class BinaryTreeNode {
        private int no;
        private BinaryTreeNode left;
        private BinaryTreeNode right;

        public BinaryTreeNode(int no) {
            this.no = no;
        }

        @Override
        public String toString() {
            return "no:" + this.no;
        }
    }

    /**
     * 添加节点
     */
    public void add(BinaryTreeNode node) {
        if (node == null) {
            return;
        }
        // 根节点为 nil，要添加的节点作为根节点
        if (this.root == null) {
            this.root = node;
            return;
        }
        this.add(this.root, node);

        // 添加节点后判断是否需要旋转

        // 右边高度超过左边 1 个高度以上，进行左旋转
        if (this.rightHeight() - this.leftHeight() > 1) {

            BinaryTreeNode rightNode = this.root.right;
            // 右子节点不为 null，并且右子节点的左子树高度大于右子节点的右子树高度
            if (rightNode != null &&
                    this.height(rightNode.left) > this.height(rightNode.right)) {
                // 将右子节点右旋转
                this.rightRotate(rightNode);
            }

            this.leftRotate(this.root);
            return;
        }

        // 左边高度超过右边 1 个高度以上，进行右旋转
        if (this.leftHeight() - this.rightHeight() > 1) {

            BinaryTreeNode leftNode = this.root.left;
            // 左子节点不为 nil, 并且左子节点的右子树高度大于左子节点的左子树高度
            if (leftNode != null &&
                    this.height(leftNode.right) > this.height(leftNode.left)) {
                // 将左子节点左旋转
                this.leftRotate(leftNode);
            }

            this.rightRotate(this.root);
        }
    }

    /**
     * 计算节点的高度
     */
    private int height(BinaryTreeNode node) {
        if (node == null) {
            return 0;
        }
        // 递归计算左子节点和右子节点的高度，返回最大的高度，然后 + 1
        return Math.max(height(node.left), height(node.right)) + 1;
    }

    /**
     * 左子树高度
     */
    private int leftHeight() {
        if (this.root == null) {
            return 0;
        }
        return this.height(this.root.left);
    }

    /**
     * 右子树高度
     */
    private int rightHeight() {
        if (this.root == null) {
            return 0;
        }
        return this.height(this.root.right);
    }

    /**
     * 左旋转
     */
    private void leftRotate(BinaryTreeNode node) {
        if (node == null) {
            return;
        }
        // 以当前节点为基础，创建一个新的节点，新节点的值等于当前节点的值
        BinaryTreeNode newNode = new BinaryTreeNode(node.no);
        // 让新节点的左子节点指向当前节点的左子节点，右子节点指向当前节点的右子节点的左子节点
        newNode.left = node.left;
        newNode.right = node.right.left;
        // 把当前节点的值替换为右子节点的值，并把当前节点右子节点指向其右子节点的右子节点
        node.no = node.right.no;
        node.right = node.right.right;
        // 让当前节点的左子节点指向新创建的节点
        node.left = newNode;
    }

    /**
     * 右旋转
     */
    private void rightRotate(BinaryTreeNode node) {
        if (node == null) {
            return;
        }
        // 以当前节点为基础，创建一个新的节点，新节点的值等于根节点的值
        BinaryTreeNode newNode = new BinaryTreeNode(node.no);
        // 让新节点的右子节点指向当前节点的右子节点，左子节点指向当前节点的左子节点的右子节点
        newNode.right = node.right;
        newNode.left = node.left.right;
        // 把当前节点的值替换为左子节点的值，并把当前节点左子节点指向其左子节点的左子节点
        node.no = node.left.no;
        node.left = node.left.left;
        // 让当前节点的右子节点指向新创建的节点
        node.right = newNode;
    }

    private void add(BinaryTreeNode root, BinaryTreeNode node) {
        // 要添加的节点小于根节点
        if (node.no < root.no) {
            // 左子节点为 null，直接添加为左子节点
            if (root.left == null) {
                root.left = node;
                return;
            }
            // 左递归
            this.add(root.left, node);
        } else {
            // 右子节点为 null，直接添加为右子节点
            if (root.right == null) {
                root.right = node;
                return;
            }
            // 右递归
            this.add(root.right, node);
        }
    }

    public void infixOrder() {
        if (this.root == null) {
            return;
        }
        this.infixOrder(this.root);
    }

    private void infixOrder(BinaryTreeNode node) {
        if (node == null) {
            return;
        }
        this.infixOrder(node.left);
        System.out.println(node);
        this.infixOrder(node.right);
    }

    public static void testLeftRotate() {
        int[] nos = new int[]{3, 2, 5, 4, 6, 7};
        AVLTree avlTree = new AVLTree();
        for (int no : nos) {
            avlTree.add(new BinaryTreeNode(no));
        }

        System.out.println("左旋转后");

        avlTree.infixOrder();

        System.out.printf("根节点 = %s\n", avlTree.root);

        System.out.printf("左子树的高度为: %d\n", avlTree.leftHeight());
        System.out.printf("右子树的高度为: %d\n", avlTree.rightHeight());
    }

    public static void testRightRotate() {
        int[] nos = new int[]{6, 4, 7, 3, 5, 2};
        AVLTree avlTree = new AVLTree();
        for (int no : nos) {
            avlTree.add(new BinaryTreeNode(no));
        }

        System.out.println("右旋转后");

        avlTree.infixOrder();

        System.out.printf("根节点 = %s\n", avlTree.root);

        System.out.printf("左子树的高度为: %d\n", avlTree.leftHeight());
        System.out.printf("右子树的高度为: %d\n", avlTree.rightHeight());
    }

    public static void testDoubleRotate() {
        int[] nos = new int[]{6, 3, 7, 2, 4, 5};
        AVLTree avlTree = new AVLTree();
        for (int no : nos) {
            avlTree.add(new BinaryTreeNode(no));
        }

        System.out.println("双旋转后");

        avlTree.infixOrder();

        System.out.printf("根节点 = %s\n", avlTree.root);

        System.out.printf("左子树的高度为: %d\n", avlTree.leftHeight());
        System.out.printf("右子树的高度为: %d\n", avlTree.rightHeight());
    }

    public static void main(String[] args) {
        // testLeftRotate();
        // testRightRotate();
        // testDoubleRotate();
    }
}
