package sparsearray

import (
    "fmt"
    "testing"
)

func TestSparseArray(t *testing.T) {
    // 定义一个二维数组
    var array [5][5]int
    // 初始化 3，6， 1，5
    array[0][2] = 3
    array[1][3] = 6
    array[2][1] = 1
    array[3][3] = 5

    fmt.Println("原二维数组：")
    printArray(array)

    // 转成稀疏数组
    sparseArray := toSparseArray(array)

    fmt.Println("转换后的稀疏数组：")
    printSparseArray(sparseArray)

    // 存储稀疏数组
    storageSparseArray(sparseArray)

    // 读取稀疏数组
    sparseArray = readSparseArray()
    fmt.Println("读取的稀疏数组：")
    printSparseArray(sparseArray)

    // 转成二维数组
    array = toArray(sparseArray)
    fmt.Println("转换后的二维数组：")
    printArray(array)
}
