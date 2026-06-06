"""
稀疏数组
在二维数组中，如果值为 0 的元素数目远远大于非 0 元素的数目，并且非 0 元素的分布没有规律，则该数组被称为稀疏数组
稀疏数组可以看做是一个压缩的数组，稀疏数组的好处有：
原数组中存在大量的无用的数据，占据了存储空间，真正有用的数据却很少
压缩后存储可以节省存储空间，在数据序列化到磁盘时，压缩存储可以提高 IO 效率

二维数组
0	0	3	0	0
0	0	0	6	0
0	1	0	0	0
0	0	0	5	0
0	0	0	0	0

稀疏数组
row	col	val
5	5	4
0	2	3
1	3	6
2	1	1
3	3	5
"""

from pathlib import Path

SPARSE_ARRAY_FILE_NAME = Path(__file__).with_name("sparse.data")


def to_sparse_array(array: list[list[int]]) -> list[list[int]]:
    """二维数组转稀疏数组"""
    # 统计非 0 的数量
    count = 0
    for i in range(len(array)):
        for j in range(len(array[i])):
            if array[i][j] != 0:
                count += 1

    sparse_array = []
    # 存储第一行信息，从左到右存储的依次是 行，列，非 0 的个数
    sparse_array.append([len(array), len(array[0]), count])

    for i in range(len(array)):
        for j in range(len(array[i])):
            if array[i][j] != 0:
                # 保存 row, col, val
                sparse_array.append([i, j, array[i][j]])
    return sparse_array


def to_array(sparse_array: list[list[int]]) -> list[list[int]]:
    """稀疏数组转二维数组"""
    row = sparse_array[0][0]
    col = sparse_array[0][1]
    # 初始化 0 值
    array = [[0] * col for _ in range(row)]
    # 从稀疏数组中第二行开始读取数据
    for i in range(1, len(sparse_array)):
        row = sparse_array[i][0]
        col = sparse_array[i][1]
        val = sparse_array[i][2]
        array[row][col] = val
    return array


def storage_sparse_array(sparse_array: list[list[int]]) -> None:
    """存储稀疏数组"""
    data = ""
    # 存储矩阵格式
    for row in sparse_array:
        for value in row:
            data += str(value) + "\t"
        data += "\n"
    SPARSE_ARRAY_FILE_NAME.write_text(data, encoding="utf-8")


def read_sparse_array() -> list[list[int]]:
    """读取稀疏数组"""
    file_lines = SPARSE_ARRAY_FILE_NAME.read_text(encoding="utf-8").split("\n")
    # 读取第一行
    strs = file_lines[0].split("\t")
    # 第一行信息
    row = int(strs[0])
    col = int(strs[1])
    val = int(strs[2])

    sparse_array = [[0, 0, 0] for _ in range(val + 1)]
    sparse_array[0] = [row, col, val]
    # 从第二行开始
    for i in range(1, len(file_lines)):
        if file_lines[i]:
            strs = file_lines[i].split("\t")

            row = int(strs[0])
            col = int(strs[1])
            val = int(strs[2])
            sparse_array[i] = [row, col, val]
    return sparse_array


def print_array(array: list[list[int]]) -> None:
    result = ""
    for row in array:
        for value in row:
            result += str(value) + "\t"
        result += "\n"
    print(result)


def main() -> None:
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


if __name__ == "__main__":
    main()
