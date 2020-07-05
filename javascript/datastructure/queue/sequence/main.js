/**
 * 队列
 * 队列是一种特殊的线性表，特殊之处在于它只允许在表的前端进行删除操作，而在表的后端进行插入操作
 * 所以队列又称为先进先出（FIFO—first in first out）
 *
 * 顺序队列
 * 顺序队列必须分配一块连续的内存，并设置两个指针进行管理，一个是队头指针 front，它指向队头元素
 * 另一个是队尾指针 rear，它指向下一个入队的位置。
 * 随着插入和删除的进行，队列元素在不断的变化，当队头指针等于队尾指针时，队列中没有任何元素，这时称为空队列
 * 顺序队列不能有效的利用已经出队的元素占用的空间
 *
 * 数组实现顺序队列
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
  // 队尾指针不能超过最大队列元素大小
  if (this.rear >= this.maxSize) {
    console.error("queue is full");
    return false;
  }
  // 把元素放入队尾，然后队尾指针加一
  this.array[this.rear++] = elem;
  return true;
}

/**
 * 取出队列元素
 */
IntQueue.prototype.take = function() {
  // 队头指针等于队尾指针表示队列为空
  if (this.front == this.rear) {
    console.error("queue is empty");
    return Number.MIN_VALUE;
  }
  // 取出当前队头指向的元素，然后队头指针加一
  return this.array[this.front++];
}

IntQueue.prototype.show = function() {
  let str = "[";
  for (let i = this.front; i < this.rear; i++) {
    str += this.array[i] + " ";
  }
  str += "]";
  console.log(str);
}

function main() {
  let intQueue = new IntQueue(3);
  intQueue.put(1);
  intQueue.put(2);
  intQueue.put(3);
  intQueue.put(4); // 队列已满，无法放入数据

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
  if (num == Number.MIN_VALUE) {
    console.log("出队失败!!!");
  }

  // 此时队列已经用完，无法放数据
  let isSuccess = intQueue.put(4);
  if (!isSuccess) {
    console.log("入队失败!!!");
  }
  console.log("intQueue:");
  intQueue.show();
}

main();
