use std::rc::Rc;
use std::cell::RefCell;
use std::fmt;

/// 哈夫曼树
/// 有 N 个权值作为 N 个叶子结点，构造一棵二叉树，如果该树的带权路径长度达到最小
/// 称这样的二叉树为最优二叉树，也称为哈夫曼树(Huffman Tree)。哈夫曼树是带权路
/// 径长度最短的树，权值较大的结点离根较近。哈夫曼树又称为最优树。

/// 构建步骤
/// 1.将 w1、w2、…、wn 看成一个序列, 每个数据可以看做一个权值
/// 2.将序列从小到大排序
/// 3.选出两个根节点的权值最小的树合并，作为一棵新树的左、右子节点，且新树的根节点权值为其左、右子树根节点权值之和
/// 4.从序列中删除选出的两个节点，并将新树加入序列
/// 5.重复 2、3、4 步，直到序列中只剩一棵树为止，该树即为所求得的哈夫曼树

pub struct Node {
    pub value: u32, // 权值
    pub left: Option<Rc<RefCell<Node>>>,  // 左子节点
    pub right: Option<Rc<RefCell<Node>>>, // 右子节点
}

impl Node {
    pub fn new(value: u32) -> Self {
        Node { value, left: None, right: None }
    }
}

impl fmt::Debug for Node {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "value: {}", self.value)
    }
}

/// 构建哈夫曼树
fn create_huffman_tree(nums: Vec<u32>) -> Option<Rc<RefCell<Node>>> {
    if nums.len() == 0 {
        return None;
    }
    let mut nodes: Vec<Rc<RefCell<Node>>> = Vec::new();
    for num in nums.iter() {
        nodes.push(Rc::new(RefCell::new(Node::new(*num))));
    }
    while nodes.len() > 1 {
        // 排序
        nodes.sort_by(|node1, node2|
            node1.borrow().value.partial_cmp(&node2.borrow().value).unwrap());
        println!("{:?}", nodes);

        let left = &nodes[0];  // 权值最小的元素
        let right = &nodes[1]; // 权值第二小的元素
        // 创建新的根节点
        let root = Rc::new(
            RefCell::new(
                Node::new(left.borrow().value + right.borrow().value)));
        // 构建二叉树
        root.borrow_mut().left = Some(Rc::clone(left));
        root.borrow_mut().right = Some(Rc::clone(right));

        // 删除处理过的节点
        nodes.remove(0);
        nodes.remove(0);

        // 将二叉树加入到 nodes
        nodes.push(root);
    }

    return Some(Rc::clone(&nodes[0]));
}

/// 前序遍历
fn pre_order(node: Option<Rc<RefCell<Node>>>) {
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

#[test]
fn test_huffman_tree() {
    let nums = vec![1, 7, 3, 8, 16];
    let root = Some(create_huffman_tree(nums));
    pre_order(root.unwrap());
}
