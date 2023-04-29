package main

import "fmt"

type AnimalIF interface { //父类接口
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct { //子类
	color string
}

func (this *Cat) Sleep() { //定义父类方法
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string { //报错返回值过多
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

func main() {
	// var animal AnimalIF
	// animal = &Cat{"Green"}	赋值，父类接口指针指向子类的具体对象，再用父类调用接口
	// animal.Sleep()
	// animal = &Dog{"Yellow"}
	// animal.Sleep()
	cat := Cat{"Green"}
	dog := Dog{"Yellow"}
	showAnimal(&cat)
	showAnimal(&dog)
}
