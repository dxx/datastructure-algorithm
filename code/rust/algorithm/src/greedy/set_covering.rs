use std::collections::HashMap;

/// 贪心算法解决集合覆盖，选择最少的广播电台的问题

/// 存在下面需要付费的广播台以及广播台信号可以覆盖的地区。如何选择最少的广播台，让所有地区都接收到信息
/// | 广播台 | 覆盖地区         |
/// | ------ | -------------- |
/// | B1     | 北京，上海，天津 |
/// | B2     | 广州，北京，深圳 |
/// | B3     | 成都，上海，杭州 |
/// | B4     | 上海，天津      |
/// | B5     | 杭州，大连      |
fn get_broadcast(broadcasts: HashMap<String, Vec<String>>) -> Vec<String> {
    let mut broadcast_keys: Vec<String> = Vec::new();

    let mut m: HashMap<String, String> = HashMap::new();
    for (key, value) in broadcasts.iter() {
        // 将 key 存入向量
        broadcast_keys.push(key.clone());
        for v in value {
            // 区域作为 key，保证不重复
            m.insert(v.clone(), "".to_string());
        }
    }

    // 对向量进行排序，避免每次遍历 broadcasts 顺序不一致，导致求出电台结果不一致
    broadcast_keys.sort();

    // 保存所有区域
    let mut all_broadcasts = Vec::new();
    for (key, _) in m {
        all_broadcasts.push(key);
    }

    let mut select_broadcasts = Vec::new();

    // 遍历所有的电台，直到所有电台被移除
    while all_broadcasts.len() > 0 {
        // 覆盖区域最多的电台
        let mut max_broadcast = "".to_string();
        for key in broadcast_keys.iter_mut() {
            let value = broadcasts.get(key);
            let retain = &mut value.unwrap().clone();
            let len = retain.len();
            // 求出交集
            retain.retain(|x| all_broadcasts.contains(x));
            // 交集
            let temp_intersection: Vec<String> = retain.clone();
            // 每次都选择覆盖最多的集合，体现出贪心算法的特点
            if temp_intersection.len() > 0 && (max_broadcast == "" || temp_intersection.len() > len)
            {
                max_broadcast = key.clone();
            }
        }
        if max_broadcast != "" {
            // 将电台添加到向量中
            select_broadcasts.push(max_broadcast.to_string());
            // 将已经覆盖的区域对应的电台移除，不参与下一次比较
            delete_element(&mut broadcast_keys, max_broadcast.to_string());
            for b in broadcasts.get(&max_broadcast).unwrap() {
                // 删除已经覆盖的区域
                delete_element(&mut all_broadcasts, b.clone());
            }
        }
    }

    return select_broadcasts;
}

fn delete_element(vec: &mut Vec<String>, delete_val: String) {
    for i in 0..vec.len() {
        if *vec.get(i).unwrap() == delete_val {
            vec.remove(i);
            // 避免删除重复的数据
            return;
        }
    }
}

#[test]
fn test_get_broadcast() {
    let broadcast_map = [
        (
            "B1".to_string(),
            vec!["北京".to_string(), "上海".to_string(), "天津".to_string()],
        ),
        (
            "B2".to_string(),
            vec!["广州".to_string(), "北京".to_string(), "深圳".to_string()],
        ),
        (
            "B3".to_string(),
            vec!["成都".to_string(), "上海".to_string(), "杭州".to_string()],
        ),
        (
            "B4".to_string(),
            vec!["上海".to_string(), "天津".to_string()],
        ),
        (
            "B5".to_string(),
            vec!["杭州".to_string(), "大连".to_string()],
        ),
    ]
    .iter()
    .cloned()
    .collect();

    let broadcasts = get_broadcast(broadcast_map);
    println!("最少选择的广播电台: {:?}", broadcasts);
}
