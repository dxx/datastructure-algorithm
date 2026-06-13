/// 插入排序
/// 1.从第二个元素开始循环，循环到元素末尾，左边分为有序列表，右边分为无序列表
/// 2.将右边无序列表的第一个元素，标记为要插入的值（insertValue），并记录要插入的位置（insertIndex）
/// 3.依次和左边的无序列表中的元素比较，如果顺序颠倒，则将有序列表中当前被比较的元素后移，同时修改 insertIndex
/// 4.最后将 insertValue 插入到 insertIndex 位置
fn insert_sort(nums: &mut [i32]) {
    for i in 1..nums.len() {
        // 要插入的下标
        let mut index= i as i32 - 1;
        // 保存插入的值
        let insert_value = nums[i];
        // 从小到大
        while index >= 0 {
            let insert_index = index as usize;
            if nums[insert_index] > insert_value {
                // 元素后移
                nums[insert_index + 1] = nums[insert_index];
                // 下标前移
                index -= 1;
            } else {
                break;
            }
        }
        // 优化：判断如果当前位置发生移动，就插入
        if index + 1 != i as i32 {
            // 插入
            nums[(index + 1) as usize] = insert_value;
        }
    }
}

#[test]
fn test_insert_sort() {
    let mut nums = [5, 1, 7, 3, 2, 4, 9, 6, 8];
    println!("排序前: {:?}", nums);
    insert_sort(&mut nums);
    println!("排序后: {:?}", nums);
}
