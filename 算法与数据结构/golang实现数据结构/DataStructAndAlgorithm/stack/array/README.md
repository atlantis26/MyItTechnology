简单的时间复杂度分析：
大O描述的是
O(1) : 与元素的规模没有关系的
O(n): n的幂是1次方，则是线性关系的
O(lgn):
O(nlogn):
0(n^2):  n的幂是2次方


添加操作： 总体为O(n)的时间复杂度
addLast(e)       O(1)时间复杂度，与数组中元素的规模没有关系的
addFirst(e)      O(n)时间复杂度，数组中每个元素向后移动一位，与元素规模呈线性关系
add(index, e)    O(n/2)= O(n)时间复杂度  ,add时触发扩容 resize  O(n)时间复杂度


删除操作： 总体为O(n)时间复杂度
removeLast(e) O(1)时间复杂度
removeFirst(e)   O(n)时间复杂度
remove(index, e) O(n/2) = O(n),  删除时触发缩容 resize  O(n)时间复杂度


修改操作：已经索引index了，为O(1)的时间复杂度，未知索引为O(n)
set(index, e)   O(1)时间复杂度

查询操作： 已经索引index了，为O(1)的时间复杂度，未知索引为O(n)
get(index)      O(1)
contains(e)     O(n)
find(e)         O(n)

数组的索引如果有意义的应用中时，比如索引作为学生学号，则能大大降低各种操作的时间复杂度，因为在已知索引index的情况下的操作，
时间复杂度是O(1)的


均摊复杂度：
addLast、removeLast方法为O(1)时间复杂度,当触发扩容和或缩容时，resize时间复杂度为O(n), 所以在分析时间复杂度

复杂度震荡：
addLast、removeLast方法为O(1)时间复杂度
当数组满了时，频繁进行addLast，removeLast，会频繁的触发扩容和缩容，resize的时间复杂度为O(n)；解决方案为当size为容量的
1/4时才触发缩容，且只缩容1/2；

