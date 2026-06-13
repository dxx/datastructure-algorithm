package horsetreads

import (
    "fmt"
    "time"
)

// 马踏棋盘算法
// 国际象棋的棋盘为 8x8 的方格棋盘，现将“马”放在任意指定的方格中，
// 按照“马”走棋的规则将“马”进行移动，要求每个方格只能进入一次，最
// 终使得“马”走遍棋盘 64 个方格

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

// 移动
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

// 骑士周游算法
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
