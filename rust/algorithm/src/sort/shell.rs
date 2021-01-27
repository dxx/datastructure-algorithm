/// 希尔排序
/// 1.计算出步长 step，step = length / 2
/// 2.从 step 开始，循环到 length
/// 3.将循环开始时的元素和比当前大 step 的元素进行比较
/// 4.发现逆序则进行交换
fn shell_sort(nums: &mut [i32]) {
    let length = nums.len();
    // 控制步长
    let mut step = length as i32 / 2;
    loop {
        if step > 0 {
            for i in step as usize..length {
                let mut j = i as i32 - step;
                loop {
                    if j >= 0 && nums[j as usize] > nums[j as usize + step as usize] {
                        // 前面的数比后面的数大，进行交换
                        swap(nums, j as usize, j as usize + step as usize);
                        j -= step;
                        continue;
                    }
                    break;
                }
            }
            step /= 2;
            continue;
        }
        break;
    }
}

fn swap(v: &mut [i32], a: usize, b: usize) {
    let temp = v[a];
    v[a] = v[b];
    v[b] = temp;
}

#[test]
fn test_shell_sort() {
    let mut nums = [5, 1, 7, 3, 2, 4, 9, 6, 8];
    println!("排序前: {:?}", nums);
    shell_sort(&mut nums);
    println!("排序后: {:?}", nums);
}
