/// 快速排序
/// 1.获取最中间的元素，然后分别将左边和右边的元素分别和最中间的元素进行比较
/// 2.左边找出比中间大的元素，右边找出比中间小的元素，交换左右的位置
/// 3.一次比较完成后，将左边和右边分别递归重复以上操作
fn quick_sort(nums: &mut Vec<i32>, start: usize, end: usize) {
    let mut l = start as i32;
    let mut r = end as i32;
    // 获取最中间的元素
    let center_value = nums[(start + end) / 2];
    while l < r {
        // 从左边找出比中间元素大的元素值
        while nums[l as usize] < center_value {
            l += 1;
        }
        // 从右边找出比中间元素小的值
        while nums[r as usize] > center_value {
            r -= 1;
        }
        if l >= r {
            break;
        }
        // 交换左边和右边的值
        swap(nums, l as usize, r as usize);

        // 交换后，nums[l] 的值等于 center_value，r 前移
        if nums[l as usize] == center_value {
            r -= 1;
        }
        // 交换后，nums[r] 的值等于 center_value，l 后移
        if nums[r as usize] == center_value {
            l += 1;
        }
    }
    if l == r {
        l += 1;
        r -= 1;
    }
    if start < r as usize {
        quick_sort(nums, start, r as usize);
    }
    if end > l as usize {
        quick_sort(nums, l as usize, end);
    }
}

fn swap(v: &mut Vec<i32>, a: usize, b: usize) {
    let temp = v[a];
    v[a] = v[b];
    v[b] = temp;
}

#[test]
fn test_quick_sort() {
    let mut nums = vec![5, 1, 8, 3, 7, 2, 9, 4, 6];
    println!("排序前: {:?}", nums);
    let len = nums.len();
    quick_sort(&mut nums, 0, len - 1);
    println!("排序后: {:?}", nums);
}
