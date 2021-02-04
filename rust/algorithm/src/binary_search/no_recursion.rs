/// 二分法查找(非递归)
fn binary_search_no_recursion(nums: &mut [i32], find_val: i32) -> i32 {
    let mut start: usize = 0;
    let mut end = nums.len() - 1;
    while start <= end {
        let mid = (start + end) / 2;
        if find_val < nums[mid] {
            // 查找的值在左边
            end = mid - 1;
        } else if find_val > nums[mid] {
            // 查找的值在右边
            start = mid + 1;
        } else {
            // 找到目标值的下标
            return mid as i32;
        }
    }
    return -1;
}

#[test]
fn test_binary_search_no_recursion() {
    let value = 100;
    let mut nums = [1, 8, 10, 89, 100, 100, 123];
    let index = binary_search_no_recursion(&mut nums, value);
    if index != -1 {
        println!("找到 {}, 下标为 {}", value, index);
    } else {
        println!("未找到 {}", value);
    }
}
