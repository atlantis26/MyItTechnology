# coding:utf-8


class Node(object):
    """
    链表节点
    """
    def __init__(self, e, next_=None):
        self.e = e
        self.next = next_

    def __repr__(self):
        return "{0}".format(self.e)


class LinkedList(object):
    """
    链表
    """
    def __init__(self):
        self.dummy_head = Node(e=None)
        self.size = 0

    def _validate_index(self, index):
        if index < 0 or index > self.size:
            raise ValueError("add failed, Illegal index")

    def is_empty(self):
        """
        链表是否为空
        :return:
        """
        return self.size == 0

    def get_size(self):
        """
        获取栈的元素个数
        :return:
        """
        return self.size

    def add(self, index, e):
        """
        在链表的‘index’位置添加元素，关键找到添加的节点的前一个节点；
        :param index:
        :param e:
        :return:
        """
        self._validate_index(index)
        prev = self.dummy_head
        for i in range(index):
            prev = prev.next
        new_node = Node(e)
        prev.next, new_node.next = new_node, prev.next
        self.size += 1

    def add_first(self, e):
        """
        在链表头添加元素节点
        :param e:
        :return:
        """
        self.add(0, e)

    def add_last(self, e):
        """
        在链表末尾添加元素节点
        :param e:
        :return:
        """
        self.add(self.size, e)

    def get(self, index):
        """
        查询链表的第index位置的元素节点
        :param index:
        :return:
        """
        self._validate_index(index)
        # 遍历链表，需知道链表的size值
        cur = self.dummy_head.next
        for i in range(index):
            cur = cur.next
        return cur

    def get_first(self):
        """
        查询链表的表头位置的元素节点
        :return:
        """
        return self.get(0)

    def get_last(self):
        """
        查询链表的末尾元素节点
        :return:
        """
        return self.get(self.size-1)

    def set(self, index, e):
        """
        修改链表第index位置的元素
        :param index:
        :param e:
        :return:
        """
        self._validate_index(index)
        cur = self.dummy_head.next
        for i in range(index):
            cur = cur.next
        cur.e = e

    def contains(self, e):
        """
        查找链表是否存在e元素
        :param e:
        :return:
        """
        cur = self.dummy_head.next
        while cur:
            if cur.e == e:
                return True
            cur = cur.next
        return False

    def remove(self, index):
        """
        删除索引index位置的元素
        :param index:
        :return:
        """
        self._validate_index(index)
        prev = self.dummy_head
        for i in range(index):
            prev = prev.next
        del_node = prev.next
        prev.next = del_node.next
        del_node.next = None
        self.size -= 1

        return del_node.e

    def remove_first(self):
        """
        删除链表表头的元素节点
        :return:
        """
        return self.remove(0)

    def remove_last(self):
        """
        删除链表末尾的元素节点
        :return:
        """
        return self.remove(self.size-1)

    def remove_element(self, e):
        """
        删除链表元素e,只删除一个，若不存在则什么事情也不干
        :param e:
        :return:
        """
        prev = self.dummy_head
        while prev.next:
            if prev.next.e == e:
                break
            prev = prev.next
        if prev.next is not None:
            del_node = prev.next
            prev.next = del_node.next
            del_node.next = None

    def __repr__(self):
        res = ""
        cur = self.dummy_head.next
        while cur:
            res += "{0}->".format(cur.e)
            cur = cur.next
        res += "None"
        return res


def test_linked_list():
    l = LinkedList()
    for i in range(10):
        l.add_last(i)
        print(l)
    l.remove_first()
    print(l)
    l.remove_first()
    print(l)
    l.remove_first()
    print(l)


if __name__ == "__main__":
    test_linked_list()
