# coding:utf-8


class Node(object):
    def __init__(self, e, next_=None):
        self.e = e
        self.next = next_


class LinkedQueue(object):
    """
    基于双向链表的队列
    """
    def __init__(self):
        self.head = None
        self.tail = None
        self.size = 0

    def is_empty(self):
        """
        队列是否为空
        :return:
        """
        return self.size == 0

    def get_size(self):
        """
        队列中元素的个数
        :return:
        """
        return self.size

    def enqueue(self, e):
        """
        元素入队, 从双向链表末尾添加
        :return:
        """
        if self.tail is None:
            self.tail = Node(e)
            self.head = self.tail
        else:
            new_node = Node(e)
            self.tail.next = new_node
            self.tail = self.tail.next
        self.size += 1

    def dequeue(self):
        """
        元素出队，从双向链表表头出队
        :return:
        """
        if self.is_empty():
            raise ValueError("dequeue failed, stack is empty")
        ret_node = self.head
        self.head = ret_node.next
        ret_node.next = None
        if self.head is None:
            self.tail = None
        self.size -= 1

        return ret_node.e

    def get_front(self):
        """
        查询队列队顶元素信息
        :return:
        """
        if self.is_empty():
            raise ValueError("queue is empty")
        return self.head.e

    def __repr__(self):
        res = "Queue: front "
        cur = self.head
        while cur:
            res += "{0}->".format(cur.e)
            cur = cur.next
        res += "None tail"

        return res


def test_linked_queue():
    q = LinkedQueue()
    for i in range(10):
        q.enqueue(i)
        print(q)
    q.dequeue()
    print(q)
    q.dequeue()
    print(q)
    q.dequeue()
    print(q)


if __name__ == "__main__":
    test_linked_queue()
