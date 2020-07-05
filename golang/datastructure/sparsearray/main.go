package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// 稀疏数组
// 在矩阵中，若数值为 0 的元素数目远远多于非 0 元素的数目，并且非 0 元素分布没有规律时，则称该矩阵为稀疏矩阵
// 稀疏数组可以看做是一个压缩的数组，稀疏数组的好处有：
// 原数组中存在大量的无效数据，占据了大量的存储空间，真正有用的数据却少之又少
// 压缩存储可以节省存储空间以避免资源的不必要的浪费，在数据序列化到磁盘时，压缩存储可以提高 IO 效率

// 二维数组
/*
0	0	3	0	0
0	0	0	6	0
0	1	0	0	0
0	0	0	5	0
0	0	0	0	0
*/
// 稀疏数组
/*
row	col	val
5	5	0
0	2	3
1	3	6
2	1	1
3	3	5
*/

// 二维数组转稀疏数组
func toSparseArray(array [5][5]int) [][3]int {
    // 统计有多少个不等于 0 的数
    //var count = 0
    //for i := 0; i < len(sequence); i++ {
    //    for j := 0; j < len(sequence[i]); j++ {
    //        if sequence[i][j] != 0 {
    //            count++
    //        }
    //    }
    //}

    // 定义稀疏数组

    // Go 语言这样写无法编译通过
    // var sparseArray [count + 1][]int
    // 使用切片来定义
    var sparseArray = make([][3]int, 0)
    sparseArray = append(sparseArray, [3]int{5, 5, 0})

    for i := 0; i < len(array); i++ {
        for j := 0; j < len(array[i]); j++ {
            if array[i][j] != 0 {
                // 保存 row, col, val
                sparseArray = append(sparseArray, [3]int{i, j, array[i][j]})
            }
        }
    }

    return sparseArray
}

// 稀疏数组转二维数组
func toArray(sparseArray [][3]int) [5][5]int {
    var array [5][5]int

    // 从稀疏数组中第二行开始读取数据
    for i := 1; i < len(sparseArray); i++ {
        row := sparseArray[i][0]
        col := sparseArray[i][1]
        val := sparseArray[i][2]
        array[row][col] = val
    }

    return array
}

var sparseArrayFileName = "./sparse.data"

// 存储稀疏数组
func storageSparseArray(sparseArray [][3]int) {
    file, err := os.Create(sparseArrayFileName)
    defer file.Close()

    if err != nil {
        fmt.Println("创建文件 sparse.data 错误:", err)
    }
    // 存储矩阵格式
    for i := 0; i < len(sparseArray); i++ {
        content := ""
        for j := 0; j < 3; j++ {
            content += strconv.Itoa(sparseArray[i][j]) + "\t"
        }
        // 行分隔符
        content += "\n"
        _, err = file.WriteString(content)
        if err != nil {
            fmt.Println("写入内容错误:", err)
        }
    }
}

// 读取稀疏数组
func readSparseArray() [][3]int {
    file, err := os.Open(sparseArrayFileName)
    defer file.Close()

    if err != nil {
        fmt.Println("打开文件 sparse.data 错误:", err)
    }
    sparseArray := make([][3]int, 0)
    reader := bufio.NewReader(file)
    for {
        // 分行读取
        content, err := reader.ReadString('\n')
        if err == io.EOF {
            break
        }
        arr := strings.Split(content, "\t")
        row, _ := strconv.Atoi(arr[0])
        col, _ := strconv.Atoi(arr[1])
        val, _ := strconv.Atoi(arr[2])
        sparseArray = append(sparseArray, [3]int{row, col, val})
    }
    return sparseArray
}

func printArray(array [5][5]int) {
    for i := 0; i < len(array); i++ {
        for j := 0; j < len(array[i]); j++ {
            fmt.Print(array[i][j], "\t")
        }
        fmt.Print("\n")
    }
}

func printSparseArray(sparseArray [][3]int) {
    for i := 0; i < len(sparseArray); i++ {
        for j := 0; j < 3; j++ {
            fmt.Print(sparseArray[i][j], "\t")
        }
        fmt.Print("\n")
    }
}

func main() {
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
