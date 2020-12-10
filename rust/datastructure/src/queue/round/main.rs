/// 队列
/// 队列是一种特殊的线性表，它只允许在线性表的前端进行删除操作，在表的后端进行插入操作
/// 所以队列又称为先进先出（FIFO—first in first out）

/// 循环队列
/// 在实际使用队列时，顺序队列空间不能重复使用，需要对顺序队列进行改进。
/// 不管是插入或删除，一旦 rear 指针增 1 或 front 指针增 1 时超出了队列所分配的空间，就
/// 让它指向起始位置。当 MaxSize - 1增 1变到 0，可用取余运算获取队头或者队尾指针增加 1 后
/// 的位置，队尾指针计算方法为 rear % MaxSize，队尾指针计算方法为 front % MaxSize。
/// 这种循环使用队列空间的队列称为循环队列。

/// 数组实现循环队列

struct IntQueue {
    array: Vec<i32>, // 存放队列元素的数组
    max_size: usize, // 最大队列元素大小
    front: usize, // 队头指针
    rear: usize // 队尾指针
}

impl IntQueue {

    pub fn new(size: usize) -> Self {
        return IntQueue {
            array: vec![0; size],
            max_size: size,
            front: 0,
            rear: 0
        };
    }

    pub fn put(&mut self, elem: i32) -> Result<i32, String> {
        if self.is_full() {
            return Err(String::from("queue is full"));
        }
        self.array[self.rear] = elem;
        // 循环累加，当 rear + 1 等于 maxSize 时变成 0，重新累加
        self.rear = (self.rear + 1) % self.max_size;
        return Ok(elem);
    }

    pub fn take(&mut self) -> Option<i32> {
        if self.is_empty() {
            return None;
        }
        let elem = self.array[self.front];
        self.front = (self.front + 1) % self.max_size;
        return Some(elem);
    }

    fn is_empty(&self) -> bool {
        // 队头指针等于队尾指针表示队列为空
        return self.front == self.rear;
    }

    fn is_full(&self) -> bool {
        // 空出一个位置，判断是否等于队头指针
        // 队尾指针指向的位置不能存放队列元素，实际上会比 max_size 指定的大小少一
        return (self.rear + 1) % self.max_size == self.front;
    }

    fn size(&self) -> usize {
        return (self.rear + self.max_size - self.front) % self.max_size;
    }

    fn show_queue(&self) {
        let mut str = String::from("[");
        let mut temp_front = self.front;
        for _ in 0..self.size() {
            str.push_str(&self.array[temp_front].to_string());
            str.push_str(" ");
            // 超过最大大小，从 0 开始
            temp_front = (temp_front + 1) % self.max_size;
        }
        str.push_str("]");
        println!("{}", str);
    }
}

fn main() {
    let mut int_queue = IntQueue::new(5);
    let _ = int_queue.put(1);
    let _ = int_queue.put(2);
    let _ = int_queue.put(3);
    let _ = int_queue.put(4);
    let _ = int_queue.put(5); // 队列已满，无法放入数据，实际上只能放 4 个元素

    int_queue.show_queue();

    let num = int_queue.take().unwrap();
    println!("取出一个元素:{}", num);
    let num = int_queue.take().unwrap();
    println!("取出一个元素:{}", num);
    let num = int_queue.take().unwrap();
    println!("取出一个元素:{}", num);
    let num = int_queue.take().unwrap();
    println!("取出一个元素:{}", num);
    let o = int_queue.take();
    match o {
        Some(data) => {
            println!("{}", data);
        },
        None => {
            println!("出队失败: queue is empty");
        }
    }

    // 取出数据后可以继续放入数据
    let _ = int_queue.put(5);
    int_queue.show_queue();
}
