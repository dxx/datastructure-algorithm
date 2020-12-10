/// 栈
/// 栈和队列一样也是一种特殊的线性表。它只能在表尾进行插入和删除操作。
/// 在进行插入和删除操作的一端被称为栈顶，另一端称为栈底。向一个栈放入
/// 新元素称为进栈、入栈或压栈，从一个栈取出元素称为出栈或退栈。每一个
/// 新元素都会放在之前放入的元素之上，删除时会删除最新的元素，所以栈有
/// 先进后出（FILO—first in last out）的特点。

pub struct Stack {
    array: Vec<String>, // 存放栈元素的向量
    max_size: usize, // 最大栈元素大小
    top: usize // 栈顶
}

impl Stack {
    pub fn new(size: usize) -> Self {
        return Stack {
            array: vec![String::from(""); size + 1],
            max_size: size + 1, // 0 下标不存数据，实际容量要多出一个
            top: 0 // 初始化为 0
        }
    }

    /// 入栈
    pub fn push(&mut self, elem: String) -> Result<String, String> {
        // 判断栈是否已满
        if self.top == self.max_size - 1 {
            return Err(String::from("stack is full"));
        }
        self.top = self.top + 1; // 栈顶加 1
        self.array[self.top] = elem;
        return Ok(self.array[self.top].clone());
    }

    /// 出栈
    pub fn pop(&mut self) -> Option<String> {
        if self.is_empty() {
            return None;
        }
        let elem = self.array[self.top].clone();
        self.top = self.top - 1; // 栈顶减 1
        return Some(elem);
    }

    /// 判断栈是否为空
    pub fn is_empty(&self) -> bool {
        return self.top == 0;
    }

    /// 窥视栈顶元素
    pub fn peek(&self) -> Option<String> {
        if self.is_empty() {
            return None;
        }
        return Some(self.array[self.top].clone());
    }

    pub fn show_stack(&self) {
        let mut str = String::from("[");
        let max = self.top + 1;
        for i in 1..max {
            str.push_str(&self.array[max - i].to_string());
            str.push_str(" ");
        }
        str.push_str("]");
        println!("{}", str);
    }
}
