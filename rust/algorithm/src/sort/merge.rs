/// 归并排序
/// 1.将要排序的序列，从中间开始递归分解
/// 2.分解到子序列只有一个元素，然后开始合并
/// 3.创建一个临时序列，大小为本次需要合并的两个有序序列长度之和
/// 4.设定两个下标，初始位置分别为两个有序序列的下标
/// 5.比较两个小标指向的元素，将较小的元素放入临时序列，并移动下标
/// 6.重复第 5 步，直到有一个下标超出序列尾部
/// 7.将另一个序列中剩余的元素保存当前顺序放入临时序列中
/// 8.最后将临时序列复制到原始序列对应的位置
fn merge_sort(nums: &mut [i32]) {
    decompose(nums, 0, nums.len() - 1);
}

/// 分解
/// start: 开始下标
/// end: 结束下标
fn decompose(nums: &mut [i32], start: usize, end: usize) {
    if start < end {
        let mid = (start + end) / 2;
        decompose(nums, start, mid); // 左边
        decompose(nums, mid + 1, end); // 右边
        merge(nums, start, mid, end);
    }
}

/// 合并
/// left: 左边有序序列起始下标
/// mid: 中间下标
/// right: 右边有序序列起始下标
fn merge(nums: &mut [i32], left: usize, mid: usize, right: usize) {
    let mut temp = Vec::new(); // 临时保存元素的向量
    let mut i = left; // 左边有序序列的起始下标
    let mut j = mid + 1; // 右边有序列表的起始下标
    // 先把左右两边有序列表中的数据，取出来比较，然后将较小的元素放入 temp 中
    while i <= mid && j <= right {
        if nums[i] <= nums[j] {
            temp.push(nums[i]);
            i += 1;
        } else {
            temp.push(nums[j]);
            j += 1;
        }
    }

    // 将剩余的左边或右边有序序列中的元素全部放入 temp 中
    while i <= mid {
        // 左边序列还有元素
        temp.push(nums[i]);
        i += 1;
    }

    while j <= right {
        // 右边序列还有元素
        temp.push(nums[j]);
        j += 1;
    }

    // 将 temp 中的元素拷贝到 nums 中
    let mut t = 0;
    let mut l = left;
    while l <= right {
        nums[l] = temp[t];
        l += 1;
        t += 1;
    }
}

#[test]
fn test_merge_sort() {
    let mut nums = [5, 0, 1, 7, 3, 2, 4, 9, 6, 8];
    println!("排序前: {:?}", nums);
    merge_sort(&mut nums);
    println!("排序后: {:?}", nums);
}
