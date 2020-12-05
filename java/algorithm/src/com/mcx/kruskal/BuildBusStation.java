package com.mcx.kruskal;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

/**
 * 克鲁斯卡尔(Kruskal)算法解决建设公交站问题
 * Kruskal 算法适用于稀疏图
 * 思想：按照权值从小到大的顺序选择 n-1 条边，并保证这 n-1 条边不构成回路
 */
public class BuildBusStation {

    /**
     * 最小生成树
     */
    public static class MinTree {

        public Graph graph;

        public MinTree(String[] vertexes, int[][] edges) {
            int numOfVertex = vertexes.length;
            String[] initialVertexes = new String[numOfVertex];
            int[][] initialMatrix = new int[numOfVertex][numOfVertex];
            for (int i = 0; i < numOfVertex; i++) {
                initialVertexes[i] = vertexes[i];
                for (int j = 0; j < numOfVertex; j++) {
                    initialMatrix[i][j] = edges[i][j];
                }
            }
            int numOfEdge = 0;
            for (int i = 0; i < numOfVertex; i++) {
                // 已经统计过的边不统计
                for (int j = i + 1; j < numOfVertex; j++) {
                    numOfEdge++;
                }
            }
            Graph graph = new Graph();
            graph.vertexes = initialVertexes;
            graph.matrix = initialMatrix;
            graph.numOfEdge = numOfEdge;
            this.graph = graph;
        }

        public void kruskal() {
            // 保存最小生成树的边
            List<Edge> edges = new ArrayList<>();

            // 保存已存在最小生成树中每个顶点对应的在树中的终点下标
            int[] endPosits = new int[this.graph.numOfEdge];

            // 获取所有边
            List<Edge> allEdges = this.getEdges();
            System.out.println("======边排序前======");
            System.out.println(allEdges);

            // 对边按照权值从小到大进行排序
            allEdges.sort(Comparator.comparingInt(a -> a.weight));

            System.out.println("======边排序后======");
            System.out.println(allEdges);

            // 遍历所以的边
            for (Edge v : allEdges) {
                // 获取边的起始顶点下标
                int startPosit = this.getVertexPosit(v.start);
                // 获取边的结束顶点下标
                int endPosit = this.getVertexPosit(v.end);

                // 获取起始顶点的终点下标
                int endPosit1 = this.getEndPosit(endPosits, startPosit);
                // 获取结束顶点的终点下标
                int endPosit2 = this.getEndPosit(endPosits, endPosit);
                // 判断是否形成回路
                if (endPosit1 != endPosit2) { // 没有构成回路
                    // 设置 endPosit1 的终点下标 endPosit2
                    // [0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0]
                    endPosits[endPosit1] = endPosit2;
                    edges.add(v); // 将改变加入最小生成树
                }
            }

            // 输出最小生成树
            for (Edge edge : edges) {
                System.out.printf("边: %s-%s => %d\n",
                        edge.start,
                        edge.end,
                        edge.weight);
            }
        }

        /**
         * 获取顶点在顶点集合中的位置
         * @param vertex 顶点
         */
        private int getVertexPosit(String vertex) {
            for (int i = 0; i < this.graph.vertexes.length; i++) {
                if (vertex.equals(this.graph.vertexes[i])) {
                    return i;
                }
            }
            return -1;
        }

        /**
         * 获取指定下标顶点的终点下标
         * posits: 存放顶点和对应终点下标，posits 的下标表示顶点下标，值表示对应顶点的终点下标
         * i: 目标顶点下标
         * posits = [0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0], i = 4 返回 5
         */
        private int getEndPosit(int[] posits, int i) {
            while (posits[i] != 0) {
                i = posits[i];
            }
            return i;
        }

        /**
         * 获取所有的边
         */
        private List<Edge> getEdges() {
            List<Edge> edges = new ArrayList<>();
            int numOfVertex = this.graph.vertexes.length;
            for (int i = 0; i < numOfVertex; i++) {
                for (int j = i + 1; j < numOfVertex; j++) {
                    // 不连通的跳过
                    if (this.graph.matrix[i][j] == Integer.MAX_VALUE) {
                        continue;
                    }
                    // 创建边
                    Edge edge = new Edge();
                    edge.start = this.graph.vertexes[i];
                    edge.end = this.graph.vertexes[j];
                    edge.weight = this.graph.matrix[i][j];
                    edges.add(edge);
                }
            }
            return edges;
        }

        public void showGraph() {
            for (int[] edges : this.graph.matrix) {
                System.out.print("[ ");
                for (int val : edges) {
                    System.out.printf("%10d ", val);
                }
                System.out.print("]\n");
            }
        }
    }

    public static class Graph {

        public String[] vertexes; // 顶点

        public int[][] matrix; // 邻接矩阵

        public int numOfEdge; // 边的条数

    }

    /**
     * 边
     */
    public static class Edge {

        public String start; // 起始顶点

        public String end; // 结束顶点

        public int weight; // 边的权值

        @Override
        public String toString() {
            return String.format("%s-%s:%d", this.start, this.end, this.weight);
        }
    }

    public static void main(String[] args) {
        String[] vertexes = new String[]{"A", "B", "C", "D", "E", "F", "G"};
        // 0-表示自己跟自己不连通，intMax-表示跟其它顶点不连通
        int[][] edges = new int[][]{
            {0, 12, Integer.MAX_VALUE, Integer.MAX_VALUE, Integer.MAX_VALUE, 16, 14},
            {12, 0, 10, Integer.MAX_VALUE, Integer.MAX_VALUE, 7, Integer.MAX_VALUE},
            {Integer.MAX_VALUE, 10, 0, 3, 5, 6, Integer.MAX_VALUE},
            {Integer.MAX_VALUE, Integer.MAX_VALUE, 3, 0, 4, Integer.MAX_VALUE, Integer.MAX_VALUE},
            {Integer.MAX_VALUE, Integer.MAX_VALUE, 5, 4, 0, 2, 8},
            {16, 7, 6, Integer.MAX_VALUE, 2, 0, 9},
            {14, Integer.MAX_VALUE, Integer.MAX_VALUE, Integer.MAX_VALUE, 8, 9, 0}
        };
        MinTree minTree = new MinTree(vertexes, edges);
        // minTree.showGraph();
        minTree.kruskal();
    }
}
