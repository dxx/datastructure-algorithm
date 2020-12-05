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

/**
 * 最小生成树
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
  let graph = new Graph(initialVertexes, initialMatrix);
  this.graph = graph;
}

function Graph(vertexes, matrix) {
  this.vertexes = vertexes; // 顶点
  this.matrix = matrix; // 领接矩阵，代表边的值
}

MinTree.prototype.prim = function(v) {
  let numOfVertex = this.graph.vertexes.length;
  // 存放已经连通的顶点集合
  let vertexMap = {};
  // 将当前顶点加入集合
  vertexMap[this.graph.vertexes[v]] = this.graph.vertexes[v];

  // 记录顶点下标
  let v1 = -1;
  let v2 = -1;
  // 记录最小边的权值，初始化成一个最大数，后续遍历中会被替换
  let minWeight = Number.MAX_VALUE;
  // n 个顶点就有 n-1 条边
  for (let k = 1; k < numOfVertex; k++) {
    // 查找已经加入集合中的顶点，和这些顶点中最近的一个顶点
    for (let i = 0; i < numOfVertex; i++) {
      for (let j = 0; j < numOfVertex; j++) {
        let weight = this.graph.matrix[i][j];
        if (vertexMap[this.graph.vertexes[i]] == this.graph.vertexes[i] && // 表示已经加入集合的顶点
              vertexMap[this.graph.vertexes[j]] == undefined && // 表示未被加入集合的顶点
              weight < minWeight) {
          v1 = i;
          v2 = j;
          minWeight = weight;
        }
      }
    }
    // 将最小的顶点加入到集合中
    vertexMap[this.graph.vertexes[v2]] = this.graph.vertexes[v2];
    // 修改最小的权值
    minWeight = Number.MAX_VALUE;
      
    console.log("边: " + this.graph.vertexes[v1] + "-" + this.graph.vertexes[v2] + " => " + this.graph.matrix[v1][v2]);
  }
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
  // minTree.showGraph();
  // 从 A 点开始
  minTree.prim(0);
}

main();
