import unittest
from main import PersonLinkedList


class Test(unittest.TestCase):
    
    def test_person_linked_list(self):
        person_linked_list = PersonLinkedList(5)
        person_linked_list.show_persons()

        person_linked_list.count(1, 3)
