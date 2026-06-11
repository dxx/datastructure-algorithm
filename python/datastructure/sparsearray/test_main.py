import unittest
from main import print_array, read_sparse_array, storage_sparse_array, to_array, to_sparse_array


class Test(unittest.TestCase):
    
    def test_sparse_array(self):
        # 定义一个 5x5 的二维数组
        array = [[0, 0, 0, 0, 0] for _ in range(5)]
        # 初始化 3，6，1，5
        array[0][2] = 3
        array[1][3] = 6
        array[2][1] = 1
        array[3][3] = 5

        print("原二维数组：")
        print_array(array)

        # 转成稀疏数组
        sparse_array = to_sparse_array(array)

        print("转换后的稀疏数组：")
        print_array(sparse_array)

        # 存储稀疏数组
        storage_sparse_array(sparse_array)

        # 读取稀疏数组
        sparse_array = read_sparse_array()
        print("读取的稀疏数组：")
        print_array(sparse_array)

        # 转成二维数组
        array = to_array(sparse_array)
        print("转换后的二维数组：")
        print_array(array)
