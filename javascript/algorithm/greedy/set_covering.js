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

function getBroadcast(broadcasts) {
  let broadcastKeys = [];
  let m = {};

  for (let key in broadcasts) {
    // 将 key 存入数组
    broadcastKeys.push(key);
    for (let j = 0; j < broadcasts[key].length; j++) {
      // 区域作为 key，保证不重复
      m[broadcasts[key][j]] = "";
    }
  }

  // 对数组进行排序，避免每次遍历 broadcasts 顺序不一致，导致求出电台结果不一致
  broadcastKeys.sort();

  let allBroadcasts = [];
  for (let b in m) {
    // 保存所有区域
    allBroadcasts.push(b);
  }

  let selectBroadcasts = [];
  // 覆盖区域最多的电台
  let maxBroadcast = "";
  // 交集
  let tempIntersection = [];
  // 遍历所有的电台，直到所有电台被移除
  while (allBroadcasts.length > 0) {
    maxBroadcast = "";
    for (let i = 0; i < broadcastKeys.length; i++) {
      let key = broadcastKeys[i];
      let value = broadcasts[key];
      // 求出交集
      tempIntersection = value.filter(function (val) { return allBroadcasts.indexOf(val) > -1 });
      // 每次都选择覆盖最多的集合，体现出贪心算法的特点
      if (tempIntersection.length > 0 &&
          (maxBroadcast == "" || tempIntersection.length > value.length)) {
        // 修改覆盖区域最多的电台
        maxBroadcast = key;
      }
    }
    
    if (maxBroadcast != "") {
      // 将电台添加到数组中
      selectBroadcasts.push(maxBroadcast);
      // 将已经覆盖的区域对应的电台移除，不参与下一次比较
      broadcastKeys = deleteElement(broadcastKeys, maxBroadcast)
      for (let j = 0; j < broadcasts[maxBroadcast].length; j++) {
        // 删除已经覆盖的区域
        allBroadcasts = deleteElement(allBroadcasts, broadcasts[maxBroadcast][j])
      }
    }
  }
  return selectBroadcasts;
}

function deleteElement(array, deleteVal) {
  for (let i = 0; i < array.length; i++) {
    if (array[i] == deleteVal) {
      array.splice(i,1);
      // 避免删除重复的数据
      return array
    }
  }
  return array
}

function main() {
  let broadcastMap = {
    "B1": ["北京", "上海", "天津"],
    "B2": ["广州", "北京", "深圳"],
    "B3": ["成都", "上海", "杭州"],
    "B4": ["上海", "天津"],
    "B5": ["杭州", "大连"]
  }
  let broadcasts = getBroadcast(broadcastMap);
  console.log("最少选择的广播电台: " + broadcasts);
}

main();
