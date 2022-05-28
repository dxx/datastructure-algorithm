/// 稀疏数组
/// 在二维数组中，如果值为 0 的元素数目远远大于非 0 元素的数目，并且非 0 元素的分布没有规律，则该数组被称为稀疏数组
/// 稀疏数组可以看做是一个压缩的数组，稀疏数组的好处有：
/// 原数组中存在大量的无用的数据，占据了存储空间，真正有用的数据却很少
/// 压缩后存储可以节省存储空间，在数据序列化到磁盘时，压缩存储可以提高 IO 效率

/// 二维数组
/// 0	0	3	0	0
/// 0	0	0	6	0
/// 0	1	0	0	0
/// 0	0	0	5	0
/// 0	0	0	0	0
///
/// 稀疏数组
/// row	col	val
/// 5	5	4
/// 0	2	3
/// 1	3	6
/// 2	1	1
/// 3	3	5
use std::fs;

const SPARSE_ARRAY_FILENAME: &str = "./sparse.data";

/// 二维数组转稀疏数组
fn to_sparse_array(array: [[i32; 5]; 5]) -> Vec<[i32; 3]> {
    // 统计非 0 的数量
    let mut count = 0;
    for i in 0..array.len() {
        for j in 0..array[i].len() {
            if array[i][j] != 0 {
                count = count + 1;
            }
        }
    }
    // 编译报错
    // let mut sparse_array = [[0; 3]; count + 1];
    let mut sparse_array: Vec<[i32; 3]> = vec![];
    // 存储第一行信息，从左到右存储的依次是 行，列，非 0 的个数
    sparse_array.push([array.len() as i32, array[0].len() as i32, count as i32]);
    for i in 0..array.len() {
        for j in 0..array[i].len() {
            if array[i][j] != 0 {
                // 保存 row, col, val
                sparse_array.push([i as i32, j as i32, array[i][j]]);
            }
        }
    }
    return sparse_array;
}

/// 稀疏数组转二维数组
fn to_array(sparse_array: Vec<[i32; 3]>) -> [[i32; 5]; 5] {
    let mut array = [[0; 5]; 5];
    let mut row;
    let mut col;
    let mut val;
    // 从稀疏数组中第二行开始读取数据
    for i in 1..sparse_array.len() {
        row = sparse_array[i][0];
        col = sparse_array[i][1];
        val = sparse_array[i][2];
        array[row as usize][col as usize] = val;
    }
    return array;
}

/// 存储稀疏数组
fn storage_sparse_array(sparse_array: Vec<[i32; 3]>) {
    let mut data = String::from("");
    // 存储矩阵格式
    for i in 0..sparse_array.len() {
        for j in 0..sparse_array[i].len() {
            let v = sparse_array[i][j].to_string();
            data.push_str(&v[0..v.len()]);
            data.push_str("\t");
        }
        data.push_str("\n");
    }

    fs::write(SPARSE_ARRAY_FILENAME, data).unwrap();
}

/// 读取稀疏数组
fn read_sparse_array() -> Vec<[i32; 3]> {
    let content = fs::read_to_string(SPARSE_ARRAY_FILENAME).unwrap();
    let file_lines: Vec<&str> = content.split("\n").collect();
    // 读取第一行
    let mut strs: Vec<&str> = file_lines[0].split("\t").collect();

    // 第一行信息
    let mut row: i32 = strs[0].parse().unwrap();
    let mut col: i32 = strs[1].parse().unwrap();
    let mut val: usize = strs[2].parse().unwrap();

    let mut sparse_array = vec![[0; 3]; val + 1];
    sparse_array[0] = [row, col, val as i32];
    // 从第二行开始
    for i in 1..file_lines.len() {
        if !file_lines[i].is_empty() {
            strs = file_lines[i].split("\t").collect();
            row = strs[0].parse().unwrap();
            col = strs[1].parse().unwrap();
            val = strs[2].parse().unwrap();
            sparse_array[i] = [row, col, val as i32];
        }
    }
    return sparse_array;
}

fn print_array(array: [[i32; 5]; 5]) {
    for i in 0..array.len() {
        for j in 0..array[i].len() {
            print!("{}\t", array[i][j]);
        }
        print!("\n");
    }
}

fn print_sparse_array(array: Vec<[i32; 3]>) {
    for i in 0..array.len() {
        for j in 0..array[i].len() {
            print!("{}\t", array[i][j]);
        }
        print!("\n");
    }
}

#[test]
fn test_sparse_array() {
    // 定义一个 5x5 的二维数组
    let mut array: [[i32; 5]; 5] = [[0; 5]; 5];
    // 初始化 3，6，1，5
    array[0][2] = 3;
    array[1][3] = 6;
    array[2][1] = 1;
    array[3][3] = 5;

    println!("原二维数组：");
    print_array(array);

    // 转成稀疏数组
    let mut sparse_array = to_sparse_array(array);

    println!("转换后的稀疏数组：");
    print_sparse_array(sparse_array.clone());

    // 存储稀疏数组
    storage_sparse_array(sparse_array.clone());

    // 读取稀疏数组
    sparse_array = read_sparse_array();
    println!("读取的稀疏数组：");
    print_sparse_array(sparse_array.clone());

    // 转成二维数组
    array = to_array(sparse_array);
    println!("转换后的二维数组：");
    print_array(array);
}
