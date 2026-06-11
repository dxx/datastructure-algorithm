import unittest
from main import test_add_employee, test_delete_employee, test_update_employee


class Test(unittest.TestCase):
    
    def test_add_employee(self):
        test_add_employee()

    def test_update_employee(self):
        test_update_employee()

    def test_delete_employee(self):
        test_delete_employee()
