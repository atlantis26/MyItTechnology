package graph

import (
	"errors"
	"fmt"
)

// 稀疏图--邻接表
type SparseGraph struct {
	N int  // 节点
	M int  // 边
	directed bool // 是否有向图
	g [][]int
}

// 新建初始化DenseGraph,g为行、列都为长度的矩阵
func NewSparseGraph(n int,  directed bool) *SparseGraph {
	graph := SparseGraph{N:n, M:0, directed:directed}
	for i:=0; i < n; i++ {
		graph.g = append(graph.g, []int {})
	}
	return &graph
}

// 加边，对v,w两个节点之间加边
func (gra *SparseGraph) AddEdge(v int, w int) (err error){
	if !(v >= 0 && v < gra.N) {
		err = errors.New("argument n is illegal")
		return err
	}
	if !(w >=0 && w < gra.N) {
		err = errors.New("argument w is illegal")
		return err
	}
	gra.g[v] = append(gra.g[v], w)
	if v != w && gra.directed == false {  //如果v,w不是同一个节点，且是无需的边，则两端都加上向来表示无向
		gra.g[w] = append(gra.g[w], v)
	}
	return err
}

func (gra *SparseGraph) String() string{
	res := ""
	for i:=0; i < gra.N; i++ {
		count := fmt.Sprintf("vertex %d:\t", i)
		for j:= 0; j < len(gra.g[i]); j++ {
			count += fmt.Sprintf("%v\t",gra.g[i][j])
		}
		count += "\n"
		res += count
	}
	return res
}

// 获取图的节点个数信息
func (gra *SparseGraph) V() int {
	return gra.N
}

// 修改设置图的节点个数信息
func (gra *SparseGraph) SetV(n int) {
	gra.N = n
}
