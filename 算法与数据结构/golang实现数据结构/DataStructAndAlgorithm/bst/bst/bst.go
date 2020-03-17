package bst

import (
	"DataStructAndAlgorithm/linkedlist/linkedqueue"
	"DataStructAndAlgorithm/linkedlist/linkedstack"
	"DataStructAndAlgorithm/utils"
	"errors"
	"fmt"
)

// 为了方便二叉搜索树中的元素插入是比较，默认Node节点元素e值为int；
// 否则将要找一个合适的方式来比较两个元素的大小，然后才能插入数据
type Node struct {
	e interface{}
	left *Node
	right * Node
}

type Bst struct {
	root *Node
	size int
}

// 新建二叉树节点
func NewNode(e interface{}) *Node {
	return &Node{e: e, left: nil, right: nil}
}

// 新建二分搜索树
func NewBst() *Bst {
	return &Bst{root: nil, size: 0}
}

// 获取bst的元素个数
func (b *Bst) GetSize() int {
	return b.size
}

// 判断bst是否为空
func (b *Bst) IsEmpty() bool {
	return b.size == 0
}

// bst添加元素
func (b *Bst) Add1(e interface{}) (err error){
	if b.root == nil {
		b.root = NewNode(e)
	} else {
		b.add1(b.root, e)
	}
	return err
}

// 插入方法1：向以node为根的bst中插入元素,递归算法
func (b *Bst) add1(node *Node, e interface{}) {
	if utils.Compare(e, node.e) == 0 {
		return
	} else if utils.Compare(e, node.e) < 0 && node.left == nil {
		node.left = NewNode(e)
		b.size ++
		return
	} else if utils.Compare(e, node.e) > 0 && node.right == nil {
		node.right = NewNode(e)
		b.size ++
		return
	}
	if utils.Compare(e, node.e) < 0 {
		b.add1(node.left, e)
	} else {
		b.add1(node.right, e)
	}
}


// bst添加元素：改进方法add1,更简洁
func (b *Bst) Add(e interface{}) (err error) {
	b.root = b.add(b.root, e)
	return err
}

// 插入方法2：改进方法add1,代码更简洁，需返回插入新节点后bst的根，仍然是递归算法
func (b *Bst) add(node *Node, e interface{}) *Node {
	if node == nil {
		b.size ++
		return NewNode(e)
	}
	if utils.Compare(e, node.e) < 0 {
		node.left = b.add(node.left, e)
	} else if utils.Compare(e, node.e) > 0 {
		node.right = b.add(node.right, e)
	}
	return node
}

// 查询bst中是否包含元素e
func (b *Bst) Contains(e interface{}) bool {
	return b.contains(b.root, e)
}

// 看以node为根的bst中是否包含元素e,递归算法
func (b *Bst) contains(node *Node, e interface{}) bool {
	if node == nil {
		return false
	}
	if e == node.e {
		return true
	} else if utils.Compare(e, node.e) < 0 {
		return b.contains(node.left, e)
	} else {
		return b.contains(node.right, e)
	}
}

// 前序遍历，遍历打印bst中的元素,先访问节点，再访问left,right子树,递归实现
func (b *Bst) PreOrder() {
	b.preOrder(b.root)
}

// 前序遍历，并打印元素的递归算法， 也可以根据业务需求，改为做修改或其它操作
func (b *Bst) preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.e)
	b.preOrder(node.left)
	b.preOrder(node.right)
}

// 前序遍历，打印bst中的元素,先访问节点，再访问left,right子树，非递归实现--使用栈来记录访问顺序，模拟系统栈的访问顺序
// 由于栈是先进后出的方式，要先访问顺序为前序遍历（node,lef,right）,那么栈元素入队后应该为(right, left),
// node要在之前就处理
func (b *Bst) PreOrderNR() {
	st, _ := linkedstack.NewLinkedListStack()
	st.Push(b.root)
	for !st.IsEmpty() {
		cur, _ := st.Pop()
		node := cur.(*Node)
		if node.right != nil {
			st.Push(node.right)
		}
		if node.left != nil {
			st.Push(node.left)
		}
	}
}

// 格式化输出树节点元素信息，基于前序遍历
func (b *Bst) String() string {
	res := ""
	b.generateBSTString(b.root, 0, &res)
	return res
}

// 生成以node为根节点，深度为depth的描述二叉树的字符串
func (b *Bst) generateBSTString(node *Node, depth int, res *string) {
	if node == nil {
		*res += b.generateDepthString(depth) + "nil\n"
		return
	}
	*res += b.generateDepthString(depth) + fmt.Sprintf("%v\n", node.e)
	b.generateBSTString(node.left, depth + 1, res)
	b.generateBSTString(node.right, depth + 1, res)
}

// 根据depth返回表示depth的深度字符串
func (b *Bst) generateDepthString(depth int) string {
	res := ""
	for i:=0; i < depth; i++ {
		res += "--"
	}
	return res
}

// 二分搜索树的中序遍历,并打印节点元素信息,先left子树，然后node节点，最后right子树
// 中序遍历的结果就是树中所有元素排序后的结果
func (b *Bst) InOrder() {
	b.inOrder(b.root)
}

// 中序遍历的递归算法实现
func (b *Bst) inOrder(node *Node) {
	if node == nil {
		return
	}
	b.inOrder(node.left)
	fmt.Println(node.e)
	b.inOrder(node.right)
}

// 二分搜索树的后续遍历,并打印节点元素信息,先left子树，然后right子树，最后node节点
func (b *Bst) PostOrder() {
	b.postOrder(b.root)
}

// 后续遍历的递归算法实现
func (b *Bst) postOrder(node *Node) {
	if node == nil {
		return
	}
	b.postOrder(node.left)
	b.postOrder(node.right)
	fmt.Println(node.e)
}

// 层序遍历，一般实现用非递归算法，需借用队列数据结构来实现
func (b *Bst) LevelOrder() {
	q := linkedqueue.NewLinkedQueue()
	q.Enqueue(b.root)
	for !q.IsEmpty() {
		cur, _ := q.Dequeue()
		node := cur.(*Node)
		fmt.Println(node.e)
		if node.left != nil {
			q.Enqueue(node.left)
		}
		if node.right != nil {
			q.Enqueue(node.right)
		}
	}
}

// 寻找bst中的最小元素
func (b *Bst) Minimum() (e interface{}, err error) {
	if b.size == 0 {
		err = errors.New("bst is empty")
		return e, err
	}
	return b.minimum(b.root).e, err
}

// 返回以node为根的bst的最小值所在的节点
func (b *Bst) minimum(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return b.minimum(node.left)
}

// 寻找bst中的最大元素
func (b *Bst) Maximum() (e interface{}, err error) {
	if b.size == 0 {
		err = errors.New("bst is empty")
		return e, err
	}
	return b.maximum(b.root).e, err
}

// 返回以node为根的bst的最大值所在的节点
func (b *Bst) maximum(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return b.maximum(node.right)
}

// 删除bst中的最小值， 从根节点向左走，最左边节点的值
func (b *Bst)RemoveMin() (ret *Node, err error) {
	// 对于最小值是最左边节点的值，该节点一定没有left子树
	ret = b.removeMin(b.root)
	return ret, err
}

// 删除掉以node为根的二分搜索树中的最小节点，返回删除节点后新的二分搜索树的根
func (b *Bst) removeMin(node *Node) *Node {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		b.size --
		return rightNode
	}
	node.left = b.removeMin(node.left)
	return node
}

// 删除bst中的最大值， 从根节点向右走，最右边节点的值
func (b *Bst)RemoveMax() (ret *Node, err error) {
	// 对于最大值是最右边节点的值，该节点一定没有right子树
	ret = b.removeMax(b.root)
	return ret, err
}

// 删除掉以node为根的二分搜索树中的最大节点，返回删除节点后新的二分搜索树的根
func (b *Bst) removeMax(node *Node) *Node {
	if node.right == nil {
		leftNode := node.left
		node.left = nil
		b.size --
		return leftNode
	}
	node.right = b.removeMax(node.right)
	return node
}

// 删除bst中的任意值
func (b *Bst) Remove(e interface{}) {
	b.root = b.remove(b.root, e)
}

// 删除以node为根的二分搜索树中值为e的节点，递归方法，返回删除节点后新的二分搜索树的根
func (b *Bst) remove(node *Node, e interface{}) *Node {
	if node == nil {
		return nil
	}
	if utils.Compare(e, node.e) < 0 {
		node.left = b.remove(node.left, e)
		return node
	} else if utils.Compare(e, node.e) > 0 {
		node.right = b.remove(node.right, e)
		return node
	} else {
		// 待删除节点left子树为空的情况
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			b.size --
			return rightNode
		}
		// 待删除节点right子树为空的情况
		if node.right == nil {
			leftNode := node.left
			node.left = nil
			b.size --
			return leftNode
		}
		// 带删除节点的左右子树均不为空的情况
		// 找到待删节点的替代节点（后继节点或前驱节点）：
		// 		1）可以找比待删除节点大的最小节点，即待删除节点的right子树的最小节点；
		// 		2）或者可以找比待删除节点小的最大节点，即待删节点的left子树中最大节点；
		// 用这个节点顶替待删除节点的位置
		successor := b.minimum(node.right) // 这里使用的第1种方式right子树的最小节点作为后继节点
		successor.right = b.removeMin(node.right)
		successor.left = node.left
		node.left = nil
		node.right = nil
		return successor
	}
}

