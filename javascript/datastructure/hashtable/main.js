/**
 * 哈希表
 * 哈希表也叫散列表，是根据关键值 key 直接进行访问的数据结构
 * 它通过把 key 映射到表中一个位置来记录，以加快查找的速度
 */

 /**
  * 员工
  */
function Employee(id, name) {
  this.id = id;
  this.name = name;
}

/**
 * 单链表
 */
function LinkNode(employee, next) {
  this.employee = employee;
  this.next = next
}

/**
 * 哈希表结，包含一个 LinkNode 数组
 */
function Hashtable(len) {
  this.linkArray = new Array(len);
}

/**
 * 添加员工的方法，按照 id 升序插入
 */
Hashtable.prototype.add = function(employee) {
  let linkNode = new LinkNode(employee, null);
  // 计算下标
  let index = employee.id % this.linkArray.length;
  let headNode = this.linkArray[index];
  // 数组没有元素
  if (!headNode) {
    // 添加到链表数组中，作为头结点
    this.linkArray[index] = linkNode;
    return;
  }
  // 判断头结点
  if (employee.id < headNode.employee.id) {
    // 新结点作为头结点
    linkNode.next = headNode;
    this.linkArray[index] = linkNode;
    return;
  } else if (employee.id === headNode.employee.id) {
    console.error("员工id重复");
    return;
  }
  // 查找后续结点中合适的位置插入
  let tempNode = headNode;
  while (true) {
    if (tempNode.next === null) {
      break;
    } else if (tempNode.next.employee.id > employee.id) {
      break;
    } else if (tempNode.next.employee.id === employee.id) {
      console.error("员工id重复");
      return;
    }
  }
  // tempNode 的下一个结点插入到 linkNode 的下一个结点
  linkNode.next = tempNode.next;
  // 将 linkNode 插入到 tempNode 后面
  tempNode.next = linkNode;
}

/**
 * 修改员工的方法
 */
Hashtable.prototype.update = function(employee) {
  let emp = this.getEmployeeById(employee.id);
  if (emp != null) {
      emp.name = employee.name;
  }
}

/**
 * 删除员工的方法
 */
Hashtable.prototype.delete = function(id) {
  // 计算下标
  let index = id % this.linkArray.length;
  let headNode = this.linkArray[index];
  // 数组没有元素
  if (!headNode) {
    return;
  }
  // 判断头结点
  if (headNode.employee.id == id) {
    this.linkArray[index] = headNode.next;
    return;
  }

  // 查找后续结点中需要删除的员工
  let tempNode = headNode;
  while (tempNode.next != null) {
    if (tempNode.next.employee.id == id) {
      // 将要删除结点的下一个结点链接到该结点的上一个结点
      tempNode.next = tempNode.next.next;
      return;
    }
    tempNode = tempNode.next;
  }
}

/**
 * 通过 id 查找员工
 */
Hashtable.prototype.getEmployeeById = function(id) {
  // 计算下标
  let index = id % this.linkArray.length;
  let headNode = this.linkArray[index];
  // 数组没有元素
  if (!headNode) {
    return null;
  }
  // 查找链表中是否存在相同 id 的员工
  let tempNode = headNode;
  while (tempNode != null) {
    if (tempNode.employee.id == id) {
        return tempNode.employee;
    }
    tempNode = tempNode.next;
  }
  return null;
}

/**
 * 显示哈希表内容的方法
 */
Hashtable.prototype.list = function() {
  for (let i = 0; i < this.linkArray.length; i++) {
    let headNode = this.linkArray[i];
    let employeeInfo = "";
    if (headNode != null) {
      employeeInfo += "["
      let tempNode = headNode;
      while (tempNode != null) {
        employeeInfo += "{id=" + tempNode.employee.id + ", name=" + tempNode.employee.name + "}";
        tempNode = tempNode.next;
      }
      employeeInfo += "]"
    }
    console.log("linkArray[" + i + "]=" + employeeInfo)
  }
}

function testAddEmployee() {
  // 创建一个哈希表
  let hashtable = new Hashtable(5);
  // 创建员工
  let employee1 = new Employee(1, "张三");
  let employee2 = new Employee(2, "李四");
  let employee3 = new Employee(5, "孙七");
  // 添加员工
  hashtable.add(employee1);
  hashtable.add(employee2);
  hashtable.add(employee3);

  console.log("添加员工后:");
  // 显示哈希表内容
  hashtable.list();
}

function testUpdateEmployee() {
  let hashtable = new Hashtable(5);
  let employee1 = new Employee(1, "张三");
  let employee2 = new Employee(2, "李四");
  let employee3 = new Employee(6, "周八");
  hashtable.add(employee1);
  hashtable.add(employee2);
  hashtable.add(employee3);

  console.log("修改员工前:");
  // 显示哈希表内容
  hashtable.list();

  // 修改员工
  let employee = new Employee(6, "菜菜");
  hashtable.update(employee);

  console.log("修改员工后:");
  hashtable.list();
}

function testDeleteEmployee() {
  let hashtable = new Hashtable(5);;
  let employee1 = new Employee(2, "李四");
  let employee2 = new Employee(5, "孙七");
  hashtable.add(employee1);
  hashtable.add(employee2);

  console.log("删除员工前:");
  hashtable.list();

  // 删除员工
  hashtable.delete(2);

  console.log("删除员工后:");
  hashtable.list();
}

function main() {
  // testAddEmployee();
  // testUpdateEmployee();
  // testDeleteEmployee();
}

main();
