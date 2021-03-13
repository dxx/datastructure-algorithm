package com.dxx.stack;

import java.util.HashMap;
import java.util.Map;
import java.util.regex.Pattern;

/**
 * 综合计算器
 */
public class Calculator {

    private final Stack numStack; // 数栈

    private final Stack operationStack; // 符号栈

    public static Map<String, Operation> operations = new HashMap<>();

    static {
        // 定义相关操作符对应的优先级和计算方法
        operations.put("+", new Operation("+", 1, (num1, num2) -> num1 + num2));
        operations.put("-", new Operation("-", 1, (num1, num2) -> num1 - num2));
        operations.put("*", new Operation("*", 2, (num1, num2) -> num1 * num2));
        operations.put("/", new Operation("/", 2, (num1, num2) -> num1 / num2));
    }

    public interface Cal {
        int cal(int num1, int num2);
    }

    public static class Operation {
        private final String opt;
        private final int priority;
        private final Cal cal;

        public Operation(String opt, int priority, Cal cal) {
            this.opt = opt;
            this.priority = priority;
            this.cal = cal;
        }

        public String getOpt() {
            return opt;
        }

        public int getPriority() {
            return priority;
        }

        public Cal getCal() {
            return cal;
        }
    }

    public Calculator() {
        this.numStack = new Stack(10);
        this.operationStack = new Stack(10);
    }

    /**
     * 判断是否是操作符号
     */
    private boolean isOperation(String opt) {
        return Calculator.operations.get(opt) != null;
    }

    /**
     * 判断是否为数字
     */
    private boolean isNum(String c) {
        return Pattern.matches("\\d+", c);
    }

    /**
     * 计算操作符的优先级
     */
    private int priority(String opt1, String opt2) {
        Operation operation1 = Calculator.operations.get(opt1);
        Operation operation2 = Calculator.operations.get(opt2);
        if (operation1 == null || operation2 == null) {
            throw new IllegalArgumentException(String.format("请检查运算符: %s, %s\\n", opt1, opt2));
        }
        return operation1.getPriority() - operation2.getPriority();
    }

    /**
     * 计算结果
     */
    private int calculateNum(int num1, int num2, String opt) {
        Operation operation = Calculator.operations.get(opt);
        if (operation != null) {
            if (operation.getOpt().equals("-") || operation.getOpt().equals("/")) {
                // 因为出栈后两数的位置颠倒，需交换两个数的位置
                num1 = num1 + num2;
                num2 = num1 - num2;
                num1 = num1 - num2;
            }
            return operation.getCal().cal(num1, num2);
        }
        return 0;
    }

    private int calculateNumFromStack() {
        // 从数栈中弹出两个数，从符号栈中弹出一个符号
        String numStr1 = this.numStack.pop();
        String numStr2 = this.numStack.pop();
        String opt = this.operationStack.pop();
        // 计算值
        return this.calculateNum(Integer.parseInt(numStr1), Integer.parseInt(numStr2), opt);
    }

    /**
     * 计算表达式的值
     */
    public void calculate(String expression) {
        if (expression == null || expression.equals("")) {
            return;
        }
        int index = 0;
        String number = "";
        while (index < expression.length()) {
            String c = expression.substring(index, index + 1);
            // 判断是否为符号
            if (this.isOperation(c)) {
                // 判断符号栈是否为空
                if (this.operationStack.isEmpty()) {
                    // 压入符号栈
                    this.operationStack.push(c);
                } else {
                    // 符号栈不为空，判断优先级
                    String opt = this.operationStack.peek();
                    // c 优先级小于等于 elem
                    if (this.priority(c, opt) <= 0) {
                        // 计算值
                        int result = this.calculateNumFromStack();
                        // 将计算结果入数栈
                        this.numStack.push(String.valueOf(result));
                    }
                    // 将当前操作符入符号栈
                    this.operationStack.push(c);
                }
            } else if (this.isNum(c)) {
                // 向后面再取一位判断是否为数字
                if (index + 1 < expression.length() && this.isNum(expression.substring(index+1, index+2))) {
                    number += c;
                    index++;
                    continue;
                }
                // 压入数栈
                this.numStack.push(number + c);
                number = "";
            } else {
                throw new IllegalStateException("无法识别的字符:" + c);
            }

            index++;
        }

        // 全部数和符号都压入对应的栈后，取出计算
        // 符号栈不为空，循环
        while (!this.operationStack.isEmpty()) {
            // 计算值
            int result = this.calculateNumFromStack();
            // 将计算结果入数栈
            this.numStack.push(String.valueOf(result));
        }
        String result = this.numStack.pop();
        System.out.printf("表达式执行结果: %s=%s\n", expression, result);
    }

    public static void main(String[] args) {
        Calculator calculator = new Calculator();

        calculator.calculate("3+5*3-6");
        calculator.calculate("30+5*3-6");
        calculator.calculate("130+5*3-6");
    }
}
