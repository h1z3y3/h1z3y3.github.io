---
title: 设计模式：简单工厂模式
date: 2018-10-21 20:58:52
tags: [设计模式,Golang]
---

这个系列是《大话设计模式》的读后感，将书中的设计模式用golang实现。     

第一个设计模式是简单工厂模式，主要用到的知识点是类的**多态**。    
**多态**表示不同的类可以执行相同的方法，但要通过它们自己的实现代码来执行。    
而在golang中没有类的概念，我们可以借助接口(interface)类型实现类的多态。   
如果一个类型实现了接口的所有方法，那么就可以说这个类型实现了这个接口。    

我们要实现两个数字的加减乘除操作作为示例
<!-- more -->
````
package simplefactory

import "fmt"

// 1. 定义一个接口类型，子类必须实现GoResult方法来实现该接口
type Operation interface {
	GetResult(a float64, b float64) (float64, error)
}

// 2. 初始化工厂类方法，传入操作符，返回对应的类
func NewOperation(oper string) Operation {
	switch oper {
	case "+":
		return &operationAdd{}
	case "-":
		return &operationSub{}
	case "*":
		return &operationMul{}
	case "/":
		return &operationDiv{}
	default:
		return nil
	}
}

// 加法
type operationAdd struct{}

func (o *operationAdd) GetResult(a float64, b float64) (float64, error) {
	return a + b, nil
}

// 减法
type operationSub struct{}

func (o *operationSub) GetResult(a float64, b float64) (float64, error) {
	return a - b, nil
}

// 乘法
type operationMul struct{}

func (o *operationMul) GetResult(a float64, b float64) (float64, error) {
	return a * b, nil
}

// 除法
type operationDiv struct{}

func (o *operationDiv) GetResult(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}

	return a / b, nil
}

````
源码以及测试源码下载地址：[https://github.com/h1z3y3/big-talk-go-design-patterns/tree/master/00_simple_factory](https://github.com/h1z3y3/big-talk-go-design-patterns/tree/master/00_simple_factory)