from stack import Stack


def main():
    # 创建一个栈
    stack = Stack(3)
    # 入栈
    stack.push("one")
    stack.push("two")
    stack.push("three")

    # 栈满，无法入栈
    is_success = stack.push("four")
    if not is_success:
        print("入栈失败!!!")
    stack.show()

    elem1 = stack.pop()
    elem2 = stack.pop()
    elem3 = stack.pop()

    print("出栈: " + str(elem1))
    print("出栈: " + str(elem2))
    print("出栈: " + str(elem3))

    elem = stack.pop()
    if elem is None or elem == "":
        print("出栈失败!!!")
    stack.show()


if __name__ == "__main__":
    main()
