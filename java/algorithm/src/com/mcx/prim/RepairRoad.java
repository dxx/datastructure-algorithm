package com.mcx.prim;

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

        public MinTree(String[] vertexes, int[][] matrix) {
            int numOfVertex = vertexes.length;
            String[] initialVertexes = new String[numOfVertex];
            int[][] initialMatrix = new int[numOfVertex][numOfVertex];
            for (int i = 0; i < numOfVertex; i++) {
                initialVertexes[i] = vertexes[i];
                for (int j = 0; j < numOfVertex; j++) {
                    initialMatrix[i][j] = matrix[i][j];
                }
            }
            Graph graph = new Graph();
            graph.vertexes = initialVertexes;
            graph.matrix = initialMatrix;
            this.graph = graph;
        }

    }

    public static class Graph {

        public String[] vertexes; // 顶点

        public int[][] matrix; // 领接矩阵，代表边的值

    }
}
