use std::rc::Rc;
use std::cell::RefCell;
use std::fmt;

/// 二叉树
/// 二叉树是每个结点最多有两个子树的树结构
/// 通常子树被称作“左子树”（left subtree）和“右子树”（right subtree）
/// 二叉树常被用于实现二叉查找树和二叉堆

pub struct BinaryTreeNode {
    pub no: u32,
    pub left: Option<Rc<RefCell<BinaryTreeNode>>>,
    pub right: Option<Rc<RefCell<BinaryTreeNode>>>,
}

impl fmt::Debug for BinaryTreeNode {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "no:{}", self.no)
    }
}

impl BinaryTreeNode {
    pub fn new(no: u32) -> BinaryTreeNode {
        BinaryTreeNode { no, left: None, right: None }
    }
}

/// 前序遍历
fn pre_order(node: Option<Rc<RefCell<BinaryTreeNode>>>) {
    if node.is_none() {
        return;
    }
    // 当前节点
    println!("{:?}", node.as_ref().unwrap().borrow());
    // 遍历左子树
    pre_order(node.as_ref().unwrap().borrow_mut().left.clone());
    // 遍历右子树
    pre_order(node.as_ref().unwrap().borrow_mut().right.clone());
}

/// 中序遍历
fn infix_order(node: Option<Rc<RefCell<BinaryTreeNode>>>) {
    if node.is_none() {
        return;
    }
    // 遍历左子树
    infix_order(node.as_ref().unwrap().borrow_mut().left.clone());
    // 当前节点
    println!("{:?}", node.as_ref().unwrap().borrow());
    // 遍历右子树
    infix_order(node.as_ref().unwrap().borrow_mut().right.clone());
}

/// 后序遍历
fn post_order(node: Option<Rc<RefCell<BinaryTreeNode>>>) {
    if node.is_none() {
        return;
    }
    // 遍历左子树
    post_order(node.as_ref().unwrap().borrow_mut().left.clone());
    // 遍历右子树
    post_order(node.as_ref().unwrap().borrow_mut().right.clone());
    // 当前节点
    println!("{:?}", node.as_ref().unwrap().borrow());
}

/// 前序查找
fn pre_order_search(
    node: Option<Rc<RefCell<BinaryTreeNode>>>,
    no: u32
)-> Option<Rc<RefCell<BinaryTreeNode>>> {
    if node.is_none() {
        return None;
    }
    println!("进入查找");
    if node.as_ref().unwrap().borrow().no == no {
        return node;
    }
    // 左边查找
    let return_node = pre_order_search(
        node.as_ref().unwrap().borrow_mut().left.clone(), no);
    if return_node.is_some() {
        // 左边找到了节点，返回
        return return_node;
    }
    // 右边查找
    return pre_order_search(
        node.as_ref().unwrap().borrow_mut().right.clone(), no);
}

/// 中序查找
fn infix_order_search(
    node: Option<Rc<RefCell<BinaryTreeNode>>>,
    no: u32
)-> Option<Rc<RefCell<BinaryTreeNode>>> {
    if node.is_none() {
        return None;
    }
    // 左边查找
    let return_node = infix_order_search(
        node.as_ref().unwrap().borrow_mut().left.clone(), no);
    if return_node.is_some() {
        // 左边找到了节点，返回
        return return_node;
    }
    println!("进入查找");
    if node.as_ref().unwrap().borrow().no == no {
        return node;
    }
    // 右边查找
    return infix_order_search(
        node.as_ref().unwrap().borrow_mut().right.clone(), no);
}

/// 后序查找
fn post_order_search(
    node: Option<Rc<RefCell<BinaryTreeNode>>>,
    no: u32
)-> Option<Rc<RefCell<BinaryTreeNode>>> {
    if node.is_none() {
        return None;
    }
    // 左边查找
    let mut return_node = post_order_search(
        node.as_ref().unwrap().borrow_mut().left.clone(), no);
    if return_node.is_some() {
        // 左边找到了节点，返回
        return return_node;
    }
    // 右边查找
    return_node = post_order_search(node.as_ref().unwrap().borrow_mut().right.clone(), no);
    if return_node.is_some() {
        // 右边找到了节点，返回
        return return_node;
    }
    println!("进入查找");
    if node.as_ref().unwrap().borrow().no == no {
        return_node = node;
    }
    return return_node;
}

fn init_node() -> Option<Rc<RefCell<BinaryTreeNode>>> {
    let root = Rc::new(RefCell::new(BinaryTreeNode::new(1)));
    let node2 = Rc::new(RefCell::new(BinaryTreeNode::new(2)));
    let node3 = Rc::new(RefCell::new(BinaryTreeNode::new(3)));
    let node4 = Rc::new(RefCell::new(BinaryTreeNode::new(4)));
    let node5 = Rc::new(RefCell::new(BinaryTreeNode::new(5)));

    // 手动建立树的关系
    root.borrow_mut().left = Some(Rc::clone(&node2));
    root.borrow_mut().right = Some(node5);
    node2.borrow_mut().left = Some(node3);
    node2.borrow_mut().right = Some(node4);

    Some(root)
}

fn test_order() {
    let root = init_node();

    println!("======前序遍历======");
    pre_order(root.clone());

    println!("======中续遍历======");
    infix_order(root.clone());

    println!("======后续遍历======");
    post_order(root.clone())
}

fn test_search() {
    let root = init_node();

    let no = 4;
    println!("======前序查找======");
    println!("查找no={}", no);
    let node = pre_order_search(root.clone(), no);
    println!("查找结果: no={}", node.as_ref().unwrap().borrow().no);

    println!("======中序查找======");
    println!("查找no={}", no);
    let node = infix_order_search(root.clone(), no);
    println!("查找结果: no={}", node.as_ref().unwrap().borrow().no);

    println!("======后序查找======");
    println!("查找no={}", no);
    let node = post_order_search(root.clone(), no);
    println!("查找结果: no={}", node.as_ref().unwrap().borrow().no);
}

fn main() {
    // test_order();
    // test_search();
}
