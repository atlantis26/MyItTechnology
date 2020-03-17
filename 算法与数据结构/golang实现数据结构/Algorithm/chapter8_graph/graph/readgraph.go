package graph

import (
	"bufio"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Graph interface {
	AddEdge(v int, w int) (err error)
	V() int
	SetV(int)
}

// 将图的点、边数据从文件中读取，并添加到实例中，参数graph可以是稀疏阵列也可以是稠密阵列
func ReadGraph(graph *Graph, filename string) *Graph {
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	rd := bufio.NewReader(fp)
	firstLine, err := rd.ReadString('\n')
	nm := strings.Split(firstLine, " ")
	_, _ = strconv.Atoi(strings.Replace(strings.Replace(nm[0], " ", "", -1),"\r\n", "", -1))
	// 获得graph的值反射对象vg
	vg := reflect.ValueOf(graph).Elem()

	//vg.FieldByName("N").Set(reflect.ValueOf(n)) // 通过反射设置图对象的节点个数n值 ?怎么会失败#todo
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		vw := strings.Split(line, " ")
		v, err := strconv.Atoi(strings.Replace(strings.Replace(vw[0], " ", "", -1),"\r\n", "", -1))
		if err != nil {
			panic(err)
		}
		w, err := strconv.Atoi(strings.Replace(strings.Replace(vw[1], " ", "", -1),"\r\n", "", -1))
		if err != nil {
			panic(err)
		}
		params := make([]reflect.Value, 2)
		params[0] = reflect.ValueOf(v)
		params[1] = reflect.ValueOf(w)
		vg.MethodByName("AddEdge").Call(params) // 通过反射添加图中v,w两点的边
	}
	return graph
}
