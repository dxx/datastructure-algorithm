/**
 * 栈
 * 栈是一种运算受限的线性表
 * 限定仅在表尾进行插入和删除操作的线性表。这一端被称为栈顶，相对地，把另一端称为栈底
 * 向一个栈插入新元素又称作进栈、入栈或压栈，它是把新元素放到栈顶元素的上面，使之成为新的栈顶元素；
 * 从一个栈删除元素又称作出栈或退栈，它是把栈顶元素删除掉，使其相邻的元素成为新的栈顶元素
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
