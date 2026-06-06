"""
哈希表
哈希表也叫散列表，是根据关键值 key 直接进行访问的数据结构
它通过把 key 映射到表中一个位置来记录，以加快查找的速度
"""


class Employee:
    """员工"""

    def __init__(self, emp_id, name):
        self.id = emp_id
        self.name = name


class LinkNode:
    """单链表"""

    def __init__(self, employee: Employee, next_node: "LinkNode | None" = None):
        self.employee = employee
        self.next = next_node


class Hashtable:
    """哈希表结，包含一个 LinkNode 数组"""

    def __init__(self, length):
        self.link_array: list[LinkNode | None] = [None] * length

    def add(self, employee):
        """添加员工的方法，按照 id 升序插入"""
        link_node = LinkNode(employee)
        # 计算下标
        index = employee.id % len(self.link_array)
        head_node = self.link_array[index]
        # 数组没有元素
        if not head_node:
            # 添加到链表数组中，作为头结点
            self.link_array[index] = link_node
            return
        # 判断头结点
        if employee.id < head_node.employee.id:
            # 新结点作为头结点
            link_node.next = head_node
            self.link_array[index] = link_node
            return
        if employee.id == head_node.employee.id:
            print("员工 id 重复")
            return
        # 查找后续结点中合适的位置插入
        temp_node = head_node
        while True:
            if temp_node.next is None:
                break
            if temp_node.next.employee.id > employee.id:
                break
            if temp_node.next.employee.id == employee.id:
                print("员工 id 重复")
                return
            temp_node = temp_node.next
        # tempNode 的下一个结点插入到 linkNode 的下一个结点
        link_node.next = temp_node.next
        # 将 linkNode 插入到 tempNode 后面
        temp_node.next = link_node

    def update(self, employee):
        """修改员工的方法"""
        emp = self.get_employee_by_id(employee.id)
        if emp is not None:
            emp.name = employee.name

    def delete(self, emp_id):
        """删除员工的方法"""
        # 计算下标
        index = emp_id % len(self.link_array)
        head_node = self.link_array[index]
        # 数组没有元素
        if not head_node:
            return
        # 判断头结点
        if head_node.employee.id == emp_id:
            self.link_array[index] = head_node.next
            return

        # 查找后续结点中需要删除的员工
        temp_node = head_node
        while temp_node.next is not None:
            if temp_node.next.employee.id == emp_id:
                # 将要删除结点的下一个结点链接到该结点的上一个结点
                temp_node.next = temp_node.next.next
                return
            temp_node = temp_node.next

    def get_employee_by_id(self, emp_id):
        """通过 id 查找员工"""
        # 计算下标
        index = emp_id % len(self.link_array)
        head_node = self.link_array[index]
        # 数组没有元素
        if not head_node:
            return None
        # 查找链表中是否存在相同 id 的员工
        temp_node = head_node
        while temp_node is not None:
            if temp_node.employee.id == emp_id:
                return temp_node.employee
            temp_node = temp_node.next
        return None

    def list(self):
        """显示哈希表内容的方法"""
        for i, head_node in enumerate(self.link_array):
            employee_info = ""
            if head_node is not None:
                employee_info += "["
                temp_node = head_node
                while temp_node is not None:
                    employee_info += "{id=" + str(temp_node.employee.id) + ", name=" + temp_node.employee.name + "}"
                    temp_node = temp_node.next
                employee_info += "]"
            print("linkArray[" + str(i) + "]=" + employee_info)


def test_add_employee():
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


def test_update_employee():
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


def test_delete_employee():
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


def main():
    # test_add_employee()
    # test_update_employee()
    # test_delete_employee()
    pass


if __name__ == "__main__":
    main()
