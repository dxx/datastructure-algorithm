extern crate array_stack;

fn main() {
    // 创建一个栈
    let mut stack = array_stack::Stack::new(3);
    // 入栈
    let _ = stack.push(String::from("one"));
    let _ = stack.push(String::from("two"));
    let _ = stack.push(String::from("three"));

    // 栈满，无法入栈
    let r = stack.push(String::from("four"));
    match r {
        Ok(data) => {
            println!("{}", data);
        },
        Err(e)  => {
            println!("{}", e);
        }
    }

    stack.show_stack();

    let elem1 = stack.pop().unwrap();
    let elem2 = stack.pop().unwrap();
    let elem3 = stack.pop().unwrap();

    println!("出栈:{}", elem1);
    println!("出栈:{}", elem2);
    println!("出栈:{}", elem3);

    // 栈空无法出栈
    let o = stack.pop();
    match o {
        Some(data) => {
            println!("{}", data);
        },
        None => {
            println!("stack is empty");
        }
    }

    stack.show_stack();

}