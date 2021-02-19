package recursion

import (
    "fmt"
    "testing"
)

func TestWalk(t *testing.T) {
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
    for _, nums := range miGongMap {
        for _, anInt := range nums {
            fmt.Printf("%d ", anInt)
        }
        fmt.Println()
    }

    // 开始探路,起点为 1, 1
    walk(miGongMap, 1, 1)

    fmt.Printf("探路之后:\n")
    for _, nums := range miGongMap {
        for _, anInt := range nums {
            fmt.Printf("%d ", anInt)
        }
        fmt.Println()
    }
}

func TestWalk2(t *testing.T) {
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
    for _, nums := range miGongMap {
        for _, anInt := range nums {
            fmt.Printf("%d ", anInt)
        }
        fmt.Println()
    }

    // 开始探路,起点为 1, 1
    walk2(miGongMap, 1, 1)

    fmt.Printf("探路之后:\n")
    for _, nums := range miGongMap {
        for _, anInt := range nums {
            fmt.Printf("%d ", anInt)
        }
        fmt.Println()
    }
}
