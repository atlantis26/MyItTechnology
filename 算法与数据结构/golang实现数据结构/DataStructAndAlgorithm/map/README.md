映射（字典）：

    1、存储键、值数据对的数据结构（key,value）
        具有的基本方法：
        	Add(key interface{}, value interface{})
        	Remove(key interface{}) interface{}
        	Contains(key interface{}) bool
        	Get(key interface{}) interface{}
        	Set(key interface{}, value interface{}) (err error)
        	GetSize() int
        	IsEmpty() bool
    
    2、有序映射和无需映射
        基于搜索树的实现映射，是有序映射，其键具有顺序性；
        基于哈希表或链表实现的映射，是无序映射；
        
    3、多重映射：
        键是可以重复的，通过改变搜索树能够存储相同元素时，则包装实现的键是可以重复键的；
    
    4、集合和映射的关系
        集合和映射的底层实现数据结构相似；集合是节点存储一个e元素信息，映射存储2个信息key和value
        
        如果已经实现了映射了，可以在映射的基础上进行包装实现集合：
            包装映射，映射的key为集合的e元素值，value为空来实现