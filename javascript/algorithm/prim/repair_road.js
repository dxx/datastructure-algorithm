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

function MinTree(vertexes, edges) {
  let numOfVertex = vertexes.length;
  let initialVertexes = new Array(numOfVertex);
  let initialMatrix = new Array(numOfVertex);
  for (let i = 0; i < numOfVertex; i++) {
    initialVertexes[i] = vertexes[i];
    initialMatrix[i] = new Array(numOfVertex);
    for (let j = 0; j < numOfVertex; j++) {
      initialMatrix[i][j] = edges[i][j];
    }
  }
  let graph = new Graph();
  graph.vertexes = initialVertexes;
  graph.matrix = initialMatrix;
  this.graph = graph;
}

function Graph(vertexes, matrix) {
  this.vertexes = vertexes; // 顶点
  this.matrix = matrix; // 领接矩阵，代表边的值
}

MinTree.prototype.showGraph = function() {
  for (let i = 0; i < this.graph.matrix.length; i++) {
    let edges = this.graph.matrix[i];
    let str = "[ ";
    for (let j = 0; j < edges.length; j++) {
      str += edges[j] + " ";
    }
    str += "]";
    console.log(str);
  }
}

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
  let minTree = new MinTree(vertexes, edges);
  minTree.showGraph();
}

main();
