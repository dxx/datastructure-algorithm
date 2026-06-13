/// 二分查找
/// 1.先找到中间值
/// 2.将中间值和查找值比较
///   查找值小于中间值, 向左进行递归查找
///   查找值大于中间值, 向右进行递归查找
///   查找值和中间值相等，返回当前下标
/// 3.如果查找时，左边的小标大于右边的下标表示未找到，返回 -1
/// 注意：使用二分查找的前提是该数组是有序的
fn binary_search(nums: &mut [i32], start: i32, end: i32, find_val: i32) -> i32 {
    if start > end {
        // 表示未找到
        return -1;
    }
    let mid = (start + end) / 2;
    if find_val < nums[mid as usize] {
        // 向左递归
        return binary_search(nums, start, mid - 1, find_val);
    } else if find_val > nums[mid as usize] {
        // 向右递归
        return binary_search(nums, mid + 1, end, find_val);
    }
    // 查找值和中间值相等，返回下标
    return mid;
}

#[test]
fn test_binary_search() {
    let value = 100;
    let nums: &mut [i32] = &mut [1, 8, 10, 89, 100, 100, 123];
    let index = binary_search(nums, 0, nums.len() as i32 - 1, value);
    if index != -1 {
        println!("找到 {}, 下标为 {}", value, index);
    } else {
        println!("未找到 {}", value);
    }
}
