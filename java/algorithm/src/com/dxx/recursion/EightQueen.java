package com.dxx.recursion;

import java.util.Arrays;

/**
 * 八皇后
 * 在 8×8 格的国际象棋上摆放 8 个皇后，使其不能互相攻击，即
 * 任意两个皇后都不能处于同一行、同一列或同一斜线上，问有多少种摆法。
 * 思路分析
 * 先将第一个皇后摆放在第一行第一列，接着将第二个皇后先摆在第二行第
 * 一列，如果不行再摆放在第二行第二列，还不行继续在摆放在下一列，依
 * 次类推，整个过程递归进行，当最后一个皇后摆放完成后，回溯到上一个皇后，继续摆放。
 * 使用一个一维数组 positions 存储象棋摆放的位置，下标表示第几个象棋，数组的值
 * 表示第几列。那么摆放第 n 个皇后时，**positions[i] == positions[n]** 表示在同一列
 * **|n - i| == |positions[n] - positions[i]|** 表示在同一斜线。
 */
public class EightQueen {

    private int[] positions; // 存储每 8 个皇后，一种摆放的位置

    public EightQueen() {
        this.positions = new int[8];
    }

    public int[] getPositions() {
        return positions;
    }

    public void setPositions(int[] positions) {
        this.positions = positions;
    }

    public void putQueen(int n) {
        // 最后一个皇后已经放置完成
        if (n == this.getPositions().length) {
            // 打印当前摆放的位置
            System.out.println(Arrays.toString(this.getPositions()));
            return;
        }
        for (int i = 0; i < this.getPositions().length; i++) {
            // i=0 时，假设当前皇后可以放在第一列
            // 如果不能放，将进行下一次循环，当前皇后放在下一个位置
            this.getPositions()[n] = i;
            // 判断是否可以放
            if (this.isCanPut(n)) {
                // 放置下一个皇后
                this.putQueen(n + 1);
            }
        }
    }

    /**
     * 判断当前皇后是否和已经摆放过的皇后冲突
     */
    private boolean isCanPut(int n) {
        int[] positions = this.getPositions();
        for (int i = 0; i < n; i++) {
            // positions[n] == positions[i] 表示在同一列
            // math.Abs(float64(n - i)) == math.Abs(float64(positions[n] - positions[i]) 表示同一斜线
            if (positions[n] == positions[i] ||
            Math.abs(n - i) == Math.abs(positions[n] - positions[i])) {
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        EightQueen eightQueen = new EightQueen();
        eightQueen.putQueen(0);
    }

}
