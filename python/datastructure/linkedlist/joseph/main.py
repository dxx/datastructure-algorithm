"""
约瑟夫问题
设编号为 1,2,...n 的 n 个人围坐一圈约定编号为 k(1<=k<=n) 的人
从 1 开始报数，数到 m 的那个人出列它的下一位又从 1 开始报数数到 m
的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列
"""


class Person:
    def __init__(self, no: int) -> None:
        self.no = no
        self.prev: Person | None = None
        self.next: Person | None = None


class PersonLinkedList:
    def __init__(self, count: int) -> None:
        if count <= 0:
            print("链表至少需要一个元素")
            self.first = None
            self.length = 0
            return
        self.first = None
        self.length = count

        prev = self.first
        # 初始化小孩，构成双向循环链表
        for i in range(1, count + 1):
            person = Person(i)
            if i == 1:
                # 初始化 First 节点
                self.first = person
                self.first.next = self.first
                self.first.prev = self.first

                # 将 prev 指向 First 节点，继续下一次循环
                prev = self.first
                continue

            if prev is None:
                return
            prev.next = person
            person.prev = prev

            # 新增加的节点的下一个节点指向第一节点
            if self.first is None:
                return
            person.next = self.first
            # 第一个节点的上一个节点指向新增加的节点
            self.first.prev = person

            prev = person

    def show_persons(self) -> None:
        if self.first is None:
            return
        current = self.first
        while True:
            print("num:" + str(current.no))
            if current.next is None:
                break
            current = current.next
            if current == self.first:
                break

    def count(self, start: int, num: int) -> None:
        if start < 1 or start > self.length:
            print("start 不能小于 1 或者不能大于 " + str(self.length))
            return
        if num > self.length:
            print("num 不能大于元素个数: " + str(self.length))
            return

        current = self.first
        if current is None:
            return

        for _ in range(1, start):
            if current.next is None:
                return
            current = current.next

        while True:
            if current.prev is None or current.next is None:
                return
            if current.prev == current and current.next == current:
                break

            for _ in range(1, num):
                if current.next is None:
                    return
                current = current.next

            prev_node = current.prev
            next_node = current.next
            if prev_node is None or next_node is None:
                return
            prev_node.next = next_node
            next_node.prev = prev_node

            print("出队人的编号: " + str(current.no))
            current = next_node
        print("最后留下人的编号: " + str(current.no))


def main() -> None:
    person_linked_list = PersonLinkedList(5)
    person_linked_list.show_persons()

    person_linked_list.count(1, 3)


if __name__ == "__main__":
    main()
