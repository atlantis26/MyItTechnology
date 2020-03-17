# coding:utf-8
from bst_.bst import Bst


class BstSet(object):
    def __init__(self):
        self.bst = Bst()

    def add(self, e):
        """
        添加集合元素
        :param e:
        :return:
        """
        self.bst.add(e)

    def remove(self, e):
        """
        移除集合元素
        :param e:
        :return:
        """
        self.bst.remove(e)

    def contains(self, e):
        """
        判断集合是否包含e元素
        :param e:
        :return:
        """
        return self.bst.contains(e)

    def get_size(self):
        """
        获取集合元素个数
        :return:
        """
        return self.bst.get_size()

    def is_empty(self):
        """
        判断集合是否为空
        :return:
        """
        return self.bst.is_empty()

    def __repr__(self):
        """
        基于底层bst前序遍历实现的string(), 格式化输出集合中的元素
        :return:
        """
        self.bst.pre_order()
        return ""


if __name__ == "__main__":
    words = ["1", "2", "2", "A", "B", "D"]
    s = BstSet()
    for i in words:
        s.add(i)
    print(s.is_empty())
    print(s.get_size())
    print(s.contains("A"))
    print(s)