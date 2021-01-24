/// 单向链表
/// 单向链表是链表的一种，其特点是链表的链接方向是单向的，对链表的访问要从头部开始
/// 链表是由结点构成，head 指针指向第一个称为表头的结点，而最后一个结点的指针指向 NULL

#[derive(Debug, Clone)]
pub struct HeroNode {
    no: i32,                     // 编号
    name: String,                // 姓名
    nickname: String,            // 昵称
    next: Option<Box<HeroNode>>, // 下一个结点
}

impl HeroNode {
    pub fn new(no: i32, name: String, nickname: String) -> Box<Self> {
        return Box::new(HeroNode {
            no,
            name,
            nickname,
            next: None,
        });
    }
}

/// 在链表尾部插入，通过 head 找到链表的尾部
fn insert_at_tail(
    mut head_node: Option<Box<HeroNode>>,
    new_node: Option<Box<HeroNode>>,
) -> Option<Box<HeroNode>> {
    let mut last_node = head_node.as_mut().unwrap();
    // 下一个结点不为空继续循环
    while last_node.next.is_some() {
        // 将下一个结点赋值给当前结点
        last_node = last_node.next.as_mut().unwrap();
    }
    // 将当前结点插入到链表的最后一个结点
    last_node.next = new_node;
    return head_node;
}

/// 按照 no 升序插入，通过 head 找到合适的插入位置
fn sort_insert_by_no(
    mut head_node: Option<Box<HeroNode>>,
    mut new_node: Option<Box<HeroNode>>,
) -> Option<Box<HeroNode>> {
    let mut temp_node = head_node.as_mut().unwrap();
    loop {
        let node = temp_node;
        // 最后一个结点，跳出循环
        if node.next.is_none() || node.next.as_ref().unwrap().no > new_node.as_ref().unwrap().no {
            // new_node 结点应该插入到 node 后面
            // node 的下一个结点插入到 new_node 的下一个结点
            new_node.as_mut().unwrap().next = node.next.clone();
            // new_node 结点插入到 new_node 的下一个结点
            node.next = Some(new_node.unwrap());
            break;
        } else if node.next.as_ref().unwrap().no == new_node.as_ref().unwrap().no {
            panic!("no 相等不能插入") // no 相等不能插入
        }
        temp_node = node.next.as_mut().unwrap();
    }
    return head_node;
}

/// 删除指定结点
fn delete_node(mut head_node: Option<Box<HeroNode>>, no: i32) -> Option<Box<HeroNode>> {
    let mut temp_node = head_node.as_mut().unwrap();
    while temp_node.next.is_some() {
        if temp_node.next.as_ref().unwrap().no == no {
            // 将下一个结点的下一个结点，链接到被删除结点的上一个结点
            temp_node.next = temp_node.next.take().unwrap().next;
            return head_node;
        }
        temp_node = temp_node.next.as_mut().unwrap();
    }
    return head_node;
}

/// 打印单链表结点内容
fn print_head_node_info(head_node: Option<Box<HeroNode>>) {
    if head_node.is_none() || head_node.as_ref().unwrap().next.is_none() {
        println!("该链表没有结点");
        return;
    }
    let mut temp_node = &head_node.as_ref().unwrap().next;
    let mut info = String::from("[");
    while temp_node.is_some() {
        let node = temp_node.as_ref().unwrap();
        info.push_str(&format!(
            "{{no:{}, name:{}, nickname:{}}}",
            node.no, node.name, node.nickname
        ));
        temp_node = &node.next;
    }
    info.push_str("]");
    println!("{}", info);
}

#[test]
fn test_insert_at_tail() {
    // 创建 head 结点，head 结点不包含数据
    let mut head_node = Some(HeroNode::new(0, String::from(""), String::from("")));
    // 创建第一个结点
    let hero_node1 = Some(HeroNode::new(
        1,
        String::from("宋江"),
        String::from("呼保义"),
    ));
    // 创建第二个结点
    let hero_node2 = Some(HeroNode::new(
        2,
        String::from("卢俊义"),
        String::from("玉麒麟"),
    ));
    // 创建第三个结点
    let hero_node3 = Some(HeroNode::new(
        3,
        String::from("吴用"),
        String::from("智多星"),
    ));

    head_node = insert_at_tail(head_node, hero_node1);
    head_node = insert_at_tail(head_node, hero_node2);
    head_node = insert_at_tail(head_node, hero_node3);

    print_head_node_info(head_node);
}

#[test]
fn test_sort_insert_by_no() {
    // 创建结点，用来做尾部插入
    let mut head = Some(HeroNode::new(0, String::from(""), String::from("")));
    let node1 = Some(HeroNode::new(
        1,
        String::from("宋江"),
        String::from("呼保义"),
    ));
    let node2 = Some(HeroNode::new(
        2,
        String::from("卢俊义"),
        String::from("玉麒麟"),
    ));
    let node3 = Some(HeroNode::new(
        3,
        String::from("吴用"),
        String::from("智多星"),
    ));

    // 将结点按照 no 升序插入
    head = insert_at_tail(head, node1);
    head = insert_at_tail(head, node3); // 将第三个结点插入到第二个位置
    head = insert_at_tail(head, node2);

    println!("尾部插入的结果:");
    print_head_node_info(head);

    // 创建 head 结点
    let mut head_node = Some(HeroNode::new(0, String::from(""), String::from("")));
    let hero_node1 = Some(HeroNode::new(
        1,
        String::from("宋江"),
        String::from("呼保义"),
    ));
    let hero_node2 = Some(HeroNode::new(
        2,
        String::from("卢俊义"),
        String::from("玉麒麟"),
    ));
    let hero_node3 = Some(HeroNode::new(
        3,
        String::from("吴用"),
        String::from("智多星"),
    ));

    // 将结点按照 no 升序插入
    head_node = sort_insert_by_no(head_node, hero_node1);
    head_node = sort_insert_by_no(head_node, hero_node3);
    head_node = sort_insert_by_no(head_node, hero_node2);

    println!("按照 no 升序插入的结果:");
    print_head_node_info(head_node);
}

#[test]
fn test_delete_node() {
    // 创建结点
    let mut head_node = Some(HeroNode::new(0, String::from(""), String::from("")));
    let hero_node1 = Some(HeroNode::new(
        1,
        String::from("宋江"),
        String::from("呼保义"),
    ));
    let hero_node2 = Some(HeroNode::new(
        2,
        String::from("卢俊义"),
        String::from("玉麒麟"),
    ));
    let hero_node3 = Some(HeroNode::new(
        3,
        String::from("吴用"),
        String::from("智多星"),
    ));
    let hero_node4 = Some(HeroNode::new(
        4,
        String::from("公孙胜"),
        String::from("入云龙"),
    ));
    let hero_node5 = Some(HeroNode::new(5, String::from("关胜"), String::from("大刀")));

    head_node = insert_at_tail(head_node, hero_node1);
    head_node = insert_at_tail(head_node, hero_node2.clone());
    head_node = insert_at_tail(head_node, hero_node3.clone());
    head_node = insert_at_tail(head_node, hero_node4.clone());
    head_node = insert_at_tail(head_node, hero_node5.clone());

    println!("删除前:");
    print_head_node_info(head_node.clone());

    // 删除 no 为 2 的结点
    head_node = delete_node(head_node, 2);
    println!("删除 no 为 2 的结点后:");
    print_head_node_info(head_node.clone());

    // 删除 no 为 3, 4 的结点
    head_node = delete_node(head_node, 3);
    head_node = delete_node(head_node, 4);
    println!("删除 no 为 3,4 的结点后:");
    print_head_node_info(head_node.clone());
}
