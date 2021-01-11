use std::cell::RefCell;
use std::cmp::max;
use std::fmt;
use std::rc::Rc;

/// AVL 树
/// 在 AVL 树中任何节点的两个子树的高度最大差别为 1 所以它也被称为高度平衡树
/// 增加和删除可能需要通过一次或多次树旋转来重新平衡这个树。AVL 树本质上是带了
/// 平衡功能的二叉排序树（二叉查找树，二叉搜索树）

/// 树节点结构体
pub struct BinaryTreeNode {
    pub no: u32,                                    // 编号
    pub left: Option<Rc<RefCell<BinaryTreeNode>>>,  // 左子节点
    pub right: Option<Rc<RefCell<BinaryTreeNode>>>, // 右子节点
}

impl fmt::Debug for BinaryTreeNode {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "no:{}", self.no)
    }
}

impl BinaryTreeNode {
    pub fn new(no: u32) -> Self {
        BinaryTreeNode {
            no,
            left: None,
            right: None,
        }
    }
}

/// AVL 树结构体
pub struct AVLTree {
    root: Option<Rc<RefCell<BinaryTreeNode>>>, // 树的根节点
}

impl AVLTree {
    pub fn new() -> Self {
        AVLTree { root: None }
    }

    /// 左子树高度
    fn left_height(&self) -> u32 {
        if self.root.is_none() {
            return 0;
        }
        return self.height(self.root.as_ref().unwrap().borrow().left.clone());
    }

    /// 左子树高度
    fn right_height(&self) -> u32 {
        if self.root.is_none() {
            return 0;
        }
        return self.height(self.root.as_ref().unwrap().borrow().right.clone());
    }

    /// 计算节点树的高度
    fn height(&self, node: Option<Rc<RefCell<BinaryTreeNode>>>) -> u32 {
        if node.is_none() {
            return 0;
        }
        let borrowed_node = node.as_ref().unwrap().borrow();
        // 递归左子节点
        let l_height = self.height(borrowed_node.left.clone());
        // 递归右子节点
        let r_height = self.height(borrowed_node.right.clone());
        // 取最大值，1 表示当前节点的高度值
        return max(l_height, r_height) + 1;
    }

    /// 左旋转
    fn left_rotate(&self, node: &mut Option<Rc<RefCell<BinaryTreeNode>>>) {
        if node.is_none() {
            return;
        }
        let mut borrowed_node = node.as_mut().unwrap().borrow_mut();
        // 以当前节点为基础，创建一个新的节点，新节点的值等于当前节点的值
        let mut new_node = BinaryTreeNode::new(borrowed_node.no);
        // 让新节点的左子节点指向当前节点的左子节点，右子节点指向当前节点的右子节点的左子节点
        new_node.left = borrowed_node.left.take();
        let node_right = borrowed_node.right.take();
        new_node.right = node_right.as_ref().unwrap().borrow_mut().left.take();
        // 把当前节点的值替换为右子节点的值，并把当前节点右子节点指向其右子节点的右子节点
        borrowed_node.no = node_right.as_ref().unwrap().borrow().no;
        borrowed_node.right = node_right.as_ref().unwrap().borrow_mut().right.take();
        // 让当前节点的左子节点指向新创建的节点
        borrowed_node.left = Some(Rc::new(RefCell::new(new_node)));
    }

    /// 右旋转
    fn right_rotate(&self, node: &mut Option<Rc<RefCell<BinaryTreeNode>>>) {
        if node.is_none() {
            return;
        }
        let mut borrowed_node = node.as_mut().unwrap().borrow_mut();
        // 以当前节点为基础，创建一个新的节点，新节点的值等于根节点的值
        let mut new_node = BinaryTreeNode::new(borrowed_node.no);
        // 让新节点的右子节点指向当前节点的右子节点，左子节点指向当前节点的左子节点的右子节点
        new_node.right = borrowed_node.right.take();
        let node_left = borrowed_node.left.take();
        new_node.left = node_left.as_ref().unwrap().borrow_mut().right.take();
        // 把当前节点的值替换为左子节点的值，并把当前节点左子节点指向其左子节点的左子节点
        borrowed_node.no = node_left.as_ref().unwrap().borrow().no;
        borrowed_node.left = node_left.as_ref().unwrap().borrow_mut().left.take();
        // 让当前节点的右子节点指向新创建的节点
        borrowed_node.right = Some(Rc::new(RefCell::new(new_node)));
    }

    /// 添加结点
    pub fn add(&mut self, mut node: Option<Rc<RefCell<BinaryTreeNode>>>) {
        if node.is_none() {
            return;
        }
        // 根节点为 None，要添加的节点作为根节点
        if self.root.is_none() {
            self.root = node;
            return;
        }
        self.recursion_add(&mut self.root.clone(), &mut node);

        // 添加节点后判断是否需要旋转

        // 右边高度超过左边 1 个高度以上，进行左旋转
        if (self.right_height() as i32 - self.left_height() as i32) > 1 {
            let mut right_node = self.root.as_ref().unwrap().borrow().right.clone();
            // 右子节点不为 None，并且右子节点的左子树高度大于右子节点的右子树高度
            if right_node.is_some()
                && self.height(right_node.as_ref().unwrap().borrow().left.clone())
                    > self.height(right_node.as_ref().unwrap().borrow().right.clone())
            {
                // 将右子节点右旋转
                self.right_rotate(&mut right_node);
            }
            self.left_rotate(&mut self.root.clone());
            return;
        }

        // 左边高度超过右边 1 个高度以上，进行右旋转
        if (self.left_height() as i32 - self.right_height() as i32) > 1 {
            let mut left_node = self.root.as_ref().unwrap().borrow().left.clone();
            // 左子节点不为 None, 并且左子节点的右子树高度大于左子节点的左子树高度
            if left_node.is_some()
                && self.height(left_node.as_ref().unwrap().borrow().right.clone())
                    > self.height(left_node.as_ref().unwrap().borrow().left.clone())
            {
                // 将左子节点左旋转
                self.left_rotate(&mut left_node);
            }
            self.right_rotate(&mut self.root.clone());
        }
    }

    fn recursion_add(
        &self,
        root: &mut Option<Rc<RefCell<BinaryTreeNode>>>,
        node: &mut Option<Rc<RefCell<BinaryTreeNode>>>,
    ) {
        // 要添加的节点小于根节点
        if node.as_ref().unwrap().borrow().no < root.as_ref().unwrap().borrow().no {
            let mut borrowed = root.as_mut().unwrap().borrow_mut();
            // 左子节点为 None，直接添加为左子节点
            if borrowed.left.is_none() {
                borrowed.left = Some(node.take().unwrap());
                return;
            }
            // 左递归
            self.recursion_add(&mut borrowed.left, node);
        } else {
            let mut borrowed = root.as_mut().unwrap().borrow_mut();
            // 右子节点为 None，直接添加为右子节点
            if borrowed.right.is_none() {
                borrowed.right = Some(node.take().unwrap());
                return;
            }
            // 右递归
            self.recursion_add(&mut borrowed.right, node);
        }
    }

    pub fn infix_order(&self) {
        if self.root.is_none() {
            return;
        }
        self.recursion_infix_order(self.root.clone());
    }

    fn recursion_infix_order(&self, node: Option<Rc<RefCell<BinaryTreeNode>>>) {
        if node.is_none() {
            return;
        }
        let borrowed_node = node.as_ref().unwrap().borrow();
        self.recursion_infix_order(borrowed_node.left.clone());
        println!("{:?}", borrowed_node);
        self.recursion_infix_order(borrowed_node.right.clone());
    }
}

fn test_left_rotate() {
    let nos: [u32; 6] = [3, 2, 5, 4, 6, 7];
    let mut avl_tree = AVLTree::new();
    for no in nos.iter() {
        avl_tree.add(Some(Rc::new(RefCell::new(BinaryTreeNode::new(*no)))))
    }

    println!("左旋转后");
    avl_tree.infix_order();

    println!("根节点={:?}", avl_tree.root.as_ref().unwrap().borrow());

    println!("左子树的高度为:{}", avl_tree.left_height());
    println!("右子树的高度为:{}", avl_tree.right_height());
}

fn test_right_rotate() {
    let nos: [u32; 6] = [6, 4, 7, 3, 5, 2];
    let mut avl_tree = AVLTree::new();
    for no in nos.iter() {
        avl_tree.add(Some(Rc::new(RefCell::new(BinaryTreeNode::new(*no)))))
    }

    println!("右旋转后");
    avl_tree.infix_order();

    println!("根节点={:?}", avl_tree.root.as_ref().unwrap().borrow());

    println!("左子树的高度为:{}", avl_tree.left_height());
    println!("右子树的高度为:{}", avl_tree.right_height());
}

fn test_double_rotate() {
    let nos: [u32; 6] = [6, 3, 7, 2, 4, 5];
    let mut avl_tree = AVLTree::new();
    for no in nos.iter() {
        avl_tree.add(Some(Rc::new(RefCell::new(BinaryTreeNode::new(*no)))))
    }

    println!("双旋转后");
    avl_tree.infix_order();

    println!("根节点={:?}", avl_tree.root.as_ref().unwrap().borrow());

    println!("左子树的高度为:{}", avl_tree.left_height());
    println!("右子树的高度为:{}", avl_tree.right_height());
}

fn main() {
    // test_left_rotate();
    // test_right_rotate();
    // test_double_rotate();
}
