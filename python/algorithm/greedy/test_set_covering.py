import unittest
from set_covering import get_broadcast


class Test(unittest.TestCase):
    
    def test_get_broadcast(self):
        broadcast_map = {
            "B1": ["北京", "上海", "天津"],
            "B2": ["广州", "北京", "深圳"],
            "B3": ["成都", "上海", "杭州"],
            "B4": ["上海", "天津"],
            "B5": ["杭州", "大连"],
        }
        broadcasts = get_broadcast(broadcast_map)
        print("最少选择的广播电台: " + ",".join(broadcasts))
