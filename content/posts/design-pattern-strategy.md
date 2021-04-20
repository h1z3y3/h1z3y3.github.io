---
title: 设计模式：策略模式
date: 2018-10-24 23:33:52
tags: [设计模式,Golang]
---

在策略模式中，我们需要创建一系列策略对象和一个能随策略对象改变而改变的Context对象，策略对象改变Context的执行方法。    

仍以两个数字的加减乘除操作作为示例
<!-- more -->

````
package strategy

import "fmt"

// Context 类
type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{
		strategy: strategy,
	}
}

func (c *Context) GetResult(a float64, b float64) (float64, error) {
	return c.strategy.GetResult(a, b)
}

// 策略接口
type Strategy interface {
	GetResult(a float64, b float64) (float64, error)
}

// 以下类实现策略接口
// 加法
type Add struct{}

func (o *Add) GetResult(a float64, b float64) (float64, error) {
	return a + b, nil
}

// 减法
type Sub struct{}

func (o *Sub) GetResult(a float64, b float64) (float64, error) {
	return a - b, nil
}

// 乘法
type Mul struct{}

func (o *Mul) GetResult(a float64, b float64) (float64, error) {
	return a * b, nil
}

// 除法
type Div struct{}

func (o *Div) GetResult(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}

	return a / b, nil
}

````
源码以及测试源码下载地址：[https://github.com/h1z3y3/big-talk-go-design-patterns/tree/master/01_strategy](https://github.com/h1z3y3/big-talk-go-design-patterns/tree/master/01_strategy)