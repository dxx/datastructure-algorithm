/// 递归
/// 递归是指在程序运行过程中调用本身的编程技巧

/// 计算阶乘
fn factorial(n: u32) -> u32 {
    if n > 0 {
        return n * factorial(n - 1);
    }
    return 1;
}

#[test]
fn test_factorial() {
    let res = factorial(5);
    println!("{}", res); // 120
}
