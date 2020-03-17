# coding: utf-8


class LoopQueue(object):
    """
    循环队列，与普通数组实现的对来来说，出队操作不会存在性能问题;
    循环队列的入队操作没有任何区别，出队操作，元素出队后，后面元素并不向前移动，而是移动一个表示队首的指针front
    """
    def __init__(self, capacity=10):
        self.data = [None] * capacity
        self.front = 0
        self.tail = 0
        self.size = 0

    def get_capacity(self):
        """
        查询队列容量
        :return:
        """
        return len(self.data)

    def get_size(self):
        """
        查询队列元素个数
        :return:
        """
        return self.size

    def _resize(self, capacity):
        """
        扩容
        :return:
        """
        new_data = [None] * capacity
        for i in range(self.size):
            new_data[i] = self.data[(i+self.front) % len(self.data)]
        self.data = new_data
        self.front = 0
        self.tail = self.size

    def is_empty(self):
        """
        查询队列是否为空
        :return:
        """
        return self.front == self.tail

    def enqueue(self, e):
        """
        元素入队
        :param e:
        :return:
        """
        # 循环队列有意浪费一个位置，如果tail+1 % len(data) == front,则表明队列已满，不能再存元素，需要扩容
        if (self.tail+1) % len(self.data) == self.front:
            self._resize(self.get_capacity() << 1)
        self.data[self.tail] = e
        self.tail = (self.tail+1) % len(self.data)
        self.size += 1

    def dequeue(self):
        """
        队顶元素出队
        :return:
        """
        if self.is_empty():
            raise ValueError("cannot dequeue from an empty queue")
        e = self.data[self.front]
        self.data[self.front] = None
        self.front = (self.front + 1) % len(self.data)
        self.size -= 1
        # 缩容
        if self.size == self.get_capacity() >> 2 and self.get_capacity() >>1 != 0:
            self._resize(self.get_capacity() >> 1)
        return e

    def get_front(self):
        """
        查询队首元素
        :return:
        """
        if self.is_empty():
            raise ValueError("queue is empty")
        return self.data[self.front]

    def __repr__(self):
        res = "Queue: size ={0}, capacity = {1}\nfront [".format(self.get_size(), self.get_capacity())
        for i in range(self.size):
            res += "{0}".format(self.data[i+self.front % len(self.data)])
            if (i + self.front) % len(self.data) != self.tail:
                res += ", "
        res += "] tail"

        return res


def test_loop_queue():
    q = LoopQueue(10)
    for i in range(10):
        q.enqueue(i)
    print(q)
    q.dequeue()
    print(q)


if __name__ == "__main__":
    test_loop_queue()



