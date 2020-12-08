/// 队列
/// 队列是一种特殊的线性表，它只允许在线性表的前端进行删除操作，在表的后端进行插入操作
/// 所以队列又称为先进先出（FIFO—first in first out）

/// 顺序队列
/// 顺序队列类似数组，它需要一块连续的内存，并有两个指针，一个是队头指针 front，它指向队头元素
/// 另一个是队尾指针 rear，它指向下一个入队的位置。
/// 不断的进行插入和删除操作，队列元素在不断的变化，当队头指针等于队尾指针时，队列中没有任何元
/// 素，没有元素的队列称为空队列。对于已经出队列的元素所占用的空间，顺序队列无法再次利用。

/// 数组实现顺序队列

struct IntQueue {
    array: Vec<i32>, // 存放队列元素的数组
    max_size: usize, // 最大队列元素大小
    front: usize, // 队头指针
    rear: usize // 队尾指针
}

impl IntQueue {
    pub fn new(size: usize) -> Self {
        return IntQueue{
            array: vec![0; size],
            max_size: size,
            front: 0,
            rear: 0
        };
    }
    pub fn put(&mut self, elem: i32) -> Result<i32, String> {
        // 队尾指针不能超过最大队列元素大小
        if self.rear >= self.max_size {
            return Err(String::from("queue is full"));
        }
        self.array[self.rear] = elem;
        self.rear = self.rear + 1; // 队尾指针加一
        return Ok(elem);
    }
    pub fn take(&mut self) -> Option<i32> {
        // 队头指针等于队尾指针表示队列为空
        if self.front == self.rear {
            return None;
        }
        let elem = self.array[self.front];
        self.front = self.front + 1; // 队头指针加一
        return Some(elem);
    }
    pub fn show_queue(&self) {
        let mut str = String::from("[");
        for i in self.front..self.rear {
            str.push_str(&self.array[i].to_string());
            str.push_str(" ");
        }
        str.push_str("]");
        println!("{}", str);
    }
}

fn main() {
    let mut int_queue = IntQueue::new(3);
    let _ = int_queue.put(1);
    let _ = int_queue.put(2);
    let _ = int_queue.put(3);
    let _ = int_queue.put(4); // 队列已满，无法放入数据

    int_queue.show_queue();

    let num = int_queue.take().unwrap();
    println!("取出一个元素:{}", num);
    let num = int_queue.take().unwrap();
    println!("取出一个元素:{}", num);
    let num = int_queue.take().unwrap();
    println!("取出一个元素:{}", num);
    let r = int_queue.take();
    match r {
        Some(data) => {
            println!("{}", data);
        },
        None  => {
            println!("出队失败: queue is empty");
        }
    }

    // 此时队列已经用完，无法放数据
    let res = int_queue.put(4); // 队列已满，无法放入数据
    match res {
        Ok(data) => {
            println!("{}", data);
        },
        Err(e)  => {
            println!("入队失败：{}", e);
        }
    }
    int_queue.show_queue();
}
