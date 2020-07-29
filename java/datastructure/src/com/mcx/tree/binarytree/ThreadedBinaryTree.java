package com.mcx.tree.binarytree;

/**
 * 线索化二叉树
 * 对于 n 个结点的二叉树，在二叉链存储结构中有 n+1 个空链域，用这些
 * 空链域存放该结点的前驱结点和后继结点的指针，这些指针称为线索，加上
 * 线索的二叉树称为线索二叉树。对二叉树以某种遍历方式（如先序、中序、
 * 后序或层次等）进行遍历，使其变为线索二叉树的过程称为对二叉树进行线索化
 */
public class ThreadedBinaryTree {

    public static class ThreadedBinaryTreeNode {
        private int no; // 编号
        private ThreadedBinaryTreeNode left; // 左子节点
        private ThreadedBinaryTreeNode right; // 右子节点

        // 增加两个标记
        private int leftTag; // 左节点标记。如果 leftTag = 0, left 表示左子节点, 如果为 leftTag = 1, left 表示前驱节点
        private int rightTag; // 右节点标记。如果 rightTag = 0, right 表示右子节点, 如果为 rightTag = 1, right 表示后继节点

        public ThreadedBinaryTreeNode(int no) {
            this.no = no;
        }

        public int getNo() {
            return no;
        }

        public void setNo(int no) {
            this.no = no;
        }

        public ThreadedBinaryTreeNode getLeft() {
            return left;
        }

        public void setLeft(ThreadedBinaryTreeNode left) {
            this.left = left;
        }

        public ThreadedBinaryTreeNode getRight() {
            return right;
        }

        public void setRight(ThreadedBinaryTreeNode right) {
            this.right = right;
        }

        public int getLeftTag() {
            return leftTag;
        }

        public void setLeftTag(int leftTag) {
            this.leftTag = leftTag;
        }

        public int getRightTag() {
            return rightTag;
        }

        public void setRightTag(int rightTag) {
            this.rightTag = rightTag;
        }

        /**
         * 从最左边至最右边查找指定节点（测试使用）
         */
        public ThreadedBinaryTreeNode search(int no) {
            ThreadedBinaryTreeNode leftChildNode = this;
            while (leftChildNode.getLeft() != null) {
                leftChildNode = leftChildNode.getLeft();
            }

            while (leftChildNode != null) {
                if (leftChildNode.getNo() == no) {
                    break;
                }
                leftChildNode = leftChildNode.getRight();
            }
            return leftChildNode;
        }
    }

    private static ThreadedBinaryTreeNode previous; // 记录遍历时的上一个结点

    /**
     * 中序线索化二叉树
     */
    public static void infixThreadTree(ThreadedBinaryTreeNode node) {
        if (node == null) {
            return;
        }
        // 线索化左子节点
        infixThreadTree(node.getLeft());

        // 线索化当前结点
        // 如果 left 为 null, 处理前驱节点
        if (node.getLeft() == null) {
            node.setLeft(previous);
            node.setLeftTag(1); // 修改标记
        }

        // 如果 right 为 null, 处理后继节点
        if (previous != null && previous.getRight() == null) {
            previous.setRight(node); // 将上一个节点的后继节点指向当前节点
            previous.setRightTag(1); // 修改标记
        }

        // 修改 previous
        previous = node;

        // 线索化右子节点
        infixThreadTree(node.getRight());
    }

    /**
     * 中序遍历线索化二叉树
     */
    public static void infixOrderThreadedTree(ThreadedBinaryTreeNode node) {
        ThreadedBinaryTreeNode currentNode = node;
        while (currentNode != null) {
            // 循环找到 leftTag = 0 的节点, 第一个找到的就是最左边的叶子节点
            while (currentNode.getLeftTag() == 0) {
                currentNode = currentNode.getLeft();
            }
            // 输出当前节点
            System.out.printf("id:%d\n", currentNode.getNo());
            // 循环输出后继节点
            while (currentNode.getRightTag() == 1) {
                currentNode = currentNode.getRight();
                System.out.printf("id:%d\n", currentNode.getNo());
            }
            // 移动当前节点
            currentNode = currentNode.getRight();
        }
    }

    public static ThreadedBinaryTreeNode initThreadedNode() {
        ThreadedBinaryTreeNode root = new ThreadedBinaryTreeNode(1);
        ThreadedBinaryTreeNode node2 = new ThreadedBinaryTreeNode(2);
        ThreadedBinaryTreeNode node3 = new ThreadedBinaryTreeNode(6);
        ThreadedBinaryTreeNode node4 = new ThreadedBinaryTreeNode(8);
        ThreadedBinaryTreeNode node5 = new ThreadedBinaryTreeNode(10);
        ThreadedBinaryTreeNode node6 = new ThreadedBinaryTreeNode(16);

        // 手动建立树的关系
        root.setLeft(node2);
        root.setRight(node3);
        node2.setLeft(node4);
        node2.setRight(node5);
        node3.setLeft(node6);

        return root;
    }

    public static void testInfixThreadedTree() {
        ThreadedBinaryTreeNode root = initThreadedNode();

        infixThreadTree(root);

        // 获取 no = 10 的结点，输出前驱和后继节点
        ThreadedBinaryTreeNode node = root.search(10);

        System.out.printf("no=%d的前驱节点为%d\n", node.getNo(), node.getLeft().getNo());
        System.out.printf("no=%d的后继节点为%d\n", node.getNo(), node.getRight().getNo());
    }

    public static void testInfixOrderThreadedTree() {
        ThreadedBinaryTreeNode root = initThreadedNode();

        infixThreadTree(root);

        infixOrderThreadedTree(root);
    }

    public static void main(String[] args) {
         testInfixThreadedTree();
        // testInfixOrderThreadedTree();
    }
}
