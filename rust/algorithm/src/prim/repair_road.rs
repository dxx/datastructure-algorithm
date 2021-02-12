use std::collections::HashMap;

/// 普里姆(Prim)算法解决村庄修路问题。
/// Prim 算法适用于稠密图
/// 算法思路：
/// 1.首先随便选一个点加入集合。
/// 2.用该点的所有边去刷新到其它点的最短路。
/// 3.找出最短路中最短的一条连接（且该点未被加入集合）。
/// 4.用该点去刷新到其他点的最短路。
/// 5.重复以上操作 n-1 次。

/// 最小生成树
pub struct MinTree {
    graph: Graph,
}

pub struct Graph {
    vertexes: Vec<String>, // 顶点
    matrix: Vec<Vec<u32>>, // 邻接矩阵，代表边的值
}

impl MinTree {
    pub fn new(vertexes: Vec<String>, edges: Vec<Vec<u32>>) -> Self {
        let graph = Graph {
            vertexes,
            matrix: edges,
        };
        MinTree { graph }
    }

    pub fn prim(&self, v: u32) {
        let graph = &self.graph;
        let num_of_vertex = graph.vertexes.len();
        // 存放已经连通的顶点集合
        let mut vertex_map = HashMap::new();
        // 将当前顶点加入集合
        vertex_map.insert(
            graph.vertexes[v as usize].clone(),
            graph.vertexes[v as usize].clone(),
        );

        // 记录顶点下标
        let mut v1: i32 = -1;
        let mut v2: i32 = -1;
        // 记录最小边的权值，初始化成一个最大数，后续遍历中会被替换
        let mut min_weight = u32::MAX;
        // n 个顶点就有 n-1 条边
        for _k in 1..num_of_vertex {
            // 查找已经加入集合中的顶点，和这些顶点中最近的一个顶点
            for i in 0..num_of_vertex {
                for j in 0..num_of_vertex {
                    let weight = graph.matrix[i][j];
                    let v = vertex_map.get(&graph.vertexes[i]);
                    if (v.is_some() && *v.unwrap() == graph.vertexes[i]) // 表示已经加入集合的顶点
                        && vertex_map.get(&graph.vertexes[j]).is_none() // 表示未被加入集合的顶点
                        && weight < min_weight
                    {
                        v1 = i as i32;
                        v2 = j as i32;
                        min_weight = weight;
                    }
                }
            }
            // 将最小的顶点加入到集合中
            vertex_map.insert(
                graph.vertexes[v2 as usize].clone(),
                graph.vertexes[v2 as usize].clone(),
            );
            // 修改最小的权值
            min_weight = u32::MAX;

            println!(
                "边：{}-{} => {}",
                graph.vertexes[v1 as usize],
                graph.vertexes[v2 as usize],
                graph.matrix[v1 as usize][v2 as usize]
            );
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
}

#[test]
fn test_repair_road() {
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
        vec![u32::MAX, 5, 7, u32::MAX, u32::MAX, u32::MAX, 2],
        vec![5, u32::MAX, u32::MAX, 9, u32::MAX, u32::MAX, 3],
        vec![7, u32::MAX, u32::MAX, u32::MAX, 8, u32::MAX, u32::MAX],
        vec![u32::MAX, 9, u32::MAX, u32::MAX, u32::MAX, 4, u32::MAX],
        vec![u32::MAX, u32::MAX, 8, u32::MAX, u32::MAX, 5, 4],
        vec![u32::MAX, u32::MAX, u32::MAX, 4, 5, u32::MAX, 6],
        vec![2, 3, u32::MAX, u32::MAX, 4, 6, u32::MAX],
    ];
    let min_tree = MinTree::new(vertexes, edges);
    min_tree.show_graph();
    // 从 A 点开始
    min_tree.prim(0);
}
