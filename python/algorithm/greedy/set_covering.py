"""
贪心算法解决集合覆盖，选择最少的广播电台的问题

存在下面需要付费的广播台以及广播台信号可以覆盖的地区。如何选择最少的广播台，让所有地区都接收到信息

 | 广播台 | 覆盖地区         |
 | ------ | -------------- |
 | B1     | 北京，上海，天津 |
 | B2     | 广州，北京，深圳 |
 | B3     | 成都，上海，杭州 |
 | B4     | 上海，天津      |
 | B5     | 杭州，大连      |
"""


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


def main() -> None:
    broadcast_map = {
        "B1": ["北京", "上海", "天津"],
        "B2": ["广州", "北京", "深圳"],
        "B3": ["成都", "上海", "杭州"],
        "B4": ["上海", "天津"],
        "B5": ["杭州", "大连"],
    }
    broadcasts = get_broadcast(broadcast_map)
    print("最少选择的广播电台: " + ",".join(broadcasts))


if __name__ == "__main__":
    main()
