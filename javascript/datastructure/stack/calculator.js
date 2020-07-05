const Stack = require("./stack");

/**
 * 综合计算器
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

function Calculator() {
  this.numStack = new Stack(10);
  this.operationStack = new Stack(10);
}

/**
 * 判断是否是操作符号
 */
Calculator.prototype._isOperation = function(opt) {
  return operations[opt] !== undefined
}

/**
 * 判断是否为数字
 */
Calculator.prototype._isNum = function(str) {
  return /\d+/g.test(str);
}

/**
 * 计算操作符的优先级
 */
Calculator.prototype._priority = function(opt1, opt2) {
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
Calculator.prototype._calculateNum = function(num1, num2, opt) {
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

Calculator.prototype._calculateNumFromStack = function() {
  // 从数栈中弹出两个数，从符号栈中弹出一个符号
  let numStr1 = this.numStack.pop();
  let numStr2 = this.numStack.pop();
  let opt = this.operationStack.pop();
  // 计算值
  return this._calculateNum(Number(numStr1), Number(numStr2), opt);
}

/**
 * 计算表达式的值
 */
Calculator.prototype.calculate = function(expression) {
  if (!expression) {
    return;
  }
  let index = 0;
  let number = "";
  while (index < expression.length) {
    let char = expression.substring(index, index + 1);
    // 判断是否为符号
    if (this._isOperation(char)) {
      // 判断符号栈是否为空
      if (this.operationStack.isEmpty()) {
        // 压入符号栈
        this.operationStack.push(char);
      } else {
        // 符号栈不为空，判断优先级
        let opt = this.operationStack.peek();
        // char 优先级小于等于 elem
        if (this._priority(char, opt) <= 0) {
          // 计算值
          let result = this._calculateNumFromStack();
          // 将计算结果入数栈
          this.numStack.push(result);
        }
        // 将当前操作符入符号栈
        this.operationStack.push(char);
      }
    } else if (this._isNum(char)) {
      // 向后面再取一位判断是否为数字
      if (index + 1 < expression.length && this._isNum(expression.substring(index + 1, index + 2))) {
        number += char;
        index++;
        continue;
      }
      this.numStack.push(number + char);
      number = "";
    } else {
      console.error("无法识别的字符:" + char);
      return;
    }
    index++;
  }

  // 全部数和符号都压入对应的栈后，取出计算
  // 符号栈不为空，循环
  while (!this.operationStack.isEmpty()) {
    // 计算值
    let result = this._calculateNumFromStack();
    // 将计算结果入数栈
    this.numStack.push(result);
  }
  let result = this.numStack.pop();
  console.log("表达式执行结果:" + expression + "=" +result);
}

function main() {
  let calculator = new Calculator();

  calculator.calculate("3+5*3-6");
  calculator.calculate("30+5*3-6");
  calculator.calculate("130+5*3-6");
}

main();
