# coding:utf-8
from array_.array import Array


class MaxHeap(object):
    def __init__(self, arr=None, capacity=10):
        if arr is None:
            self._array = Array(capacity)
        else:
            if not isinstance(arr, Array):
                raise ValueError(u"arr is not sub instance of Array")
            self._array = arr
            # 循环siftdown每个节点元素， 若其大于其父节点则进行swap
            for i in range(self.parent(arr.get_size()-1), -1, -1):
                self._siftdown(i)

    def get_size(self):
        """
        返回堆中的元素个数
        :return:
        """
        return self._array.get_size()

    def is_empty(self):
        """
        判断堆中元素是否为空
        :return:
        """
        return self._array.is_empty()

    @staticmethod
    def parent(index):
        """
        返回完全二叉树的数组表示中,一个索引所表示的元素的父节点的索引
        :return:
        """
        if index == 0:
            return -1
        return (index - 1) >> 1

    @staticmethod
    def left_child(index):
        """
        返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
        :param index:
        :return:
        """
        return index << 1 + 1

    @staticmethod
    def right_child(index):
        """
        返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
        :param index:
        :return:
        """
        return index << 1 + 2

    def _siftup(self, k):
        """
        堆中添加元素后的上浮操作，k表示k节点的index
        :param k:
        :return:
        """
        while True:
            pk = self.parent(k)  # 获取k节点的父节点的索引
            if pk < 0:
                return
            k_val = self._array.get(k)  # 获取k节点的值
            pk_val = self._array.get(pk)  # 获取k父节点的值
            if k > 0 and pk_val < k_val:  # 如果k节点的值大于其父节点，那么进行swap
                self._swap(k, pk)
                k = pk
            else:
                break

    def _siftdown(self, k):
        """
        堆中添加元素后的下沉操作，k表示k节点的index
        :param k:
        :return:
        """
        # 对于数组堆来说，下沉操作，k的值从0开始，小于最大索引值
        while self.left_child(k) < self._array.get_size():
            j = self.left_child(k)  # k的左孩子的索引
            e1 = self._array.get(j+1)  # 获取right孩子的值
            e2 = self._array.get(j)  # 获取left孩子的值
            if j+1 < self._array.get_size() and e1 > e2:
                j = self.right_child(k)
            # 左右孩子进行比较，得到最大值的孩子节点，j记录最大值孩子节点的索引
            # 再将k节点与其最大子节点的值进行比较，如果k的值不小于子孩子的值，则结束，否则swap进行下沉替换操作
            ek = self._array.get(k)
            ej = self._array.get(j)
            if ek >= ej:
                break
            else:
                self._swap(k, j)
                k = j

    def _swap(self, i, j):
        """
        交换两个索引代表的节点的位置
        :param i:
        :param j:
        :return:
        """
        if i < 0 or j < 0 or i >= self._array.get_size() or j >= self._array.get_size():
            raise ValueError(u"index is illegal")
        i_val = self._array.get(i)
        j_val = self._array.get(j)
        self._array.set(i, j_val)
        self._array.set(j, i_val)

    def add(self, e):
        """
        向堆中添加元素，先默认加到最后位置（满足特性1），然后进行上否操作，确保新元素比其父节点小或等
        :param e:
        :return:
        """
        self._array.add_last(e)  # 对新增节点值先加入到最后位置
        self._siftup(self.get_size()-1)  # 对新增节点即最后一个节点进行上浮操作

    def find_max(self):
        """
        查看堆中的最大元素
        :return:
        """
        if self.get_size() == 0:
            raise ValueError(u"cannot find max when heap is empty")
        return self._array.get(0)

    def extract_max(self):
        """
        从堆中取出元素（最大）（对于堆二叉树是根节点，对于数组表示来说是index=0的元素）
        :return:
        """
        ret = self.find_max()  # 找到堆中最大元素
        self._swap(0, self._array.get_size()-1)  # 将堆最大元素（堆顶）与最末尾元素进行swap
        self._array.remove_last()  # 将swap后最后元素删除
        self._siftdown(0)  # 最后元素被替换到堆的根位置，需要进行下沉操作
        return ret

    def replace(self, e):
        """
        取出堆中的最大元素，并且替换成元素e
        :param e:
        :return:
        """
        ret = self.find_max()
        self._array.set(0, e)
        self._siftdown(0)
        return ret


def test_array_max_heap():
    heap = MaxHeap()
    n = 10
    for i in range(n):
        heap.add(i)
    arr = [None] * n
    for i in range(n):
        t = heap.extract_max()
        print(t)
        arr[i] = t
    for i in range(1, n-1):
        if arr[i-1] < arr[i]:
            raise ValueError("err")
    print("test max heap completed")


if __name__ == "__main__":
    test_array_max_heap()
