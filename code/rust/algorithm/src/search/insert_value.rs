/// 插值查找
/// 基于二分法查找，步骤和二分法查找一样
/// 中间下标计算公式: mid = start + (end - start) * (value - array[start]) / (array[end] - array[start])
fn insert_val_search(nums: &mut [i32], start: i32, end: i32, find_val: i32) -> i32 {
    if start > end {
        return -1;
    }
    // 根据 find_val 自适应计算中间下标
    let mid = start + (end - start) * (find_val - nums[start as usize]) / (nums[end as usize] - nums[start as usize]);
    println!("mid: {}", mid);
    if find_val < nums[mid as usize] {
        // 向左递归
        return insert_val_search(nums, start, mid - 1, find_val);
    } else if find_val > nums[mid as usize] {
        // 向右递归
        return insert_val_search(nums, mid + 1, end, find_val);
    }
    // 查找值和中间值相等，返回下标
    mid
}

#[test]
fn test_insert_val_search() {
    let nums: &mut [i32] = &mut [0; 100];
    for n in 1..=100 {
        nums[n - 1] = n as i32;
    }
    let value = 58;
    let index = insert_val_search(nums, 0, nums.len() as i32 - 1, value);
    if index != -1 {
        println!("找到 {}, 下标为 {}", value, index);
    } else {
        println!("未找到 {}", value);
    }
}
