## 贪心算法

>各种语言实现代码：[Go](./golang/algorithm/greedy)   [Java](./java/algorithm/src/com/mcx/greedy)   [JavaScript](./javascript/algorithm/greedy)   [Rust](./rust/algorithm/src/greedy)
>
>默认使用 **Go** 语言实现。

### 简介

贪心算法（贪婪算法）是指在对问题求解时，总是做出在当前看来是最好的选择。也就是说，不从整体最优上加以考虑，算法得到的是在某种意义上的局部最优解。

贪心算法不是对所有问题都能得到整体最优解，关键是贪心策略的选择，选择的贪心策略必须具备无后效性，即某个状态以前的过程不会影响以后的状态，只与当前状态有关。

### 集合覆盖

存在下面需要付费的广播台以及广播台信号可以覆盖的地区。如何选择最少的广播台，让所有地区都能接收到信号

| 广播台 | 覆盖地区         |
| ------ | ---------------- |
| B1     | 北京，上海，天津 |
| B2     | 广州，北京，深圳 |
| B3     | 成都，上海，杭州 |
| B4     | 上海，天津       |
| B5     | 杭州，大连       |

使用穷举法实现，列出每个可能的广播台的集合，这被称为幂集。假设总的有 n 个广播台，则广播台的组合总共有 2^n -1 个。

使用贪心算法，则可以得到非常接近的解，并且效率高。使用贪心算法求解思路如下：

1. 遍历所有广播电台，将所有的区域放入一个不重复的数组或集合中。

2. 找出一个覆盖最多地区的电台（此电台可能包含一些已经覆盖的地区，但是没有关系）。

3. 将覆盖最多地区的电台保存起来，放入数组或者集合中，去掉该电台，下一次不参与比较。

   同时将已经覆盖的区域从所有区域的数组或集合中移除。

4. 重复步骤 1，直到所有地区被覆盖。

代码实现：

```go
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
```

```go
func TestGetBroadcast(t *testing.T) {
    broadcastMap := map[string][]string {
        "B1": {"北京", "上海", "天津"},
        "B2": {"广州", "北京", "深圳"},
        "B3": {"成都", "上海", "杭州"},
        "B4": {"上海", "天津"},
        "B5": {"杭州", "大连"},
    }
    broadcasts := getBroadcast(broadcastMap)
    fmt.Printf("最少选择的广播电台: %v\n", broadcasts)
}
```

运行：

```shell
golang/algorithm>go test -v -run ^TestGetBroadcast$ ./greedy
=== RUN   TestGetBroadcast
最少选择的广播电台: [B1 B2 B3 B5]
```
