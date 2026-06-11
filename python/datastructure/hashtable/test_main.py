import unittest
from main import Employee, Hashtable


class Test(unittest.TestCase):
    
    def test_add_employee(self):
        # 创建一个哈希表
        hashtable = Hashtable(5)
        # 创建员工
        employee1 = Employee(1, "张三")
        employee2 = Employee(2, "李四")
        employee3 = Employee(5, "孙七")
        # 添加员工
        hashtable.add(employee1)
        hashtable.add(employee2)
        hashtable.add(employee3)

        print("添加员工后:")
        # 显示哈希表内容
        hashtable.list()

    def test_update_employee(self):
        hashtable = Hashtable(5)
        employee1 = Employee(1, "张三")
        employee2 = Employee(2, "李四")
        employee3 = Employee(6, "周八")
        hashtable.add(employee1)
        hashtable.add(employee2)
        hashtable.add(employee3)

        print("修改员工前:")
        # 显示哈希表内容
        hashtable.list()

        # 修改员工
        employee = Employee(6, "菜菜")
        hashtable.update(employee)

        print("修改员工后:")
        hashtable.list()

    def test_delete_employee(self):
        hashtable = Hashtable(5)
        employee1 = Employee(2, "李四")
        employee2 = Employee(5, "孙七")
        hashtable.add(employee1)
        hashtable.add(employee2)

        print("删除员工前:")
        hashtable.list()

        # 删除员工
        hashtable.delete(2)

        print("删除员工后:")
        hashtable.list()
