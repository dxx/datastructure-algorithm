# 数据结构和算法

![rust](https://img.shields.io/badge/language-rust-c99272.svg)
![golang](https://img.shields.io/badge/language-golang-00add8.svg)
![java](https://img.shields.io/badge/language-java-b07219.svg)
![javascript](https://img.shields.io/badge/language-javascript-yellow.svg)

> 数据结构和算法图解，使用多种语言实现

## 数据结构介绍

数据结构是计算机存储、组织数据的方式。数据结构是组织数据元素的集合，这些数据元素之间存在一种或多种特定关系。在解决某种问题时，选择合适的数据结构可以带来更高的运行或者存储效率。

## 算法介绍

算法是一系列解决问题的代码指令，算法代表着用系统的方法描述解决问题的策略。也就是说，能够对一定符合规定的输入，在有限时间内获得所要求的输出。不同的算法完成同样的任务所花费的时间或空间效率是不同的。一个算法的优劣可以用空间复杂度与时间复杂度来衡量。

## 数据结构和算法的关系

关系：

* 数据结构是底层，算法高层
* 数据结构为算法提供服务
* 算法围绕数据结构操作

程序等于数据结构 + 算法

数据结构是算法实现的基础，算法总是要依赖于某种数据结构来实现的。往往是在发展一种算法的时候，构建了适合于这种算法的数据结构。

当然两者也是有一定区别的，算法更加的抽象一些，侧重于对问题的建模，而数据结构则是具体实现方面的问题了，两者是相辅相成的。

## 线性结构和非线性结构

数据结构包括线性结构和非线性结构。

线性结构是最常用的数据结构，特点是数据元素之间存在一对一的关系。线性结构有两种存储方式，一种是顺序存储叫做顺序表，其存储的元素在内存中是连续的，另外一种叫做链表，其存储的元素不一定是连续的，元素节点中存放数据元素以及相邻元素的地址信息。常见的线性结构有数组、队列、链表和栈。

非线性结构包括二维数组、多维数组、广义表、树和图。

## 数据结构

[稀疏数组](./datastructure_01_稀疏数组.md)

[队列](./datastructure_02_队列.md)

[栈](./datastructure_03_栈.md)

[链表](./datastructure_04_链表.md)

[哈希表](./datastructure_05_哈希表.md)

[树](./datastructure_06_树.md)

[图](./datastructure_07_图.md)

## 算法

[递归](./algorithm_01_递归.md)

* [迷宫回溯](./algorithm_01_递归.md#迷宫回溯)

* [八皇后](./algorithm_01_递归.md#八皇后)

[排序](./algorithm_02_排序.md)

* [冒泡排序](./algorithm_02_排序.md#冒泡排序)
* [选择排序](./algorithm_02_排序.md#选择排序)
* [插入排序](./algorithm_02_排序.md#插入排序)
* [快速排序](./algorithm_02_排序.md#快速排序)
* [希尔排序](./algorithm_02_排序.md#希尔排序)
* [归并排序](./algorithm_02_排序.md#归并排序)
* [基数排序](./algorithm_02_排序.md#基数排序)
* [堆排序](./datastructure_06_树.md#堆排序)

[查找](./algorithm_03_查找.md)

* [线性查找](./algorithm_03_查找.md#线性查找)
* [二分法查找](./algorithm_03_查找.md#二分法查找)
* [二分法查找(非递归)](./algorithm_03_查找.md#二分法查找非递归)
* [插值查找](./algorithm_03_查找.md#插值查找)
* [斐波那契查找](./algorithm_03_查找.md#斐波那契查找)

[分治算法](./algorithm_04_分治算法.md)

* [汉诺塔](./algorithm_04_分治算法.md#汉诺塔)

[动态规划](./algorithm_05_动态规划.md)

* [背包问题](./algorithm_05_动态规划.md#背包问题)

[KMP算法](./algorithm_06_KMP算法.md)

* [暴力匹配](./algorithm_06_KMP算法.md#暴力匹配)
* [KMP匹配](./algorithm_06_KMP算法.md#KMP匹配)

[贪心算法](./algorithm_07_贪心算法.md)

[普里姆(Prim)算法](./algorithm_08_普里姆(Prim)算法.md)

[克鲁斯卡尔算法](./algorithm_09_克鲁斯卡尔算法.md)

[马踏棋盘算法](./algorithm_10_马踏棋盘算法.md)