/// KMP 匹配
/// 核心是利用匹配失败后的信息，尽量减少模式串与主串的匹配次数以达到快速匹配的目的
/// 具体实现就是通过一个 next() 函数实现，函数本身包含了模式串的局部匹配信息
fn kmp_search(str: &str, m: &str) -> i32 {
    let next = get_next(m);

    let str_bytes = str.as_bytes();
    let match_bytes = m.as_bytes();

    let str_length = str_bytes.len();
    let match_length = match_bytes.len();

    let mut j = 0;
    for i in 0..str_length {
        // 算法核心点
        while j > 0 && str_bytes[i] != match_bytes[j as usize] {
            // 根据部分匹配表，更新 j
            j = next[j as usize - 1]
        }

        if str_bytes[i] == match_bytes[j as usize] {
            j += 1;
        }
        // 判断是否找到了
        if j == match_length as i32 {
            return i as i32 - (j - 1);
        }
    }
    return -1;
}

/// 计算出匹配的部分信息
fn get_next(m: &str) -> Vec<i32> {
    let m_bytes = m.as_bytes();
    let mut next = vec![0; m_bytes.len()];

    // 第一个字符的值为 0
    next[0] = 0;

    let mut j: i32 = 0;
    for i in 1..m_bytes.len() {

        // 核心，比较直到相等
        while j > 0 && m_bytes[i] != m_bytes[j as usize] {
            // 更新 j
            j = next[j as usize - 1];
        }

        // 相等
        if m_bytes[i] == m_bytes[j as usize] {
            j += 1;
        }
        next[i] = j;
    }
    return next;
}

#[test]
fn test_kmp_search() {
    let str = "CBC DCABCABABCABD BBCCA";
    let m = "ABCABD";
    let index = kmp_search(str, m);
    println!("{} 在 {} 中的位置为 {}", m, str, index);
}
