package com.mcx.hashtable;

/**
 * 哈希表
 * 哈希表也叫散列表，是根据关键值 key 直接进行访问的数据结构
 * 它通过把 key 映射到表中一个位置来记录，以加快查找的速度
 */
public class Main {

    /**
     * 员工
     */
    public static class Employee {
        private int id;
        private String name;

        public Employee(int id, String name) {
            this.id = id;
            this.name = name;
        }

        public int getId() {
            return id;
        }

        public void setId(int id) {
            this.id = id;
        }

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }
    }

    /**
     * 单链表
     */
    public static class LinkNode {
        private Employee employee;
        private LinkNode next;
        public LinkNode(Employee employee, LinkNode next) {
            this.employee = employee;
            this.next = next;
        }

        public Employee getEmployee() {
            return employee;
        }

        public LinkNode getNext() {
            return next;
        }

        public void setNext(LinkNode next) {
            this.next = next;
        }
    }

    /**
     * 哈希表结，包含一个 LinkNode 数组
     */
    public static class Hashtable {

        private final LinkNode[] linkArray;

        public Hashtable(int len) {
            linkArray = new LinkNode[len];
        }

        public LinkNode[] getLinkArray() {
            return linkArray;
        }

        /**
         * 添加员工的方法，按照 id 升序插入
         */
        public void add(Employee employee) {
            LinkNode linkNode = new LinkNode(employee, null);
            // 计算下标
            int index = employee.getId() % this.getLinkArray().length;
            LinkNode headNode = this.getLinkArray()[index];
            // 数组没有元素
            if (headNode == null) {
                // 添加到链表数组中，作为头结点
                this.getLinkArray()[index] = linkNode;
                return;
            }
            // 判断头结点
            if (employee.getId() < headNode.getEmployee().getId()) {
                // 新结点作为头结点
                linkNode.setNext(headNode);
                this.getLinkArray()[index] = linkNode;
                return;
            } else if (employee.getId() == headNode.getEmployee().getId()) {
                throw new IllegalStateException("员工id重复");
            }
            // 查找后续结点中合适的位置插入
            LinkNode tempNode = headNode;
            while (true) {
                if (tempNode.getNext() == null) { // 最后一个结点
                    break;
                } else if (tempNode.getNext().getEmployee().getId() > employee.getId()) {
                    break;
                } else if (tempNode.getNext().getEmployee().getId() == employee.getId()) {
                    throw new IllegalStateException("员工id重复");
                }
            }
            // tempNode 的下一个结点插入到 linkNode 的下一个结点
            linkNode.setNext(tempNode.getNext());
            // 将 linkNode 插入到 tempNode 后面
            tempNode.setNext(linkNode);
        }

        /**
         * 修改员工的方法
         */
        public void update(Employee employee) {
            Employee emp = this.getEmployeeById(employee.getId());
            if (emp != null) {
                emp.setName(employee.getName());
            }
        }

        /**
         * 删除员工的方法
         */
        public void delete(int id) {
            // 计算下标
            int index = id % this.getLinkArray().length;
            LinkNode headNode = this.getLinkArray()[index];
            // 数组没有元素
            if (headNode == null) {
                return;
            }
            // 判断头结点
            if (headNode.getEmployee().getId() == id) {
                this.getLinkArray()[index] = headNode.getNext();
                return;
            }

            // 查找后续结点中需要删除的员工
            LinkNode tempNode = headNode;
            while (tempNode.getNext() != null) {
                if (tempNode.getNext().getEmployee().getId() == id) {
                    // 将要删除结点的下一个结点链接到该结点的上一个结点
                    tempNode.setNext(tempNode.getNext().getNext());
                    return;
                }
                tempNode = tempNode.getNext();
            }
        }

        /**
         * 通过 id 查找员工
         */
        public Employee getEmployeeById(int id) {
            // 计算下标
            int index = id % this.getLinkArray().length;
            LinkNode headNode = this.getLinkArray()[index];
            // 数组没有元素
            if (headNode == null) {
                return null;
            }
            // 查找链表中是否存在相同 id 的员工
            LinkNode tempNode = headNode;
            while (tempNode != null) {
                if (tempNode.getEmployee().getId() == id) {
                    return tempNode.getEmployee();
                }
                tempNode = tempNode.getNext();
            }
            return null;
        }

        /**
         * 显示哈希表内容的方法
         */
        public void list() {
            for (int i = 0; i < this.getLinkArray().length; i++) {
                LinkNode headNode = this.getLinkArray()[i];
                StringBuilder sb = new StringBuilder();
                if (headNode != null) {
                    sb.append("[");
                    LinkNode tempNode = headNode;
                    while (tempNode != null) {
                        sb.append(String.format("{id=%d, name=%s}",
                                tempNode.getEmployee().getId(), tempNode.getEmployee().getName()));
                        tempNode = tempNode.next;
                    }
                    sb.append("]");
                }
                System.out.printf("linkArray[%d]=%s\n", i, sb.toString());
            }
        }
    }

    public static void testAddEmployee() {
        // 创建一个哈希表
        Hashtable hashtable = new Hashtable(5);
        // 创建员工
        Employee employee1 = new Employee(1, "张三");
        Employee employee2 = new Employee(2, "李四");
        Employee employee3 = new Employee(5, "孙七");
        // 添加员工
        hashtable.add(employee1);
        hashtable.add(employee2);
        hashtable.add(employee3);

        System.out.println("添加员工后:");
        // 显示哈希表内容
        hashtable.list();
    }

    public static void testUpdateEmployee() {
        Hashtable hashtable = new Hashtable(5);
        Employee employee1 = new Employee(1, "张三");
        Employee employee2 = new Employee(2, "李四");
        Employee employee3 = new Employee(6, "周八");
        hashtable.add(employee1);
        hashtable.add(employee2);
        hashtable.add(employee3);

        System.out.println("修改员工前:");
        // 显示哈希表内容
        hashtable.list();

        // 修改员工
        Employee employee = new Employee(6, "菜菜");
        hashtable.update(employee);

        System.out.println("修改员工后:");
        hashtable.list();
    }

    public static void testDeleteEmployee() {
        Hashtable hashtable = new Hashtable(5);;
        Employee employee1 = new Employee(2, "李四");
        Employee employee2 = new Employee(5, "孙七");
        hashtable.add(employee1);
        hashtable.add(employee2);

        System.out.println("删除员工前:");
        hashtable.list();

        // 删除员工
        hashtable.delete(2);

        System.out.println("删除员工后:");
        hashtable.list();
    }

    public static void main(String[] args) {
        // testAddEmployee();
        // testUpdateEmployee();
        // testDeleteEmployee();
    }

}
