# coding: utf-8
from linkedlist_.linkedlist import LinkedList


class LinkedSet(object):
    def __init__(self):
        self.link = LinkedList()

    def add(self, e):
        """
        添加集合元素
        :param e:
        :return:
        """
        if not self.link.contains(e):
            self.link.add_first(e)

    def remove(self, e):
        """
        移除集合元素
        :param e:
        :return:
        """
        self.link.remove_element(e)

    def contains(self, e):
        """
        判断集合是否包含e元素
        :param e:
        :return:
        """
        return self.link.contains(e)

    def get_size(self):
        """
        获取集合元素个数
        :return:
        """
        return self.link.get_size()

    def is_empty(self):
        """
        判断集合是否为空
        :return:
        """
        return self.link.is_empty()

    def __repr__(self):
        """
        格式化输出集合中的元素
        :return:
        """
        return self.link.__repr__()


if __name__ == "__main__":
    words = ["1", "2", "2", "A", "B", "D"]
    s = LinkedSet()
    for i in words:
        s.add(i)
    print(s.is_empty())
    print(s.get_size())
    print(s.contains("A"))
    print(s)

