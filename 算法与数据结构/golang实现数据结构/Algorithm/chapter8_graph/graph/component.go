package graph

import (
	"reflect"
)

type Component struct {
	G *Graph
	visited []bool
	ccount int
}

// 初始化Component对象
func NewComponent(graph *Graph) *Component {
	val :=  reflect.ValueOf(graph).Interface().(Graph)
	visited := make([]bool, val.V())
	for i:=0; i < val.V(); i++ {
		visited[i] = false
	}
	return &Component{G: graph, visited: visited, ccount: 0}
}

// dfs计算联通分量(深度优先遍历方式)，即图中的没有相连的子图有多少个
func (c *Component) Dfs(v int) {
	c.visited[v] = true
	adj := AdjIterator{}
	for i:= adj.Begin(); adj.End()==false; i=adj.Next(){
		if c.visited[i] == false {
			c.Dfs(i)
		}
	}
}

// 返回有多少个联通分量
func (c *Component) Count() int{
	return c.ccount
}