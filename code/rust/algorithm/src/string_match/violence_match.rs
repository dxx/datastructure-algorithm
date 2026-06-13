/// 暴力匹配
/// 如果当前字符匹配成功（即 S[i] == P[j]），则 i++，j ++，继续匹配下一个字符
/// 如果失配（即 S[i] != P[j]），令i = i - (j - 1)，j = 0。相当于每次匹配失败时，i 回溯，j 被置为 0
fn violence_search(str: &str, m: &str) -> i32 {
    let str_bytes = str.as_bytes();
    let match_bytes = m.as_bytes();

    let str_length = str_bytes.len();
    let match_length = match_bytes.len();

    if str_length < match_length {
        return -1;
    }

    let mut i: i32 = 0;
    let mut j: i32 = 0;
    // 当字符串长度并且子字符串长度超出范围
    while i < str_length as i32 && j < match_length as i32 {
        if str_bytes[i as usize] == match_bytes[j as usize] {
            i += 1;
            j += 1;
        } else {
            // i 重置到上一次第一个相等字符的位置 + 1
            i = i - (j - 1);
            // j 重置为 0
            j = 0;
        }
    }

    return if j == match_length as i32 {
        return i - j; // 匹配到，然后返回索引
    } else {
        -1
    }
}

#[test]
fn test_violence_search() {
    let str = "CBC DCABCABABCABD BBCCA";
    let m = "ABCABD";
    let index = violence_search(str, m);
    println!("{} 在 {} 中的位置为 {}", m, str, index);
}
