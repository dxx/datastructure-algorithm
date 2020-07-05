package com.mcx.stack;

import java.util.*;
import java.util.regex.Pattern;

/**
 * 逆波兰计算器
 */
public class ReversePoland {
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

    /**
     * 判断是否为数字
     */
    private boolean isNum(String s) {
        return Pattern.compile("\\d+").matcher(s).find();
    }

    /**
     * 计算操作符的优先级
     */
    private int priority(String opt1, String opt2) {
        Operation operation1 = ReversePoland.operations.get(opt1);
        Operation operation2 = ReversePoland.operations.get(opt2);
        if (operation1 == null || operation2 == null) {
            throw new IllegalArgumentException(String.format("请检查运算符: %s, %s\\n", opt1, opt2));
        }
        return operation1.getPriority() - operation2.getPriority();
    }

    /**
     * 计算结果
     */
    private int calculateNum(int num1, int num2, String opt) {
        Operation operation = operations.get(opt);
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

    /**
     * 后缀表达式计算
     * 1. 循环读取每个字符，判断是否是数字
     * 2. 如果是数字直接入栈
     * 3. 如果是运算符，从栈中弹出两个数，计算表达式的值，将结果压入栈中
     */
    public int calSuffixExpression(String[] expr) {
        Stack stack = new Stack(expr.length);
        for (String str : expr) {
            if (this.isNum(str)) {
                stack.push(str);
                continue;
            }
            Operation operation = ReversePoland.operations.get(str);
            if (operation == null) {
                throw new IllegalStateException("无效的运算符:" + str);
            }
            // 计算
            String numStr1 = stack.pop();
            String numStr2 = stack.pop();
            int result = this.calculateNum(Integer.parseInt(numStr1), Integer.parseInt(numStr2), str);
            // 入栈
            stack.push(String.valueOf(result));
        }
        // 弹出最后结果
        return Integer.parseInt(stack.pop());
    }

    /**
     * 将表达式转换成数组
     */
    public String[] exprToArray(String expr) {
        if (expr == null || expr.equals("")) {
            return new String[0];
        }
        List<String> strList = new ArrayList<>();
        for (int i = 0; i < expr.length(); i++) {
            StringBuilder sb = new StringBuilder(expr.substring(i, i + 1));
            if (this.isNum(sb.toString())) {
                // 向后面继续判断是否为数字
                while (i + 1 < expr.length() && this.isNum(expr.substring(i + 1, i + 2))) {
                    sb.append(expr, i + 1, i + 2);
                    i++;
                }
            }
            strList.add(sb.toString());
        }

        return strList.toArray(new String[]{});
    }

    /**
     * 中缀表达式转后缀表达式
     */
    public String[] infixToSuffix(String[] infix) {
        if (infix == null) {
            throw new IllegalArgumentException("中缀表达式不能为空");
        }
        // 初始化两个栈，一个运算符栈 stack1 和另一个储存中间结果的栈 stack2
        Stack stack = new Stack(infix.length);
        // 由于中间结果栈不需要弹出元素，可以使用集合来保存
        List<String> suffixes = new ArrayList<>();
        // 循环表达式
        for (String str : infix) {
            // 遇到数字时，将其放入 suffixes
            if (this.isNum(str)) {
                suffixes.add(str);
                continue;
            }
            // 如果是 ( 直接入栈
            if (str.equals("(")) {
                stack.push(str);
                continue;
            }
            if (str.equals(")")) {
                while (!stack.isEmpty() && !stack.peek().equals("(")) {
                    // 弹出 stack 中栈顶的元素，并添加到 suffixes
                    suffixes.add(stack.pop());
                }
                // 弹出 (，消除一对 ( )
                stack.pop();
                continue;
            }
            Operation operation = ReversePoland.operations.get(str);
            if (operation != null) {
                if (stack.isEmpty() || stack.peek().equals("(")) {
                    // 如果 stack 为空或栈顶运算符为左括号 "("，则直接将此运算符入栈
                    stack.push(str);
                    continue;
                }
                // 栈不为空，并且当前字符串的优先级小于等于栈顶的元素
                while (!stack.isEmpty() && this.priority(str, stack.peek()) <= 0) {
                    // 将栈顶的元素添加到 suffixes
                    suffixes.add(stack.pop());
                }
                // 入栈
                stack.push(str);
            } else {
                throw new IllegalStateException("无法识别的字符:" + str);
            }
        }
        // 将 stack 中剩余的运算符依次添加到 suffixes
        while (!stack.isEmpty()) {
            suffixes.add(stack.pop());
        }
        // 因为这里用的是集合，它里面元素的顺序就是栈元素出栈后逆序排列的顺序
        return suffixes.toArray(new String[]{});
    }

    public static void main(String[] args) {
        ReversePoland reversePoland = new ReversePoland();
        String expr = "3 5 3 * + 2 -";
        // 假设数和数或符号之间有空格
        String[] expressions = expr.split(" ");
        int result = reversePoland.calSuffixExpression(expressions);
        System.out.printf("后缀表达式 %s 的计算结果为:%d\n", expr, result);

        expr = "1+((2+3)*4)-5";
        expressions = reversePoland.exprToArray(expr);
        System.out.printf("将中缀表达式放入数组, 结果为:%s\n", Arrays.toString(expressions));

        expressions = reversePoland.infixToSuffix(expressions);
        System.out.printf("中缀表达式转换成后缀表达式, 结果为:%s\n", Arrays.toString(expressions));

        result = reversePoland.calSuffixExpression(expressions);
        System.out.printf("计算表达式%s, 结果为:%s\n", expr, result);
    }
}
