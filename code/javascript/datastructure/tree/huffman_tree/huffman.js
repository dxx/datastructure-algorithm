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

function Node(value) {
  this.value = value;
  this.left = null;
  this.right = null;
}

function createHuffmanTree(nums) {
  if (!nums) {
    return null;
  }
  let nodes = [];
  for (let i = 0; i < nums.length; i++) {
    nodes.push(new Node(nums[i]));
  }
  while (nodes.length > 1) {
    // 排序
    nodes.sort((node1, node2) => node1.value - node2.value);
    let left = nodes[0]; // 权值最小的元素
    let right = nodes[1]; // 权值第二小的元素
    // 创建新的根节点
    let root = new Node(left.value + right.value);
    // 构建二叉树
    root.left = left;
    root.right = right;

    // 删除处理过的节点
    nodes.splice(0, 1);
    nodes.splice(0, 1);

    // 将二叉树加入到 nodes
    nodes.push(root);
  }
  return nodes[0];
}

function preOrder(node) {
  if (node == null) {
    return;
  }
  console.log(node.value);
  preOrder(node.left);
  preOrder(node.right);
}

function main() {
  let nums = [1, 7, 3, 8, 16];
  let root = createHuffmanTree(nums);
  preOrder(root);
}

main();
