## Hash
___
Hash 是一个键值对组合，其value形式如：value = [{filed1,value},...{filedN,valueN}].Hash特别适合用于存储对象。其内部主要由哈希表实现。
![hash.png](image%2Fhash.png)

### 应用场景
1. 缓存对象  
   Hash 类型的 （key，field， value） 的结构与对象的（对象id， 属性， 值）的结构相似，也可以用来存储对象
2. 购物车  
   以用户 id 为 key，商品 id 为 field，商品数量为 value，恰好构成了购物车的3个要素。  
- 添加商品：HSET cart:{用户id} {商品id} 1
- 添加数量：HINCRBY cart:{用户id} {商品id} 1
- 商品总数：HLEN cart:{用户id}
- 删除商品：HDEL cart:{用户id} {商品id}
- 获取购物车所有商品：HGETALL cart:{用户id}