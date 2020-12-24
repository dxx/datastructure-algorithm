package com.mcx.horsetreads;

import java.util.Comparator;
import java.util.LinkedList;

/**
 * 马踏棋盘算法
 * 国际象棋的棋盘为 8x8 的方格棋盘，现将“马”放在任意指定的方格中，
 * 按照“马”走棋的规则将“马”进行移动，要求每个方格只能进入一次，最
 * 终使得“马”走遍棋盘 64 个方格
 *
 *
 * 棋盘
 */
public class Chessboard {

    public int row; // 表示棋盘的行数

    public int col; // 表示棋盘的列数

    public boolean[] visited; // 标记点是否被访问过

    public int[][] steps; // 存放步数

    public boolean finished; // 表示是否已经走完

    /**
     * 位置
     */
    public static class Point {

        public int x; // X 下标

        public int y; // Y 下标

        public Point(int x, int y) {
            this.x = x;
            this.y = y;
        }

        public String toString() {
            return String.format("x:%d, y:%d", this.x, this.y);
        }
    }

    public Chessboard(int row, int col) {
        int length = row * col;
        this.row = row;
        this.col = col;
        this.visited = new boolean[length];
        this.steps = new int[row][col];
        this.finished = false;
    }

    /**
     * 移动
     * @param startX 起始横坐标。从 0 开始
     * @param startY 起始纵坐标。从 0 开始
     */
    public void move(int startX, int startY) {
        System.out.println("开始走马踏棋");
        long startTime = System.nanoTime();

        this.traversal(startX, startY, 1);

        System.out.println("马踏棋结束");
        long endTime = System.nanoTime();
        System.out.println(String.format("耗时:%fs", (float)(endTime - startTime) / 1000000000));

        // 打印所有走过的步数
        for (int i = 0; i < this.row; i++) {
            for (int j = 0; j < this.col; j++) {
                System.out.print(String.format("%2d ", this.steps[i][j]));
            }
            System.out.println();
        }
    }

    /**
     * 骑士周游算法
     */
    public void traversal(int x, int y, int step) {
        // 将当前位置标记已访问。y = 4, col = 8, x = 4 => 4 * 8 + 4 = 36
        this.visited[y * this.col + x] = true;
        // 记录步数
        this.steps[y][x] = step;
        // 获取下一步可以走的所有位置
        LinkedList<Point> points = this.nextPoints(new Point(x, y));
        // 排序优化，优先选择下一步最少可走数目的位置，减少回溯，体现出贪心算法的特点
        // 去掉此方法，算法的耗时会很久，根据走法不同，结果也会不同
        points.sort(Comparator.comparingInt(point -> this.nextPoints(point).size()));

        while (points.size() != 0) {
            // 取出第一个点
            Point p = points.removeFirst();
            // 该点未被访问过
            if (!this.visited[p.y * this.col + p.x]) {
                // 继续往下走
                this.traversal(p.x, p.y, step + 1);
            }
        }
        // 比较已经走的步数和应该走的步数，如果不相等表示没有走完，将棋盘当前的位置重置
        if (step < this.row * this.col && !this.finished) {
            this.visited[y * this.col + x] = false;
            this.steps[y][x] = 0;
        } else {
            this.finished = true;
        }
    }

    /**
     * 获取当位置的下一步可走位置的集合，最多可有 8 个位置
     */
    public LinkedList<Point> nextPoints(Point point) {
        LinkedList<Point> points = new LinkedList<>();
        // 判断是否可以走 1 的位置
        if (point.x - 1 >= 0 && point.y - 2 >= 0) {
            points.add(new Point(point.x - 1, point.y - 2));
        }
        // 判断是否可以走 2 的位置
        if (point.x + 1 < this.col && point.y - 2 >= 0) {
            points.add(new Point(point.x + 1, point.y - 2));
        }
        // 判断是否可以走 3 的位置
        if (point.x + 2 < this.col && point.y - 1 >= 0) {
            points.add(new Point(point.x + 2, point.y - 1));
        }
        // 判断是否可以走 4 的位置
        if (point.x + 2 < this.col && point.y + 1 < this.row) {
            points.add(new Point(point.x + 2, point.y + 1));
        }
        // 判断是否可以走 5 的位置
        if (point.x + 1 < this.col && point.y + 2 < this.row) {
            points.add(new Point(point.x + 1, point.y + 2));
        }
        // 判断是否可以走 6 的位置
        if (point.x - 1 >= 0 && point.y + 2 < this.row) {
            points.add(new Point(point.x - 1, point.y + 2));
        }
        // 判断是否可以走 7 的位置
        if (point.x - 2 >= 0 && point.y + 1 < this.row) {
            points.add(new Point(point.x - 2, point.y + 1));
        }
        // 判断是否可以走 8 的位置
        if (point.x - 2 >= 0 && point.y - 1 >= 0) {
            points.add(new Point(point.x - 2, point.y - 1));
        }
        return points;
    }

    public static void main(String[] args) {
        Chessboard chessboard = new Chessboard(8, 8);
        chessboard.move(4, 4);
    }
}
