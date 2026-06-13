/**
 * 二叉树
 * 二叉树是每个结点最多有两个子树的树结构
 * 通常子树被称作“左子树”（left subtree）和“右子树”（right subtree）
 * 二叉树常被用于实现二叉查找树和二叉堆
 */

function BinaryTreeNode(no) {
  this.no = no;
  this.left = null;
  this.right = null;
}

/**
 * 前序遍历 
 */
function preOrder(node) {
  if (!node) {
    return;
  }
  // 当前节点
  console.log(node.no);
  // 遍历左子树
  preOrder(node.left);
  // 遍历右子树
  preOrder(node.right);
}

/**
 * 中序遍历 
 */
function infixOrder(node) {
  if (!node) {
    return;
  }
  // 遍历左子树
  infixOrder(node.left);
  // 当前节点
  console.log(node.no);
  // 遍历右子树
  infixOrder(node.right);
}

/**
 * 后序遍历 
 */
function postOrder(node) {
  if (!node) {
    return;
  }
  // 遍历左子树
  postOrder(node.left);
  // 遍历右子树
  postOrder(node.right);
  // 当前节点
  console.log(node.no);
}

/**
 * 前序查找 
 */
function preOrderSearch(node, no) {
  if (!node) {
    return null;
  }
  console.log("进入查找");
  if (node.no === no) {
    return node;
  }
  // 左边查找
  let returnNode = preOrderSearch(node.left, no);
  if (returnNode) {
    // 左边找到了节点，返回
    return returnNode
  }
  // 右边查找
  return preOrderSearch(node.right, no);
}

/**
 * 中序查找 
 */
function infixOrderSearch(node, no) {
  if (!node) {
    return null;
  }
  // 左边查找
  let returnNode = infixOrderSearch(node.left, no);
  if (returnNode) {
    // 左边找到了节点，返回
    return returnNode
  }
  console.log("进入查找");
  if (node.no === no) {
    return node;
  }
  // 右边查找
  return infixOrderSearch(node.right, no);
}

/**
 * 后序查找
 */
function postOrderSearch(node, no) {
  if (!node) {
    return null;
  }
  // 左边查找
  let returnNode = postOrderSearch(node.left, no);
  if (returnNode) {
    // 左边找到了节点，返回
    return returnNode
  }
  // 右边查找
  returnNode = postOrderSearch(node.right, no);
  if (returnNode !== null) {
    // 右边找到了节点，返回
    return returnNode;
  }
  console.log("进入查找");
  if (node.no === no) {
      returnNode = node;
  }
  return returnNode;
}

function initNode() {
  let root = new BinaryTreeNode(1);
  let node2 = new BinaryTreeNode(2);
  let node3 = new BinaryTreeNode(3);
  let node4 = new BinaryTreeNode(4);
  let node5 = new BinaryTreeNode(5);

  // 手动建立树的关系
  root.left = node2;
  root.right = node5;
  node2.left = node3;
  node2.right = node4;

  return root;
}

function testOrder() {
  let root = initNode();

  console.log("======前序遍历======");
  preOrder(root);

  console.log("======中续遍历======");
  infixOrder(root);

  console.log("======后续遍历======");
  postOrder(root);
}

function testSearch() {
  let root = initNode();

  let no = 4;

  console.log("======前序查找======");
  console.log("查找no=" + no);
  let node = preOrderSearch(root, no);
  console.log("查找结果: no=" + node.no);

  no = 4;
  console.log("======中序查找======");
  console.log("查找no=" + no);
  node = infixOrderSearch(root, no);
  console.log("查找结果: no=" + node.no);

  no = 4;
  console.log("======后序查找======");
  console.log("查找no=" +  no);
  node = postOrderSearch(root, no);
  console.log("查找结果: no=" + node.no);
}

function main() {
  // testOrder();
  // testSearch();
}

main();