/**
 * 单向链表
 * 单向链表是链表的一种，其特点是链表的链接方向是单向的，对链表的访问要从头部开始
 * 链表是由结点构成，head 指针指向第一个称为表头的结点，而最后一个结点的指针指向 NULL
 */

function HeroNode(no, name, nickname) {
  this.no = no; // 编号
  this.name = name; // 姓名
  this.nickname = nickname; // 昵称
  this.next = null; // 下一个节点
}

/**
 * 在链表尾部插入，通过 head 找到链表的尾部
 */
function insertAtTail(headNode, newNode) {
  let lastNode = headNode;
  // 下一个结点不为空继续循环
  while (lastNode.next !== null) {
    // 将下一个结点赋值给当前结点
    lastNode = lastNode.next;
  }
  // 将当前结点插入到链表的最后一个结点
  lastNode.next = newNode;
}

/**
 * 按照 no 升序插入，通过 head 找到合适的插入位置
 */
function sortInsertByNo(headNode, newNode) {
  let tempNode = headNode;
  while (true) {
    if (tempNode.next === null) {
      break;
    } else if (tempNode.next.no > newNode.no) {
      break;
    } else if (tempNode.next.no == newNode.no) {
      console.error("no 相等不能插入");
      return;
    }
    tempNode = tempNode.next;
  }
  // tempNode 的下一个结点插入到 newNode 的下一个结点
  newNode.next = tempNode.next;
  // newNode 结点插入到 tempNode 的下一个结点
  tempNode.next = newNode;
}

/**
 * 删除指定结点
 */
function deleteNode(headNode, node) {
  let tempNode = headNode;
  while (tempNode.next != null) {
    if (tempNode.next.no == node.no) {
      // 将下一个结点的下一个结点，链接到被删除结点的上一个结点
      tempNode.next = tempNode.next.next;
      return;
    }
    tempNode = tempNode.next;
  }
}

/**
 * 打印单链表结点内容
 */
function printHeadNodeInfo(headNode) {
  if (!headNode.next) {
    console.log("该链表没有节点");
    return;
  }
  let str = "[";
  let tempNode = headNode.next;
  while (tempNode !== null) {
    str += "{no:" + tempNode.no + ", name:" + tempNode.name + ", nickname:" + tempNode.nickname + "}";
    tempNode = tempNode.next;
  }
  str += "]";
  console.log(str);
}

function testInsertAtTail() {
  // 创建 head 结点，head 结点不包含数据
  let headNode = new HeroNode();
  // 创建第一个结点
  let heroNode1 = new HeroNode(1, "宋江", "呼保义");
  // 创建第二个结点
  let heroNode2 = new HeroNode(2, "卢俊义", "玉麒麟");
  // 创建第三个结点
  let heroNode3 = new HeroNode(3, "吴用", "智多星");

  // 将结点添加到链表尾部
  insertAtTail(headNode, heroNode1);
  insertAtTail(headNode, heroNode2);
  insertAtTail(headNode, heroNode3);

  printHeadNodeInfo(headNode);
}

function testSortInsertByNo() {
  // 创建结点，用来做尾部插入
  let head = new HeroNode();
  let node1 = new HeroNode(1, "宋江", "呼保义");
  let node2 = new HeroNode(2, "卢俊义", "玉麒麟");
  let node3 = new HeroNode(3, "吴用", "智多星");

  insertAtTail(head, node1);
  insertAtTail(head, node3); // 将第三个结点插入到第二个位置
  insertAtTail(head, node2);

  console.log("尾部插入的结果:");
  printHeadNodeInfo(head);

  // 创建 head 结点
  let headNode = new HeroNode();
  // 创建第一个结点
  let heroNode1 = new HeroNode(1, "宋江", "呼保义");
  // 创建第二个结点
  let heroNode2 = new HeroNode(2, "卢俊义", "玉麒麟");
  // 创建第三个结点
  let heroNode3 = new HeroNode(3, "吴用", "智多星");

  // 将结点按照 no 升序插入
  sortInsertByNo(headNode, heroNode1);
  sortInsertByNo(headNode, heroNode3);
  sortInsertByNo(headNode, heroNode2);

  console.log("按照 no 升序插入的结果:");
  printHeadNodeInfo(headNode);
}

function testDeleteNode() {
  // 创建结点
  let headNode = new HeroNode();
  let heroNode1 = new HeroNode(1, "宋江", "呼保义");
  let heroNode2 = new HeroNode(2, "卢俊义", "玉麒麟");
  let heroNode3 = new HeroNode(3, "吴用", "智多星");
  let heroNode4 = new HeroNode(4, "公孙胜", "入云龙");
  let heroNode5 = new HeroNode(5, "关胜", "大刀");

  // 插入结点
  insertAtTail(headNode, heroNode1);
  insertAtTail(headNode, heroNode2);
  insertAtTail(headNode, heroNode3);
  insertAtTail(headNode, heroNode4);
  insertAtTail(headNode, heroNode5);

  console.log("删除前:");
  printHeadNodeInfo(headNode);

  // 删除 no 为 2 的结点
  deleteNode(headNode, heroNode2);
  console.log("删除 no 为 2 的结点后:");
  printHeadNodeInfo(headNode);

  // 删除 no 为 3, 4 的结点
  deleteNode(headNode, heroNode3);
  deleteNode(headNode, heroNode4);
  console.log("删除 no 为 3,4 的结点后:");
  printHeadNodeInfo(headNode);
}

function main() {
  // testInsertAtTail();
  // testSortInsertByNo();
  // testDeleteNode();
}

main();
