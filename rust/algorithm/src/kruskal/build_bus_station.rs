use std::fmt;

/// 克鲁斯卡尔(Kruskal)算法解决建设公交站问题
/// Kruskal 算法适用于稀疏图
/// 思想：按照权值从小到大的顺序选择 n-1 条边，并保证这 n-1 条边不构成回路

/// 最小生成树
pub struct MinTree {
    graph: Graph,
}

pub struct Graph {
    vertexes: Vec<String>, // 顶点
    matrix: Vec<Vec<u32>>, // 邻接矩阵
    num_of_edge: usize,    // 边的条数
}

pub struct Edge {
    start: String, // 起始顶点
    end: String,   // 结束顶点
    weight: u32,   // 边的权值
}

impl Edge {
    pub fn new(start: String, end: String, weight: u32) -> Self {
        Edge { start, end, weight }
    }
}

impl fmt::Debug for Edge {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}-{}:{}", self.start, self.end, self.weight)
    }
}

impl MinTree {
    pub fn new(vertexes: Vec<String>, edges: Vec<Vec<u32>>) -> Self {
        let num_of_vertex = vertexes.len();
        let mut num_of_edge = 0;
        for i in 0..num_of_vertex {
            for _j in i + 1..num_of_vertex {
                num_of_edge += 1;
            }
        }
        let graph = Graph {
            vertexes,
            matrix: edges,
            num_of_edge,
        };
        MinTree { graph }
    }

    pub fn kruskal(&self) {
        let mut edges = Vec::new();

        let mut end_posits = vec![0; self.graph.num_of_edge];

        let mut all_edges = self.get_edges();
        println!("======边排序前======");
        println!("{:?}", all_edges);

        all_edges.sort_by(|edge1, edge2| edge1.weight.cmp(&edge2.weight));
        println!("======边排序后======");
        println!("{:?}", all_edges);

        // 遍历所以的边
        for v in all_edges.iter() {
            // 获取边的起始顶点下标
            let start_posit = self.get_vertex_posit(v.start.clone());
            // 获取边的结束顶点下标
            let end_posit = self.get_vertex_posit(v.end.clone());

            // 获取起始顶点的终点下标
            let end_posit1 = self.get_end_posit(&end_posits, start_posit as usize);
            // 获取结束顶点的终点下标
            let end_posit2 = self.get_end_posit(&end_posits, end_posit as usize);
            // 判断是否形成回路
            if end_posit1 != end_posit2 {
                // 没有构成回路
                // 设置 end_posit1 的终点下标 end_posit2
                // [0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0]
                end_posits[end_posit1] = end_posit2 as u32;
                edges.push(v); // 将改变加入最小生成树
            }
        }

        // 输出最小生成树
        for edge in edges {
            println!("边: {}-{} => {}", edge.start, edge.end, edge.weight);
        }
    }

    pub fn show_graph(&self) {
        for edges in self.graph.matrix.iter() {
            print!("[ ");
            for val in edges {
                print!("{:12}", val);
            }
            println!("]\n");
        }
    }

    /// 获取顶点在顶点集合中的位置
    /// vertex: 顶点
    fn get_vertex_posit(&self, vertex: String) -> i32 {
        let vertexes = &self.graph.vertexes;
        for i in 0..vertexes.len() {
            if vertex == vertexes[i] {
                return i as i32;
            }
        }
        return -1;
    }

    /// 获取指定下标顶点的终点下标
    /// posits: 存放顶点和对应终点下标，posits 的下标表示顶点下标，值表示对应顶点的终点下标
    /// i: 目标顶点下标
    /// posits = [0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0], i = 4 返回 5
    fn get_end_posit(&self, posits: &Vec<u32>, i: usize) -> usize {
        let mut index = i;
        while posits[index] != 0 {
            index = posits[index] as usize;
        }
        return index;
    }

    /// 获取所有边
    fn get_edges(&self) -> Vec<Edge> {
        let mut edges = Vec::new();
        let graph = &self.graph;
        let num_of_vertex = graph.vertexes.len();
        for i in 0..num_of_vertex {
            for j in i + 1..num_of_vertex {
                // 不连通的跳过
                if graph.matrix[i][j] == u32::MAX {
                    continue;
                }
                // 创建边
                edges.push(Edge::new(
                    graph.vertexes[i].clone(),
                    graph.vertexes[j].clone(),
                    graph.matrix[i][j],
                ));
            }
        }
        edges
    }
}

#[test]
fn test_build_bus_station() {
    let vertexes = vec![
        "A".to_string(),
        "B".to_string(),
        "C".to_string(),
        "D".to_string(),
        "E".to_string(),
        "F".to_string(),
        "G".to_string(),
    ];
    let edges = vec![
        vec![0, 12, u32::MAX, u32::MAX, u32::MAX, 16, 14],
        vec![12, 0, 10, u32::MAX, u32::MAX, 7, u32::MAX],
        vec![u32::MAX, 10, 0, 3, 5, 6, u32::MAX],
        vec![u32::MAX, u32::MAX, 3, 0, 4, u32::MAX, u32::MAX],
        vec![u32::MAX, u32::MAX, 5, 4, 0, 2, 8],
        vec![16, 7, 6, u32::MAX, 2, 0, 9],
        vec![14, u32::MAX, u32::MAX, u32::MAX, 8, 9, 0],
    ];
    let min_tree = MinTree::new(vertexes, edges);
    min_tree.show_graph();
    min_tree.kruskal();
}
