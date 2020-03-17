# coding:utf-8


class Array(object):
    """
    自定义动态数组，动态扩容，动态缩容
    """
    def __init__(self, capacity=10):
        self.data = [None] * capacity
        self.size = 0

    def get_size(self):
        """
        获取数组中的元素个数
        :return:
        """
        return self.size

    def get_capacity(self):
        return len(self.data)

    def is_empty(self):
        """
        返回数组是否为空
        :return:
        """
        return self.size == 0

    def _resize(self, capacity):
        """
        进行扩容或缩容
        :return:
        """
        new_data = [None] * capacity
        for index in range(self.size):
            new_data[index] = self.data[index]
        self.data = new_data

    def add(self, index, e):
        """
        在第index个索引位置插入一个新的元素e,原index位置元素及以后的元素向后赋值，腾出index位置供新元素e赋值
        :return:
        """
        # 判断索引是否符合数组的索引值区间, index=self.size表示在末尾追加
        if index < 0 or index > self.size:
            raise ValueError("add failed, Require index >= 0 and < Array.size")
        # 如果array满了，则进行二倍扩容
        if self.size == len(self.data):
            self._resize(len(self.data) << 1)
        # 将index即以后的元素向后移动一个位置
        for i in range(self.size, index-1, -1):
            self.data[i] = self.data[i-1]
        # 将新元素e加入到原index位置，覆盖原位置的元素，原位置的元素已经向后赋值
        self.data[index] = e
        self.size += 1

    def add_last(self, e):
        """
        在数组末尾添加元素e
        :return:
        """
        self.add(self.size, e)

    def add_first(self, e):
        """
        在数组首部添加元素
        :param e:
        :return:
        """
        self.add(0, e)

    def set(self, index, e):
        """
        修改index索引位置的元素
        :param index:
        :param e:
        :return:
        """
        if index < 0 or index >= self.size:
            raise ValueError("get failed, index is illegal")
        self.data[index] = e

    def get(self, index):
        """
        获取index索引位置的元素
        :return:
        """
        if index < 0 or index >= self.size:
            raise ValueError("get failed, index is illegal")
        return self.data[index]

    def get_last(self):
        """
        获取数组最后一个元素
        :return:
        """
        return self.data[self.size-1]

    def get_first(self):
        """
        获取数组第一个元素
        :return:
        """
        return self.data[0]

    def contains(self, e):
        """
        查找数组中是否存在元素e
        :return:
        """
        for i in self.data:
            if i == e:
                return True
        return False

    def find(self, e):
        """
        查找数组中元素e所在的索引，如果不存在元素e,则返回-1
        :return:
        """
        for i in range(self.size):
            if self.data[i] == e:
                return i
        return -1

    def remove(self, index):
        """
        删除index索引位置的元素，index后面的元素向前移动一个位置
        :param index:
        :return:
        """
        e = self.get(index)
        for i in range(index+1, self.size):
            self.data[i-1] = self.data[i]
        self.size -= 1
        self.data[self.size] = None
        # 检查删除元素后是否需要缩容
        if self.size == len(self.data) >> 2:
            self._resize(len(self.data) >> 1)
        return e

    def remove_last(self):
        """
        删除数组最后一个元素
        :return:
        """
        return self.remove(self.size-1)

    def remove_first(self):
        """
        删除数组第一个元素
        :return:
        """
        return self.remove(0)

    def remove_element(self, e):
        """
        从数组中删除元素e,如果存在e，则删除e,只删除一个；否则什么事情也不干
        :return:
        """
        index = self.find(e)
        if index != -1:
            self.remove(index)

    def __repr__(self):
        res = "Array: size = {0}, capacity = {1}\n".format(self.size, len(self.data))
        res += "["
        for i in range(self.size):
            res += "{0}".format(self.data[i])
            if i != self.size - 1:
                res += ", "
        res += "]"

        return res


def new_array(capacity):
    """
    构造函数，传入数组的容量capacity构造Array
    :return:
    """
    return Array(capacity=capacity)


def new_array1():
    """
    无参数的构造函数，默认数组的容量capacity=10
    :return:
    """
    return Array(capacity=10)


def new_array2(arr):
    """
    将静态数组重构为动态数组
    :return:
    """
    a = Array(capacity=len(arr))
    a.data = arr

    return a


def test_array():
    arr = new_array(10)
    for i in range(10):
        arr.add_last(i)

    # 添加时，触发扩容，再完成添加
    print(arr)
    arr.add(1, 100)
    print(arr)
    arr.remove(0)
    arr.remove(0)
    arr.remove(0)
    arr.remove(0)
    arr.remove(0)
    arr.remove(0)
    # 删除完成后，触发缩容
    print(arr)


if __name__ == "__main__":
    test_array()
