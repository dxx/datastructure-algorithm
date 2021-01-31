/**
 * 约瑟夫问题
 * 设编号为 1,2,...n 的 n 个人围坐一圈约定编号为 k(1<=k<=n) 的人
 * 从 1 开始报数，数到 m 的那个人出列它的下一位又从 1 开始报数数到 m
 * 的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列
 */

function Person(no) {
  this.no = no;
  this.prev = null;
  this.next = null;
}

function PersonLinkedList(count) {
  if (count <= 0) {
    console.error("链表至少需要一个元素");
    return;
  }
  this.first = null;
  this.length = count;

  let prev = this.first;
  // 初始化小孩，构成双向循环链表
  for (let i = 1; i <= count; i++) {
    let person = new Person(i);
    if (i === 1) {
      // 初始化 First 节点
      this.first = person;
      this.first.next = this.first;
      this.first.prev = this.first;
      
      // 将 prev 指向 First 节点，继续下一次循环
      prev = this.first;
      continue;
    }

    prev.next = person;
    person.prev = prev;
  
    // 新增加的节点的下一个节点指向第一节点
    person.next = this.first;
    // 第一个节点的上一个节点指向新增加的节点
    this.first.prev = person;

    prev = person;
  }
}

PersonLinkedList.prototype.showPersons = function() {
  if (this.first === null) {
    return;
  }
  let current = this.first;
  while(true) {
    console.log("num:" + current.no);
    current = current.next;
    if (current === this.first) {
      break;
    }
  }
}

PersonLinkedList.prototype.count = function(start, num) {
  if (start < 1 || start > this.length) {
    console.log("start 不能小于 1 或者不能大于 " + this.length);
    return;
  }
  if (num > this.length) {
    console.log("num 不能大于元素个数: " + this.length);
    return;
  }

  let current = this.first;

  for (let i = 1; i <= start - 1; i++) {
    current = current.next;
  }

  while (true) {
    if (current.prev === current && current.next === current) {
      break;
    }

    for (let i = 1; i <= num - 1; i++) {
      current = current.next;
    }

    current.prev.next = current.next;
    current.next.prev = current.prev;

    console.log("出队人的编号: " + current.no);
    current = current.next;
  }
  console.log("最后留下人的编号: " + current.no);
}

function main() {
  let personLinkedList = new PersonLinkedList(5);
  personLinkedList.showPersons();

  personLinkedList.count(1, 3);
}

main();
