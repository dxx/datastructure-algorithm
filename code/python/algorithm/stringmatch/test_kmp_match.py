import unittest
from kmp_match import kmp_search


class Test(unittest.TestCase):
    
    def test_kmp_search(self):
        string = "CBC DCABCABABCABD BBCCA"
        match = "ABCABD"
        index = kmp_search(string, match)
        print(match + " 在 " + string + " 中的位置为 " + str(index))
