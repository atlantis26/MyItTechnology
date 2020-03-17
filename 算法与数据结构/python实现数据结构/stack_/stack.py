# coding:utf-8
from array_.array import Array


class ArrayStack(object):
    """
    基于动态数组基础上实现的栈
    """
    def __init__(self, capacity):
        self.array = Array(capacity)

    def push(self, e):
        """
        想栈中加入新元素
        :return:
        """
        return self.array.add_last(e)

    def pop(self):
        """
        从栈中取出栈顶元素
        :return:
        """
        return self.array.remove_last()

    def peek(self):
        """
        查询栈顶元素个数
        :return:
        """
        return self.array.get_last()

    def get_size(self):
        """
        查询栈元素个数
        :return:
        """
        return self.array.get_size()

    def get_capacity(self):
        """
        查询栈容量
        :return:
        """
        return self.array.get_capacity()

    def is_empty(self):
        """
        判断栈是否为空
        :return:
        """
        return self.array.is_empty()

    def __str__(self):
        res = "Stack: ["
        for i in range(self.array.get_size()):
            res += "{0}".format(self.array.get(i))
            if i != self.array.get_size() - 1:
                res += ", "
        res += "] top"

        return res


def test_array_stack():
    st = ArrayStack(10)
    for i in ["test1", "test2", "test3"]:
        st.push(i)
    print(st)
    st.pop()
    print(st)
    st.pop()
    print(st)
    st.pop()
    print(st)


if __name__ == "__main__":
    test_array_stack()
