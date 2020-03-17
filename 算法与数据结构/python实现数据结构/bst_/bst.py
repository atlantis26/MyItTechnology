# coding:utf-8
from linkedlist_.linkedstack import LinkedStack
from linkedlist_.linkedqueue import LinkedQueue


class Node(object):
    """
    二分搜索树的节点
    """
    def __init__(self, e, left=None, right=None):
        self.e = e
        self.left = left
        self.right = right


class Bst(object):
    """
    二分搜索树
    """
    def __init__(self, root=None):
        self.root = root
        self.size = 0

    def is_empty(self):
        """
        判断bst是否为空
        :return:
        """
        return self.size == 0

    def get_size(self):
        """
        获取bst的元素个数
        :return:
        """
        return self.size

    def add1(self, e):
        """
        bst添加元素
        :return:
        """
        if self.root is None:
            self.root = Node(e)
        else:
            self._add1(self.root, e)

    def _add1(self, node, e):
        """
        向以node为根的bst中插入元素，递归算法
        :param node:
        :param e:
        :return:
        """
        if e == node.e:
            return
        elif e < node.e:
            if node.left is None:
                node.left = Node(e)
                self.size += 1
            else:
                self._add1(node.left, e)
        else:
            if node.right is Node:
                node.right = Node(e)
                self.size += 1
            else:
                self._add1(node.right, e)

    def add(self, e):
        """
        bst添加元素，改进方法，，更简洁
        :param e:
        :return:
        """
        self.root = self._add(self.root, e)

    def _add(self, node, e):
        """
        改进方法add1,返回插入新节点后bst的根，递归算法
        :param node:
        :param e:
        :return:
        """
        if node is None:
            self.size += 1
            return Node(e)
        if e < node.e:
            node.left = self._add(node.left, e)
        elif e > node.e:
            node.right = self._add(node.right, e)
        return node

    def contains(self, e):
        """
        查询bst中是否包含元素e
        :param e:
        :return:
        """
        return self._contains(self.root, e)

    def _contains(self, node, e):
        """
        以node为根的bst中是否包含元素e,递归算法
        :param node:
        :param e:
        :return:
        """
        if node is None:
            return False
        if e == node.e:
            return True
        elif e < node.e:
            return self._contains(node.left, e)
        else:
            return self._contains(node.right, e)

    def pre_order(self):
        """
        前序遍历，遍历打印bst中的元素，先访问节点，再访问left,right子树，递归实现
        :return:
        """
        self._pre_order(self.root)

    def _pre_order(self, node):
        """
        前序遍历，并打印元素的递归算法，也可以根据业务需求，改为做修改或其他操作
        :param node:
        :return:
        """
        if node is None:
            return
        print(node.e)
        self._pre_order(node.left)
        self._pre_order(node.right)

    def pre_order_nr(self):
        """
        非递归实现前序遍历，打印bst中的元素，需借助栈来记录访问顺序，模拟系统栈的访问顺序；
        由于栈是先进后出的方式，要实现node,left,right顺序，所以，子树的入队顺序为right，left;
        :return:
        """
        st = LinkedStack()
        st.push(self.root)
        while not st.is_empty():
            node = st.pop()
            if node.right is not None:
                st.push(node.right)
            if node.left is not None:
                st.push(node.left)

    def __repr__(self):
        """
        格式化输出树节点元素信息，基于前序遍历
        :return:
        """
        res = ""
        self._generate_bst_string(self.root, 0, res)
        return res

    def _generate_bst_string(self, node, depth, res):
        """
        生成以node为根节点，深度为depth的描述二叉树的字符串
        :param node:
        :param depth:
        :param res:
        :return:
        """
        if node is None:
            res += self._generate_depth_string(depth) + "None\n"
            return
        res += self._generate_depth_string(depth) + "{0}".format(node.e)
        self._generate_bst_string(node.left, depth + 1, res)
        self._generate_bst_string(node.right, depth + 1, res)

    @staticmethod
    def _generate_depth_string(depth):
        """
        根据depth返回表示depth的深度字符串
        :param depth:
        :return:
        """
        res = ""
        for i in range(depth):
            res += "--"
        return res

    def in_order(self):
        """
        bst的中序遍历，并打印节点元素信息，先left子树，然后node节点，最后right子树；
        中序遍历的结果就是树中所有元素排序后的结果
        :return:
        """
        self._in_order(self.root)

    def _in_order(self, node):
        """
        中序遍历的递归算法实现
        :param node:
        :return:
        """
        if node is None:
            return
        self._in_order(node.left)
        print(node.e)
        self._in_order(node.right)

    def post_order(self):
        """
        二分搜索树的后续遍历，并打印节点元素信息，先left子树，然后right子树，最后node节点
        :return:
        """
        self._post_order(self.root)

    def _post_order(self, node):
        """
        后续遍历的递归算法实现
        :param node:
        :return:
        """
        if node is None:
            return
        self._post_order(node.left)
        self._post_order(node.right)
        print(node.e)

    def level_order(self):
        """
        层序遍历，一般实现用非递归算法，需借用队列数据结构来实现
        :return:
        """
        q = LinkedQueue()
        q.enqueue(self.root)
        while not q.is_empty():
            node = q.dequeue()
            print(node.e)
            if node.left is not None:
                q.enqueue(node.left)
            if node.right is not None:
                q.enqueue(node.right)

    def minimum(self):
        """
        寻找bst中的最小元素
        :return:
        """
        if self.get_size() == 0:
            raise ValueError("bst is empty")
        return self._minimum(self.root)

    def _minimum(self, node):
        """
        返回以node为根的bst的最小值所在的节点
        :param node:
        :return:
        """
        if node.left is None:
            return node
        return self._minimum(node.left)

    def remove_min(self):
        """
        删除bst中的最小值，从根节点向左走，最左边节点的值；
        对于最小值是最左边节点的值，该节点一定没有left子树
        :return:
        """
        ret = self._remove_min(self.root)
        return ret

    def _remove_min(self, node):
        """
        删除掉以node为根的bst中的最小节点，返回删除节点后新的bst的根
        :param node:
        :return:
        """
        if node.left is None:
            right_node = node.right
            node.right = None
            self.size -= 1
            return right_node
        node.left = self._remove_min(node.left)
        return node

    def maximum(self):
        """
        寻找bst中的最大元素
        :return:
        """
        if self.get_size() == 0:
            raise ValueError("bst is empty")
        return self._maximum(self.root)

    def _maximum(self, node):
        """
        返回node为根的bst的最大值所在的节点
        :return:
        """
        if node.right is None:
            return node
        return self._maximum(self.root)

    def remove_max(self):
        """
        删除bst中的最大值，从根节点向右走，最右边节点的值；
        对于最大值是最右边节点的值，该节点一定没有right子树
        :return:
        """
        ret = self._remove_max(self.root)
        return ret

    def _remove_max(self, node):
        """
        删除掉以node为根的bst的最大节点，返回删除节点后新的bst的根
        :param node:
        :return:
        """
        if node.right is None:
            left_node = node.left
            node.left = None
            self.size -= 1
            return left_node
        node.right = self._remove_max(node.right)
        return node

    def remove(self, e):
        """
        删除bst中的任意值
        :param e:
        :return:
        """
        self.root = self._remove(self.root, e)

    def _remove(self, node, e):
        """
        删除以node为根的bst中为e的节点，递归算法，返回删除节点后新的二分搜索树的根
        :param e:
        :return:
        """
        if node is None:
            return None
        if e < node.e:
            node.left = self._remove(node.left, e)
            return node
        elif e > node.e:
            node.right = self._remove(node.right, e)
            return node
        else:
            # 待删除节点left子树为空的情况
            if node.left is None:
                right_node = node.right
                node.right = None
                self.size -= 1
                return right_node
            # 待删除节点right子树为空的情况
            if node.right is None:
                left_node = node.left
                node.left = None
                self.size -= 1
                return left_node
            # 待删除节点的左右子树均不为空的情况：
            # 找到待删除节点的替代节点（后继节点或前驱节点）：
            # 1）可以找比待删除节点大的最小节点，即待删除节点的right子树的最小节点；
            # 2）或者可以找比待删除节点小的最大节点，即待删节点的left子树中最大节点；
            # 用这个节点顶替待删除节点的位置
            successor = self._minimum(node.right)
            successor.right = self._remove_min(node.right)
            successor.left = node.left
            node.left = None
            node.right = None

            return successor


def test_bst():
    b = Bst()
    nums = [5, 3, 6, 8, 4, 2]
    for i in nums:
        b.add(i)
    b.pre_order()
    b.pre_order_nr()
    b.level_order()
    b.in_order()
    b.post_order()
    print(b)


if __name__ == "__main__":
    test_bst()
