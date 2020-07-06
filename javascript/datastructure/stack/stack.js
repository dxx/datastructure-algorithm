/**
 * 栈
 * 栈和队列一样也是一种特殊的线性表。它只能在表尾进行插入和删除操作。
 * 在进行插入和删除操作的一端被称为栈顶，另一端称为栈底。向一个栈放入
 * 新元素称为进栈、入栈或压栈，从一个栈取出元素称为出栈或退栈。每一个
 * 新元素都会放在之前放入的元素之上，删除时会删除最新的元素，所以栈有
 * 先进后出（FILO—first in last out）的特点。
 */

function Stack(size) {
  this.array = new Array(size); // 存放栈元素
  this.maxSize = size; // 最大栈元素大小
  this.top = -1; // 栈顶
}

/**
 * 入栈
 */
Stack.prototype.push = function(elem) {
  // 判栈是否已满
  if (this.top == this.maxSize - 1) {
    console.error("stack is full");
    return false;
  }
  // 栈顶加 1，将元素放入栈顶
  this.array[++this.top] = elem;
  return true;
}

/**
 * 出栈
 */
Stack.prototype.pop = function() {
  if (this.top == -1) {
    console.error("stack is empty");
    return "";
  }
  // 取出栈顶元素，然后加 1
  return this.array[this.top--];
}

/**
 * 判断栈是否为空
 */
Stack.prototype.isEmpty = function() {
  return this.top == -1;
}

/**
 * 窥视栈顶元素
 */
Stack.prototype.peek = function() {
  if (this.isEmpty()) {
    return undefined;
  }
  return this.array[this.top];
}

Stack.prototype.show = function() {
  let str = "[";
  for (let i = this.top; i >= 0; i --) {
    str += this.array[i] + " ";
  }
  str += "]";
  console.log(str);
}

module.exports = Stack;
