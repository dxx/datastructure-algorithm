use std::collections::LinkedList;

pub struct Graph {
    pub vertexes: Vec<String>, // 顶点
    pub matrix: Vec<Vec<u32>>, // 邻接矩阵。0-不通，1-通
    pub num_of_edge: usize,    // 边的数目
}

impl Graph {
    /// 创建图
    pub fn new(num: usize) -> Self {
        let vertexes: Vec<String> = Vec::new();
        let matrix = vec![vec![0; num]; num];
        Graph {
            vertexes,
            matrix,
            num_of_edge: 0,
        }
    }

    /// 添加顶点
    pub fn add_vertex(&mut self, vertex: String) {
        self.vertexes.push(vertex);
    }

    /// 添加边
    /// i1: 第一个顶点下标
    /// i2: 第二个顶点下标
    /// weight: 权值. 0-表示不通, 1-表示通
    pub fn add_edge(&mut self, i1: usize, i2: usize, weight: u32) {
        // 在二维向量中设置权值，因为无方向图，所以两个位置都需要设置
        self.matrix[i1][i2] = weight;
        self.matrix[i2][i1] = weight;
        self.num_of_edge += 1;
    }

    /// 获取顶点数量
    pub fn get_num_of_vertex(&self) -> usize {
        self.vertexes.len()
    }

    /// 获取边的数量
    pub fn get_num_of_edge(&self) -> usize {
        self.num_of_edge
    }

    /// 深度优先遍历
    pub fn dfs(&self) {
        let mut is_visited = vec![false; self.get_num_of_vertex()];
        // 遍历所有的顶点，进行深度优先遍历
        for i in 0..self.get_num_of_vertex() {
            if is_visited[i] == false {
                self.dfs_recursion(&mut is_visited, i);
            }
        }
    }

    // 递归遍历
    // 1.访问初始顶点 v，并标记顶点 v 为已访问
    // 2.查找顶点 v 的第一个邻接顶点 w
    // 3.如果 w 存在，则继续执行第 4 步。如果 w 不存在，则回到第 1 步，将从 v 的下一个顶点继续访问
    // 4.如果 w 未被访问，对 w 进行深度优先遍历递归， 继续进行步骤 1、2、3
    // 5.查找顶点 v 的 w 邻接顶点的下一个邻接顶点，重复步骤 3
    fn dfs_recursion(&self, is_visited: &mut Vec<bool>, v: usize) {
        print!("{}->", self.vertexes[v]);

        // 标记已被访问
        is_visited[v] = true;

        // 获取第一个邻接顶点
        let mut w = self.get_first_vertex(v);
        // 存在则继续调用
        while w != -1 {
            // 未被访问
            if is_visited[w as usize] == false {
                // 继续遍历
                self.dfs_recursion(is_visited, w as usize);
            }
            // 查找顶点 v 的 w 邻接顶点的下一个邻接顶点
            w = self.get_next_vertex(v, w as usize);
        }
    }

    /// 广度优先遍历
    pub fn bfs(&self) {
        let mut is_visited = vec![false; self.get_num_of_vertex()];
        // 遍历所有的顶点，进行广度优先遍历
        for i in 0..self.get_num_of_vertex() {
            if is_visited[i] == false {
                self.bfs2(&mut is_visited, i);
            }
        }
    }

    /// 1.访问初始顶点 v 并标记顶点 v 为已访问。
    /// 2.顶点 v 入队列。
    /// 3.当队列非空时，继续执行，否则结束。
    /// 4.出队列，取得队头结点 u。
    /// 5.查找结点 u 的第一个邻接顶点 w。
    /// 6.若顶点 u 的邻接顶点 w 不存在，则转到步骤 3，否则循环执行以下三个步骤:
    /// 7.若顶点 w 尚未被访问，则访问顶点 w 并标记为已访问。
    /// 8.将顶点 w 入队列。
    /// 9.查找顶点 u 的继 w 邻接顶点后的下一个邻接顶点 w，转到步骤 6。
    fn bfs2(&self, is_visited: &mut Vec<bool>, v: usize) {
        print!("{}->", self.vertexes[v]);

        let mut queue = LinkedList::new();

        // 标记已被访问
        is_visited[v] = true;

        // 将顶点入队列
        queue.push_back(v);
        while !queue.is_empty() {
            // 取出头结点下标
            let u = queue.pop_front();
            // 获取第一个邻接节点的下标
            let mut w = self.get_first_vertex(u.unwrap());
            while w != -1 {
                // 未被访问
                if is_visited[w as usize] == false {
                    print!("{}->", self.vertexes[w as usize]);
                    // 标记已被访问
                    is_visited[w as usize] = true;
                    // 入队列
                    queue.push_back(w as usize);
                }
                // 获取顶点 u 的继 w 邻接顶点后的下一个邻接顶点
                w = self.get_next_vertex(u.unwrap(), w as usize);
            }
        }
    }

    /// 获取第一个邻接顶点下标
    fn get_first_vertex(&self, i: usize) -> i32 {
        for j in 0..self.vertexes.len() {
            if self.matrix[i][j] > 0 {
                return j as i32;
            }
        }
        return -1;
    }

    /// 获取下一个邻接顶点下标
    fn get_next_vertex(&self, i1: usize, i2: usize) -> i32 {
        for j in i2 + 1..self.vertexes.len() {
            if self.matrix[i1][j] > 0 {
                return j as i32;
            }
        }
        return -1;
    }

    /// 显示邻接矩阵
    pub fn show_edges(&self) {
        for i in 0..self.matrix.len() {
            print!("[");
            for j in 0..self.matrix[i].len() {
                print!(" {} ", self.matrix[i][j]);
            }
            println!("]");
        }
    }
}

fn main() {
    let vertexes = ["A", "B", "C", "D", "E"];
    let mut graph = Graph::new(5);

    for v in vertexes.iter() {
        graph.add_vertex(v.to_string());
    }

    // A-B
    graph.add_edge(0, 1, 1);
    // A-C
    graph.add_edge(0, 2, 1);
    // B-C
    graph.add_edge(1, 2, 1);
    // B-E
    graph.add_edge(1, 4, 1);
    // C-D
    graph.add_edge(2, 3, 1);

    graph.show_edges();

    println!("======深度优先遍历======");
    graph.dfs();

    println!();

    println!("======广度优先遍历======");
    graph.bfs();
}
