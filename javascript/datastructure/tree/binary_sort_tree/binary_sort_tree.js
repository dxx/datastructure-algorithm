/**
 * 二叉排序树
 * 二叉排序树（Binary Sort Tree），又称二叉查找树（Binary Search Tree），也叫二叉搜索树
 * 在一般情况下，查询效率比链表结构要高。对于任何一个非叶子节点，它的左子节点的值小于自身的值，
 * 它的右子节点的值大于自身的值，具有这样性质的二叉树称为二叉排序树
 */

function BinaryTreeNode(no) {
  this.no = no;
  this.left = null;
  this.right = null;
}

function BinarySortTree() {
  this.root = null; // 树的根节点
}

/**
 * 添加结点
 */
BinarySortTree.prototype.add = function(node) {
  if (node === null) {
    return;
  }
  // 根节点为 nil，要添加的节点作为根节点
  if (this.root === null) {
    this.root = node;
    return;
  }
  this._add(this.root, node);
}

BinarySortTree.prototype._add = function(root, node) {
  // 要添加的节点小于根节点
  if (node.no < root.no) {
    // 左子节点为 null，直接添加为左子节点
    if (root.left === null) {
      root.left = node;
      return;
    }
    // 左递归
    this._add(root.left, node);
  } else {
    // 右子节点为 null，直接添加为左子节点
    if (root.right === null) {
      root.right = node;
      return;
    }
    // 右递归
    this._add(root.right, node);
  }
}

BinarySortTree.prototype.search = function(no) {
  if (this.root === null) {
    return null;
  }
  // 查找的节点就是根节点
  if (this.root.no === no) {
    return [null, this.root];
  }
  return this._recursionSearch(this.root, no);
}

/**
 * 递归查找指定节点
 * 返回查找到的父节点和查找到的节点
 */
BinarySortTree.prototype._recursionSearch = function(node, no) {
  if (node === null) {
    return null;
  }
  if (node.left !== null && node.left.no === no) {
    return [node, node.left];
  }
  if (node.right !== null && node.right.no === no) {
    return [node, node.right];
  }
  // 判断是往左边还是往右边查找
  if (no < node.no) {
    return this._recursionSearch(node.left, no);
  } else {
    return this._recursionSearch(node.right, no);
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
BinarySortTree.prototype.delete = function(no) {
  let binaryTreeNodes = this.search(no);
  // 没有找到要删除的节点
  if (!binaryTreeNodes) {
    return;
  }
  let parentNode = binaryTreeNodes[0];
  let node = binaryTreeNodes[1];
  // 当前节点为叶子节点
  if (node.left === null && node.right === null) {
    // 被删除的节点为根节点
    if (parentNode === null) {
      this.root = null;
      return;
    }
    // 当前节点为父节点的左子节点
    if (parentNode.left !== null && parentNode.left.no === no) {
      parentNode.left = null;
    }
    // 当前节点为父节点的右子节点
    if (parentNode.right !== null && parentNode.right.no === no) {
      parentNode.right = null;
    }
    return
  }
  // 当前节点有两颗子树
  if (node.left !== null && node.right != null) {
    // 把右子节点作为根节点，从左边开始遍历到最后一个叶子节点
    let leftChildNode = node.right;
    while (leftChildNode.left !== null) {
      leftChildNode = leftChildNode.left;
    }
    // 删除最小的叶子节点
    this.delete(leftChildNode.no);

    // 替换掉被删除节点的值
    node.no = leftChildNode.no;
  } else { // 当前节点只有一颗子树
    let replaceNode = null;
    if (node.left !== null) {
      replaceNode = node.left;
    }
    if (node.right !== null) {
      replaceNode = node.right;
    }

    // 父节点为 null，表示根节点
    if (parentNode === null) {
      this.root = replaceNode;
    }
    // 当前节点为父节点的左子节点
    if (parentNode.left !== null && parentNode.left.no === no) {
      parentNode.left = replaceNode;
    }
    // 当前节点为父节点的右子节点
    if (parentNode.right !== null && parentNode.right.no === no) {
      parentNode.right = replaceNode;
    }
  }
}

BinarySortTree.prototype.infixOrder = function() {
  this._infixOrder(this.root);
}

BinarySortTree.prototype._infixOrder = function(node) {
  if (node === null) {
    return;
  }
  this._infixOrder(node.left);
  console.log("no:" + node.no);
  this._infixOrder(node.right);
}

function main() {
  let nos = [8, 5, 10, 3, 6, 9, 12, 2];
  let binarySortTree = new BinarySortTree();
  for (let i = 0; i < nos.length; i++) {
    binarySortTree.add(new BinaryTreeNode(nos[i]));
  }
  console.log("======中序遍历======");
  binarySortTree.infixOrder();

  binarySortTree.delete(6);

  console.log("======删除叶子节点 6======");
  binarySortTree.infixOrder();

  binarySortTree.delete(5);

  console.log("======删除只有一颗子树的节点 5======");
  binarySortTree.infixOrder();

  binarySortTree.delete(10);

  console.log("======删除有两颗子树的节点 10======");
  binarySortTree.infixOrder();
}

main()
