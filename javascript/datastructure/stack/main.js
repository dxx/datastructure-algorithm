const Stack = require("./stack");

function main() {
  // 创建一个栈
  let stack = new Stack(3);
  // 入栈
  stack.push("one");
  stack.push("two");
  stack.push("three");

  // 栈满，无法入栈
  let isSuccess = stack.push("four");
  if (!isSuccess) {
    console.log("入栈失败!!!");
  }
  stack.show();

  let elem1 = stack.pop();
  let elem2 = stack.pop();
  let elem3 = stack.pop();

  console.log("出栈:" + elem1);
  console.log("出栈:" + elem2);
  console.log("出栈:" + elem3);

  let elem = stack.pop();
  if (elem == null) {
    console.log("出栈失败!!!");
  }
  stack.show();
}

main();
