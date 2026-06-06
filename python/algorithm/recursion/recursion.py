"""
递归
递归是指在程序运行过程中调用本身的编程技巧
"""


def factorial(n: int) -> int:
    if n > 0:
        return n * factorial(n - 1)
    return 1


def main() -> None:
    res = factorial(5)
    print(res)  # 120


if __name__ == "__main__":
    main()
