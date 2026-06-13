/// 斐波那契数组长度
const MAX_SIZE: usize = 20;

/// 获取斐波那契数列
fn fibonacci() -> Vec<u32> {
    let mut fib = vec![0; MAX_SIZE];
    fib[0] = 1;
    fib[1] = 1;
    for i in 2..MAX_SIZE {
        fib[i] = fib[i - 1] + fib[i - 2];
    }
    return fib;
}

/// 斐波那契查找
/// 中间值计算公式: mid = start + F[k - 1] - 1
fn fibonacci_search(nums: &mut [i32], find_val: i32) -> i32 {
    let mut start: usize = 0;
    let mut end = nums.len() - 1;
    let mut k: usize = 0; // 斐波那契分割数的下标
    let f = fibonacci(); // 斐波那契数列

    // 获取斐波那契分割数的下标
    while end > f[k] as usize - 1 {
        k += 1;
    }

    // 创建一个向量，长度为斐波那契分割数的值
    let dst_length = f[k] as usize;
    let mut dst: Vec<i32> = vec![0; dst_length];
    // 将要查找的数组复制到 dst 中
    for i in 0..nums.len() {
        dst[i] = nums[i];
    }
    // 将最后一个元素填充到 dst 元素为 0 的位置
    // [1, 8, 10, 89, 100, 100, 123] => [1 8 10 89 100 100 123 123]
    for i in end + 1..dst_length {
        dst[i] = nums[end];
    }

    // 使用循环来寻找 find_val
    while start <= end {
        let mid = start + f[k - 1] as usize - 1; // 中间值
        if find_val < dst[mid] {
            // 向左查找
            end = mid - 1;
            // 全部元素 = 前面的元素 + 后边元素
            // 即f[k] = f[k-1] + f[k-2]
            // 因为前面有 f[k-1]个元素,所以可以继续拆分 f[k-1] = f[k-2] + f[k-3]
            // 即 在 f[k-1] 的前面继续查找 k--
            // 即下次循环 mid = f[k-1-1]-1
            k -= 1
        } else if find_val > dst[mid] {
            // 向右查找
            start = mid + 1;
            // 全部元素 = 前面的元素 + 后边元素
            // f[k] = f[k-1] + f[k-2]
            // 因为后面我们有f[k-2] 所以可以继续拆分 f[k-1] = f[k-3] + f[k-4]
            // 即在f[k-2] 的前面进行查找 k -=2
            // 即下次循环 mid = f[k - 1 - 2] - 1
            k -= 2;
        } else {
            return if mid <= end {
                mid as i32 // 若相等则说明 mid 即为查找到的位置
            } else {
                end as i32 // 若 mid > end 则说明是扩展的数值，返回 end
            };
        }
    }
    return -1;
}

#[test]
fn test_fibonacci_search() {
    let value = 100;
    let mut nums = [1, 8, 10, 89, 100, 100, 123];
    let index = fibonacci_search(&mut nums, value);
    if index != -1 {
        println!("找到 {}, 下标为 {}", value, index);
    } else {
        println!("未找到 {}", value);
    }
}
