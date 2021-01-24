/// 选择排序
/// 1.假定第一元素为最大或最小的元素
/// 2.找出最大或最小的元素的小标，循环 length - 1 次
/// 3.每次循环完成后将最大值或最小值和本次循环的第一个元素交换
fn select_sort(nums: &mut Vec<i32>) {
    let length = nums.len();
    for i in 0..length - 1 {
        // 记录最小值的下标
        let mut min_index = i;
        for j in i + 1..length {
            if nums[min_index] > nums[j] {
                // 修改最小值下标
                min_index = j;
            }
        }
        // 优化：判断是否需要交换
        if min_index != i {
            swap(nums, i, min_index);
        }
    }
}

fn swap(v: &mut Vec<i32>, a: usize, b: usize) {
    let temp = v[a];
    v[a] = v[b];
    v[b] = temp;
}

#[test]
fn test_select_sort() {
    let mut nums = vec![3, 5, 7, 1, 2, 4, 9, 6, 8];
    println!("排序前: {:?}", nums);
    select_sort(&mut nums);
    println!("排序后: {:?}", nums);
}
