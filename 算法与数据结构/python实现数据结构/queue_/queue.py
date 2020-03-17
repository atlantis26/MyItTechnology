# coding:utf-8
from array_.array import Array


class ArrayQueue(object):
    def __init__(self, capacity):
        self._array = Array(capacity)

    def enqueue(self, e):
        """
        元素入队
        :param e:
        :return:
        """
        return self._array.add_last(e)

    def dequeue(self):
        """
        元素出队
        :return:
        """
        return self._array.remove_first()

    def get_front(self):
        """
        查询队顶元素
        :return:
        """
        return self._array.get_first()

    def get_size(self):
        """
        获取队列中元素个数
        :return:
        """
        return self._array.get_size()

    def get_capacity(self):
        """
        获取队列的容量
        :return:
        """
        return self._array.get_capacity()

    def is_empty(self):
        """
        判断队列是否为空
        :return:
        """
        return self._array.is_empty()

    def __repr__(self):
        res = "Queue: front ["
        for i in range(self._array.get_size()):
            res += "{0}".format(self._array.get(i))
            if i != self._array.get_size() - 1:
                res += ", "
        res += "] tail"

        return res


def test_array_queue():
    q = ArrayQueue(10)
    for i in range(10):
        q.enqueue(i)
    print(q)
    q.dequeue()
    print(q)


if __name__ == "__main__":
    test_array_queue()
