/// 哈希表
/// 哈希表也叫散列表，是根据关键值 key 直接进行访问的数据结构
/// 它通过把 key 映射到表中一个位置来记录，以加快查找的速度

/// 员工结构体
#[derive(Clone)]
pub struct Employee {
    pub id: u32,
    pub name: String,
}

impl Employee {
    pub fn new(id: u32, name: String) -> Self {
        return Employee { id, name };
    }
}

/// 单链表结构体
#[derive(Clone)]
pub struct LinkNode {
    pub employee: Box<Employee>,
    pub next: Option<Box<LinkNode>>,
}

impl LinkNode {
    pub fn new(employee: Box<Employee>) -> Self {
        return LinkNode {
            employee,
            next: None,
        };
    }
}

/// 哈希表结构体，包含一个 LinkNode 向量
pub struct Hashtable {
    pub link_array: Vec<Option<Box<LinkNode>>>,
}

impl Hashtable {
    /// 创建哈希表
    pub fn new(len: usize) -> Self {
        return Hashtable {
            link_array: vec![None; len],
        };
    }

    /// 添加员工的方法，按照 id 升序插入
    pub fn add(&mut self, employee: Box<Employee>) -> Result<bool, &str> {
        let id = employee.id;
        let link_array = &mut self.link_array;
        let mut link_node = Box::new(LinkNode::new(employee));
        // 计算下标
        let index = id as usize % link_array.len();
        // 添加到链表向量中，作为头结点
        if link_array[index].is_none() {
            link_array[index] = Some(link_node);
            return Ok(true);
        }

        let head_node = link_array[index].as_mut().unwrap();
        // 判断头结点
        if head_node.employee.id > id {
            link_node.next = link_array[index].take();
            link_array[index] = Some(link_node);
            return Ok(true);
        } else if head_node.employee.id == id {
            return Err("员工id重复");
        }

        // 查找后续结点中合适的位置插入
        let mut temp_node = head_node;
        loop {
            if temp_node.next.is_none() {
                break;
            } else if temp_node.next.as_ref().unwrap().employee.id > id {
                break;
            } else if temp_node.next.as_ref().unwrap().employee.id == id {
                return Err("员工id重复");
            }
            temp_node = temp_node.next.as_mut().unwrap();
        }
        // temp_node 的下一个结点插入到 link_node 的下一个结点
        link_node.next = temp_node.next.take();
        // 将 link_node 插入到 temp_node 后面
        temp_node.next = Some(link_node);
        return Ok(true);
    }

    /// 修改员工的方法
    pub fn update(&mut self, employee: Box<Employee>) {
        let mut emp = self.get_employee_by_id(employee.id);
        if emp.is_some() {
            emp.as_mut().unwrap().name = employee.name;
        }
    }

    /// 删除员工的方法
    pub fn delete(&mut self, id: u32) {
        let link_array = &mut self.link_array;
        // 计算下标
        let index = id as usize % link_array.len();
        // 数组没有元素
        if link_array[index].is_none() {
            return;
        }
        let head_node = link_array[index].as_mut().unwrap();
        // 判断头结点
        if head_node.employee.id == id {
            link_array[index] = head_node.next.take();
            return;
        }

        // 查找后续结点中需要删除的员工
        let mut temp_node = head_node;
        loop {
            let next = match temp_node.next.as_mut() {
                Some(node) => {
                    if node.employee.id == id {
                        // 将要删除结点的下一个结点链接到该结点的上一个结点
                        node.next = node.next.as_mut().unwrap().next.take();
                        break;
                    }
                    node
                }
                None => {
                    break;
                }
            };
            temp_node = next;
        }
    }

    /// 通过 id 查找员工
    fn get_employee_by_id(&mut self, id: u32) -> Option<&mut Box<Employee>> {
        // 计算下标
        let index = id as usize % self.link_array.len();
        let head_node = &mut self.link_array[index];
        // 向量没有元素
        if head_node.is_none() {
            return None;
        }
        // 查找链表中是否存在相同 id 的员工
        let mut temp_node = head_node.as_mut().unwrap();
        if temp_node.employee.id == id {
            return Some(&mut temp_node.employee);
        }
        loop {
            let next = match temp_node.next.as_mut() {
                Some(node) => {
                    if node.employee.id == id {
                        return Some(&mut node.employee);
                    }
                    node
                }
                None => break,
            };
            temp_node = next;
        }
        return None;
    }

    /// 显示哈希表内容的方法
    pub fn list(&self) {
        let link_array = &self.link_array;
        for i in 0..link_array.len() {
            let elem = &link_array[i];
            let mut str = String::from("");
            if elem.is_some() {
                str.push_str("[");
                let mut temp_node = elem.as_ref().unwrap();
                str.push_str(&format!(
                    "{{id={}, name={}}}",
                    temp_node.employee.id, temp_node.employee.name
                ));
                loop {
                    let next = match temp_node.next.as_ref() {
                        Some(node) => {
                            str.push_str(&format!(
                                "{{id={}, name={}}}",
                                node.employee.id, node.employee.name
                            ));
                            node
                        }
                        None => {
                            break;
                        }
                    };
                    temp_node = next;
                }
                str.push_str("]");
            }
            println!("linkArray[{}]={}", i, str)
        }
    }
}

#[test]
fn test_add_employee() {
    // 创建一个哈希表
    let mut hashtable = Hashtable::new(5);
    // 创建员工
    let employee1 = Box::new(Employee::new(1, String::from("张三")));
    let employee2 = Box::new(Employee::new(2, String::from("李四")));
    let employee3 = Box::new(Employee::new(5, String::from("孙七")));
    // 添加员工
    let _ = hashtable.add(employee1);
    let _ = hashtable.add(employee2);
    let _ = hashtable.add(employee3);

    println!("添加员工后:");
    // 显示哈希表内容
    hashtable.list();
}

#[test]
fn test_update_employee() {
    let mut hashtable = Hashtable::new(5);
    let employee1 = Box::new(Employee::new(1, String::from("张三")));
    let employee2 = Box::new(Employee::new(2, String::from("李四")));
    let employee3 = Box::new(Employee::new(6, String::from("周八")));
    // 添加员工
    let _ = hashtable.add(employee1);
    let _ = hashtable.add(employee2);
    let _ = hashtable.add(employee3);

    println!("添加员工后:");
    // 显示哈希表内容
    hashtable.list();

    // 修改员工
    let employee = Box::new(Employee::new(6, String::from("菜菜")));
    hashtable.update(employee);

    println!("修改员工后:");
    hashtable.list();
}

#[test]
fn test_delete_employee() {
    let mut hashtable = Hashtable::new(5);
    let employee1 = Box::new(Employee::new(2, String::from("李四")));
    let employee2 = Box::new(Employee::new(5, String::from("孙七")));
    let _ = hashtable.add(employee1);
    let _ = hashtable.add(employee2);

    println!("删除员工前:");
    hashtable.list();

    hashtable.delete(2);

    println!("删除员工后:");
    hashtable.list();
}
