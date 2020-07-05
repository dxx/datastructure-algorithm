package com.mcx.sparsearray;

import java.io.*;
import java.nio.ByteBuffer;
import java.nio.channels.FileChannel;

/**
 * 稀疏数组
 * 在矩阵中，若数值为 0 的元素数目远远多于非 0 元素的数目，并且非 0 元素分布没有规律时，则称该矩阵为稀疏矩阵
 * 稀疏数组可以看做是一个压缩的数组，稀疏数组的好处有：
 * 原数组中存在大量的无效数据，占据了大量的存储空间，真正有用的数据却少之又少
 * 压缩存储可以节省存储空间以避免资源的不必要的浪费，在数据序列化到磁盘时，压缩存储可以提高 IO 效率
 *
 * 二维数组
 * 0	0	3	0	0
 * 0	0	0	6	0
 * 0	1	0	0	0
 * 0	0	0	5	0
 * 0	0	0	0	0
 *
 * 稀疏数组
 * row	col	val
 * 5	5	4
 * 0	2	3
 * 1	3	6
 * 2	1	1
 * 3	3	5
 */
public class Main {

    private static final String sparseArrayFileName = "./sparse.data";

    /**
     * 二维数组转稀疏数组
     */
    public static int[][] toSparseArray(int[][] array) {
        // 统计非 0 的数量
        int count = 0;
        for (int[] ints : array) {
            for (int anInt : ints) {
                if (anInt != 0) {
                    count++;
                }
            }
        }

        int[][] sparseArray = new int[count + 1][3];
        // 存储第一行信息，从左到右存储的依次是 行，列，非 0 的个数
        sparseArray[0] = new int[]{array.length, array[0].length, count};
        for (int i = 0, row = 1; i < array.length; i++) {
            for (int j = 0; j < array[i].length; j++) {
                if (array[i][j] != 0) {
                    sparseArray[row][0] = i; // row
                    sparseArray[row][1] = j; // col
                    sparseArray[row][2] = array[i][j]; // val
                    row++;
                }
            }
        }
        return sparseArray;
    }

    /**
     * 稀疏数组转二维数组
     */
    public static int[][] toArray(int[][] sparseArray) {
        int row = sparseArray[0][0];
        int col = sparseArray[0][1];
        int[][] array = new int[row][col];
        int val;
        // 从稀疏数组中第二行开始读取数据
        for (int i = 1; i < sparseArray.length; i++) {
            row = sparseArray[i][0];
            col = sparseArray[i][1];
            val = sparseArray[i][2];
            array[row][col] = val;
        }
        return array;
    }

    /**
     * 存储稀疏数组
     */
    public static void storageSparseArray(int[][] sparseArray) {
        File file = new File(sparseArrayFileName);
        try (FileChannel fileChannel = new FileOutputStream(file).getChannel()) {
            StringBuilder sb = new StringBuilder();
            // 存储矩阵格式
            for (int[] ints : sparseArray) {
                for (int anInt : ints) {
                    sb.append(anInt).append("\t");
                }
                sb.append("\n");
            }
            fileChannel.write(ByteBuffer.wrap(sb.toString().getBytes()));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    /**
     * 读取稀疏数组
     */
    public static int[][] readSparseArray() {
        int[][] sparseArray = {};
        File file = new File(sparseArrayFileName);
        try (FileInputStream fis = new FileInputStream(file);
             BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(fis))
        ) {
            // 读取第一行
            String str = bufferedReader.readLine();
            String[] strs = str.split("\t");
            int row = Integer.parseInt(strs[0]);
            int col = Integer.parseInt(strs[1]);
            int val = Integer.parseInt(strs[2]);
            sparseArray = new int[val + 1][3];
            // 第一行信息
            sparseArray[0][0] = row;
            sparseArray[0][1] = col;
            sparseArray[0][2] = val;
            // 从第二行开始
            int i = 1;
            while ((str = bufferedReader.readLine()) != null) {
                str = str.replace("\r\n", "");
                strs = str.split("\t");
                row = Integer.parseInt(strs[0]);
                col = Integer.parseInt(strs[1]);
                val = Integer.parseInt(strs[2]);
                sparseArray[i][0] = row;
                sparseArray[i][1] = col;
                sparseArray[i][2] = val;
                i++;
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        return sparseArray;
    }

    public static void printArray(int[][] array) {
        for (int[] ints : array) {
            for (int anInt : ints) {
                System.out.print(anInt + "\t");
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        // 定义一个 5x5 的二维数组
        int[][] array = new int[5][5];
        // 初始化 3，6， 1，5
        array[0][2] = 3;
        array[1][3] = 6;
        array[2][1] = 1;
        array[3][3] = 5;

        System.out.println("原二维数组：");
        printArray(array);

        // 转成稀疏数组
        int[][] sparseArray = toSparseArray(array);

        System.out.println("转换后的稀疏数组：");
        printArray(sparseArray);

        // 存储稀疏数组
        storageSparseArray(sparseArray);

        // 读取稀疏数组
        sparseArray = readSparseArray();
        System.out.println("读取的稀疏数组：");
        printArray(sparseArray);

        // 转成二维数组
        array = toArray(sparseArray);
        System.out.println("转换后的二维数组：");
        printArray(array);
    }
}
