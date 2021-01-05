/// 哈希表
/// 哈希表也叫散列表，是根据关键值 key 直接进行访问的数据结构
/// 它通过把 key 映射到表中一个位置来记录，以加快查找的速度

/// 员工结构体
pub struct Employee {
    pub id: u32,
    pub name: String,
}

/// 单链表结构体
pub struct LinkNode {
    pub employee: Box<Employee>,
    pub next: Option<Box<LinkNode>>,
}

/// 哈希表结构体，包含一个 LinkNode 向量
pub struct Hashtable {
    pub link_array: Vec<LinkNode>,
}

impl Hashtable {
    /// 创建哈希表
    pub fn new() -> Self {
        return Hashtable {
            link_array: Vec::new(),
        };
    }
}

fn main() {}
