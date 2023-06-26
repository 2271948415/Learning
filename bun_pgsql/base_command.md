### bun连接数据库
```
dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"

sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

db := bun.NewDB(sqldb, pgdialect.New())
```
### bun基于模型创建和删除表
```
res, err := db.NewCreateTable().Model((*User)(nil)).Exec(ctx)

res, err := db.NewDropTable().Model((*User)(nil)).Exec(ctx)
```
### bun更新数据
```
插入行
user := &User{Name: "admin"}
res, err := db.NewInsert().Model(user).Exec(ctx) 

更新行
user := &User{ID: 1, Name: "admin"}
res, err := db.NewUpdate().Model(user).Column("name").WherePK().Exec(ctx)

删除行
user := &User{ID: 1}
res, err := db.NewDelete().Model(user).WherePK().Exec(ctx)
```
### bun扫描查询结果到结构体等
```
user := new(User)
err := db.NewSelect().Model(user).Limit(1).Scan(ctx)

扫描到标量
var id int64
var name string
err := db.NewSelect().Model((*User)(nil)).Column("id", "name").Limit(1).Scan(ctx, &id, &name)

扫描到map
var m map[string]interface{}
err := db.NewSelect().Model((*User)(nil)).Limit(1).Scan(ctx, &m)
```