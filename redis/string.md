## string
___
### 介绍
String 是最基本的 key-value 结构，key 是唯一标识，value 是具体的值
### 内部实现
String 类型的底层的数据结构实现主要是 int 和 SDS，SDS和我们认识的C字符串不太一样。
- SDS不仅可以保存文本数据，还可以保存二进制数据
- SDS获取字符串的时间复杂度是o(1)
- Redis 的 SDS API 是安全的，拼接字符串不会造成缓冲区溢出。因为 SDS 在拼接字符串之前会检查 SDS 空间是否满足要求，如果空间不够会自动扩容，所以不会导致缓冲区溢出的问题
字符串内部编码有三种:int、raw、embstr
  如果一个字符串对象保存的是整数值，并且这个整数值可以用long类型来表示，那么字符串对象会将整数值保存在字符串对象结构的ptr属性里面（将void*转换成 long），并将字符串对象的编码设置为int。
  如果字符串对象保存的是一个字符串，并且这个字符申的长度小于等于 39 字节（redis 7.+版本），那么字符串对象将使用一个简单动态字符串（SDS）来保存这个字符串，并将对象的编码设置为embstr， embstr编码是专门用于保存短字符串的一种优化编码方式：
  如果字符串对象保存的是一个字符串，并且这个字符串的长度大于 39 字节，那么字符串对象将使用一个简单动态字符串（SDS）来保存这个字符串，并将对象的编码设置为raw：
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
- 计数
- 分布式锁
- 共享session信息