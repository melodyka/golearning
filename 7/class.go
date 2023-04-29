package main

import "fmt"

type Hero struct {
	name string
	sex  string
}

func (this *Hero) Show() {
	fmt.Println("name= ", this.name)
}

func (this *Hero) Getname() string {
	return this.name
}

func (this *Hero) Setname(newname string) {
	this.name = newname
}

type SuperMan struct {
	Hero  //继承Hero类方法
	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name= ", this.name)
	fmt.Println("sex= ", this.sex)
	fmt.Println("level= ", this.level)
}

func main() {
	hero := Hero{"zhang3", "female"}

	hero.Show()
	hero.Setname("li4")
	hero.Show()

	//s := SuperMan{Hero{"li4", "female"}, 88}	//容易重复
	var s SuperMan
	s.name = "li4"
	s.sex = "male"
	s.level = 88
	s.Show()
	s.Fly()
	s.Print()
}
