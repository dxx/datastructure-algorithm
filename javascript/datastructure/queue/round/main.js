/**
 * 队列
 * 队列是一种特殊的线性表，它只允许在线性表的前端进行删除操作，在表的后端进行插入操作
 * 所以队列又称为先进先出（FIFO—first in first out）
 *
 * 循环队列
 * 在实际使用队列时，顺序队列空间不能重复使用，需要对顺序队列进行改进。
 * 不管是插入或删除，一旦 rear 指针增 1 或 front 指针增 1 时超出了队列所分配的空间，就
 * 让它指向起始位置。当 MaxSize - 1增 1变到 0，可用取余运算获取队头或者队尾指针增加 1 后
 * 的位置，队尾指针计算方法为 rear % MaxSize，队尾指针计算方法为 front % MaxSize。
 * 这种循环使用队列空间的队列称为循环队列。
 *
 * 数组实现循环队列
 */

function IntQueue(size) {
  this.array = new Array(size); // 存放队列元素的数组
  this.maxSize = size; // 最大队列元素大小
  this.front = 0; // 队头指针
  this.rear = 0; // 队尾指针
}

/**
 * 放入队列元素
 */
IntQueue.prototype.put = function(elem) {
  if (this.isFull()) {
    console.error("queue is full");
    return false;
  }
  // 把元素放入队尾
  this.array[this.rear] = elem;
  // 循环累加，当 rear + 1 等于 maxSize 时变成 0，重新累加
  this.rear = (this.rear + 1) % this.maxSize;
  return true;
}

/**
 * 取出队列元素
 */
IntQueue.prototype.take = function() {
  if (this.isEmpty()) {
    console.error("queue is empty");
    return Number.MIN_VALUE;
  }
  // 取出当前队头指向的元素
  elem = this.array[this.front];
  this.front = (this.front + 1) % this.maxSize;
  return elem;
}

IntQueue.prototype.isEmpty = function() {
  // 队头指针等于队尾指针表示队列为空
  return this.front == this.rear;
}

IntQueue.prototype.isFull = function() {
  // 空出一个位置，判断是否等于队头指针
  // 队尾指针指向的位置不能存放队列元素，实际上会比 maxSize 指定的大小少一
  return (this.rear + 1) % this.maxSize == this.front;
}

IntQueue.prototype.size = function() {
  return (this.rear + this.maxSize - this.front) % this.maxSize;
}

IntQueue.prototype.show = function() {
  let str = "[";
  let tempFront = this.front;
  for (let i = 0; i < this.size(); i++) {
    str += this.array[tempFront] + " ";
    tempFront = (tempFront + 1) % this.maxSize;
  }
  str += "]";
  console.log(str);
}

function main() {
  let intQueue = new IntQueue(5);
  intQueue.put(1);
  intQueue.put(2);
  intQueue.put(3);
  intQueue.put(4);
  intQueue.put(5); // 队列已满，无法放入数据，实际上只能放 4 个元素

  console.log("intQueue:");
  intQueue.show();

  let num = intQueue.take();
  console.log("取出一个元素:" + num);
  num = intQueue.take();
  console.log("取出一个元素:" + num);
  num = intQueue.take();
  console.log("取出一个元素:" + num);
  num = intQueue.take();
  console.log("取出一个元素:" + num);
  num = intQueue.take();
  console.log("取出一个元素:" + num);
  if (num == Number.MIN_VALUE) {
    console.log("出队失败!!!");
  }

  // 取出数据后可以继续放入数据
  intQueue.put(5);
  console.log("intQueue:");
  intQueue.show();
}

main();
