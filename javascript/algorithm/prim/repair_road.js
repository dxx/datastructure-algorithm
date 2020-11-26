/**
 * 普里姆(Prim)算法解决村庄修路问题。
 * Prim 算法适用于稠密图
 * 算法思路：
 * 1.首先随便选一个点加入集合。
 * 2.用该点的所有边去刷新到其它点的最短路。
 * 3.找出最短路中最短的一条连接（且该点未被加入集合）。
 * 4.用该点去刷新到其他点的最短路。
 * 5.重复以上操作 n-1 次。
 */

function main() {
  let vertexes = ["A", "B", "C", "D", "E", "F", "G"];
  let edges = [
    [Number.MAX_VALUE, 5, 7, Number.MAX_VALUE, Number.MAX_VALUE, Number.MAX_VALUE, 2],
    [5, Number.MAX_VALUE, Number.MAX_VALUE, 9, Number.MAX_VALUE, Number.MAX_VALUE, 3],
    [7, Number.MAX_VALUE, Number.MAX_VALUE, Number.MAX_VALUE, 8, Number.MAX_VALUE, Number.MAX_VALUE],
    [Number.MAX_VALUE, 9, Number.MAX_VALUE, Number.MAX_VALUE, Number.MAX_VALUE, 4, Number.MAX_VALUE],
    [Number.MAX_VALUE, Number.MAX_VALUE, 8, Number.MAX_VALUE, Number.MAX_VALUE, 5, 4],
    [Number.MAX_VALUE, Number.MAX_VALUE, Number.MAX_VALUE, 4, 5, Number.MAX_VALUE, 6],
    [2, 3, Number.MAX_VALUE, Number.MAX_VALUE, 4, 6, Number.MAX_VALUE]
  ];
}

main();
