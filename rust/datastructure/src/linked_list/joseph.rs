use std::cell::RefCell;
use std::rc::Rc;

/// 约瑟夫问题
/// 设编号为 1,2,...n 的 n 个人围坐一圈约定编号为 k(1<=k<=n) 的人
/// 从 1 开始报数，数到 m 的那个人出列它的下一位又从 1 开始报数数到 m
/// 的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列

#[derive(Eq, PartialEq, Debug)]
pub struct Person {
    pub no: u32,                           // 编号
    pub prev: Option<Rc<RefCell<Person>>>, // 前一个人
    pub next: Option<Rc<RefCell<Person>>>, // 后一个人
}

impl Person {
    pub fn new(no: u32) -> Self {
        Person {
            no,
            prev: None,
            next: None,
        }
    }
}

pub struct PersonLinkedList {
    pub first: Option<Rc<RefCell<Person>>>, // 第一个小孩
    pub length: usize,                      // 小孩数量
}

impl PersonLinkedList {
    pub fn new(count: usize) -> Self {
        if count < 1 {
            panic!("链表至少需要一个元素");
        }

        let mut person_linked_list = PersonLinkedList {
            first: None,
            length: count,
        };

        // 初始化小孩，构成双向环形链表
        let mut prev: Rc<RefCell<Person>> = Rc::new(RefCell::new(Person::new(0)));
        for i in 1..=count {
            let person = Rc::new(RefCell::new(Person::new(i as u32)));
            if i == 1 {
                // 初始化 first 节点
                person_linked_list.first = Some(Rc::clone(&person));
                let first = person_linked_list.first.as_ref().unwrap();
                first.borrow_mut().next = Some(Rc::clone(first));
                first.borrow_mut().prev = Some(Rc::clone(first));

                // 将 prev 指向 first 节点，继续下一次循环
                prev = Rc::clone(first);
                continue;
            }
            prev.borrow_mut().next = Some(Rc::clone(&person));
            person.borrow_mut().prev = Some(Rc::clone(&prev));

            let first = person_linked_list.first.as_ref().unwrap();
            // 新增加的节点的下一个节点指向第一节点
            person.borrow_mut().next = Some(Rc::clone(first));
            // 第一个节点的上一个节点指向新增加的节点
            first.borrow_mut().prev = Some(Rc::clone(&person));

            prev = person;
        }
        return person_linked_list;
    }

    pub fn count(&self, start: i32, num: i32) {
        if start < 1 || start > self.length as i32 {
            println!("start 不能小于 1 或者不能大于 {}", self.length);
            return;
        }
        if num > self.length as i32 {
            println!("num 不能大于元素个数: {}", self.length);
            return;
        }
        let mut current = Rc::clone(self.first.as_ref().unwrap());

        // 获取从 current 开始的第 start - 1 个节点
        current = get_next_num_node(current, (start - 1) as u32);

        loop {
            // 表示只有一个节点
            if current.borrow().prev.as_ref().unwrap().borrow().eq(&current.borrow())
                && current.borrow().next.as_ref().unwrap().borrow().eq(&current.borrow()) {
                break;
            }
            // 获取从 current 开始的第 num - 1 个节点
            current = get_next_num_node(current, (num - 1) as u32);
            // 删除元素
            let p = Rc::clone(current.borrow().prev.as_ref().unwrap());
            let n = Rc::clone(current.borrow().next.as_ref().unwrap());
            p.borrow_mut().next = Some(Rc::clone(&n));
            n.borrow_mut().prev = Some(Rc::clone(&p));

            println!("出队人的编号: {}", current.borrow().no);
            let next = match current.borrow().next.as_ref() {
                Some(node) => Rc::clone(node),
                None => break,
            };
            current = next;
        }
        println!("最后留下人的编号: {}", current.borrow().no);
    }

    pub fn show_persons(&self) {
        if self.first.is_none() {
            return;
        }
        let first = self.first.as_ref().unwrap();
        if first.borrow().eq(&first.borrow().next.as_ref().unwrap().borrow()) {
            println!("num: {}", first.borrow().no);
            return;
        }

        let mut current = Rc::clone(first);
        loop {
            println!("num: {}", current.borrow().no);
            let next = match current.borrow().next.as_ref() {
                Some(node) => Rc::clone(node),
                None => break,
            };
            current = next;

            if current.borrow().eq(&first.borrow()) {
                break;
            }
        }
    }
}

fn get_next_num_node(node: Rc<RefCell<Person>>, num: u32) -> Rc<RefCell<Person>> {
    let mut current = node;
    let mut i = 1;
    loop {
        // 循环 num 次
        if i > num {
            break;
        }
        let next = match current.borrow().next.as_ref() {
            Some(node) => Rc::clone(node),
            None => break,
        };
        current = next;
        i += 1;
    }
    return current;
}

#[test]
fn test_joseph() {
    let person_linked_list = PersonLinkedList::new(5);
    person_linked_list.show_persons();

    person_linked_list.count(1, 3);
}
