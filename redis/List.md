## List 
___
List是简单的字符串列表，可从头部或尾部插入元素。每个列表支持超过40亿个元素。
### 内部实现
List 类型的底层数据结构是由双向链表或压缩列表实现的
- 如果列表的元素个数小于512，列表每个元素的值都小于64字节，Redis会使用压缩列表。否则则会采用双向列表。
- LinkedList
    双向链表,pre指向前一个节点,next指向后一个节点,value保存当前节点的数据, List 中有head,tail,dup,free,,match,len,其中match比较两节点value是否相同，相同返回1，不同返回0.
- zipList
    压缩列表，为了节约内存而开发，zipList是由连续的内存组成。
- quickList
    quickList 是由前面两者组成的混合体,它将linkedList按段切分，每⼀段使⽤zipList来紧凑存储，多个zipList之间使⽤双向指针串接起来。
### 常用命令
```
# 将一个或多个value值插入到key列表的表头。
LPUSH key value [value ...]

# 将一个或多个value值插入到key列表的表尾。
RPUSH key value [value ...]

# 移除并返回key列表的头元素
LPOP key     

# 移除并返回key列表的尾元素
RPOP key

# 返回列表key中指定区间内的元素
LRANGE key start stop

# 从key列表表头/表尾弹出一个元素，没有就阻塞timeout秒
BLPOP key [key ...] timeou
BRPOP key [key ...] timeou
```
### 应用场景
    消息队列，但list不支持多个消费者同时处理一条消息。