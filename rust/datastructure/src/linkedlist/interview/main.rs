#[derive(Clone)]
pub struct Node {
    name: String,
    next: Option<Box<Node>>
}

impl Node {
    pub fn new(name: String) -> Self {
        return Node {name, next: None}
    }
}

fn insert_at_tail(
    mut head_node: Option<Box<Node>>,
    new_node: Option<Box<Node>>
) -> Option<Box<Node>> {
    let mut last_node = head_node.as_mut().unwrap();
    while let Some(ref mut node) = last_node.next {
        last_node = node;
    }
    last_node.next = new_node;
    return head_node;
}

fn print_node_info(head_node: Option<Box<Node>>) {
    if head_node.is_none() {
        println!("该链表没有结点");
        return;
    }
    let mut temp_node = &head_node;
    let mut info = String::from("[");
    while let Some(node) = temp_node {
        info.push_str(&format!("{{name:{}}}", node.name));
        temp_node = &node.next;
    }
    info.push_str("]");
    println!("{}", info);
}

/// 获取单链表有效的结点数
/// 1.遍历结点数
/// 2.定义一个长度遍历，每遍历一次长度 +1
fn get_node_length(head_node: Option<Box<Node>>) -> u32 {
    if head_node.is_none() {
        return 0;
    }
    let mut length = 0;
    let mut node = &head_node;
    while let Some(unwrap_node) = node {
        length += 1;
        node = &unwrap_node.next;
    }
    return length
}

/// 获取倒数第 n 个结点
/// 遍历
/// 1.获取链表结点数 length
/// 2.遍历到 length - n 个结点
/// 3.然后返回
fn get_last_index_node(
    mut head_node: Option<Box<Node>>,
    mut index: usize
) -> Option<Box<Node>> {
    // 头结点为空
    if head_node.is_none() {
        return None;
    }
    let length = get_node_length(head_node.clone());
    if index <= 0 || index > (length as usize) {
        return None;
    }
    let mut last_node = &mut head_node;
    for _ in 0..(length as usize - index) {
        last_node = &mut last_node.as_mut().unwrap().next;
    }
    return last_node.take();
}

/// 单链表反转
/// 1.定义一个新的头结点 reverseHead
/// 2.遍历链表，每遍历一个结点，将其取出，放在新的头结点 reverseHead 的后面
/// 3.最后将头结点的 next 结点指向 reverseHead 的 next 结点
fn reverse_node(head_node: Option<Box<Node>>) -> Option<Box<Node>> {
    if head_node.is_none() {
        return None;
    }
    let mut reverse_head = Some(Node::new(String::from("")));
    let mut current = head_node;
    while current.is_some() {
        let mut node = current.unwrap();
        let mut next = node.next.take();
        node.next = reverse_head.as_mut().unwrap().next.take();
        reverse_head.as_mut().unwrap().next = Some(node);
        current = next;
    }
    return reverse_head.take().unwrap().next;
}

fn test_get_length() {
    let mut head_node = Some(Box::new(Node::new(String::from("node1"))));
    let node2 = Some(Box::new(Node::new(String::from("node2"))));
    let node3 = Some(Box::new(Node::new(String::from("node3"))));
    head_node = insert_at_tail(head_node, node2);
    head_node = insert_at_tail(head_node, node3);

    let length = get_node_length(head_node);
    println!("单链表结点个数为: {}\n", length);
}

fn test_get_last_index_node() {
    let mut head_node = Some(Box::new(Node::new(String::from("node1"))));
    let node2 = Some(Box::new(Node::new(String::from("node2"))));
    let node3 = Some(Box::new(Node::new(String::from("node3"))));
    head_node = insert_at_tail(head_node, node2);
    head_node = insert_at_tail(head_node, node3);

    let index = 2;
    let last_node = get_last_index_node(head_node, index);
    println!("单链表结点中倒数第{}个结点为: {}", index, last_node.as_ref().unwrap().name);
}

fn test_reverse_node() {
    let mut head_node = Some(Box::new(Node::new(String::from("node1"))));
    let node2 = Some(Box::new(Node::new(String::from("node2"))));
    let node3 = Some(Box::new(Node::new(String::from("node3"))));
    head_node = insert_at_tail(head_node, node2);
    head_node = insert_at_tail(head_node, node3);

    println!("反转前:");
    print_node_info(head_node.clone());

    head_node = reverse_node(head_node);
    println!("反转后:");
    print_node_info(head_node.clone());
}

fn main() {
    // test_get_length();
    // test_get_last_index_node();
	// test_reverse_node();
}