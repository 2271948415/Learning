# postgresSQL
***
### postgressql操作命令
- 启动或结束postgresql
```
pg_ctl -D /usr/local/var/postgres start/end 
```
- 连接到 PostgreSQL 服务器
```
psql -h <hostname> -p <port> -U <username> -d <database>
```
- 列出数据库
```
\l 
```
- 列出表
```
\d or \dt
```
- 显示表结构
```
\d <table> 或 \dt <table> 
```
- 连接数据库
```
\c <database> 
```
- 退出
```
\q 
```