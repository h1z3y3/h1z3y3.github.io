---
title: 设计模式：装饰器模式
date: 2018-11-06 21:33:52
tags: [设计模式,Golang]
---

装饰器模式主要解决要动态的给一个类添加一些新功能，而又不想让这个类变得庞大。
这种模式需要创建一个装饰类来包装扩展原有的类，并且在保证原有的类保持结构一致的前提下，提供额外的功能。

下面是给一个人装饰衣服的实例：
<!-- more -->

````
package decorator

import "fmt"

type Person interface {
	Show()
}

// 具体实现
type ConcreteComponent struct {
}

func (c *ConcreteComponent) Show() {
	fmt.Print("A Person wears sunglasses; ")
}

// 男人
type Man struct{}

func (m *Man) Show() {
	fmt.Print("A man wear a hat!")
}

// 女人
type Woman struct{}

func (w *Woman) Show() {
	fmt.Print("A woman wear a skirt!")
}

// TShirt
type TShirtDecorator struct {
	Person
	Color string
}

func (t *TShirtDecorator) Show() {
	t.Person.Show() // 调用父类的 Show() 方法
	// "装饰": 增加自己特有的属性
	fmt.Print(fmt.Sprintf("Color: %s, TShirt; ", t.Color))
}

func WearTShirt(p Person, c string) Person {
	return &TShirtDecorator{
		Person: p,
		Color:  c,
	}
}

// Pants
type PantsDecorator struct {
	Person
	Length int64
}

func (p *PantsDecorator) Show() {
	p.Person.Show()
	fmt.Print(fmt.Sprintf("Lenght: %dcm, Pants.; ", p.Length))
}

func WearPants(p Person, l int64) Person {
	return &PantsDecorator{
		Person: p,
		Length: l,
	}
}

// Shoes
type ShoesDecorator struct {
	Person
	Size int64
}

func (s *ShoesDecorator) Show() {
	s.Person.Show()
	fmt.Print(fmt.Sprintf("Size: %d, Shoes; ", s.Size))
}

func WearShoes(p Person, s int64) Person {
	return &ShoesDecorator{
		Person: p,
		Size:   s,
	}
}

// Examples
func ExampleConcrete_Wear() {
	var p Person = &ConcreteComponent{}
	p = WearTShirt(p, "Blue")
	p = WearPants(p, 100)
	p = WearShoes(p, 42)
	p.Show()

	// Output: A Person wears sunglasses; Color: Blue, TShirt; Lenght: 100cm, Pants.; Size: 42, Shoes;
}

func ExampleMan_Show() {
	var xiaoming Person = &Man{}
	xiaoming = WearShoes(xiaoming, 43)
	xiaoming = WearTShirt(xiaoming, "White")

	xiaoming.Show()

	// Output:A man wear a hat!Size: 43, Shoes; Color: White, TShirt;
}

func ExampleWoman_Show() {
	var xiaohong Person = &Woman{}
	xiaohong = WearTShirt(xiaohong, "Red")
	xiaohong = WearShoes(xiaohong, 38)

	xiaohong.Show()

	// Output: A woman wear a skirt!Color: Red, TShirt; Size: 38, Shoes;
}

````
源码以及测试源码下载地址：[https://github.com/h1z3y3/big-talk-go-design-patterns/tree/master/02_decorator](https://github.com/h1z3y3/big-talk-go-design-patterns/tree/master/02_decorator)