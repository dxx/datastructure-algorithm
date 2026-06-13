package com.dxx.tree.huffmantree;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

/**
 * 哈夫曼树
 * 有 N 个权值作为 N 个叶子结点，构造一棵二叉树，如果该树的带权路径长度达到最小
 * 称这样的二叉树为最优二叉树，也称为哈夫曼树(Huffman Tree)。哈夫曼树是带权路
 * 径长度最短的树，权值较大的结点离根较近。哈夫曼树又称为最优树。
 *
 * 构建步骤
 * 1.将 w1、w2、…、wn 看成一个序列, 每个数据可以看做一个权值
 * 2.将序列从小到大排序
 * 3.选出两个根节点的权值最小的树合并，作为一棵新树的左、右子节点，且新树的根节点权值为其左、右子树根节点权值之和
 * 4.从序列中删除选出的两个节点，并将新树加入序列
 * 5.重复 2、3、4 步，直到序列中只剩一棵树为止，该树即为所求得的哈夫曼树
 */
public class Huffman {

    public static class Node {
        private int value;
        private Node left;
        private Node right;

        public Node(int value) {
            this.value = value;
        }

        public int getValue() {
            return value;
        }

        public void setValue(int value) {
            this.value = value;
        }

        public Node getLeft() {
            return left;
        }

        public void setLeft(Node left) {
            this.left = left;
        }

        public Node getRight() {
            return right;
        }

        public void setRight(Node right) {
            this.right = right;
        }

    }

    public static Node createHuffmanTree(int[] nums) {
        if (nums == null) {
            return null;
        }
        List<Node> nodes = new ArrayList<>();
        for (int num : nums) {
            nodes.add(new Node(num));
        }

        while (nodes.size() > 1) {
            // 排序
            nodes.sort(Comparator.comparingInt(Node::getValue));

            Node left = nodes.get(0); // 权值最小的元素
            Node right = nodes.get(1); // 权值第二小的元素
            // 创建新的根节点
            Node root = new Node(left.getValue() + right.getValue());
            // 构建二叉树
            root.setLeft(left);
            root.setRight(right);

            // 删除处理过的节点
            nodes.remove(left);
            nodes.remove(right);

            // 将二叉树加入到 nodes
            nodes.add(root);
        }
        return nodes.get(0);
    }

    public static void preOrder(Node node) {
        if (node == null) {
            return;
        }
        System.out.println(node.getValue());
        preOrder(node.getLeft());
        preOrder(node.getRight());
    }

    public static void main(String[] args) {
        int[] nums = new int[]{1, 7, 3, 8, 16};
        Node root = createHuffmanTree(nums);
        preOrder(root);
    }
}
