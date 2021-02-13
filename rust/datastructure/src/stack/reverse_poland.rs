use crate::array_stack::Stack;
use std::collections::HashMap;
use regex::Regex;

/// 逆波兰计算器

#[derive(Clone, Debug)]
pub struct Operation {
    opt: String,
    priority: i8,
    opt_func: fn(i32, i32) -> i32,
}

pub struct ReversePoland {
    operations: HashMap<String, Operation>,
}

impl ReversePoland {
    pub fn new() -> Self {
        return ReversePoland {
            operations: [
                (
                    String::from("+"),
                    Operation {
                        opt: String::from("+"),
                        priority: 1,
                        opt_func: |num1: i32, num2: i32| -> i32 { num1 + num2 },
                    },
                ),
                (
                    String::from("-"),
                    Operation {
                        opt: String::from("-"),
                        priority: 1,
                        opt_func: |num1: i32, num2: i32| -> i32 { num1 - num2 },
                    },
                ),
                (
                    String::from("*"),
                    Operation {
                        opt: String::from("*"),
                        priority: 2,
                        opt_func: |num1: i32, num2: i32| -> i32 { num1 * num2 },
                    },
                ),
                (
                    String::from("/"),
                    Operation {
                        opt: String::from("/"),
                        priority: 2,
                        opt_func: |num1: i32, num2: i32| -> i32 { num1 / num2 },
                    },
                ),
            ]
            .iter()
            .cloned()
            .collect(),
        };
    }

    /// 判断是否为数字
    fn is_num(&self, char: &str) -> bool {
        let num_regex = Regex::new("\\d+").unwrap();
        return num_regex.is_match(char);
    }

    /// 计算操作符的优先级
    fn priority(&self, opt1: &str, opt2: &str) -> i8 {
        let option1 = self.operations.get(opt1);
        let option2 = self.operations.get(opt2);
        if option1.is_some() && option2.is_some() {
            return option1.unwrap().priority - option2.unwrap().priority;
        }
        panic!("请检查运算符: {}, {}", opt1, opt2)
    }

    /// 计算结果
    fn calculate_num(&self, mut num1: i32, mut num2: i32, opt: &str) -> i32 {
        let opt_func = self.operations.get(opt).unwrap().opt_func;
        if opt == "-" || opt == "/" {
            // 因为出栈后两数的位置颠倒，需交换两个数的位置
            num1 = num1 + num2;
            num2 = num1 - num2;
            num1 = num1 - num2;
        }
        return opt_func(num1, num2);
    }

    /// 后缀表达式计算
    /// 1. 循环读取每个字符，判断是否是数字
    /// 2. 如果是数字直接入栈
    /// 3. 如果是运算符，从栈中弹出两个数，计算表达式的值，将结果压入栈中
    pub fn cal_suffix_expression(&self, expr: Vec<&str>) -> i32 {
        let mut stack = Stack::new(expr.len());
        for i in 0..expr.len() {
            let str: &str = expr.get(i).unwrap();
            if self.is_num(str) {
                let _ = stack.push(String::from(str));
                continue;
            }
            let option = self.operations.get(str);
            if option.is_none() {
                panic!("无效的运算符: {}", str)
            }
            // 计算
            // 从数栈中弹出两个数，从符号栈中弹出一个符号
            let num_str1 = stack.pop().unwrap();
            let num_str2 = stack.pop().unwrap();
            let result = self.calculate_num(
                num_str1.parse::<i32>().unwrap(),
                num_str2.parse::<i32>().unwrap(),
                str,
            );
            // 入栈
            let _ = stack.push(result.to_string());
        }
        // 弹出最后结果
        return stack.pop().unwrap().parse::<i32>().unwrap();
    }

    /// 将表达式转换成向量
    pub fn expr_to_vec(&self, expr: &str) -> Vec<String> {
        let mut expressions: Vec<String> = Vec::new();
        if expr == "" {
            return expressions;
        }
        for mut i in 0..expr.len() {
            let mut char = String::from(&expr[i..i + 1]);
            if self.is_num(&char) {
                // 向后面继续判断是否为数字
                while i + 1 < expr.len() && self.is_num(&expr[i + 1..i + 2]) {
                    char += &expr[i + 1..i + 2];
                    i += 1;
                }
            }
            expressions.push(char);
        }
        return expressions;
    }

    /// 中缀表达式转后缀表达式
    pub fn infix_to_suffix(&self, infix: Vec<String>) -> Vec<String> {
        // 初始化两个栈，一个运算符栈 stack1 和另一个储存中间结果的栈 stack2
        let mut stack = Stack::new(infix.len());
        // 由于中间结果栈不需要弹出元素，可以使用向量来保存
        let mut suffixes: Vec<String> = Vec::new();
        // 循环表达式
        for i in 0..infix.len() {
            let str: &str = infix.get(i).unwrap();
            // 遇到数字时，将其放入 suffixes
            if self.is_num(str) {
                suffixes.push(String::from(str));
                continue;
            }
            if str == "(" {
                // 如果是 ( 直接入栈
                let _ = stack.push(String::from(str));
                continue;
            }
            if str == ")" {
                while stack.peek().unwrap() != "(" {
                    // 弹出 stack 中栈顶的元素，并追加到 suffixes
                    let elem = stack.pop().unwrap();
                    suffixes.push(elem);
                }
                // 弹出 (，消除一对 ( )
                let _ = stack.pop();
                continue;
            }
            // 如果是运算符
            let option = self.operations.get(str);
            if option.is_some() {
                if stack.is_empty() || stack.peek().unwrap() == "(" {
                    // 如果 stack 为空或栈顶运算符为左括号 "("，则直接将此运算符入栈
                    let _ = stack.push(String::from(str));
                    continue;
                }
                // 栈不为空，并且当前字符串的优先级小于等于栈顶的元素
                while !stack.is_empty()
                    && self.priority(&option.unwrap().opt, &stack.peek().unwrap()) <= 0
                {
                    let elem = stack.pop().unwrap();
                    // 将栈顶的元素追加到 suffixes
                    suffixes.push(elem);
                }
                // 直接入栈
                let _ = stack.push(String::from(str));
            } else {
                panic!("无法识别的字符: {}", str)
            }
        }
        while !stack.is_empty() {
            // 将 stack 中剩余的运算符依次追加到 suffixes
            let elem = stack.pop().unwrap();
            suffixes.push(elem);
        }

        return suffixes;
    }
}

#[test]
fn test_reverse_poland() {
    let reverse_poland = ReversePoland::new();
    let expr = "3 5 3 * + 2 -";
    // 假设数和数或符号之间有空格
    let expressions = expr.split(" ").collect();
    let result = reverse_poland.cal_suffix_expression(expressions);
    println!("后缀表达式 {} 的计算结果为: {}", expr, result);

    let expr = "1+((2+3)*4)-5";
    let expressions = reverse_poland.expr_to_vec(expr);
    println!("将中缀表达式放入向量, 结果为: {:?}", expressions);

    let expressions = reverse_poland.infix_to_suffix(expressions);
    println!("中缀表达式转换成后缀表达式, 结果为: {:?}", expressions);

    let result = reverse_poland.cal_suffix_expression(expressions.iter().map(|s| s as &str).collect());
    println!("计算表达式 {:?}, 结果为: {}", expr, result)
}
