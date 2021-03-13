## 马踏棋盘算法

>各种语言实现代码：[Go](./golang/algorithm/horsetreads)   [Java](./java/algorithm/src/com/mcx/horsetreads)   [JavaScript](./javascript/algorithm/horsetreads)   [Rust](./rust/algorithm/src/horse_treads)
>
>默认使用 **Go** 语言实现。

### 简介

国际象棋的棋盘为 8x8 的方格棋盘，现将“马”放在任意指定的方格中，按照“马”走棋的规则将“马”进行移动，要求每个方格只能进入一次，最终使得“马”走遍棋盘 64 个方格。

![algorithm_chessboard](https://dxx.github.io/static-resource/datastructure-algorithm/images/algorithm_chessboard.png)

### 实现

骑士周游问题就可以使用图的深度优先搜索（回溯）来解决。

解题思路：

* 用一个二维数组表示棋盘。
* 将当前位置标记为已访问，当前位置对应的值记为当前是第几步，根据当前位置计算出马可以走的所有位置，最多八个位置。
* 遍历所有的位置，然后取出一个位置，递归进行，如果当前位置没有下一步可以走的位置就回溯。
* 比较已经走的步数和应该走的步数，如果不相等表示没有走完，将棋盘当前的位置重置。

> 注意：马从不同的位置开始走，会得到不同的结果，效率也会有影响，使用贪心算法进行优化

代码实现：

定义棋盘和代表一个位棋盘上的一个点

```go
// 棋盘
type Chessboard struct {
    row      int     // 表示棋盘的行数
    col      int     // 表示棋盘的列数
    visited  []bool  // 标记点是否被访问过
    steps    [][]int // 存放步数
    finished bool    // 表示是否已经走完
}

// 位置
type Point struct {
    x int // X 下标
    y int // Y 下标
}

func (point *Point) String() string {
    return fmt.Sprintf("x:%d, y:%d", point.x, point.y)
}

func NewChessboard(row, col int) *Chessboard {
    length := row * col
    visited := make([]bool, length)
    steps := make([][]int, row)
    for i := 0; i < len(steps); i++ {
        steps[i] = make([]int, col)
    }
    return &Chessboard{row, col, visited, steps, false}
}
```

以上面图中的点为例，根据当前位置找出下一步可以走的所有位置：

```go
// 获取当位置的下一步可走位置的集合，最多可有 8 个位置
func (chessboard *Chessboard) nextPoints(point *Point) []*Point {
    points := make([]*Point, 0)
    // 判断是否可以走 1 的位置
    if point.x-1 >= 0 && point.y-2 >= 0 {
        points = append(points, &Point{x: point.x - 1, y: point.y - 2})
    }
    // 判断是否可以走 2 的位置
    if point.x+1 < chessboard.col && point.y-2 >= 0 {
        points = append(points, &Point{x: point.x + 1, y: point.y - 2})
    }
    // 判断是否可以走 3 的位置
    if point.x+2 < chessboard.col && point.y-1 >= 0 {
        points = append(points, &Point{x: point.x + 2, y: point.y - 1})
    }
    // 判断是否可以走 4 的位置
    if point.x+2 < chessboard.col && point.y+1 < chessboard.row {
        points = append(points, &Point{x: point.x + 2, y: point.y + 1})
    }
    // 判断是否可以走 5 的位置
    if point.x+1 < chessboard.col && point.y+2 < chessboard.row {
        points = append(points, &Point{x: point.x + 1, y: point.y + 2})
    }
    // 判断是否可以走 6 的位置
    if point.x-1 >= 0 && point.y+2 < chessboard.row {
        points = append(points, &Point{x: point.x - 1, y: point.y + 2})
    }
    // 判断是否可以走 7 的位置
    if point.x-2 >= 0 && point.y+1 < chessboard.row {
        points = append(points, &Point{x: point.x - 2, y: point.y + 1})
    }
    // 判断是否可以走 8 的位置
    if point.x-2 >= 0 && point.y-1 >= 0 {
        points = append(points, &Point{x: point.x - 2, y: point.y - 1})
    }
    return points
}
```

定义排序方法用来优化选择下一步棋子的走的位置：

```go
// 对位置进行排序，根据下一步可走的位置数量从小到大排序
func (chessboard *Chessboard) sort(points []*Point) {
    if points == nil {
        return
    }
    for i := 1; i < len(points); i++ {
        insertIndex := i - 1
        insertValue := points[i]
        nextPoints := chessboard.nextPoints(points[i])
        for insertIndex >= 0 && len(chessboard.nextPoints(points[insertIndex])) > len(nextPoints) {
            points[insertIndex+1] = points[insertIndex]
            insertIndex--
        }
        if insertIndex+1 != i {
            points[insertIndex+1] = insertValue
        }
    }
}
```

> 这里使用插入排序，也可以使用其它排序算法，或者语言核心库中的排序方式来实现

骑士周游核心算法：

```go
func (chessboard *Chessboard) traversal(x, y, step int) {
    // 将当前位置标记已访问。y = 4, col = 8, x = 4 => 4 * 8 + 4 = 36
    chessboard.visited[y*chessboard.col+x] = true
    // 记录步数
    chessboard.steps[y][x] = step

    // 获取下一步可以走的所有位置
    points := chessboard.nextPoints(&Point{x, y})

    // 排序优化，优先选择下一步最少可走数目的位置，减少回溯，体现出贪心算法的特点
    chessboard.sort(points) // 去掉此方法，算法的耗时会很久，根据走法不同，结果也会不同

    for len(points) != 0 {
        // 取出第一个点
        p := points[0]
        // 移除第一个位置
        points = points[1:]
        // 该点未被访问过
        if chessboard.visited[p.y*chessboard.col+p.x] == false {
            // 继续往下走
            chessboard.traversal(p.x, p.y, step+1)
        }
    }

    // 比较已经走的步数和应该走的步数，如果不相等表示没有走完，将棋盘当前的位置重置
    if step < chessboard.row*chessboard.col && !chessboard.finished {
        chessboard.visited[y*chessboard.col+x] = false
        chessboard.steps[y][x] = 0
    } else {
        chessboard.finished = true
    }
}
```

走棋子方法：

```go
// startX: 起始横坐标。从 0 开始
// startY: 起始纵坐标。从 0 开始
func (chessboard *Chessboard) Move(startX, startY int) {

    fmt.Printf("开始走马踏棋\n")
    startTime := time.Now().Nanosecond()

    chessboard.traversal(startX, startY, 1)

    fmt.Printf("马踏棋结束\n")
    endTime := time.Now().Nanosecond()
    fmt.Printf("耗时:%fs\n", float32(endTime-startTime)/1000000000)

    // 打印所有走过的步数
    for i := 0; i < chessboard.row; i++ {
        for j := 0; j < chessboard.col; j++ {
            fmt.Printf("%2d ", chessboard.steps[i][j])
        }
        fmt.Println()
    }
}
```

测试代码：

```go
func TestChessboard(t *testing.T) {
    chessboard := NewChessboard(8, 8)
    // 从 4,4 的位置开始走
    chessboard.Move(4, 4)
}
```

运行：

```shell
golang/algorithm>go test -v -run ^TestChessboard$ ./horsetreads
=== RUN   TestChessboard
开始走马踏棋
马踏棋结束
耗时:0.000998s
27 48 11 58 25  4  9  6
12 59 26 53 10  7 24  3
49 28 47 62 57 52  5  8
60 13 56 51 54 63  2 23
29 50 61 46  1 44 19 40
14 35 32 55 64 41 22 43
33 30 37 16 45 20 39 18
36 15 34 31 38 17 42 21
```

### 马踏棋盘游戏

有一个马踏棋盘游戏地址：[http://www.4399.com/flash/146267_2.htm](http://www.4399.com/flash/146267_2.htm)。

使用上述代码将 8, 8 改成 6, 6：

```go
func TestChessboard(t *testing.T) {
    chessboard := NewChessboard(6, 6)
    chessboard.Move(0, 0)
}
```

结果：

```
 1 22  9 26  3 24
10 35  2 23 16 27
21  8 31 36 25  4
32 11 34 17 28 15
 7 20 13 30  5 18
12 33  6 19 14 29
```

得到走棋的步骤，按照步骤就可以通过啦。