use std::cell::RefCell;
use std::rc::Rc;

/// 双向链表
/// 双向链表也叫双链表，是链表的一种，它的每个数据结点中都有两个指针，分别指向前一个和后一个结点
/// 所以，从双向链表中的任意一个结点开始，都可以很方便地访问它的前一个结点和后一个结点。

#[derive(Debug)]
pub struct HeroNode {
    no: i32,                             // 编号
    name: String,                        // 姓名
    nickname: String,                    // 昵称
    prev: Option<Rc<RefCell<HeroNode>>>, // 上一个结点
    next: Option<Rc<RefCell<HeroNode>>>, // 下一个结点
}

impl HeroNode {
    pub fn new(no: i32, name: String, nickname: String) -> Rc<RefCell<Self>> {
        return Rc::new(RefCell::new(HeroNode {
            no,
            name,
            nickname,
            prev: None,
            next: None,
        }));
    }
}

/// 在链表尾部插入，通过 head 找到链表的尾部
fn insert_at_tail(
    head_node: Option<Rc<RefCell<HeroNode>>>,
    new_node: Option<Rc<RefCell<HeroNode>>>,
) -> Option<Rc<RefCell<HeroNode>>> {
    let mut last_node = Rc::clone(head_node.as_ref().unwrap());
    loop {
        let next = match last_node.borrow().next.as_ref() {
            Some(n) => Rc::clone(n),
            None => {
                break;
            }
        };
        // 将下一个结点赋值给当前结点
        last_node = next;
    }
    let node = new_node.unwrap();
    // 将当前结点插入到链表的最后一个结点
    last_node.borrow_mut().next = Some(Rc::clone(&node));
    // 将新结点的上一个结点指向当前结点
    node.borrow_mut().prev = Some(Rc::clone(&last_node));
    return head_node;
}

/// 删除指定结点
fn delete_node(head_node: Option<Rc<RefCell<HeroNode>>>, no: i32) -> Option<Rc<RefCell<HeroNode>>> {
    let mut temp_node = Rc::clone(head_node.as_ref().unwrap());
    let mut deleted_node = HeroNode::new(0, String::from(""), String::from(""));
    loop {
        let next = match temp_node.borrow().next.as_ref() {
            Some(rc_node) => {
                if rc_node.borrow_mut().no == no {
                    deleted_node = Rc::clone(rc_node);
                    break;
                }
                Rc::clone(rc_node)
            }
            None => {
                break;
            }
        };
        temp_node = next;
    }
    let borrow_deleted_node = deleted_node.borrow_mut();
    // 未找到删除节点，直接返回
    if borrow_deleted_node.no == head_node.as_ref().unwrap().borrow().no {
        return head_node;
    }
    // 最后一个结点的 next 为 None
    if borrow_deleted_node.next.is_some() {
        // 将查找到的结点的上一个结点的下一个结点指针指向当前结点的下一个结点
        borrow_deleted_node.prev.as_ref().unwrap().borrow_mut().next =
            Some(Rc::clone(borrow_deleted_node.next.as_ref().unwrap()));
        // 将查找到的结点的下一个结点的上一个结点指针指向当前指针的上一个结点
        borrow_deleted_node.next.as_ref().unwrap().borrow_mut().prev =
            Some(Rc::clone(&borrow_deleted_node.prev.as_ref().unwrap()));
    } else {
        borrow_deleted_node.prev.as_ref().unwrap().borrow_mut().next = None;
    }
    return head_node;
}

/// 打印链表结点内容
fn print_head_node_info(head_node: Option<Rc<RefCell<HeroNode>>>) {
    let mut borrow_head_node = Rc::clone(head_node.as_ref().unwrap());
    if borrow_head_node.borrow().next.is_none() {
        println!("该链表没有结点");
        return;
    }
    let mut temp_node = borrow_head_node;
    let mut info = String::from("[");
    loop {
        let next = match temp_node.borrow().next.as_ref() {
            Some(rc_node) => {
                let info_node = rc_node.borrow();
                info.push_str(&format!(
                    "{{no:{}, name:{}, nickname:{}}}",
                    info_node.no, info_node.name, info_node.nickname
                ));
                Rc::clone(rc_node)
            }
            None => {
                break;
            }
        };
        temp_node = next;
    }
    info.push_str("]");
    println!("{}", info);
}

fn test_insert_at_tail() {
    // 创建 head 结点，head 结点不包含有效数据
    let mut head_node = Some(HeroNode::new(0, String::from(""), String::from("")));
    // 创建第一个结点
    let hero_node1 = Some(HeroNode::new(
        3,
        String::from("吴用"),
        String::from("智多星"),
    ));
    // 创建第二个结点
    let hero_node2 = Some(HeroNode::new(
        6,
        String::from("林冲"),
        String::from("豹子头"),
    ));
    // 创建第三个结点
    let hero_node3 = Some(HeroNode::new(
        7,
        String::from("秦明"),
        String::from("霹雳火"),
    ));

    // 将结点添加到链表尾部
    head_node = insert_at_tail(head_node, hero_node1);
    head_node = insert_at_tail(head_node, hero_node2);
    head_node = insert_at_tail(head_node, hero_node3);

    print_head_node_info(head_node);
}

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

    // 插入结点
    head_node = insert_at_tail(head_node, hero_node1);
    head_node = insert_at_tail(head_node, hero_node2);
    head_node = insert_at_tail(head_node, hero_node3);
    head_node = insert_at_tail(head_node, hero_node4);
    head_node = insert_at_tail(head_node, hero_node5);

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

fn main() {
    // test_insert_at_tail();
    // test_delete_node();
}
