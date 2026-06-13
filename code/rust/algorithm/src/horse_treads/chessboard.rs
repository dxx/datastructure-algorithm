use std::collections::VecDeque;
use std::fmt::{Debug, Formatter, Result};
use std::time::{SystemTime, UNIX_EPOCH};

/// 马踏棋盘算法
/// 国际象棋的棋盘为 8x8 的方格棋盘，现将“马”放在任意指定的方格中，
/// 按照“马”走棋的规则将“马”进行移动，要求每个方格只能进入一次，最
/// 终使得“马”走遍棋盘 64 个方格

/// 棋盘
pub struct Chessboard {
    row: usize,           // 表示棋盘的行数
    col: usize,           // 表示棋盘的列数
    visited: Vec<bool>,   // 标记点是否被访问过
    steps: Vec<Vec<u32>>, // 存放步数
    finished: bool,       // 表示是否已经走完
}

/// 位置
pub struct Point {
    x: usize, // X 下标
    y: usize, // Y 下标
}

impl Point {
    pub fn new(x: usize, y: usize) -> Self {
        Point { x, y }
    }
}

impl Debug for Point {
    fn fmt(&self, f: &mut Formatter) -> Result {
        write!(f, "x:{}, y:{}", self.x, self.y)
    }
}

impl Chessboard {
    pub fn new(row: usize, col: usize) -> Self {
        let length = row * col;
        let visited = vec![false; length];
        let steps = vec![vec![0; col]; row];
        Chessboard {
            row,
            col,
            visited,
            steps,
            finished: false,
        }
    }

    /// 移动
    /// startX: 起始横坐标。从 0 开始
    /// startY: 起始纵坐标。从 0 开始
    pub fn move_chessboard(&mut self, start_x: usize, start_y: usize) {
        println!("开始走马踏棋");
        let start_time = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_nanos();

        self.traversal(start_x, start_y, 1);

        println!("马踏棋结束");
        let end_time = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_nanos();

        println!(
            "耗时: {}s", (end_time - start_time) as f64 / 1000000000 as f64
        );

        // 打印所有走过的步数
        for i in 0..self.row {
            for j in 0..self.col {
                print!("{:2} ", self.steps[i][j]);
            }
            println!();
        }
    }

    /// 骑士周游算法
    fn traversal(&mut self, x: usize, y: usize, step: u32) {
        // 将当前位置标记已访问。y = 4, col = 8, x = 4 => 4 * 8 + 4 = 36
        self.visited[y * self.col + x] = true;
        // 记录步数
        self.steps[y][x] = step;
        // 获取下一步可以走的所有位置
        let mut points = self.next_points(&Point::new(x, y));

        // 排序优化，优先选择下一步最少可走数目的位置，减少回溯，体现出贪心算法的特点
        // 去掉此方法，算法的耗时会很久，根据走法不同，结果也会不同
        points.make_contiguous().sort_by(|point1, point2| {
            self.next_points(point1)
                .len()
                .cmp(&self.next_points(point2).len())
        });

        while points.len() != 0 {
            // 取出第一个点
            let p = points.pop_front().unwrap();
            // 该点未被访问过
            if !self.visited[p.y * self.col + p.x] {
                // 继续往下走
                self.traversal(p.x, p.y, step + 1);
            }
        }
        // 比较已经走的步数和应该走的步数，如果不相等表示没有走完，将棋盘当前的位置重置
        if (step as usize) < self.row * self.col && !self.finished {
            self.visited[y * self.col + x] = false;
            self.steps[y][x] = 0;
        } else {
            self.finished = true;
        }
    }

    /// 获取当位置的下一步可走位置的集合，最多可有 8 个位置
    fn next_points(&self, point: &Point) -> VecDeque<Point> {
        let mut points = VecDeque::new();
        // 判断是否可以走 1 的位置
        if point.x as i32 - 1 >= 0 && point.y as i32 - 2 >= 0 {
            points.push_back(Point::new(point.x - 1, point.y - 2));
        }
        // 判断是否可以走 2 的位置
        if point.x + 1 < self.col && point.y as i32 - 2 >= 0 {
            points.push_back(Point::new(point.x + 1, point.y - 2));
        }
        // 判断是否可以走 3 的位置
        if point.x + 2 < self.col && point.y as i32 - 1 >= 0 {
            points.push_back(Point::new(point.x + 2, point.y - 1));
        }
        // 判断是否可以走 4 的位置
        if point.x + 2 < self.col && point.y + 1 < self.row {
            points.push_back(Point::new(point.x + 2, point.y + 1));
        }
        // 判断是否可以走 5 的位置
        if point.x + 1 < self.col && point.y + 2 < self.row {
            points.push_back(Point::new(point.x + 1, point.y + 2));
        }
        // 判断是否可以走 6 的位置
        if point.x as i32 - 1 >= 0 && point.y + 2 < self.row {
            points.push_back(Point::new(point.x - 1, point.y + 2));
        }
        // 判断是否可以走 7 的位置
        if point.x as i32 - 2 >= 0 && point.y + 1 < self.row {
            points.push_back(Point::new(point.x - 2, point.y + 1));
        }
        // 判断是否可以走 8 的位置
        if point.x as i32 - 2 >= 0 && point.y as i32 - 1 >= 0 {
            points.push_back(Point::new(point.x - 2, point.y - 1));
        }
        points
    }
}

#[test]
fn test_chessboard() {
    let mut chessboard = Chessboard::new(8, 8);
    // 从 4,4 的位置开始走
    chessboard.move_chessboard(4, 4);
}
