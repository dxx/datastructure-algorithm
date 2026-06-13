/**
 * 双向链表
 * 双向链表也叫双链表，是链表的一种，它的每个数据结点中都有两个指针，分别指向前一个和后一个结点
 * 所以，从双向链表中的任意一个结点开始，都可以很方便地访问它的前一个结点和后一个结点。
 */

function HeroNode(no, name, nickname) {
  this.no = no; // 编号
  this.name = name; // 姓名
  this.nickname = nickname; // 昵称
  this.prev = null; // 上一个节点
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
  // 将新结点的上一个结点指向当前结点
  newNode.prev = lastNode;
}

/**
 * 删除指定结点
 */
function deleteNode(headNode, node) {
  let tempNode = headNode.next;
  while (tempNode !== null) {
    if (tempNode.no === node.no) {
      // 将查找到的结点的上一个结点的下一个结点指针指向当前结点的下一个结点
      tempNode.prev.next = tempNode.next;
      // 最后一个结点的 next 指向空
      if (tempNode.next !== null) {
        // 将查找到的结点的下一个结点的上一个结点指针指向当前指针的上一个结点
        tempNode.next.prev = tempNode.prev;
      }
    }
    tempNode = tempNode.next;
  }
}

/**
 * 打印链表结点内容
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
  let heroNode1 = new HeroNode(3, "吴用", "智多星");
  // 创建第二个结点
  let heroNode2 = new HeroNode(6, "林冲", "豹子头");
  // 创建第三个结点
  let heroNode3 = new HeroNode(7, "秦明", "霹雳火");

  // 将结点添加到链表尾部
  insertAtTail(headNode, heroNode1);
  insertAtTail(headNode, heroNode2);
  insertAtTail(headNode, heroNode3);

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
  // testDeleteNode();
}

main();
