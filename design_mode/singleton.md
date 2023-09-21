### 单例模式
单例模式保证一个类只有一个实例，单例模式又分为懒汉模式和饿汉模式。

#### 懒汉模式
```go
var (
    instance *singleton
    once sync.Once
)

type singleton struct {
	
}

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}
```

#### 饿汉模式
```go
type singleton struct {

}

var instance = new(singleton)

func GetInstance() *singleton {
    return instance
}



