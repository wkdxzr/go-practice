package main

import "fmt"

type Pet interface {
	GetName() string
	GetAge() int
	MakeSound() string
	GetPrice() float64
}

// 定义猫咪结构体
type Cat struct {
	Name  string
	Age   int
	Color string
	Price float64
}

func (c Cat) GetName() string {
	return c.Name
}
func (c Cat) GetAge() int {
	return c.Age
}
func (c Cat) MakeSound() string {
	return "喵喵喵"
}
func (c Cat) GetPrice() float64 {
	return c.Price
}

// 定义狗狗结构体
type Dog struct {
	Name  string
	Age   int
	Breed string
	Price float64
}

func (d Dog) GetName() string {
	return d.Name
}
func (d Dog) GetAge() int {
	return d.Age
}
func (d Dog) MakeSound() string {
	return "汪汪汪"
}
func (d Dog) GetPrice() float64 {
	return d.Price
}

type PetShop struct {
	Pets []Pet
}

// 添加⼀个宠物到商店
func (shop *PetShop) AddPet(pet Pet) {
	shop.Pets = append(shop.Pets, pet)
}

// 显示所有宠物信息
func (shop *PetShop) ShowAllPets() {
	// 请⾃⾏实现遍历所有宠物
	for _, pet := range shop.Pets {
		switch p := pet.(type) {
		case Cat:
			fmt.Printf("宠物信息: %s %d %s %.2f\n", p.Name, p.Age, p.Color, p.Price)
		case Dog:
			fmt.Printf("宠物信息: %s %d %s %.2f\n", p.Name, p.Age, p.Breed, p.Price)
		}
	}
}

func main() {
	// 创建宠物商店
	shop := PetShop{}

	// 创建⼀些宠物实例
	cat := Cat{
		Name:  "咪咪",
		Age:   2,
		Color: "橘⾊",
		Price: 1000.0,
	}

	dog := Dog{
		Name:  "旺财",
		Age:   3,
		Breed: "柴⽝",
		Price: 2000.0,
	}

	// 添加宠物到商店
	shop.AddPet(cat)
	shop.AddPet(dog)

	// 显示所有宠物信息
	shop.ShowAllPets()
}
