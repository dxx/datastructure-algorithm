package horsetreads

import "testing"

func TestChessboard(t *testing.T) {
    chessboard := NewChessboard(8, 8)
    // 从 4,4 的位置开始走
    chessboard.Move(4, 4)
}
