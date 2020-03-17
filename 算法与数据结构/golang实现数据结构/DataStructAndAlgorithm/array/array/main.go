package array

import (
	"errors"
	"fmt"
)

type Student struct {
	name string
	score int
}

// Array为自定义动态数组，动态扩容，动态缩容
type Array struct {
	data []interface{}
	size int
}

// 构造函数, 传入数组的容量capacity构造Array
func NewArray(capacity int) (arr *Array) {
	arr = new(Array)
	arr.data = make([]interface{}, capacity)
	arr.size = 0
	return arr
}

// 无参数的构造函数，默认数组的容量capacity=10
func NewArray1()  (arr *Array) {
	arr = new(Array)
	arr.data = make([]interface{}, 10)
	arr.size = 0
	return arr
}

// 将静态数组重构为动态数组
func NewArray2(e []interface{}) (arr *Array) {
	arr = new(Array)
	arr.data = e
	arr.size = len(e)
	return arr
}

// 获取数组中的元素个数
func (arr *Array) GetSize() int{
	return arr.size
}

func (arr *Array) GetCapacity() int{
	return len(arr.data)
}

// 返回数组是否为空
func (arr *Array) IsEmpty() bool {
	return arr.size == 0
}

// 向数组末添加元素,size为末尾空的索引，只需对size索引的位置赋值，并维护size
func (arr *Array) AddLast(e interface{}) (err error) {
	err = arr.Add(arr.size, e)
	return err
}

// 在所有元素前添加一个新元素
func (arr * Array) AddFirst(e interface{}) (err error) {
	err = arr.Add(0, e)
	return err
}

// 在第index个位置插入一个新的元素e
func (arr * Array) Add(index int, e interface{}) (err error){
	// 保障有足够的空间来存储新元素
	//if arr.size == len(arr.data){
	//	return errors.New("add failed, Array is full")
	//}

	// 判断index索引要符合素组的索引值区间
	if index < 0 || index >= arr.size {
		err = errors.New("add failed, Require index >= 0 and < Array.size")
		return err
	}

	// 如果array满了，则进行2倍扩容
	if arr.size == len(arr.data) {
		arr.resize(len(arr.data) << 1)
	}

	// 将索引index及以后的元素都向后移动一个位置，空出index位置用于赋值新元素e
	for i := arr.size-1; i >= index; i-- {
		arr.data[i+1] = arr.data[i]
	}
	// 此时index处元素仍然存在，复制一份其副本放在了index+1处，可以放心覆盖掉index处的值了
	arr.data[index] = e
	arr.size += 1
	return  err
}

// 修改index索引位置的元素
func (arr *Array) Set(index int, e interface{}) (err error){
	if index < 0 || index >= arr.size {
		err = errors.New("get failed, index is illegal")
		return err
	}
	arr.data[index] = e
	return err
}

// 获取index索引位置的元素
func (arr *Array) Get(index int) (data interface{}, err error){
	if index < 0 || index >= arr.size {
		err = errors.New("get failed, index is illegal")
		return nil, err
	}
	data = arr.data[index]

	return data, err
}
// 获取最后一个位置的元素
func (arr *Array) GetLast() (interface{}, error){
	return arr.Get(arr.size - 1)
}

// 获取第一个位置的元素
func (arr *Array) GetFirst() (interface{}, error) {
	return arr.Get(0)
}

// 打印自定义结构对象信息
func (arr *Array) String() (res string) {
	res = fmt.Sprintf("Array: size = %d, capacity = %d\n", arr.size, len(arr.data))
	res += "["
	for i:=0; i <arr.size; i++{
		res += fmt.Sprintf("%v", arr.data[i])
		if i != arr.size -1 {
			res += ", "
		}
	}
	res += "]"
	return res
}

// 查找数组中是否有元素e
func (arr *Array) Contains(e interface{}) bool {
	for i:= 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return true
		}
	}
	return false
}

// 查找数组中元素e所在的索引，如果不存在元素e,则返回-1
func (arr *Array) Find(e interface{}) int {
	for i :=0; i<arr.size; i++ {
		if arr.data[i] == e {
			return i
		}
	}
	return -1
}

// 删除指定位置元素： 算法是将索引右侧的元素向左挪动一个位置，维护size减小一个数字
func (arr *Array) Remove(index int) (ret interface{}, err error) {
	if index < 0 || index >= arr.size {
		err = errors.New("remove failed, index is illegal")
		return nil, err
	}
	ret = arr.data[index]
	for i:= index + 1; i < arr.size; i++ {
		arr.data[i-1] = arr.data[i]
	}
	arr.size -= 1
	// 这一步不是必须的，因为size重新指定后，size+1的位置虽然还有值但是不会再被看到，再新的addLast函数执行时会被覆盖掉
	arr.data[arr.size] = nil
	// lazy方式的缩容，如果删除元素后，size是capacity的1/4，则进行缩容，缩容1/2
	if arr.size == len(arr.data) >> 2 {
		arr.resize(len(arr.data) >> 1)
	}
	return ret, err
}

// 从数组中删除第一个元素，返回删除的元素
func (arr *Array) RemoveFirst() (interface{}, error) {
	return arr.Remove(0)
}

// 从数组中删除最后一个元素，返回删除的元素
func (arr *Array) RemoveLast() (interface{}, error) {
	return arr.Remove(arr.size - 1)
}

// 从数组中删除元素e,如果存在e则删除e，只删除一个e，否则什么事情都不干
func (arr *Array) RemoveElement(e interface{}) {
	index := arr.Find(e)
	if index != -1 {
		_, err := arr.Remove(index)
		if nil != err {
			fmt.Println(err)
		}
	}
}

// 动态数组扩容，在add时如果检查到容量已经满了，可以选择add失败；或者本方法的resize进行扩容再add
func (arr *Array) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < arr.size; i++ {
		newData[i]= arr.data[i]
	}
	arr.data = newData
}
