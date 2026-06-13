/**
 * 循环链表
 * 循环链表的特点是表中最后一个结点的指针域指向头结点，整个链表形成一个环
 *
 * 双向循环链表
 */

function PersonNode(no, name) {
  this.no = no;
  this.name = name;
  this.prev = null;
  this.next = null;
}

/**
 * 插入结点
 */
function insertNode(headNode, newNode) {
  // 判断是否第一次插入
  if (headNode.next === null) {
    headNode.no = newNode.no;
    headNode.name = newNode.name;
    headNode.prev = headNode;
    headNode.next = headNode;
    return;
  }
  let lastNode = headNode;
  // 下一个结点不等于头结点继续循环
  while (lastNode.next !== headNode) {
    lastNode = lastNode.next;
  }
  // 将新结点添加到链表末尾
  lastNode.next = newNode;
  newNode.prev = lastNode;
  // 将新结点下一个结点指针指向头结点
  newNode.next = headNode;
  headNode.prev = newNode;
}

/**
 * 删除指定结点，返回头结点
 */
function deleteNode(headNode, node) {
  // 没有结点 或者 只有一个头结点
  if (headNode.next === null || headNode.next === headNode) {
    // 头结点就是要删除的结点
    if (headNode.no === node.no) {
      headNode.prev = null;
      headNode.next = null;
    }
    return headNode;
  }

  let tempNode = headNode.next;
  let isExist = false;
  while (true) {
    if (tempNode === headNode) { // 最后一个结点
      if (tempNode.no === node.no) {
        isExist = true;
        // 头结点删除了，将头结点的下一个结点作为头结点
        headNode = tempNode.next;
      }
      break;
    } else if (tempNode.no === node.no) {
      isExist = true;
      break;
    }
    tempNode = tempNode.next;
  }
  // 存在需要删除的结点
  if (isExist) {
    // 将查找到的结点的上一个结点的下一个结点指针指向当前结点的下一个结点
    tempNode.prev.next = tempNode.next;
    // 将查找到的结点的下一个结点的上一个结点指针指向当前指针的上一个结点
    tempNode.next.prev = tempNode.prev;
  }
  return headNode;
}

/**
 * 打印循环链表的信息
 */
function printRoundNodeInfo(headNode) {
  if (!headNode.next) {
    console.log("该链表没有节点");
    return;
  }
  let str = "[";
  let tempNode = headNode;
  while (true) {
    str += "{no:" + tempNode.no + ", name:" + tempNode.name + "}";
    // 表示最后一个结点
    if (tempNode.next === headNode) {
      break;
    }
    tempNode = tempNode.next;
  }
  str += "]";
  console.log(str);
}

function testInsertNode() {
  // 创建 head 结点，head 结点不初始化数据，等到添加了第一个结点后才初始化数据
  let headNode = new PersonNode();
  // 创建第一个结点
  let personNode1 = new PersonNode(1, "张三");
  // 创建第二个结点
  let personNode2 = new PersonNode(2, "李四");
  // 创建第三个结点
  let personNode3 = new PersonNode(3, "王五");

  // 插入结点
  insertNode(headNode, personNode1);
  insertNode(headNode, personNode2);
  insertNode(headNode, personNode3);

  printRoundNodeInfo(headNode);
}

function testDeleteNode() {
  // 创建结点
  let headNode = new PersonNode();
  let personNode1 = new PersonNode(1, "张三");
  let personNode2 = new PersonNode(2, "李四");
  let personNode3 = new PersonNode(3, "王五");
  let personNode4 = new PersonNode(4, "赵六");
  let personNode5 = new PersonNode(5, "孙七");

  // 插入结点
  insertNode(headNode, personNode1);
  insertNode(headNode, personNode2);
  insertNode(headNode, personNode3);
  insertNode(headNode, personNode4);
  insertNode(headNode, personNode5);

  console.log("删除前:");
  printRoundNodeInfo(headNode);

  // 删除 no 为 2 的结点
  headNode = deleteNode(headNode, personNode2);
  console.log("删除 no 为 2 的结点后:");
  printRoundNodeInfo(headNode);

  let newNode = new PersonNode(6, "周八");

  insertNode(headNode, newNode);
  console.log("插入新结点:");
  printRoundNodeInfo(headNode);

  // 删除 no 为 1，3 的结点
  headNode = deleteNode(headNode, personNode1);
  headNode = deleteNode(headNode, personNode3);
  console.log("删除 no 为 1,3 的结点后:");
  printRoundNodeInfo(headNode);
}

function main() {
  // testInsertNode();
  // testDeleteNode();
}

main();
