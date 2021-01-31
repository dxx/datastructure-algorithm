const Stack = require("./stack");

/**
 * 逆波兰计算器
 */

function Operation(opt, priority, optFun) {
  this.opt = opt;
  this.priority = priority;
  this.optFun = optFun;
}

// 定义相关操作符对应的优先级和计算方法
const operations = {
  "+": new Operation("+", 1, (num1, num2) => num1 + num2),
  "-": new Operation("-", 1, (num1, num2) => num1 - num2),
  "*": new Operation("*", 2, (num1, num2) => num1 * num2),
  "/": new Operation("/", 2, (num1, num2) => num1 / num2)
}

function ReversePoland() {
}

/**
 * 判断是否为数字
 */
ReversePoland.prototype._isNum = function(str) {
  return /\d+/g.test(str);
}

/**
 * 计算操作符的优先级
 */
ReversePoland.prototype._priority = function(opt1, opt2) {
  let operation1 = operations[opt1];
  let operation2 = operations[opt2];
  if (!operation1 || !operation2) {
    console.error("请检查运算符: " + opt1 + "," + opt2);
    return;
  }
  return operation1.priority - operation2.priority;
}

/**
 * 计算结果
 */
ReversePoland.prototype._calculateNum = function(num1, num2, opt) {
  let operation = operations[opt];
  if (operation) {
    if (operation.opt === "-" || operation.opt === "/") {
      // 因为出栈后两数的位置颠倒，需交换两个数的位置
      num1 = num1 + num2;
      num2 = num1 - num2;
      num1 = num1 - num2;
    }
    return operation.optFun(num1, num2);
  }
  return 0;
}

/**
 * 后缀表达式计算
 * 1. 循环读取每个字符，判断是否是数字
 * 2. 如果是数字直接入栈
 * 3. 如果是运算符，从栈中弹出两个数，计算表达式的值，将结果压入栈中
 */
ReversePoland.prototype.calSuffixExpression = function(expr) {
  let stack = new Stack(expr.length);
  for (let i = 0; i < expr.length; i++) {
    let str = expr[i];
    if (this._isNum(str)) {
      stack.push(str);
      continue;
    }
    let operation = operations[str];
    if (!operation) {
      console.error("无效的运算符: " + str);
      return;
    }
    // 计算
    let numStr1 = stack.pop();
    let numStr2 = stack.pop();
    let result = this._calculateNum(Number(numStr1), Number(numStr2), str);
    stack.push(result);
  }
  // 弹出最后结果
  return Number(stack.pop());
}

/**
 * 将表达式转换成数组
 */
ReversePoland.prototype.exprToArray = function(expr) {
  if (!expr) {
    return [];
  }
  let expressions = [];
  for (let i = 0; i < expr.length; i++) {
    let s = expr[i];
    if (this._isNum(s)) {
      // 向后面继续判断是否为数字
      while(i + 1 < expr.length && this._isNum(expr.substring(i + 1, i + 2))) {
        s += expr.substring(i + 1, i + 2);
        i++;
      }
    }
    expressions.push(s);
  }
  return expressions;
}

/**
 * 中缀表达式转后缀表达式
 */
ReversePoland.prototype.infixToSuffix = function(infix) {
  if (!infix) {
    return;
  }
  // 初始化两个栈，一个运算符栈 stack1 和另一个储存中间结果的栈 stack2
  let stack = new Stack(infix.length);
  // 由于中间结果栈不需要弹出元素，可以使用集合来保存
  let suffixes = [];
  // 循环表达式
  for (let i = 0; i < infix.length; i++) {
    let str = infix[i];
    // 遇到数字时，将其放入 suffixes
    if (this._isNum(str)) {
      suffixes.push(str);
      continue;
    }
    // 如果是 ( 直接入栈
    if (str === "(") {
      stack.push(str);
      continue;
    }
    if (str === ")") {
      while(!stack.isEmpty() && stack.peek() !== "(") {
        // 弹出 stack 中栈顶的元素，并添加到 suffixes
        suffixes.push(stack.pop());
      }
      // 弹出 (，消除一对 ( )
      stack.pop();
      continue;
    }
    let operation = operations[str];
    if (operation) {
      if (stack.isEmpty() || stack.peek() === "(") {
        // 如果 stack 为空或栈顶运算符为左括号 "("，则直接将此运算符入栈
        stack.push(str);
        continue;
      }
      // 栈不为空，并且当前字符串的优先级小于等于栈顶的元素
      while (!stack.isEmpty() && this._priority(str, stack.peek()) <= 0) {
        // 将栈顶的元素添加到 suffixes
        suffixes.push(stack.pop());
      }
      // 入栈
      stack.push(str);
    } else {
      console.error("无法识别的字符: " + str);
      return;
    }
  }
  // 将 stack 中剩余的运算符依次添加到 suffixes
  while(!stack.isEmpty()) {
    suffixes.push(stack.pop());
  }
  // 因为这里用的是集合，它里面元素的顺序就是栈元素出栈后逆序排列的顺序
  return suffixes;
}

function main() {
  let reversePoland = new ReversePoland();
  let expr = "3 5 3 * + 2 -";
  // 假设数和数或符号之间有空格
  let expressions = expr.split(" ");
  let result = reversePoland.calSuffixExpression(expressions);
  console.log("后缀表达式 " + expr + " 的计算结果为: " + result);

  expr = "1+((2+3)*4)-5";
  expressions = reversePoland.exprToArray(expr);
  console.log("将中缀表达式放入数组, 结果为: " + expressions);

  expressions = reversePoland.infixToSuffix(expressions);
  console.log("中缀表达式转换成后缀表达式, 结果为: " + expressions);

  result = reversePoland.calSuffixExpression(expressions);
  console.log("计算表达式 " + expr + ", 结果为: " + result);
}

main();
