# coding:utf-8
from linkedlist_.linkedlist import LinkedList


class LinkedStack(object):
    def __init__(self):
        self.link = LinkedList()

    def push(self, e):
        """
        入栈
        :return:
        """
        self.link.add_first(e)

    def pop(self):
        """
        出栈
        :return:
        """
        return self.link.remove_first()

    def peek(self):
        """
        获取栈顶元素信息
        :return:
        """
        return self.link.get_first()

    def is_empty(self):
        """
        栈元素是否为空
        :return:
        """
        return self.link.is_empty()

    def get_size(self):
        """
        获取栈元素的个数
        :return:
        """
        return self.link.get_size()

    def __repr__(self):
        res = "Stack: top ["
        for i in range(self.get_size()):
            res += "{0}".format(self.link.get(i))
            if i != self.get_size()-1:
                res += ", "
        res += "]"
        return res


def test_linked_stack():
    st = LinkedStack()
    for i in range(10):
        st.push(i)
        print(st)
    st.pop()
    print(st)
    st.pop()
    print(st)
    st.pop()
    print(st)


if __name__ == "__main__":
    test_linked_stack()

