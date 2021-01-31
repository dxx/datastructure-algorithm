/**
 * AVL 树
 * 在 AVL 树中任何节点的两个子树的高度最大差别为 1 所以它也被称为高度平衡树
 * 增加和删除可能需要通过一次或多次树旋转来重新平衡这个树。AVL 树本质上是带了
 * 平衡功能的二叉排序树（二叉查找树，二叉搜索树）
 */

function BinaryTreeNode(no) {
  this.no = no;
  this.left = null;
  this.right = null;
}

function AVLTree() {
  this.root = null; // 树的根节点
}

/**
 * 添加结点
 */
AVLTree.prototype.add = function(node) {
  if (node === null) {
    return;
  }
  // 根节点为 nil，要添加的节点作为根节点
  if (this.root === null) {
    this.root = node;
    return;
  }
  this._add(this.root, node);

  // 添加节点后判断是否需要旋转

  // 右边高度超过左边 1 个高度以上，进行左旋转
  if (this._rightHeight() - this._leftHeight() > 1) {

    let rightNode = this.root.right;
    // 右子节点不为 null，并且右子节点的左子树高度大于右子节点的右子树高度
    if (rightNode != null &&
        this._height(rightNode.left) > this._height(rightNode.right)) {
        // 将右子节点右旋转
        this._rightRotate(rightNode);  
    }
    
    this._leftRotate(this.root);
    return;
  }

  if (this._leftHeight() - this._rightHeight() > 1) {

    let leftNode = this.root.left;
    // 左子节点不为 nil, 并且左子节点的右子树高度大于左子节点的左子树高度
    if (leftNode != null &&
      this._height(leftNode.right) - this._height(leftNode.left)) {
      // 将左子节点左旋转
      this._leftRotate(leftNode);
    }

    this._rightRotate(this.root);
  }
}

/**
 * 计算节点的高度
 */
AVLTree.prototype._height = function(node) {
  if (!node) {
    return 0;
  }
  return Math.max(this._height(node.left), this._height(node.right)) + 1;
}

/**
 * 左子树高度
 */
AVLTree.prototype._leftHeight = function() {
  if (this.root == null) {
    return 0;
  }
  return this._height(this.root.left);
}

/**
 * 右子树高度
 */
AVLTree.prototype._rightHeight = function() {
  if (this.root == null) {
    return 0;
  }
  return this._height(this.root.right);
}

/**
 * 左旋转
 */
AVLTree.prototype._leftRotate = function(node) {
  if (!node) {
    return;
  }
  // 以当前节点为基础，创建一个新的节点，新节点的值等于当前节点的值
  let newNode = new BinaryTreeNode(node.no);
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
AVLTree.prototype._rightRotate = function(node) {
  if (!node) {
    return;
  }
  // 以当前节点为基础，创建一个新的节点，新节点的值等于根节点的值
  let newNode = new BinaryTreeNode(node.no);
  // 让新节点的右子节点指向当前节点的右子节点，左子节点指向当前节点的左子节点的右子节点
  newNode.right = node.right;
  newNode.left = node.left.right;
  // 把当前节点的值替换为左子节点的值，并把当前节点左子节点指向其左子节点的左子节点
  node.no = node.left.no;
  node.left = node.left.left;
  // 让当前节点的右子节点指向新创建的节点
  node.right = newNode;
}

AVLTree.prototype._add = function(root, node) {
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


AVLTree.prototype.infixOrder = function() {
  this._infixOrder(this.root);
}

AVLTree.prototype._infixOrder = function(node) {
  if (node === null) {
    return;
  }
  this._infixOrder(node.left);
  console.log("no:" + node.no);
  this._infixOrder(node.right);
}

function testLeftRotate() {
  let nos = [3, 2, 5, 4, 6, 7];
  let avlTree = new AVLTree();
  for (let i = 0; i < nos.length; i++) {
      avlTree.add(new BinaryTreeNode(nos[i]));
  }

  console.log("左旋转后");

  avlTree.infixOrder();

  console.log("根节点 = ", avlTree.root.no);

  console.log("左子树的高度为: ", avlTree._leftHeight());
  console.log("右子树的高度为: ", avlTree._rightHeight());
}

function testRightRotate() {
  let nos = [6, 4, 7, 3, 5, 2];
  let avlTree = new AVLTree();
  for (let i = 0; i < nos.length; i++) {
      avlTree.add(new BinaryTreeNode(nos[i]));
  }

  console.log("右旋转后");

  avlTree.infixOrder();

  console.log("根节点 = ", avlTree.root.no);

  console.log("左子树的高度为: ", avlTree._leftHeight());
  console.log("右子树的高度为: ", avlTree._rightHeight());
}

function testDoubleRotate() {
  let nos = [6, 3, 7, 2, 4, 5];
  let avlTree = new AVLTree();
  for (let i = 0; i < nos.length; i++) {
      avlTree.add(new BinaryTreeNode(nos[i]));
  }

  console.log("双旋转后");

  avlTree.infixOrder();

  console.log("根节点 = ", avlTree.root.no);

  console.log("左子树的高度为: ", avlTree._leftHeight());
  console.log("右子树的高度为: ", avlTree._rightHeight());
}

function main() {
  // testLeftRotate();
  // testRightRotate();
  // testDoubleRotate();
}

main()
  