use std::cell::RefCell;
use std::rc::Rc;

/// 循环链表
/// 循环链表的特点是表中最后一个结点的指针域指向头结点，整个链表形成一个环

/// 双向循环链表
#[derive(Eq, PartialEq)]
pub struct PersonNode {
    no: i32,
    name: String,
    prev: Option<Rc<RefCell<PersonNode>>>,
    next: Option<Rc<RefCell<PersonNode>>>,
}

impl PersonNode {
    pub fn new(no: i32, name: String) -> Rc<RefCell<Self>> {
        return Rc::new(RefCell::new(PersonNode {
            no,
            name,
            prev: None,
            next: None,
        }));
    }
}

/// 插入结点
fn insert_node(
    head_node: Option<Rc<RefCell<PersonNode>>>,
    new_node: Option<Rc<RefCell<PersonNode>>>,
) -> Option<Rc<RefCell<PersonNode>>> {
    let mut last_node = Rc::clone(head_node.as_ref().unwrap());
    loop {
        let next = match last_node.borrow().next.as_ref() {
            Some(rc_node) => {
                // 下一个结点等于头结点跳出循环
                if rc_node == head_node.as_ref().unwrap() {
                    break;
                }
                Rc::clone(rc_node)
            }
            None => break,
        };
        last_node = next;
    }
    let borrow_new_node = new_node.unwrap();
    let borrow_head_node = head_node.as_ref().unwrap();
    // 将新结点添加到链表末尾
    last_node.borrow_mut().next = Some(Rc::clone(&borrow_new_node));
    borrow_new_node.borrow_mut().prev = Some(Rc::clone(&last_node));
    // 将新结点下一个结点指针指向头结点
    borrow_new_node.borrow_mut().next = Some(Rc::clone(borrow_head_node));
    borrow_head_node.borrow_mut().prev = Some(Rc::clone(&borrow_new_node));
    return head_node;
}

// 删除指定结点，返回头结点
fn delete_node(
    head_node: Option<Rc<RefCell<PersonNode>>>,
    no: i32,
) -> Option<Rc<RefCell<PersonNode>>> {
    let mut temp_node = Rc::clone(head_node.as_ref().unwrap());
    let mut deleted_node: Option<Rc<RefCell<PersonNode>>> = None;
    // 头结点就是要删除的结点
    if temp_node.borrow().no == no {
        deleted_node = Some(Rc::clone(&temp_node));
    } else {
        loop {
            let next = match temp_node.borrow().next.as_ref() {
                Some(rc_node) => {
                    // 找到删除的节点
                    if rc_node.borrow().no == no {
                        deleted_node = Some(Rc::clone(rc_node));
                        break;
                    }
                    // 下一个结点等于头结点跳出循环
                    if rc_node.borrow().no == head_node.as_ref().unwrap().borrow().no {
                        break;
                    }
                    Rc::clone(rc_node)
                }
                None => break,
            };
            temp_node = next;
        }
    }
    // 存在需要删除的结点
    if deleted_node.is_some() {
        let node = deleted_node.as_ref().unwrap();
        // 将查找到的结点的上一个结点的下一个结点指针指向当前结点的下一个结点
        node.borrow().prev.as_ref().unwrap().borrow_mut().next =
            Some(Rc::clone(node.borrow().next.as_ref().unwrap()));
        // 将查找到的结点的下一个结点的上一个结点指针指向当前指针的上一个结点
        node.borrow().next.as_ref().unwrap().borrow_mut().prev =
            Some(Rc::clone(node.borrow().prev.as_ref().unwrap()));
        if node.borrow().no == temp_node.borrow().no {
            // 头结点删除了，将头结点的下一个结点作为头结点
            return Some(Rc::clone(node.borrow().next.as_ref().unwrap()));
        }
    }
    return head_node;
}

/// 打印循环链表的信息
fn print_round_node_info(head_node: Option<Rc<RefCell<PersonNode>>>) {
    if head_node.is_none() {
        println!("该循环链表没有节点");
        return;
    }
    let mut info = String::from("[");
    let info_node = Rc::clone(head_node.as_ref().unwrap());
    info.push_str(&format!(
        "{{no:{}, name:{}}}",
        info_node.borrow().no,
        info_node.borrow().name
    ));
    let mut temp_node = Rc::clone(head_node.as_ref().unwrap());
    loop {
        let next = match temp_node.borrow().next.as_ref() {
            Some(rc_node) => {
                if rc_node == head_node.as_ref().unwrap() {
                    break;
                }
                info.push_str(&format!("{{no:{}, name:{}}}", rc_node.borrow().no, rc_node.borrow().name));
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

#[test]
fn test_insert_node() {
    // 创建第一个结点
    let mut head_node = Some(PersonNode::new(1, String::from("张三")));
    // 创建第二个结点
    let person_node2 = Some(PersonNode::new(2, String::from("李四")));
    // 创建第三个结点
    let person_node3 = Some(PersonNode::new(3, String::from("王五")));

    head_node = insert_node(head_node, person_node2);
    head_node = insert_node(head_node, person_node3);

    print_round_node_info(head_node);
}

#[test]
fn test_delete_node() {
    // 创建结点
    let mut head_node = Some(PersonNode::new(1, String::from("张三")));
    let hero_node2 = Some(PersonNode::new(2, String::from("李四")));
    let hero_node3 = Some(PersonNode::new(3, String::from("王五")));
    let hero_node4 = Some(PersonNode::new(4, String::from("赵六")));
    let hero_node5 = Some(PersonNode::new(5, String::from("孙七")));

    // 插入结点
    head_node = insert_node(head_node, hero_node2);
    head_node = insert_node(head_node, hero_node3);
    head_node = insert_node(head_node, hero_node4);
    head_node = insert_node(head_node, hero_node5);

    println!("删除前:");
    print_round_node_info(head_node.clone());

    // 删除 no 为 2 的结点
    head_node = delete_node(head_node, 2);
    println!("删除 no 为 2 的结点后:");
    print_round_node_info(head_node.clone());

    let new_node = Some(PersonNode::new(6, String::from("周八")));
    head_node = insert_node(head_node, new_node);
    println!("插入新结点:");
    print_round_node_info(head_node.clone());

    // 删除 no 为 1，3 的结点
    head_node = delete_node(head_node, 1);
    head_node = delete_node(head_node, 3);
    println!("删除 no 为 1,3 的结点后:");
    print_round_node_info(head_node.clone());
}
