# coding:utf-8


class Node(object):
    def __init__(self, key, value, next_=None):
        self.key = key
        self.value = value
        self.next = next_


class LinkedMap(object):
    def __init__(self):
        self.dummy_head = Node(None, None)
        self.size = 0

    def _get_node(self, key):
        """
        获取关于key的节点
        :param key:
        :return:
        """
        cur = self.dummy_head.next
        while cur:
            if cur.key == key:
                return cur
            cur = cur.next
        return None

    def add(self, key, value):
        """
        映射新增键值对元素，底层链表表头添加元素
        :param key:
        :param value:
        :return:
        """
        node = self._get_node(key)
        if node is None:
            self.dummy_head.next = Node(key, value, self.dummy_head.next)
            self.size += 1
        else:
            node.value = value

    def remove(self, key):
        """
        映射删除元素
        :param key:
        :return:
        """
        prev = self.dummy_head
        while prev.next is not None:
            if prev.next.key == key:
                break
            prev = prev.next
        if prev.next is not None:
            del_node = prev.next
            prev.next = del_node.next
            del_node.next = None
            self.size -= 1
            return del_node.value
        return None

    def contains(self, key):
        """
        映射是否包含元素key
        :param key:
        :return:
        """
        return self._get_node(key) is not None

    def get(self, key):
        """
        映射获取key的value值
        :param key:
        :return:
        """
        node = self._get_node(key)
        if node is None:
            return None
        else:
            return node.value

    def set(self, key, value):
        """
        映射修改key的value值
        :param key:
        :param value:
        :return:
        """
        node = self._get_node(key)
        if node is None:
            raise ValueError("{0} doesn't exist".format(key))
        node.value = value

    def get_size(self):
        """
        获取映射中的键值对个数
        :return:
        """
        return self.size

    def is_empty(self):
        """
        判断映射是否为空
        :return:
        """
        return self.size == 0


def test_linked_map():
    m = LinkedMap()
    m.add(1, "测试1")
    m.add(2, "测试2")
    print(m.get_size())
    print(m.get(1))
    print(m.get(2))
    print(m.get(3))


if __name__ == "__main__":
    test_linked_map()
