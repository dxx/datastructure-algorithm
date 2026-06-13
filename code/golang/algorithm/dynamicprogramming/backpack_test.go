package dynamicprogramming

import (
    "fmt"
    "testing"
)

func TestMaxValue(t *testing.T) {
    // 物品重量(kg)
    w := []int{
        1, 2, 1,
    }
    // 物品价值
    v := []int{
        500, 5000, 3000,
    }
    // 背包容量
    c := 3
    max := findMaxValue(w, v, c)
    fmt.Printf("最大价值总和为: %d\n", max)
}

func TestMaxValue2(t *testing.T) {
    // 物品重量(kg)
    w := []int{
        1, 2, 1,
    }
    // 物品价值
    v := []int{
        500, 5000, 3000,
    }
    // 背包容量
    c := 3
    max := findMaxValue2(w, v, c)
    fmt.Printf("最大价值总和为: %d\n", max)
}
