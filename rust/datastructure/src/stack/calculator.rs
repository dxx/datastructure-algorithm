extern crate regex;

use array_stack::Stack;
use std::collections::HashMap;

/// 综合计算器

#[derive(Clone, Debug)]
pub struct Operation {
    opt: String,
    priority: i8,
    opt_func: fn(i32, i32) -> i32,
}

/// 计算器结构体
pub struct Calculator {
    num_stack: Stack,
    operation_stack: Stack,
    operations: HashMap<String, Operation>,
}

impl Calculator {
    pub fn new() -> Self {
        return Calculator {
            num_stack: Stack::new(10),
            operation_stack: Stack::new(10),
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

    /// 判断是否是操作符号
    fn is_operation(&self, opt: &str) -> bool {
        let option = self.operations.get(opt);
        return option.is_some();
    }

    /// 判断是否为数字
    fn is_num(&self, char: &str) -> bool {
        let num_regex = regex::Regex::new("\\d+").unwrap();
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
            let temp = num1;
            num1 = num2;
            num2 = temp;
        }
        return opt_func(num1, num2);
    }

    fn calculate_num_from_stack(&mut self) -> i32 {
        // 从数栈中弹出两个数，从符号栈中弹出一个符号
        let num_str1 = self.num_stack.pop().unwrap();
        let num_str2 = self.num_stack.pop().unwrap();
        let opt = self.operation_stack.pop().unwrap();
        // 计算值
        return self.calculate_num(
            num_str1.parse::<i32>().unwrap(),
            num_str2.parse::<i32>().unwrap(),
            &opt,
        );
    }

    /// 计算表达式的值
    pub fn calculate(&mut self, expression: &str) {
        if expression == "" {
            return;
        }
        let mut index: usize = 0;
        let mut number = String::from("");
        while index < expression.chars().count() {
            let char = &expression[index..index + 1];
            // 判断是否为符号
            if self.is_operation(char) {
                // 判断符号栈是否为空
                if self.operation_stack.is_empty() {
                    // 压入符号栈
                    let _ = self.operation_stack.push(String::from(char));
                } else {
                    // 符号栈不为空，判断优先级
                    let opt = self.operation_stack.peek().unwrap();
                    // char 优先级小于等于 opt
                    if self.priority(char, &opt) <= 0 {
                        // 计算值
                        let result = self.calculate_num_from_stack();
                        // 将计算结果压入数栈
                        let _ = self.num_stack.push(result.to_string());
                    }
                    // 将当前操作符入符号栈
                    let _ = self.operation_stack.push(String::from(char));
                }
            } else if self.is_num(char) {
                // 向后面再取一位判断是否为数字
                if index < expression.chars().count() - 1
                    && self.is_num(&expression[index + 1..index + 2])
                {
                    number = number + &char.to_owned();
                    index += 1;
                    continue;
                }
                // 压入数栈
                let _ = self.num_stack.push(number.to_owned() + char);
                number = String::from("");
            } else {
                panic!("无法识别的字符串:{}", char);
            }

            index = index + 1;
        }

        // 全部数和符号都压入对应的栈后，取出计算
        // 符号栈为空，跳出循环
        while !self.operation_stack.is_empty() {
            // 计算值
            let result = self.calculate_num_from_stack();
            // 将计算结果入数栈
            let _ = self.num_stack.push(result.to_string());
        }
        // 弹出最终结果
        let result = self.num_stack.pop().unwrap();
        println!("表达式执行结果：{}={}", expression, result);
    }
}

#[test]
fn test_stack_calculator() {
    let mut calculator = Calculator::new();

    calculator.calculate("3+5*3-6");
    calculator.calculate("30+5*3-6");
    calculator.calculate("130+5*3-6");
}
