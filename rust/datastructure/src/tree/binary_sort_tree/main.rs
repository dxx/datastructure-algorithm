use std::cell::RefCell;
use std::fmt;
use std::rc::Rc;

/// 二叉排序树
/// 二叉排序树（Binary Sort Tree），又称二叉查找树（Binary Search Tree），也叫二叉搜索树
/// 在一般情况下，查询效率比链表结构要高。对于任何一个非叶子节点，它的左子节点的值小于自身的值，
/// 它的右子节点的值大于自身的值，具有这样性质的二叉树称为二叉排序树

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

/// 二叉排序树结构体
pub struct BinarySortTree {
    root: Option<Rc<RefCell<BinaryTreeNode>>>, // 树的根节点
}

impl BinarySortTree {
    pub fn new() -> Self {
        BinarySortTree { root: None }
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
        self.recursion_add(&mut self.root.clone(), &mut node)
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

    pub fn search(
        &self,
        no: u32,
    ) -> (
        Option<Rc<RefCell<BinaryTreeNode>>>,
        Option<Rc<RefCell<BinaryTreeNode>>>,
    ) {
        if self.root.is_none() {
            return (None, None);
        }
        if self.root.as_ref().unwrap().borrow().no == no {
            return (None, self.root.clone());
        }
        return self.recursion_search(self.root.clone(), no);
    }

    /// 递归查找指定节点
    /// 返回查找到的父节点和查找到的节点
    fn recursion_search(
        &self,
        node: Option<Rc<RefCell<BinaryTreeNode>>>,
        no: u32,
    ) -> (
        Option<Rc<RefCell<BinaryTreeNode>>>,
        Option<Rc<RefCell<BinaryTreeNode>>>,
    ) {
        if node.is_none() {
            return (None, None);
        }
        let borrowed_node = node.as_ref().unwrap().borrow();
        if borrowed_node.left.is_some() && borrowed_node.left.as_ref().unwrap().borrow().no == no {
            return (node.clone(), borrowed_node.left.clone());
        }
        if borrowed_node.right.is_some() && borrowed_node.right.as_ref().unwrap().borrow().no == no
        {
            return (node.clone(), borrowed_node.right.clone());
        }
        // 判断是往左边还是往右边查找
        return if no < borrowed_node.no {
            self.recursion_search(borrowed_node.left.clone(), no)
        } else {
            self.recursion_search(borrowed_node.right.clone(), no)
        };
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

fn main() {
    let nos: [u32; 8] = [8, 5, 10, 3, 6, 9, 12, 2];
    let mut binary_sort_tree = BinarySortTree::new();
    for no in nos.iter() {
        binary_sort_tree.add(Some(Rc::new(RefCell::new(BinaryTreeNode::new(*no)))))
    }

    println!("======中序遍历======");
    binary_sort_tree.infix_order();

    println!("======查找节点======");
    let (_parent, node) = binary_sort_tree.search(6);
    println!("查找no={}", node.as_ref().unwrap().borrow().no);
}
