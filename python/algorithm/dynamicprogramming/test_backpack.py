import unittest
from backpack import find_max_value2


class Test(unittest.TestCase):
    
    def test_find_max_value2(self):
        # 物品重量(kg)
        w = [1, 2, 1]
        # 物品价值
        v = [500, 5000, 3000]
        # 背包容量
        c = 3
        max_value = find_max_value2(w, v, c)
        print("最大价值总和为:", max_value)
