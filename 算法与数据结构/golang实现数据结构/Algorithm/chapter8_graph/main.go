package main

import (
	"Algorithm/chapter8_graph/graph"
	"fmt"
	"math/rand"
)


// 测试稀疏阵列图的相邻边的迭代器
func TestSparseGraph() {
	N := 20
	M := 100
	g1 := graph.NewSparseGraph(N, false)
	for i:=0; i < M; i++ {
		a := rand.Intn(N)
		b := rand.Intn(N)
		g1.AddEdge(a, b)
	}
	for v:=0; v < N; v++ {
		var adj graph.AdjIterator
		adjInstance := graph.NewSparseGraphAdjIterator(g1, v) // 生成关于v节点相邻边的迭代器
		adj = adjInstance
		res := fmt.Sprintf("%d : ", v)
		for w := adj.Begin();adj.End() == false; w = adj.Next() {
			res += fmt.Sprintf("%d ", w)
		}
		fmt.Println(res)
	}
}
// 测试相邻边的迭代器
func TestDenseGraph() {
	N := 20
	M := 100
	g1 := graph.NewDenseGraph(N, false)
	for i:=0; i < M; i++ {
		a := rand.Intn(N)
		b := rand.Intn(N)
		g1.AddEdge(a, b)
	}
	for v:=0; v < N; v++ {
		var adj graph.AdjIterator
		adjInstance := graph.NewDenseGraphAdjIterator(g1, v) // 生成关于v节点相邻边的迭代器
		adj = adjInstance
		res := fmt.Sprintf("%d : ", v)
		for w := adj.Begin();adj.End() == false; w = adj.Next() {
			res += fmt.Sprintf("%d ", w)
		}
		fmt.Println(res)
	}
}
// 测试从文件中读取图点、边数据，并生成稀疏阵列
func TestReadGraphWithSparse() {
	filename := "E:/gopath/src/Algorithm/chapter8_graph/testG1.txt"
	var gra graph.Graph
	sparse := graph.NewSparseGraph(13, false)
	gra = sparse
	g1 := graph.ReadGraph(&gra, filename)
	fmt.Println(*g1)
}

// 测试从文件中读取图点、边数据，并生成稠密阵列
func TestReadGraphWithDense() {
	filename := "E:/gopath/src/Algorithm/chapter8_graph/testG1.txt"
	var gra graph.Graph
	sparse := graph.NewDenseGraph(13, false)
	gra = sparse
	g1 := graph.ReadGraph(&gra, filename)
	fmt.Println(*g1)
}

func main() {
	//sparsegraph.TestSparseGraph()
	//densegraph.TestDenseGraph()
	TestReadGraphWithSparse()
	TestReadGraphWithDense()
}
