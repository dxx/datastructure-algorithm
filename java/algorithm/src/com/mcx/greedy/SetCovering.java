package com.mcx.greedy;

import java.util.*;

/**
 * 贪心算法解决集合覆盖，选择最少的广播电台的问题
 *
 * 存在下面需要付费的广播台以及广播台信号可以覆盖的地区。如何选择最少的广播台，让所有地区都接收到信息
 * | 广播台 | 覆盖地区         |
 * | ------ | -------------- |
 * | B1     | 北京，上海，天津 |
 * | B2     | 广州，北京，深圳 |
 * | B3     | 成都，上海，杭州 |
 * | B4     | 上海，天津      |
 * | B5     | 杭州，大连      |
 */
public class SetCovering {

    public static List<String> getBroadcast(Map<String, List<String>> broadcasts) {
        List<String> broadcastKeys = new ArrayList<>();

        Map<String, String> m = new HashMap<>();
        for (String key : broadcasts.keySet()) {
            // 将 key 存入集合
            broadcastKeys.add(key);
            for (String val : broadcasts.get(key)) {
                // 区域作为 key，保证不重复
                m.put(val, "");
            }
        }

        // 对切片进行排序，避免每次遍历 broadcasts 顺序不一致，导致求出电台结果不一致
        Collections.sort(broadcastKeys);

        // 保存所有区域
        List<String> allBroadcasts = new ArrayList<>(m.keySet());

        List<String> selectBroadcasts = new ArrayList<>();

        // 覆盖区域最多的电台
        String maxBroadcast;
        // 交集
        List<String> tempIntersection;
        // 遍历所有的电台，直到所有电台被移除
        while (allBroadcasts.size() > 0) {
            maxBroadcast = "";
            for (String key : broadcastKeys) {
                List<String> value = broadcasts.get(key);
                // 求出交集
                value.retainAll(allBroadcasts);
                tempIntersection = value;
                // 每次都选择覆盖最多的集合，体现出贪心算法的特点
                if (tempIntersection.size() > 0 &&
                        (maxBroadcast.equals("") || tempIntersection.size() > value.size())) {
                    // 修改覆盖区域最多的电台
                    maxBroadcast = key;
                }
            }
            if (!maxBroadcast.equals("")) {
                // 将电台添加到切片中
                selectBroadcasts.add(maxBroadcast);
                // 将已经覆盖的区域对应的电台移除，不参与下一次比较
                broadcastKeys.remove(maxBroadcast);
                for (String b : broadcasts.get(maxBroadcast)) {
                    // 删除已经覆盖的区域
                    allBroadcasts.remove(b);
                }
            }
        }
        return selectBroadcasts;
    }

    /**
     * 求交集
     */
    public static List<String> intersect(List<String> list1, List<String> list2) {
        list1.retainAll(list2);
        return list1;
    }

    public static void main(String[] args) {
        Map<String, List<String>> broadcastMap = new HashMap<>();
        broadcastMap.put("B1", new ArrayList<>(Arrays.asList("北京", "上海", "天津")));
        broadcastMap.put("B2", new ArrayList<>(Arrays.asList("广州", "北京", "深圳")));
        broadcastMap.put("B3", new ArrayList<>(Arrays.asList("成都", "上海", "杭州")));
        broadcastMap.put("B4", new ArrayList<>(Arrays.asList("上海", "天津")));
        broadcastMap.put("B5", new ArrayList<>(Arrays.asList("杭州", "大连")));

        List<String> broadcasts = getBroadcast(broadcastMap);
        System.out.printf("最少选择的广播电台: %s", broadcasts);
    }
}
