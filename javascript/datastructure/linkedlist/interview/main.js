const { Node, getNodeLength, getLastIndexNode,reverseNode } = require("./single_node");

function testGetLength() {
  let headNode = new Node();
  let node1 = new Node("node1");
  let node2 = new Node("node2");
  let node3 = new Node("node3");

  headNode.next = node1;
  node1.next = node2;
  node2.next = node3;

  let length = getNodeLength(headNode);
  console.log("单链表结点个数为: " + length);
}

function testGetLastIndexNode() {
  let headNode = new Node();
  let node1 = new Node("node1");
  let node2 = new Node("node2");
  let node3 = new Node("node3");
  headNode.next = node1;
  node1.next = node2;
  node2.next = node3;
  let index = 2;
  let lastNode = getLastIndexNode(headNode, index);
  console.log("单链表结点中倒数第" + index + "个结点为: " + lastNode.name);
}

function testReverseNode() {
  let headNode = new Node();
  let node1 = new Node("node1");
  let node2 = new Node("node2");
  let node3 = new Node("node3");
  headNode.next = node1;
  node1.next = node2;
  node2.next = node3;

  console.log("反转前:");
  headNode.printNodeInfo();

  reverseNode(headNode);

  console.log("反转后:");
  headNode.printNodeInfo();
}

function main() {
  // testGetLength();
  // testGetLastIndexNode();
  // testReverseNode();
}

main();
