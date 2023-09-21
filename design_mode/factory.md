### 工厂模式
工厂模式是面向对象中创建对象常用的模式。

#### 简单工厂模式
简单工厂模式可以根据参数的不同返回不同的实例，使用者不需要关注创建的细节。
```go
type Gun struct {
    name string
    power int
}

func (g Gun) shot() {
    fmt.Printf("%s Shoots bullets with a firepower of %d\n", g.name, g.power)
}

func NewGun(name string, power int) *Gun {
    return &Gun{
        name: name,
        power: power,
    }
}
```

#### 抽象工厂模式
IFactory是一个抽象工厂，它定义了创建枪支和弹药的接口。Ak47Factory是一个具体的工厂，它实现了IFactory接口，可以创建AK47枪支和7.62mm口径的弹药。 这样，如果你想添加新的枪支和弹药类型，你只需要创建一个新的工厂，实现IFactory接口，然后在新的工厂中创建新的枪支和弹药实例。你不需要修改现有的代码，这就是抽象工厂模式的好处
```go
// 定义枪支接口
type IGun interface {
    Name() string
    Power() int
}

// 定义弹药接口
type IAmmo interface {
    Caliber() string
}

// AK47枪支实现
type Ak47 struct {
    name  string
    power int
}

func (g *Ak47) Name() string {
    return g.name
}

func (g *Ak47) Power() int {
    return g.power
}

// 7.62口径弹药实现
type SevenPointSixTwo struct {
    caliber string
}

func (a *SevenPointSixTwo) Caliber() string {
    return a.caliber
}

// 抽象工厂接口
type IFactory interface {
    MakeGun() IGun
    MakeAmmo() IAmmo
}

// AK47工厂实现
type Ak47Factory struct{}

func (f *Ak47Factory) MakeGun() IGun {
    return &Ak47{name: "AK47", power: 47}
}

func (f *Ak47Factory) MakeAmmo() IAmmo {
    return &SevenPointSixTwo{caliber: "7.62mm"}
}

func main() {
    factory := &Ak47Factory{}
    gun := factory.MakeGun()
    ammo := factory.MakeAmmo()

    fmt.Println("Gun:", gun.Name(), "with power of", gun.Power())
    fmt.Println("Ammo caliber:", ammo.Caliber())
}
```

#### 工厂方法模式
DogFactory和CatFactory都是工厂，它们的任务是创建Dog和Cat对象。这就是工厂方法模式的基本思想：将对象的创建过程封装在一个方法中，然后通过这个方法来创建对象。这样做的好处是，如果我们需要改变对象的创建过程，我们只需要修改这个方法，而不需要修改使用这个对象的代码
```go
// Animal 是一个接口，它有一个方法 Speak
type Animal interface {
	Speak() string
}

// Dog 是一个实现了 Animal 接口的结构体
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Cat 是一个实现了 Animal 接口的结构体
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// AnimalFactory 是一个接口，它有一个方法 CreateAnimal
type AnimalFactory interface {
	CreateAnimal() Animal
}

// DogFactory 是一个实现了 AnimalFactory 接口的结构体
type DogFactory struct{}

func (df DogFactory) CreateAnimal() Animal {
	return Dog{}
}

// CatFactory 是一个实现了 AnimalFactory 接口的结构体
type CatFactory struct{}

func (cf CatFactory) CreateAnimal() Animal {
	return Cat{}
}

func main() {
	var dogFactory AnimalFactory = DogFactory{}
	dog := dogFactory.CreateAnimal()
	fmt.Println(dog.Speak()) // 输出: Woof!

	var catFactory AnimalFactory = CatFactory{}
	cat := catFactory.CreateAnimal()
	fmt.Println(cat.Speak()) // 输出: Meow!
}


```




