### 创建数据库
```
createdb dbname
```
### 访问数据库
```
psql dbname
```
### 创建数据库表
```
CREATE TABLE Weather (
    city varchar(80),
    temp_lo int,
    temp_hi int,
    prcp real,
    date date
);
```
### 插入行
```
INSERT INTO weather VALUES ('San Francisco', 46, 50, 0.25, '1994-11-27');
```
### 查询数据
```
SELECT * FROM weather;
或者你可以自定义列
SELECT city, (temp_hi+temp_lo)/2 AS temp_avg, date FROM weather;
或者按排序顺序返回查询结果 
SELECT * FROM weather ORDER BY city;
或者从请求中删除重复值
SELECT DISTINCT city FROM weather;
或者选择最大值
SELECT max(temp_lo) FROM weather;
统计某一值出现的次数
SELECT city, count(*), max(temp_lo)
    FROM weather
    GROUP BY city;
统计某一值出现的次数加限制条件
SELECT city, count(*), max(temp_lo)
    FROM weather
    GROUP BY city
    HAVING max(temp_lo) < 40;
```
### 更新数据
```
UPDATE weather
    SET temp_hi = temp_hi - 2,  temp_lo = temp_lo - 2
    WHERE date > '1994-11-28';
```
### 删除数据库
```
dropdb dbname
```
### 删除数据库表
```
drop table name
```
### 删除数据库表内容
```
DELETE FROM weather WHERE city = 'Hayward';
```