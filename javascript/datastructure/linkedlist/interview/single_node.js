function Node(name) {
  this.name = name;
  this.next = null;
}

Node.prototype.printNodeInfo = function() {
  if (!this.next) {
    console.log("该链表没有节点");
    return;
  }
  let str = "[";
  let tempNode = this.next;
  while (tempNode !== null) {
    str += "{name:" + tempNode.name + "}";
    tempNode = tempNode.next;
  }
  str += "]";
  console.log(str);
}

/**
 * 获取单链表有效的结点数
 * 1.遍历结点数
 * 2.定义一个长度遍历，每遍历一次长度 +1
 */
function getNodeLength(headNode) {
  // 头结点为空，返回 0
  if (!headNode) {
    return 0;
  }
  let length = 0;
  let node = headNode.next;
  while (node !== null) {
    length++;
    node = node.next;
  }
  return length;
}

/**
 * 获取倒数第 n 个结点
 *
 * 快慢指针
 * 1.定义快指针 fast 和 慢指针 slow
 * 2.快慢指针的初始值指向头结点
 * 3.快指针先走 index 步
 * 4.慢指针开始走直到快指针指向了末尾结点
 * 5.此时慢指针就是倒数第 n 个结点
 */
function getLastIndexNode(headNode, index) {
  // 头结点为空，index 小于等于 0 返回空
  if (headNode == null || index <= 0) {
      return null;
  }
  let fast = headNode;
  let slow = headNode;
  while (fast != null) {
      if (index > 0) {
          fast = fast.next;
          index--;
          continue;
      }
      // 快慢指针同时走
      fast = fast.next;
      slow = slow.next;
  }
  // index 超过了链表的长度
  if (slow == headNode) {
      return null;
  }
  return slow;
}

/**
 * 获取倒数第 n 个结点
 * 
 * 遍历
 * 1.获取链表结点数 length
 * 2.遍历到 length - n 个结点
 * 3.然后返回
 */
function getLastIndexNode2(headNode, index) {
  // 头结点为空，返回空
  if (!headNode) {
    return null;
  }
  let length = getNodeLength(headNode);
  if (index <= 0 || index > length) {
    return null;
  }
  let lastNode = headNode.next;
  for (let i = 0; i < length - index; i++) {
    lastNode = lastNode.next;
  }
  return lastNode;
}

/**
 * 单链表反转
 * 1.定义一个新的头结点 reverseHead
 * 2.遍历链表，每遍历一个结点，将其取出，放在新的头结点 reverseHead 的后面
 * 3.最后将头结点的 next 结点指向 reverseHead 的 next 结点
 */
function reverseNode(headNode) {
  if (headNode == null || headNode.next == null) {
    return;
  }
  let reverseHead = new Node();
  let current = headNode.next;
  let next;
  while (current != null) {
    // 保存当前结点的下一个结点
    next = current.next
    // 将 reverseHead 结点的下一个结点放在当前结点的下一个结点
    current.next = reverseHead.next;
    // 当前结点放在 reverseHead 后面
    reverseHead.next = current;
    // 移动当前结点
    current = next;
  }
  // 将头结点的 next 结点指向 reverseHead 的 next 结点
  headNode.next = reverseHead.next;
}

module.exports = {
  Node,
  getNodeLength,
  getLastIndexNode,
  reverseNode
}
