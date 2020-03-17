package graph

// 寻路算法
type Path struct {
	G *Graph
	s int
	visited []bool
	from []int
}

//// 初始化Path
//func  NewPath(graph *Graph, s int) {
//	p := Path{}
//	if !(s >= 0 && s < p.G.V()) {
//		panic("s is error")
//	}
//	p.visited =
//
//}

// 原点到w是否有路径
func (p *Path) HasPath(w int) {

}

// 原点到w是否有路径，存到vec中
func path(w int, vec []int) {

}

// 格式化输出路径
func ShowPath(w int) {

}
