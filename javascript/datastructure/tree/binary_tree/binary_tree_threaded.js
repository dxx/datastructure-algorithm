/**
 * 二叉树线索化
 * 对于 n 个结点的二叉树，在二叉链存储结构中有 n+1 个空链域，用这些
 * 空链域存放该结点的前驱结点和后继结点的指针，这些指针称为线索，加上
 * 线索的二叉树称为线索二叉树。对二叉树以某种遍历方式（如先序、中序、
 * 后序或层次等）进行遍历，使其变为线索二叉树的过程称为对二叉树进行线索化
 */

function ThreadedBinaryTreeNode(no) {
  this.no = no; // 编号
  this.left = null; // 左子节点
  this.right = null; // 右子节点

  // 增加两个标记
  this.leftTag = 0; // 左节点标记。如果 leftTag = 0, left 表示左子节点, 如果为 leftTag = 1, left 表示前驱节点
  this.rightTag = 0; // 右节点标记。如果 rightTag = 0, right 表示右子节点, 如果为 rightTag = 1, right 表示后继节点
}
/**
 * 从最左边至最右边查找指定节点（测试使用）
 */
ThreadedBinaryTreeNode.prototype.search = function(no) {
  let leftChildNode = this;
  while (leftChildNode.left !== null) {
    leftChildNode = leftChildNode.left;
  }

  while (leftChildNode !== null) {
    if (leftChildNode.no === no) {
        break;
    }
    leftChildNode = leftChildNode.right;
  }
  return leftChildNode;
}

let previous = null; // 记录遍历时的上一个结点

/**
 * 中序线索化二叉树
 */
function infixThreadTree(node) {
  if (!node) {
    return;
  }
  // 线索化左子节点
  infixThreadTree(node.left);

  // 线索化当前结点
  // 如果 left 为 null, 处理前驱节点
  if (node.left === null) {
      node.left = previous;
      node.leftTag = 1; // 修改标记
  }

  // 如果 right 为 null, 处理后继节点
  if (previous !== null && previous.right === null) {
      previous.right = node; // 将上一个节点的后继节点指向当前节点
      previous.rightTag = 1; // 修改标记
  }

  // 修改 previous
  previous = node;

  // 线索化右子节点
  infixThreadTree(node.right);
}

/**
 * 中序遍历线索化二叉树
 */
function infixOrderThreadedTree(node) {
  let currentNode = node;
  while (currentNode !== null) {
    // 循环找到 leftTag = 0 的节点, 第一个找到的就是最左边的叶子节点
    while (currentNode.leftTag === 0) {
        currentNode = currentNode.left;
    }
    // 输出当前节点
    console.log("id:" + currentNode.no);
    // 循环输出后继节点
    while (currentNode.rightTag === 1) {
        currentNode = currentNode.right;
        console.log("id:" + currentNode.no);
    }
    // 移动当前节点
    currentNode = currentNode.right;
  }
}

function initThreadedNode() {
  let root = new ThreadedBinaryTreeNode(1);
  let node2 = new ThreadedBinaryTreeNode(2);
  let node3 = new ThreadedBinaryTreeNode(6);
  let node4 = new ThreadedBinaryTreeNode(8);
  let node5 = new ThreadedBinaryTreeNode(10);
  let node6 = new ThreadedBinaryTreeNode(16);

  // 手动建立树的关系
  root.left = node2;
  root.right = node3;
  node2.left = node4;
  node2.right = node5;
  node3.left = node6;

  return root
}

function testInfixThreadedTree() {
  let root = initThreadedNode();

  infixThreadTree(root);

  // 获取 no = 10 的结点，输出前驱和后继节点
  let node = root.search(10);

  console.log("no=" + node.no + "的前驱节点为" + node.left.no);
  console.log("no=" + node.no + "的后继节点为" + node.right.no);
}

function testInfixOrderThreadedTree() {
  let root = initThreadedNode();

  infixThreadTree(root);

  infixOrderThreadedTree(root);
}

function main() {
  // testInfixThreadedTree();
  // testInfixOrderThreadedTree();
}

main();
