/// 背包问题（01背包）
/// 有 n 件物品和一个容量为 v 的背包。第 i 件物品的重量是 w[i]，价值是 v[i]。
/// 求解将哪些物品装入背包可使这些物品的重量总和不超过背包容量，且价值总和最大。
/// 背包问题可以分 01 背包、完全背包、多重背包等。

/// 01 背包问题要求每种物品仅有一件，可以选择放或不放。

/// 假设现有如下物品
///
/// | 物品名称 | 重量(kg) | 价值 |
/// | ------- | -------- | ---- |
/// | 音响     | 1      | 500  |
/// | 电脑     | 2      | 5000 |
/// | 手机     | 1      | 3000 |
///
/// 给定一个容量为 3 的背包，不超过背包容量时，放入物品使得总价值最大，最大总价值是多少。

/// 思路
/// 将问题转换成填充一个二维表格。定义一个二维数组 value，value\[i\]\[j\] 表示放入前 i 件物品容量为 j 的最大价值总和。
///
/// 第 0 件物品的任意容量总价值都是 0，容量为 0 时，放入任意一件物品的总价值也是 0。
/// 如果当前物品的容量大于当前背包的容量，那么最大价值总和只能是前 i - 1 件物品放入
/// 当前背包时的最大价值总和，最大价值为 value[i -1][j]。如果当前物品的容量
/// 小于或等于当前背包的容量，那么最大价值总和就是第 i 件物品放入当前背包的价值 v[i]
/// 加上前 i - 1 件物品放入剩余容量的最大值和 value[i - 1][v - w[i]]，
/// 即 v[i] + value[i - 1][j - w[i]]，有可能放入前 i 件物品的最大价值和
/// 比放入前 i - 1件最大价值和小，这里就要取两者中的最大值。
///
/// 于是就有状态转移方程：
/// 1. i == 0, j == 0 时，value[0][j] = value[i][0] = 0
/// 2. w[i] > j 时，value[i][j] = value[i -1][j]
/// 3. w[i] <= j 时，max{value[i -1][j], v[i] + value[i - 1][j - w[i]]}

fn find_max_value(w: Vec<u32>, v: Vec<u32>, c: u32) -> u32 {
    let n = w.len();
    // 初始化二维数组，二维数组比实际物品数多出一行一列
    let mut value: Vec<Vec<u32>> = vec![vec![0; (c + 1) as usize]; n + 1];

    // // 初始第一列，默认为 0，也可以不初始化
    // for i in 0..value.len() {
    //     value[i][0] = 0;
    // }
    // // 初始第一行
    // for j in 0..value.len() {
    //     value[0][j] = 0;
    // }
    // 遍历所有物品
    for i in 1..=n {
        // 遍历所有容量
        for j in 1..=(c as usize) {
            value[i][j] = if w[i - 1] > j as u32 { // 实际物品下标从 0 开始
                value[i - 1][j]
            } else {
                max(value[i - 1][j], v[i - 1] + value[i - 1][j - w[i - 1] as usize])
            }
        }
    }

    println!("填表如下：");

    for i in 0..value.len() {
        for val in value[i].iter() {
            print!("{:5}", val);
        }
        println!();
    }
    value[n][c as usize]
}

fn find_max_value2(w: Vec<u32>, v: Vec<u32>, c: u32) -> u32 {
    // 数组比实际物品总数多出一列
    let mut value = vec![0; (c + 1) as usize];

    // 遍历所有物品
    for i in 0..w.len() {
        // 遍历所有容量
        // 因为每个物品只能使用一次，给每一行填表时，采用倒序遍历
        for j in (w[i] as usize..=c as usize).rev() {
            // 滑动数组优化空间复杂度
            value[j] = max(value[j], v[i] + value[j - w[i] as usize]);
        }
    }
    return value[c as usize];
}

fn max(a: u32, b: u32) -> u32 {
    return if a > b {
        a
    } else {
        b
    }
}

#[test]
fn test_backpack() {
    // 物品重量(kg)
    let w = vec![1, 2, 1];
    // 物品价值
    let v = vec![500, 5000, 3000];
    // 背包容量
    let c = 3;
    let max = find_max_value(w, v, c);
    println!("最大价值总和为: {}", max);
}

#[test]
fn test_backpack2() {
    // 物品重量(kg)
    let w = vec![1, 2, 1];
    // 物品价值
    let v = vec![500, 5000, 3000];
    // 背包容量
    let c = 3;
    let max = find_max_value2(w, v, c);
    println!("最大价值总和为: {}", max);
}
