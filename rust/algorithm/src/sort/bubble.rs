/// 冒泡排序
/// 1.从当前元素起，向前依次比较每一对相邻元素，若逆序则交换
/// 2.对所有元素均重复以上步骤，直至最后一个元素
fn bubble_sort(nums: &mut Vec<i32>) {
    let length = nums.len();
    // 外循环为排序趟数，length 个数进行 length - 1 趟
    for i in 0..length - 1 {
        // 内循环为每趟比较的次数，第 i 趟比较 length - i 次
        for j in 0..length - 1 - i {
            // 递减循环
            let index = length - 1 - j;
            // 相邻元素比较比较大小，然后交换位置
            if nums[index] < nums[index - 1] {
                swap(nums, index, index - 1);
            }
        }
        println!("第 {} 趟排序结果:{:?}", i + 1, nums);
    }
}

/// 优化
/// 在某次循环中，如果发现没有发生交换，则终止循环
fn optimize_bubble_sort(nums: &mut Vec<i32>) {
    let length = nums.len();
    let mut is_change = false; // 标记是否发生交换
    for i in 0..length - 1 {
        for j in 0..length - 1 - i {
            // 递减循环
            let index = length - 1 - j;
            if nums[index] < nums[index - 1] {
                swap(nums, index, index - 1);
                is_change = true; // 发生交换
            }
        }
        println!("第 {} 趟排序结果:{:?}", i + 1, nums);
        if !is_change {
            break; // 跳出循环，终止比较
        } else {
            is_change = false; // 重置
        }
    }
}

fn swap(v: &mut Vec<i32>, a: usize, b: usize) {
    let temp = v[a];
    v[a] = v[b];
    v[b] = temp;
}

#[test]
fn test_bubble_sort() {
    let mut nums = vec![1, 5, 7, 3, 2, 4, 9, 6, 8];
    println!("交换前: {:?}", nums);
    bubble_sort(&mut nums);
    println!("交换后: {:?}", nums);

    let mut nums = vec![1, 5, 7, 3, 2, 4, 9, 6, 8];
    println!("优化前: {:?}", nums);
    optimize_bubble_sort(&mut nums);
    println!("优化后: {:?}", nums);
}
