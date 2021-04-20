---
title: 设计模式：工厂方法模式
date: 2018-11-25 18:01:29
tags: [设计模式,Golang]
---

工厂方法模式是简单工厂的升级。他创建一个用于实例化类的接口，并由工厂的子类决定实例化哪个类。工厂方法模式使得一个类的实例化延迟到子类。

下面仍然以“两个数字的运算”作为例子

<!-- more -->

operations.go // 运算类

````
package factory_method

// 运算
type Operation interface {
	SetA(float64)
	SetB(float64)
	GetResult() (float64, error)
}

// 运算基类，实现公共的方法
type OperationBase struct {
	a float64
	b float64
}

func (oper *OperationBase) SetA(a float64) {
	oper.a = a
}

func (oper *OperationBase) SetB(b float64) {
	oper.b = b
}

// 加法运算
type AddOperation struct {
	*OperationBase
}

func (oper *AddOperation) GetResult() (float64, error) {
	return oper.a + oper.b, nil
}

// 减法运算
type SubOperation struct {
	*OperationBase
}

func (oper *SubOperation) GetResult() (float64, error) {
	return oper.a - oper.b, nil
}

// 乘法元算
type MulOperation struct {
	*OperationBase
}

func (oper *MulOperation) GetResult() (float64, error) {
	return oper.a * oper.b, nil
}

````

factory_method.go // 工厂类

````
package factory_method

// 工厂类
type OperationFactory interface {
	CreateOperation() Operation
}

// 加法工厂
type AddFactory struct {
}

func (f *AddFactory) CreateOperation() Operation {
	return &AddOperation{
		OperationBase: &OperationBase{},
	}
}

// 减法工厂
type SubFactory struct {
}

func (f *SubFactory) CreateOperation() Operation {
	return &SubOperation{
		OperationBase: &OperationBase{},
	}
}

// 乘法工厂
type MulFactory struct {
}

func (f *MulFactory) CreateOperation() Operation {
	return &MulOperation{
		OperationBase: &OperationBase{},
	}
}

````


源码以及测试源码下载地址：[https://github.com/h1z3y3/big-talk-go-design-patterns/tree/master/04_factory_method](https://github.com/h1z3y3/big-talk-go-design-patterns/tree/master/04_factory_method)