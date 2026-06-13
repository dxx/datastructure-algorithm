package greedy

import (
    "sort"
)

// 贪心算法解决集合覆盖，选择最少的广播电台的问题

// 存在下面需要付费的广播台以及广播台信号可以覆盖的地区。如何选择最少的广播台，让所有地区都接收到信息
// | 广播台 | 覆盖地区         |
// | ------ | -------------- |
// | B1     | 北京，上海，天津 |
// | B2     | 广州，北京，深圳 |
// | B3     | 成都，上海，杭州 |
// | B4     | 上海，天津      |
// | B5     | 杭州，大连      |

func getBroadcast(broadcasts map[string][]string) []string {
    broadcastKeys := make([]string, 0)

    m := make(map[string]string)
    for key, val := range broadcasts {
        // 将 key 存入切片
        broadcastKeys = append(broadcastKeys, key)
        for _, broadcast := range val {
            // 区域作为 key，保证不重复
            m[broadcast] = ""
        }
    }

    // 对切片进行排序，避免每次遍历 broadcasts 顺序不一致，导致求出电台结果不一致
    sort.Strings(broadcastKeys)

    allBroadcasts := make([]string, 0)
    for b, _ := range m {
        // 保存所有区域
        allBroadcasts = append(allBroadcasts, b)
    }

    selectBroadcasts := make([]string, 0)

    // 覆盖区域最多的电台
    var maxBroadcast string
    // 交集
    var tempIntersection []string
    // 遍历所有的电台，直到所有电台被移除
    for len(allBroadcasts) > 0 {
        maxBroadcast = ""
        for _, key := range broadcastKeys {
            value := broadcasts[key]
            // 求出交集
            tempIntersection = intersect(value, allBroadcasts)
            // 每次都选择覆盖最多的集合，体现出贪心算法的特点
            if len(tempIntersection) > 0 &&
                (maxBroadcast == "" || len(tempIntersection) > len(value)) {
                // 修改覆盖区域最多的电台
                maxBroadcast = key
            }
        }
        if maxBroadcast != "" {
            // 将电台添加到切片中
            selectBroadcasts = append(selectBroadcasts, maxBroadcast)
            // 将已经覆盖的区域对应的电台移除，不参与下一次比较
            broadcastKeys = deleteElement(broadcastKeys, maxBroadcast)
            for _, b := range broadcasts[maxBroadcast] {
                // 删除已经覆盖的区域
                allBroadcasts = deleteElement(allBroadcasts, b)
            }
        }
    }
    return selectBroadcasts
}

// 求交集
func intersect(slice1, slice2 []string) []string {
    m := make(map[string]int)
    nn := make([]string, 0)
    for _, v := range slice1 {
        m[v] = 1
    }
    for _, v := range slice2 {
        times, _ := m[v]
        if times == 1 {
            nn = append(nn, v)
        }
    }
    return nn
}

func deleteElement(slice []string, deleteVal string) []string {
    for i := 0; i < len(slice); i++ {
        if slice[i] == deleteVal {
            slice = append(slice[:i], slice[i+1:]...)
            // 避免删除重复的数据
            return slice
        }
    }
    return slice
}
