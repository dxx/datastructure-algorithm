package com.mcx.tree.binarysorttree;

/**
 * 二叉排序树
 * 二叉排序树（Binary Sort Tree），又称二叉查找树（Binary Search Tree），也叫二叉搜索树
 * 在一般情况下，查询效率比链表结构要高。对于任何一个非叶子节点，它的左子节点的值小于自身的值，
 * 它的右子节点的值大于自身的值，具有这样性质的二叉树称为二叉排序树
 */
public class BinarySortTree {

    private BinaryTreeNode root; // 树的根节点

    public static class BinaryTreeNode {
        private int no;
        private BinaryTreeNode left;
        private BinaryTreeNode right;

        public BinaryTreeNode(int no) {
            this.no = no;
        }

        public int getNo() {
            return no;
        }

        public void setNo(int no) {
            this.no = no;
        }

        public BinaryTreeNode getLeft() {
            return left;
        }

        public void setLeft(BinaryTreeNode left) {
            this.left = left;
        }

        public BinaryTreeNode getRight() {
            return right;
        }

        public void setRight(BinaryTreeNode right) {
            this.right = right;
        }

        @Override
        public String toString() {
            return "no:" + this.getNo();
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
    }

    private void add(BinaryTreeNode root, BinaryTreeNode node) {
        // 要添加的节点小于根节点
        if (node.getNo() < root.getNo()) {
            // 左子节点为 null，直接添加为左子节点
            if (root.getLeft() == null) {
                root.setLeft(node);
                return;
            }
            // 左递归
            this.add(root.getLeft(), node);
        } else {
            // 右子节点为 null，直接添加为右子节点
            if (root.getRight() == null) {
                root.setRight(node);
                return;
            }
            // 右递归
            this.add(root.getRight(), node);
        }
    }

    public BinaryTreeNode[] search(int no) {
        if (this.root == null) {
            return null;
        }
        // 查找的节点就是根节点
        if (this.root.getNo() == no) {
            return new BinaryTreeNode[]{null, this.root};
        }
        return this.recursionSearch(this.root, no);
    }

    /**
     * 递归查找指定节点
     * 返回查找到的父节点和查找到的节点
     */
    public BinaryTreeNode[] recursionSearch(BinaryTreeNode node, int no) {
        if (node == null) {
            return null;
        }
        if (node.getLeft() != null && node.getLeft().getNo() == no) {
            return new BinaryTreeNode[]{node, node.getLeft()};
        }
        if (node.getRight() != null && node.getRight().getNo() == no) {
            return new BinaryTreeNode[]{node, node.getRight()};
        }
        // 判断是往左边还是往右边查找
        if (no < node.getNo()) {
            return this.recursionSearch(node.getLeft(), no);
        } else {
            return this.recursionSearch(node.getRight(), no);
        }
    }

    /**
     * 删除节点
     * 1.节点是叶子节点直接删除
     * 2.节点是子节点且只有一颗子树，左子树或右子树。如果被删除的节点是父节点的
     *   左子节点，将父节点的左子节点指向该删除节点的子树，如果是父节点的右子节
     *   点，则将父节点的右子节点指向该删除节点的子树
     * 3.节点是子节点且只有两颗子树。从被删除节点的右子节点的左子树中找到最小值的节点，将
     *   其删除，然后将该节点的值赋值给被删除的节点
     */
    public void delete(int no) {
        BinaryTreeNode[] binaryTreeNodes = this.search(no);
        // 没有找到要删除的节点
        if (binaryTreeNodes == null) {
            return;
        }
        BinaryTreeNode parentNode = binaryTreeNodes[0];
        BinaryTreeNode node = binaryTreeNodes[1];
        // 当前节点为叶子节点
        if (node.getLeft() == null && node.getRight() == null) {
            // 被删除的节点为根节点
            if (parentNode == null) {
                this.root = null;
                return;
            }
            // 当前节点为父节点的左子节点
            if (parentNode.getLeft() != null && parentNode.getLeft().getNo() == no) {
                parentNode.setLeft(null);
            }
            // 当前节点为父节点的右子节点
            if (parentNode.getRight() != null && parentNode.getRight().getNo() == no) {
                parentNode.setRight(null);
            }
            return;
        }
        // 当前节点有两颗子树
        if (node.getLeft() != null && node.getRight() != null) {
            // 把右子节点作为根节点，从左边开始遍历到最后一个叶子节点
            BinaryTreeNode leftChildNode = node.getRight();
            while (leftChildNode.getLeft() != null) {
                leftChildNode = leftChildNode.getLeft();
            }
            // 删除最小的叶子节点
            this.delete(leftChildNode.getNo());

            // 替换掉被删除节点的值
            node.setNo(leftChildNode.getNo());
        } else { // 当前节点只有一颗子树
            BinaryTreeNode replaceNode = null;
            if (node.getLeft() != null) {
                replaceNode = node.getLeft();
            }
            if (node.getRight() != null) {
                replaceNode = node.getRight();
            }

            // 父节点为 null，表示根节点
            if (parentNode == null) {
                this.root = replaceNode;
                return;
            }
            // 当前节点为父节点的左子节点
            if (parentNode.getLeft() != null && parentNode.getLeft().getNo() == no) {
                parentNode.setLeft(replaceNode);
            }
            // 当前节点为父节点的右子节点
            if (parentNode.getRight() != null && parentNode.getRight().getNo() == no) {
                parentNode.setRight(replaceNode);
            }
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
        this.infixOrder(node.getLeft());
        System.out.println(node);
        this.infixOrder(node.getRight());
    }

    public static void main(String[] args) {
        int[] nos = new int[]{8, 5, 10, 3, 6, 9, 12, 2};
        BinarySortTree binarySortTree = new BinarySortTree();
        for (int no : nos) {
            binarySortTree.add(new BinaryTreeNode(no));
        }

        System.out.println("======中序遍历======");
        binarySortTree.infixOrder();

        binarySortTree.delete(6);

        System.out.println("======删除叶子节点 6======");
        binarySortTree.infixOrder();

        binarySortTree.delete(5);

        System.out.println("======删除只有一颗子树的节点 5======");
        binarySortTree.infixOrder();

        binarySortTree.delete(10);

        System.out.println("======删除有两颗子树的节点 10======");
        binarySortTree.infixOrder();
    }
}
