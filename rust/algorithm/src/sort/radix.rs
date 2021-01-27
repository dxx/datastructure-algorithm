/// 基数排序
/// 1.创建一个二维数组，数组长度为 10，并初始化 10 个一维数组，每个一维数组的长度为待排序数组的长度
/// 2.遍历待排序的数组，从最低位开始，求出每个元素的个位数作为二维数组的下标，将其放入到二维数组对应的数组中
/// 3.从二维数组中依次取出所有元素放入原数组中
/// 4.重复步骤 2，依次计算个位、十位、百位等，作为下标，直到待排序数组中的最大位数
fn radix_sort(nums: &mut [i32]) {
    if nums.len() == 0 {
        return;
    }
    let mut max = nums[0];
    for num in nums.iter() {
        if max < *num {
            max = *num;
        }
    }
    let max_length = max.to_string().len();

    let mut bucket = vec![vec![0; nums.len()]; 10]; // 桶数组，二维数组中的元素使用切片替代
    let mut order = [0; 10]; // 存放每个桶真实存放数据的长度

    let mut n = 1; // 控制元素的位数
    for _i in 0..max_length {
        for num in nums.iter() {
            let bucket_index = (*num / n % 10) as usize; // 计算桶的下表
            bucket[bucket_index][order[bucket_index]] = *num;
            order[bucket_index] += 1; // 尾下标 + 1
        }

        let mut num_index = 0;
        // 从桶数组中依次取出所有元素放入原数组
        for i in 0..order.len() {
            let bucket_length = order[i];
            if bucket_length != 0 {
                for j in 0..bucket_length {
                    nums[num_index] = bucket[i][j];
                    num_index += 1;
                }
                order[i] = 0; // 重置当前桶下标
            }
        }

        n *= 10;
    }
}

#[test]
fn test_radix_sort() {
    let mut nums = [5, 1, 7, 13, 21, 32, 9, 66, 8, 20];
    println!("排序前: {:?}", nums);
    radix_sort(&mut nums);
    println!("排序后: {:?}", nums);
}
