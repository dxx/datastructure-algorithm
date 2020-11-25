package com.mcx.prim;

import java.util.HashMap;
import java.util.Map;

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
public class RepairRoad {

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
            Graph graph = new Graph();
            graph.vertexes = initialVertexes;
            graph.matrix = initialMatrix;
            this.graph = graph;
        }

        public void prim(int v) {
            int numOfVertex = this.graph.vertexes.length;
            // 存放已经连通的顶点集合
            Map<String, String> vertexMap = new HashMap<>(numOfVertex);
            // 将当前顶点加入集合
            vertexMap.put(this.graph.vertexes[v], this.graph.vertexes[v]);

            // 记录顶点下标
            int v1 = -1;
            int v2 = -1;
            // 记录最小边的权值，初始化成一个最大数，后续遍历中会被替换
            int minWeight = Integer.MAX_VALUE;
            // n 个顶点就有 n-1 条边
            for (int k = 1; k < numOfVertex; k++) {
                // 查找已经加入集合中的顶点，和这些顶点中最近的一个顶点
                for (int i = 0; i < numOfVertex; i++) {
                    for (int j = 0; j < numOfVertex; j++) {
                        int weight = this.graph.matrix[i][j];
                        if (this.graph.vertexes[i].equals(vertexMap.get(this.graph.vertexes[i])) && // 表示已经加入集合的顶点
                                vertexMap.get(this.graph.vertexes[j]) == null && // 表示未被加入集合的顶点
                                weight < minWeight) {
                            v1 = i;
                            v2 = j;
                            minWeight = weight;
                        }
                    }
                }
                // 将最小的顶点加入到集合中
                vertexMap.put(this.graph.vertexes[v2], this.graph.vertexes[v2]);
                // 修改最小的权值
                minWeight = Integer.MAX_VALUE;

                System.out.printf("边:%s-%s => %d\n",
                        this.graph.vertexes[v1],
                        this.graph.vertexes[v2],
                        this.graph.matrix[v1][v2]);
            }
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

        public int[][] matrix; // 领接矩阵，代表边的值

    }

    public static void main(String[] args) {
        String[] vertexes = new String[]{"A", "B", "C", "D", "E", "F", "G"};
        int[][] edges = new int[][]{
            {Integer.MAX_VALUE, 5, 7, Integer.MAX_VALUE, Integer.MAX_VALUE, Integer.MAX_VALUE, 2},
            {5, Integer.MAX_VALUE, Integer.MAX_VALUE, 9, Integer.MAX_VALUE, Integer.MAX_VALUE, 3},
            {7, Integer.MAX_VALUE, Integer.MAX_VALUE, Integer.MAX_VALUE, 8, Integer.MAX_VALUE, Integer.MAX_VALUE},
            {Integer.MAX_VALUE, 9, Integer.MAX_VALUE, Integer.MAX_VALUE, Integer.MAX_VALUE, 4, Integer.MAX_VALUE},
            {Integer.MAX_VALUE, Integer.MAX_VALUE, 8, Integer.MAX_VALUE, Integer.MAX_VALUE, 5, 4},
            {Integer.MAX_VALUE, Integer.MAX_VALUE, Integer.MAX_VALUE, 4, 5, Integer.MAX_VALUE, 6},
            {2, 3, Integer.MAX_VALUE, Integer.MAX_VALUE, 4, 6, Integer.MAX_VALUE},
        };
        MinTree minTree = new MinTree(vertexes, edges);
        // minTree.showGraph();
        // 从 A 点开始
        minTree.prim(0);
    }
}
