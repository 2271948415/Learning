## string
___
### 介绍
String 是最基本的 key-value 结构，key 是唯一标识，value 是具体的值
### 内部实现
String 类型的底层的数据结构实现主要是 int 和 SDS，SDS和我们认识的C字符串不太一样。
#### sds定义
![string.png](..%2Fimages%2Fredis%2Fstring.png)
- SDS不仅可以保存文本数据，还可以保存二进制数据
- SDS杜绝缓冲区溢出，在c中对字符串进行扩容之前必须考虑容量是否足够，不足够得去重新分配内存。SDS的API里面也有一个用于执行拼接操作的sdscat函数，它可以将一个C字符串拼接到给定SDS所保存的字符串的后面，但是在执行拼接操作之前，sdscat会先检查给定SDS的空间是否足够，如果不够的话，sdscat就会先扩展SDS的空间，然后才执行拼接操作
- SDS获取字符串长度的时间复杂度是o(1),sds结构中的len属性记录了字符串的长度。
- Redis 的 SDS API 是安全的，拼接字符串不会造成缓冲区溢出。因为 SDS 在拼接字符串之前会检查 SDS 空间是否满足要求，如果空间不够会自动扩容，所以不会导致缓冲区溢出的问题
字符串内部编码有三种:int、raw、embstr
  如果一个字符串对象保存的是整数值，并且这个整数值可以用long类型来表示，那么字符串对象会将整数值保存在字符串对象结构的ptr属性里面（将void*转换成 long），并将字符串对象的编码设置为int。
  如果字符串对象保存的是一个字符串，并且这个字符申的长度小于等于 39 字节，那么字符串对象将使用一个简单动态字符串（SDS）来保存这个字符串，并将对象的编码设置为embstr， embstr编码是专门用于保存短字符串的一种优化编码方式：
  如果字符串对象保存的是一个字符串，并且这个字符串的长度大于 39 字节，那么字符串对象将使用一个简单动态字符串（SDS）来保存这个字符串，并将对象的编码设置为raw.
  embstr 会通过一次内存分配函数来分配一块连续的内存空间来保存redisobject和sds，而raw编码则会通过两次内存分配来函数来分别分配两块空间来保存redisobject和sds。
### 常用命令
- 基本命令
```
# 设置key-value的值
> SET name bo

# 根据key获得value
> GET name

# 判断key是否存在
> EXISTS NAME

# 删除某个key的值
> DEL name

# 批量设置key—value
> MSET key1 value1 key2 value2

# 获取多个key
> MGET key1 key2

#设置过期时间
> SET key value EX 60
> SETEX key 60 value
```
- 计数器
```
> SET number 0

# 将key中的数字加一
> INCR number

# 将 key 中储存的数字值减一
> DECR number

# 将key中存储的数字值加 10
> INCRBY number 10

# 将key中存储的数字值键 10
> DECRBY number 10
```
### 应用场景
- 缓存对象   
比如：验证码信息，用户登录状态等等。
- 计数
- 分布式锁   
分布式锁是一种用于分布式系统中实现资源的互斥访问的机制。在分布式环境中，多个节点同时访问共享资源时，为了保证数据的一致性和正确性，需要保证同一时间只有一个节点可以对资源进行操作。   
  Redis的SETNX命令是一种用于设置键值对的原子性操作。它在键不存在的情况下设置键的值，并返回设置成功与否的结果。
  执行SETNX命令时，会进行以下操作： 如果键key不存在，则将键 key 的值设置为 value。 如果键 key 已经存在，则不进行任何操作，返回0。
- 共享session信息