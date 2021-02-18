## 稀疏数组

> 各种语言实现代码：[Go](./golang/datastructure/sparsearray)   [Java](./java/datastructure/src/com/mcx/sparsearray)   [JavaScript](./javascript/datastructure/sparsearray)   [Rust](./rust/datastructure/src/sparse_array)
>
> 默认使用 **Go** 语言实现。

### 简介

在二维数组中，如果值为 0 的元素数目远远大于非 0 元素的数目，并且非 0 元素的分布没有规律，则该数组被称为稀疏数组。如果非 0 元素数目占大多数，则称该数组为稠密数组。数组的稠密度指的是非零元素的总数比上数组所有元素的总数。

下图是一个 0 值远大于非 0 值的二维数组

![data_structure_sparsearray_01](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/data_structure_sparsearray_01.png)

稀疏数组可以看做是一个压缩的数组，稀疏数组的好处有：

* 原数组中存在大量的无用的数据，占据了存储空间，真正有用的数据却很少
* 压缩后存储可以节省存储空间，在数据序列化到磁盘时，压缩存储可以提高 IO 效率

采用稀疏数组的存储方式为第一行存储原始数据总行数，总列数，默认值 0，接下来每一行都存储非0数所在行，所在列，和具体值。上图中的二维数组转成稀疏数组后如下：

![data_structure_sparsearray_02](https://code-mcx.github.io/static-resource/datastructure-algorithm/images/data_structure_sparsearray_02.png)



下面使用稀疏数组存储上述的二维数组，把稀疏数组保存在文件中，并且可以重新恢复成二维数组。

### 创建二维数组并初始化

```go
func printArray(array [5][5]int) {
    for i := 0; i < len(array); i++ {
        for j := 0; j < len(array[i]); j++ {
            fmt.Print(array[i][j], "\t")
        }
        fmt.Print("\n")
    }
}
```

```go
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
}
```

运行：

```shell
golang/datastructure>go test -v -run ^TestSparseArray$ ./sparsearray
=== RUN   TestSparseArray
原二维数组：
0       0       3       0       0
0       0       0       6       0
0       1       0       0       0
0       0       0       5       0
0       0       0       0       0
```

### 二维数组转稀疏数组

```go
// 二维数组转稀疏数组
func toSparseArray(array [5][5]int) [][3]int {
    // Go 语言这样写无法编译通过
    // var sparseArray [count + 1][]int
    // 使用切片来定义
    var sparseArray = make([][3]int, 0)
    sparseArray = append(sparseArray, [3]int{ 5, 5, 0})

    for i := 0; i < len(array); i++ {
        for j := 0; j < len(array[i]); j++ {
            if array[i][j] != 0 {
                // 保存 row, col, val
                sparseArray = append(sparseArray, [3]int{ i, j, array[i][j]})
            }
        }
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
```

```go
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
}
```

运行：

```shell
golang/datastructure>go test -v -run ^TestSparseArray$ ./sparsearray
=== RUN   TestSparseArray
原二维数组：
0       0       3       0       0
0       0       0       6       0
0       1       0       0       0
0       0       0       5       0
0       0       0       0       0
转换后的稀疏数组：
5       5       0
0       2       3
1       3       6
2       1       1
3       3       5
```

### 存储和读取稀疏数组

```go
var sparseArrayFileName = "./sparse.data"
// 存储稀疏数组
func storageSparseArray(sparseArray [][3]int) {
    file, err := os.Create(sparseArrayFileName)
    defer file.Close()

    if err != nil {
        fmt.Println("创建文件 sparse.data 错误:", err)
        return
    }
    // 存储矩阵格式
    for i := 0 ; i < len(sparseArray); i++ {
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
        return nil
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
        sparseArray = append(sparseArray, [3]int { row, col, val})
    }
    return sparseArray
}
```

```go
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
}
```

运行：

```shell
golang/datastructure>go test -v -run ^TestSparseArray$ ./sparsearray
=== RUN   TestSparseArray
...
读取的稀疏数组：
5       5       0
0       2       3
1       3       6
2       1       1
3       3       5
```

运行以上代码后，打开 `sparse.data` 文件，内容如下：

```
5	5	0	
0	2	3	
1	3	6	
2	1	1	
3	3	5	
```

### 稀疏数组转二维数组

```go
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
```

```go
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
```

运行：

```shell
golang/datastructure>go test -v -run ^TestSparseArray$ ./sparsearray
=== RUN   TestSparseArray
...
转换后的二维数组：
0       0       3       0       0
0       0       0       6       0
0       1       0       0       0
0       0       0       5       0
0       0       0       0       0
```