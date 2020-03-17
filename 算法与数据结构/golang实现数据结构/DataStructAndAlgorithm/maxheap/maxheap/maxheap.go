package maxheap

import (
	"DataStructAndAlgorithm/array/array"
	"DataStructAndAlgorithm/utils"
	"errors"
)

type MaxHeap struct {
	array *array.Array
}

// 新建基于数组的最大堆实例,设置基础容量
func NewMaxHeap(capacity int) *MaxHeap {
	return & MaxHeap{array:array.NewArray(capacity)}
}

// 新建基于数组的最大堆实例, 基础容量使用默认值
func NewMaxHeap1() *MaxHeap {
	return & MaxHeap{array:array.NewArray1()}
}

// 将任意数组整理堆的结构，并返回生成的堆对象，（会修改原数组e的元素为堆的顺序）
func HeapifyNewMaxHeap (e []interface{}) *MaxHeap {
	data := MaxHeap{array:array.NewArray2(e)}
	// 循环siftDown每个节点元素，若其大于与其父节点进行swap
	for i, _ := data.Parent(len(e)-1); i >=0; i-- {
		data.siftDown(i)
	}
	return &data
}

// 返回堆中的元素个数
func (h *MaxHeap) GetSize() int {
	return h.array.GetSize()
}

// 返回布尔值，表示堆中是否为空
func (h *MaxHeap) IsEmpty() bool{
	return h.array.IsEmpty()
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的父节点的索引
func (h *MaxHeap) Parent(index int) (int, error){
	if index == 0 {
		err := errors.New("index-0 doesn't have parent")
		return 0, err
	}
	return (index - 1) >> 1, nil
}

// 返回万全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func (h *MaxHeap) LeftChild(index int) int {
	return index << 1 + 1
}

// 返回万全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func (h *MaxHeap) RightChild(index int) int {
	return index << 1 + 2
}

// 向堆中添加元素,先默认加到最后位置（满足对特性1）；然后进行上浮操作，确保新元素比其父节点小或等
func (h *MaxHeap) Add(e interface{}) {
	h.array.AddLast(e)  // 对新增节点值先加入到最后位置
	h.siftUp(h.GetSize() - 1)  // 对新增节点即最后一个节点进行上浮操作
}

// 交换两个索引代表的节点的位置
func (h *MaxHeap) swap(i int,j int) error {
	if i < 0 || j < 0 || i >= h.array.GetSize() || j >= h.array.GetSize() {
		err := errors.New("index is illegal")
		return err
	}
	ti, _ := h.array.Get(i)
	tj, _ := h.array.Get(j)
	h.array.Set(i, tj)
	h.array.Set(j, ti)
	return nil
}

// 堆中添加元素后的上浮操作, k表示k节点的index
func (h *MaxHeap) siftUp(k int) {
	for {
		pk,err := h.Parent(k)  // 获取k节点的父节点的索引
		if err != nil {
			return
		}
		kVal,_ := h.array.Get(k) // 获取k节点的值
		pVal,_ := h.array.Get(pk) //获取k父节点的值
		if k > 0 && utils.Compare(pVal, kVal) < 0 {  // 如果k节点的值大于其父节点，那么进行swap
			h.swap(k, pk)
			k = pk
		} else {
			break
		}
	}
}

// 堆中取出堆顶元素后的下沉操作, k表示k节点的index
func (h *MaxHeap) siftDown(k int) {
	for h.LeftChild(k) < h.array.GetSize(){  // 对于数组堆来说，下沉操作，k的值从0开始，小于最大索引值
		j := h.LeftChild(k)  // k的左孩子的索引
		e1, _ := h.array.Get(j+1) // 获得right孩子的值
		e2, _ := h.array.Get(j) // 获取left孩子的值
		if j +1 < h.array.GetSize() && utils.Compare(e1, e2) > 0 {
			j = h.RightChild(k)
		}
		// 左右孩子进行比较，得到最大值的孩子节点，j记录最大值孩子节点的索引
		// 再将k节点与其最大子节点值进行比较，如果k的值不小于子孩子的值，则结束，否则swap进行下沉替换操作
		ek, _ := h.array.Get(k)
		ej, _ := h.array.Get(j)
		if utils.Compare(ek, ej) >= 0 {
			break
		} else {
			h.swap(k, j)
			k = j
		}
	}
}

// 查看堆中的最大元素
func (h *MaxHeap) FindMax() (e interface{}, err error) {
	if h.array.GetSize() == 0 {
		err := errors.New("cannot find max when heap is empty")
		return nil, err
	}
	return h.array.Get(0)
}

// 堆中取出元素(最大)（对于堆二叉树是根节点，对于数组表示来说是index=0的元素）
func (h *MaxHeap) ExtractMax() (ret interface{}, err error) {
	ret, err = h.FindMax() // 找到堆中最大元素
	if err != nil {
		return ret, err
	}
	h.swap(0, h.array.GetSize() - 1) //将堆中最大元素与最末尾元素进行swap
	h.array.RemoveLast() // 将swap后的最后元素删除
	h.siftDown(0) //最后元素被替换到堆的根位置，需要进行下沉操作
	return ret, err
}

// 取出堆中的最大元素，并且替换成元素e
func (h *MaxHeap) Replace(e interface{}) interface{} {
	ret, _ := h.FindMax()
	h.array.Set(0, e)
	h.siftDown(0)
	return ret
}

