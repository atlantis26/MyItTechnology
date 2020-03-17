# coding:utf-8


class Node(object):
    def __init__(self, key, value):
        self.key = key
        self.value = value
        self.left = None
        self.right = None


class BstMap(object):
    def __init__(self):
        self.root = None
        self.size = 0

    def _get_node(self, node, key):
        """
        获取key的节点，使用前序遍历查找
        :param node:
        :param key:
        :return:
        """
        if node is None:
            return None
        if key == node.key:
            return node
        elif key < node.key:
            return self._get_node(node.left, key)
        else:
            return self._get_node(node.right, key)

    def _add(self, node, key, value):
        """
        向以node为根节点的插入数据元素节点（key,value）, 返回插入新节点后的二分搜索树的根，递归算法
        :param node:
        :param key:
        :param value:
        :return:
        """
        if node is None:
            self.size += 1
            return Node(key, value)
        if key < node.key:
            node.left = self._add(node.left, key, value)
        elif key > node.key:
            node.right = self._add(node.right, key, value)
        else:
            node.value = value

        return node

    def add(self, key, value):
        """
        映射新增键值对元素
        :param key:
        :param value:
        :return:
        """
        self.root = self._add(self.root, key, value)

    def remove(self, key):
        """
        映射删除元素
        :param key:
        :return:
        """
        node = self._get_node(self.root, key)
        if node is not None:
            self.root = self._remove(self.root, key)
            return node.value
        return None


    def _remove(self, node, key):
        """
        删除掉以node为根的二分搜索树中为key的节点，递归算法；
        返回删除节点后最新的二分搜索树的根
        :param node:
        :param key:
        :return:
        """
        if node is None:
            return None
        if key < node.key:
            node.left = self._remove(node.left, key)
        elif key > node.key:
            node.right = self._remove(node.right, key)
        else:
            # 待删除节点的left子树为空的情况
            if node.left is None:
                right_node = node.right
                node.right = None
                self.size -= 1
                return right_node
            if node.right is None:
                left_node = node.left
                node.left = None
                self.size -= 1
                return left_node
            # 待删除节点的左右子树均不为空的情况
            # 找到比待删除节点大的最小节点，即待删除节点右子树的最小节点
            successor = self._minimum(node.right)
            successor.right = self._remove_min(node.right)
            successor.left = node.left
            node.left = None
            node.right = None

        return node

    def _minimum(self, node):
        """
        找到bst树中的最小节点，即最左边的节点
        :return:
        """
        if node.left is None:
            return node
        return self._minimum(node.left)

    def _remove_min(self, node):
        """
        删除bst树中的最小节点，返回删除后的根节点
        :param node:
        :return:
        """
        if node.left is None:
            right_node = node.right
            node.right = None
            self.size -= 1
            return right_node
        node.left = self._remove_min(node.left)

        return node

    def contains(self, key):
        """
        映射是否包含元素key
        :param key:
        :return:
        """
        return self._get_node(self.root, key) is not None

    def get(self, key):
        """
        映射获取key的value值
        :param key:
        :return:
        """
        node = self._get_node(self.root, key)
        if node is None:
            return None
        return node.value

    def set(self, key, value):
        """
        映射修改key的value值
        :param key:
        :param value:
        :return:
        """
        node = self._get_node(self.root, key)
        if node is None:
            raise ValueError("{0} doesn't exist".format(key))
        else:
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


def test_bst_map():
    m = BstMap()
    m.add(1, "测试1")
    m.add(2, "测试2")
    print(m.get_size())
    print(m.get(1))
    print(m.get(2))
    print(m.get(3))


if __name__ == "__main__":
    test_bst_map()
