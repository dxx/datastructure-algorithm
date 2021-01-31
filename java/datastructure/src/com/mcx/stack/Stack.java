package com.mcx.stack;

/**
 * 栈
 * 栈和队列一样也是一种特殊的线性表。它只能在表尾进行插入和删除操作。
 * 在进行插入和删除操作的一端被称为栈顶，另一端称为栈底。向一个栈放入
 * 新元素称为进栈、入栈或压栈，从一个栈取出元素称为出栈或退栈。每一个
 * 新元素都会放在之前放入的元素之上，删除时会删除最新的元素，所以栈有
 * 先进后出（FILO—first in last out）的特点。
 */
public class Stack {
    private final String[] array; // 存放栈元素
    private final int maxSize; // 最大栈元素大小
    private int top; // 栈顶

    public Stack(int size) {
        this.array = new String[size];
        this.maxSize = size;
        this.top = -1;
    }

    /**
     * 入栈
     */
    public boolean push(String elem) {
        // 判栈是否已满
        if (this.top == this.maxSize - 1) {
            System.err.println("stack is full");
            return false;
        }
        // 栈顶加 1，将元素放入栈顶
        this.array[++this.top] = elem;
        return true;
    }

    /**
     * 出栈
     */
    public String pop() {
        if (this.top == -1) {
            System.err.println("stack is empty");
            return "";
        }
        // 取出栈顶元素，然后加 1
        return this.array[this.top--];
    }

    /**
     * 判断栈是否为空
     */
    public boolean isEmpty() {
        return this.top == -1;
    }

    /**
     * 窥视栈顶元素
     */
    public String peek() {
        if (this.isEmpty()) {
            return null;
        }
        return this.array[this.top];
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append("[");
        for (int i = this.top; i >= 0; i--) {
            sb.append(this.array[i]).append(" ");
        }
        sb.append("]");
        return sb.toString();
    }

    public static void main(String[] args) {
        // 创建一个栈
        Stack stack = new Stack(3);
        // 入栈
        stack.push("one");
        stack.push("two");
        stack.push("three");

        // 栈满，无法入栈
        boolean isSuccess = stack.push("four");
        if (!isSuccess) {
            System.out.println("入栈失败!!!");
        }
        System.out.println(stack);

        String elem1 = stack.pop();
        String elem2 = stack.pop();
        String elem3 = stack.pop();

        System.out.println("出栈: " + elem1);
        System.out.println("出栈: " + elem2);
        System.out.println("出栈: " + elem3);

        String elem = stack.pop();
        if (elem == null) {
            System.out.println("出栈失败!!!");
        }
        System.out.println(stack);
    }
}
