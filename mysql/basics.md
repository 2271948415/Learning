mysql架构分为两层：server层和存储引擎层，server层主要负责建立连接、分析和执行sql。主要包括连接器、查缓存。解析器、预处理器、优化器、执行器等。存储引擎层负责数据的存取。
### mysql的执行流程
1. 连接器
```
mysql -u$user -p 来连接mysql服务，连接的过程先经过TCP三次握手，mysql基于TCP协议进行传输。
show processlist 来查看服务有多少的客户连接，mysql空闲连接最大时长由wait_timeout参数控制，默认八小时。
mysql的链接也有长连接和短连接的概念：
// 短连接
连接 mysql 服务（TCP 三次握手）
执行sql
断开 mysql 服务（TCP 四次挥手）

// 长连接
连接 mysql 服务（TCP 三次握手）
执行sql
执行sql
执行sql
....
断开 mysql 服务（TCP 四次挥手）

长连接会更多的占用内存，一般有两种解决方式：
1.定期断开长连接
2.客户端主动重置连接，调用mysql_reset_connection
```
2.查询缓存
```
客户端向服务端发送了sql语句，服务端会解析sql语句的第一个字段，看是什么类型的语句。如果是select语句就会去查询缓存，在缓存里查找看之前是否执行过这一条命令，这个查询缓存是以 key-value 形式保存在内存中的，key 为 sql 查询语句，value 为 sql 语句查询的结果。
```