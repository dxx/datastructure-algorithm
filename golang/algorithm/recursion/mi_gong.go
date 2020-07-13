package main

import "fmt"

// 迷宫回溯
// 假设一个二维数组中 0 表示通道，1 表示为墙壁，求一个到终点的最短路径
// 思路分析
// 使用递归算法，判断小球是否能够继续往下走，如果可以走就继续，不可用走就退回寻找其它通道。
// 需要指定一个走路策略即走路的方向是从下 -> 右 -> 左 -> 上，还是从上 -> 右 -> 左 -> 下。
// 测试每一种走路策略，对比找出最少的步骤就是最短路径。

// 终点
const x = 6
const y = 6

// 走路函数
// 如果小球走到 6,6 的位置，表示已经走完
// 假设 0-未走过，1-墙，2-走过，3-走不通
// 规定策略：下 -> 右 -> 上 -> 左
// 如果走不通再回溯
func walk(miGongMap [][]int, i, j int) bool {
    // 到达终点
    if miGongMap[x][y] == 2 {
        return true
    }
    if miGongMap[i][j] == 0 {
        // 设置为已走过
        miGongMap[i][j] = 2
        // 下 -> 右 -> 上 -> 左
        if walk(miGongMap, i + 1, j) { // 向下走
            return true
        } else if walk(miGongMap, i, j + 1) { // 向右走
            return true
        } else if walk(miGongMap, i - 1, j) { // 向上走
            return true
        } else if walk(miGongMap, i, j - 1) { // 向左走
            return true
        } else {
            // 走过但路不通
            miGongMap[i][j] = 3
        }
    }
    // 可能为 1, 2 , 3
    return false
}

// 规定策略：上 -> 右 -> 下 -> 左
func walk2(miGongMap [][]int, i, j int) bool {
    // 到达终点
    if miGongMap[x][y] == 2 {
        return true
    }
    if miGongMap[i][j] == 0 {
        // 设置为已走过
        miGongMap[i][j] = 2
        // 上 -> 右 -> 下 -> 左
        if walk2(miGongMap, i - 1, j) { // 向上走
            return true
        } else if walk2(miGongMap, i, j + 1) { // 向右走
            return true
        } else if walk2(miGongMap, i + 1, j) { // 向下走
            return true
        } else if walk2(miGongMap, i, j - 1) { // 向左走
            return true
        } else {
            // 走过但路不通
            miGongMap[i][j] = 3
        }
    }
    // 可能为 1, 2 , 3
    return false
}

func main() {
    // 初始化地图，0 表示通道，1 表示墙
    miGongMap := [][]int{
        {1, 1, 1, 1, 1, 1 , 1, 1},
        {1, 0, 0, 0, 0, 0 , 0, 1},
        {1, 0, 0, 0, 0, 0 , 0, 1},
        {1, 1, 1, 0, 0, 0 , 0, 1},
        {1, 0, 0, 0, 0, 0 , 0, 1},
        {1, 0, 0, 0, 0, 0 , 0, 1},
        {1, 0, 0, 0, 0, 0 , 0, 1},
        {1, 1, 1, 1, 1, 1 , 1, 1},
    }
    fmt.Printf("探路之前:\n")
    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            fmt.Printf("%d ", miGongMap[i][j])
        }
        fmt.Println()
    }

    // 开始探路,起点为 1, 1
    walk2(miGongMap, 1, 1)

    fmt.Printf("探路之后:\n")
    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            fmt.Printf("%d ", miGongMap[i][j])
        }
        fmt.Println()
    }
}
