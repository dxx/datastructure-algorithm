package com.mcx.graph;

import java.util.LinkedList;

/**
 * 图
 */
public class Main {

    public static class Graph {
        private final String[] vertexes; // 顶点
        private final int[][] matrix; //  邻接矩阵。0-不通，1-通
        private int numOfVertex; // 顶点数目
        private int numOfEdge; // 边的数目

        public Graph(int num) {
            this.vertexes = new String[num];
            this.matrix = new int[num][num];
        }

        /**
         * 添加顶点
         */
        public void addVertex(String vertex) {
            this.vertexes[this.numOfVertex++] = vertex;
        }

        /**
         * 添加边
         * i1: 第一个顶点下标
         * i2: 第二个顶点下标
         * weight: 权值. 0-表示不通, 1-表示通
         */
        public void addEdge(int i1, int i2, int weight) {
            // 在二维数组中设置权值，因为无方向图，所以两个位置都需要设置
            this.matrix[i1][i2] = weight;
            this.matrix[i2][i1] = weight;
            this.numOfEdge++;
        }

        /**
         * 获取顶点数量
         */
        public int getNumOfVertex() {
            return this.numOfVertex;
        }

        /**
         * 获取边的数量
         */
        public int getNumOfEdge() {
            return this.numOfEdge;
        }

        /**
         * 深度优先遍历
         */
        public void dfs() {
            boolean[] isVisited = new boolean[this.getNumOfVertex()];
            // 遍历所有的顶点，进行深度优先遍历
            for (int i = 0; i < this.getNumOfVertex(); i++) {
                if (!isVisited[i]) {
                    this.dfsRecursion(isVisited, i);
                }
            }
        }

        /**
         * 递归遍历
         * 1.访问初始顶点 v，并标记顶点 v 为已访问
         * 2.查找顶点 v 的第一个邻接顶点 w
         * 3.如果 w 存在，则继续执行第 4 步。如果 w 不存在，则回到第 1 步，将从 v 的下一个顶点继续访问
         * 4.如果 w 未被访问，对 w 进行深度优先遍历递归， 继续进行步骤 1、2、3
         * 5.查找顶点 v 的 w 邻接顶点的下一个邻接顶点，重复步骤 3
         */
        public void dfsRecursion(boolean[] isVisited, int v) {
            System.out.printf("%s->", this.vertexes[v]);

            // 标记已被访问
            isVisited[v] = true;

            // 获取第一个邻接顶点
            int w = this.getFirstVertex(v);
            // 存在则继续调用
            while (w != -1) {
                // 未被访问
                if (!isVisited[w]) {
                    // 继续遍历
                    this.dfsRecursion(isVisited, w);
                }
                // 查找顶点 v 的 w 邻接顶点的下一个邻接顶点
                w = this.getNextVertex(v, w);
            }
        }

        /**
         * 广度优先遍历
         */
        public void bfs() {
            boolean[] isVisited = new boolean[this.getNumOfVertex()];
            // 遍历所有的顶点，进行广度优先遍历
            for (int i = 0; i < this.getNumOfVertex(); i++) {
                if (!isVisited[i]) {
                    this.bfs(isVisited, i);
                }
            }
        }

        /**
         * 1.访问初始顶点 v 并标记顶点 v 为已访问。
         * 2.顶点 v 入队列。
         * 3.当队列非空时，继续执行，否则结束。
         * 4.出队列，取得队头结点 u。
         * 5.查找结点 u 的第一个邻接顶点 w。
         * 6.若顶点 u 的邻接顶点 w 不存在，则转到步骤 3，否则循环执行以下三个步骤:
         * 7.若顶点 w 尚未被访问，则访问顶点 w 并标记为已访问。
         * 8.将顶点 w 入队列。
         * 9.查找顶点 u 的继 w 邻接顶点后的下一个邻接顶点 w，转到步骤 6。
         */
        public void bfs(boolean[] isVisited, int v) {
            System.out.printf("%s->", this.vertexes[v]);

            LinkedList<Integer> queue = new LinkedList<>();
            // 标记已被访问
            isVisited[v] = true;
            // 将顶点入队列
            queue.addFirst(v);
            while (!queue.isEmpty()) {
                // 取出头结点下标
                int u = queue.removeLast();
                // 获取第一个邻接节点的下标
                int w = this.getFirstVertex(u);
                while (w != -1) {
                    // 未被访问
                    if (!isVisited[w]) {
                        System.out.printf("%s->", this.vertexes[w]);
                        // 标记已被访问
                        isVisited[w] = true;
                        // 入队列
                        queue.addFirst(w);
                    }
                    // 获取顶点 u 的继 w 邻接顶点后的下一个邻接顶点
                    w = this.getNextVertex(u, w);
                }
            }
        }

        /**
         * 获取第一个邻接顶点下标
         */
        public int getFirstVertex(int i) {
            for (int j = 0; j < this.getNumOfVertex(); j++) {
                if (this.matrix[i][j] > 0) {
                    return j;
                }
            }
            return -1;
        }

        /**
         * 获取下一个邻接顶点下标
         */
        public int getNextVertex(int i1, int i2){
            for (int j = i2 + 1; j < this.getNumOfVertex(); j++) {
                if (this.matrix[i1][j] > 0) {
                    return j;
                }
            }
            return -1;
        }

        /**
         * 显示邻接矩阵
         */
        public void showEdges() {
            for (int i = 0; i < this.matrix.length; i++) {
                System.out.print("[");
                for (int j = 0; j < this.matrix[i].length; j++) {
                    System.out.printf(" %d ", this.matrix[i][j]);
                }
                System.out.println("]");
            }
        }
    }

    public static void main(String[] args) {
        String[] vertexes = new String[]{"A", "B", "C", "D", "E"};
        Graph graph = new Graph(5);

        for (String vertex : vertexes) {
            graph.addVertex(vertex);
        }

        // A-B
        graph.addEdge(0, 1, 1);
        // A-C
        graph.addEdge(0, 2, 1);
        // B-C
        graph.addEdge(1, 2, 1);
        // B-E
        graph.addEdge(1, 4, 1);
        // C-D
        graph.addEdge(2, 3, 1);

        graph.showEdges();

        System.out.println("======深度优先遍历======");
        graph.dfs();

        System.out.println();

        System.out.println("======广度优先遍历======");
        graph.bfs();
    }

}
