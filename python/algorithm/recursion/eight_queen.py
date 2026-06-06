"""
八皇后
在 8×8 格的国际象棋上摆放 8 个皇后，使其不能互相攻击，即
任意两个皇后都不能处于同一行、同一列或同一斜线上，问有多少种摆法。
思路分析
先将第一个皇后摆放在第一行第一列，接着将第二个皇后先摆在第二行第
一列，如果不行再摆放在第二行第二列，还不行继续在摆放在下一列，依
次类推，整个过程递归进行，当最后一个皇后摆放完成后，回溯到上一个皇后，继续摆放。
使用一个一维数组 positions 存储象棋摆放的位置，下标表示第几个象棋，数组的值
表示第几列。那么摆放第 n 个皇后时，**positions[i] == positions[n]** 表示在同一列
**|n - i| == |positions[n] - positions[i]|** 表示在同一斜线。
"""


class EightQueen:
    def __init__(self) -> None:
        self.positions = [0 for _ in range(8)]

    def put_queen(self, n: int) -> None:
        # 最后一个皇后已经放置完成
        if n == len(self.positions):
            # 打印当前摆放的位置
            print(self.positions)
            return
        for i in range(len(self.positions)):
            # i=0 时，假设当前皇后可以放在第一列
            # 如果不能放，将进行下一次循环，当前皇后放在下一个位置
            self.positions[n] = i
            # 判断是否可以放
            if self._is_can_put(n):
                # 放置下一个皇后
                self.put_queen(n + 1)

    def _is_can_put(self, n: int) -> bool:
        """判断当前皇后是否和已经摆放过的皇后冲突"""
        for i in range(n):
            if self.positions[n] == self.positions[i] or abs(n - i) == abs(self.positions[n] - self.positions[i]):
                return False
        return True


def main() -> None:
    eight_queen = EightQueen()
    eight_queen.put_queen(0)


if __name__ == "__main__":
    main()
