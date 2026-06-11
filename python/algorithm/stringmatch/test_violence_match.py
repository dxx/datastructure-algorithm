import unittest
from violence_match import violence_search


class Test(unittest.TestCase):
    
    def test_violence_search(self):
        string = "CBC DCABCABABCABD BBCCA"
        match = "ABCABD"
        index = violence_search(string, match)
        print(match + " 在 " + string + " 中的位置为 " + str(index))
