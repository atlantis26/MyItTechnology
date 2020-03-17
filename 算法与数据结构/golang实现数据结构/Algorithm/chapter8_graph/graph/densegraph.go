package graph

import (
	"errors"
	"fmt"
)

// 稠密图--邻接矩阵
type DenseGraph struct {
	N int  // 节点
	M int  // 边
	directed bool // 是否有向图
	g [][]bool
}

// 新建初始化DenseGraph,g为行、列都为长度的矩阵
func NewDenseGraph(n int,  directed bool) *DenseGraph {
	graph := DenseGraph{N:n, M:0, directed:directed}
	for i:=0; i < n; i++ {
		v := make([]bool, n)
		for j:=0; j < n; j++ {
			v[j] = false
		}
		graph.g = append(graph.g, v)
	}
	return &graph
}

// 建立边，在v节点和w节点之间
func (gra *DenseGraph) AddEdge(v int, w int) (err error) {
	if !(v >= 0 && v < gra.N) {
		err = errors.New("argument n is illegal")
		return err
	}
	if !(w >=0 && w < gra.N) {
		err = errors.New("argument w is illegal")
		return err
	}
	// 判断两点之间是否已经有边了，有了则不再加边和m++操作
	if ret, _ := gra.HasEdge(v, w); ret == true {
		return err
	}
	gra.g[v][w] = true
	if gra.directed == false { // 如果是无向图，则没有指向directed；等价于边的两端都有向的
		gra.g[w][v] = true
	}
	gra.M ++
	return err
}

// 判断两点之间是否有边
func (gra *DenseGraph) HasEdge(v int, w int) (ret bool, err error) {
	if !(v >= 0 && v < gra.N) {
		err = errors.New("argument n is illegal")
		return ret, err
	}
	if !(w >=0 && w < gra.N) {
		err = errors.New("argument w is illegal")
		return ret, err
	}
	return gra.g[v][w], err
}

func (gra *DenseGraph) String() string{
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
func (gra *DenseGraph) V() int {
	return gra.N
}

// 修改设置图的节点个数信息
func (gra *DenseGraph) SetV(n int) {
	gra.N = n
}