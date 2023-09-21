package main

import "fmt"

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

