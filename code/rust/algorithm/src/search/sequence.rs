/// 线性查找
fn sequence_search(nums: &mut [i32], num: i32) -> i32 {
    for i in 0..nums.len() {
        if nums[i] == num {
            return i as i32;
        }
    }
    return -1;
}

#[test]
fn test_sequence_search() {
    let value = 8;
    let mut nums = [2, 5, 1, 7, 8, 16];
    let index = sequence_search(&mut nums, value);
    if index != -1 {
        println!("{} 在 nums 中的下标为：{}", value, index);
    }
}
