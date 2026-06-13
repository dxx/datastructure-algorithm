## 贪心算法

>各种语言实现代码：[Go](../code/golang/algorithm/greedy)   [Java](../code/java/algorithm/src/com/dxx/greedy)   [JavaScript](../code/javascript/algorithm/greedy)   [TypeScript](../code/typescript/algorithm/greedy)   [Python](../code/python/algorithm/greedy)   [Rust](../code/rust/algorithm/src/greedy)
>
>默认使用 **Python** 语言实现。

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

```python
def get_broadcast(broadcasts: dict[str, list[str]]) -> list[str]:
    broadcast_keys = []
    area_map = {}

    for key, value in broadcasts.items():
        # 将 key 存入数组
        broadcast_keys.append(key)
        for area in value:
            # 区域作为 key，保证不重复
            area_map[area] = ""

    # 对数组进行排序，避免每次遍历 broadcasts 顺序不一致，导致求出电台结果不一致
    broadcast_keys.sort()
    # 保存所有区域
    all_broadcasts = list(area_map.keys())
    select_broadcasts = []

    # 遍历所有的电台，直到所有电台被移除、所有区域被覆盖
    while len(all_broadcasts) > 0:
        # 覆盖区域最多的电台
        max_broadcast = ""
        # 交集
        max_intersection_length = 0
        for key in broadcast_keys:
            value = broadcasts[key]
            # 求出交集
            temp_intersection = [area for area in value if area in all_broadcasts]
            # 每次都选择覆盖最多的集合，体现出贪心算法的特点
            if len(temp_intersection) > max_intersection_length:
                # 修改覆盖区域最多的电台
                max_broadcast = key
                max_intersection_length = len(temp_intersection)

        if max_broadcast != "":
            # 将电台添加到数组中
            select_broadcasts.append(max_broadcast)
            # 将已经覆盖的区域对应的电台移除，不参与下一次比较
            broadcast_keys = delete_element(broadcast_keys, max_broadcast)
            for area in broadcasts[max_broadcast]:
                # 删除已经覆盖的区域
                all_broadcasts = delete_element(all_broadcasts, area)
        else:
            break
    return select_broadcasts


def delete_element(array: list[str], delete_val: str) -> list[str]:
    for i, value in enumerate(array):
        if value == delete_val:
            del array[i]
            # 避免删除重复的数据
            return array
    return array
```

测试代码：

```python
class Test(unittest.TestCase):
    
    def test_get_broadcast(self):
        broadcast_map = {
            "B1": ["北京", "上海", "天津"],
            "B2": ["广州", "北京", "深圳"],
            "B3": ["成都", "上海", "杭州"],
            "B4": ["上海", "天津"],
            "B5": ["杭州", "大连"],
        }
        broadcasts = get_broadcast(broadcast_map)
        print("最少选择的广播电台: " + ",".join(broadcasts))
```

运行：

```shell
❯ python -m unittest test_set_covering.Test.test_get_broadcast
最少选择的广播电台: B1,B2,B3,B5
```
