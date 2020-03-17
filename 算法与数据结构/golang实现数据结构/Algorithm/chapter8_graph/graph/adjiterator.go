package graph


type AdjIterator interface {
	Begin() int
	Next() int
	End() bool
}

// 稠密阵列图的相邻边的迭代器
type DenseGraphAdjIterator struct {
	G *DenseGraph // G是一个二维数组
	v int // v是节点的索引，即第一层数组的索引
	index int // v节点也是一个数组切片，index是这个数组切片的索引，即第二层数组的索引
}

// 稀疏阵列的相邻边的迭代器
type SparseGraphAdjIterator struct {
	G *SparseGraph  // G是一个二维数组
	v int // v是节点的索引，即第一层数组的索引
	index int   // v节点也是一个数组切片，index是这个数组切片的索引，即第二层数组的索引
}

// 新建稀疏阵列图AdjIterator实例对象
func NewSparseGraphAdjIterator(graph *SparseGraph, v int) *SparseGraphAdjIterator {
	return &SparseGraphAdjIterator{G:graph, v: v, index:0}
}

// 新建稠密阵列图AdjIterator实例对象
func NewDenseGraphAdjIterator(graph *DenseGraph, v int) *DenseGraphAdjIterator {
	return &DenseGraphAdjIterator{G:graph, v: v, index:-1}
}

// 开始迭代
func (a *DenseGraphAdjIterator) Begin() int {
	a.index = -1
	return a.Next()
}

// 迭代下一个元素
func (a *DenseGraphAdjIterator) Next() int {
	for a.index+= 1; a.index < len(a.G.g[a.v]); a.index ++ {
		if a.G.g[a.v][a.index] == true {
			return a.index
		}
	}
	return -1
}

// 判断是否已迭代完
func (a *DenseGraphAdjIterator) End() bool {
	return a.index >= len(a.G.g)
}


// 开始迭代
func (a *SparseGraphAdjIterator) Begin() int {
	if len(a.G.g[a.v]) > 0 {
		return a.G.g[a.v][a.index]
	}
	return -1
}

// 迭代下一个元素
func (a *SparseGraphAdjIterator) Next() int {
	a.index ++
	if a.index < len(a.G.g[a.v]) &&  len(a.G.g[a.v]) > 0 {
		return a.G.g[a.v][a.index]
	}
	return -1
}

// 判断是否已迭代完
func (a *SparseGraphAdjIterator) End() bool {
	return a.index >= len(a.G.g[a.v])
}
