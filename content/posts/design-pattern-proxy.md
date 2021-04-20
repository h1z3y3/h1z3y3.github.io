---
title: 设计模式：代理模式
date: 2018-11-25 16:50:58
tags: [设计模式,Golang]
---

代理模式即在真实类的基础上封装一层代理类，由代理类完成对真实类的调用。    
以便可以在代理类中做一些额外的工作，如进行访问权限校验、保存Cache缓存等操作。

下面以"读取图片资源"为例说明代理模式：

<!-- more -->

````
package proxy

// 接口，代理类和真实类都要实现
type Image interface {
	Get() string
}

// 真实的图片类
type RealImage struct {}

func (r *RealImage) Get() string {
	return "real_image_url"
}

// 代理类
type ImageProxy struct {
	realImage RealImage
}

// 由代理类进行原类的调用，从而能在原类基础上做一些操作
func (r *ImageProxy) Get() string {

	var res string
	// pre: 权限检查、查看是否有cache等
	res += "pre:"

	res += r.realImage.Get()

	// after: 保存cache、格式化结果等
	res += ":after"

	return res
}


````
源码以及测试源码下载地址：[https://github.com/h1z3y3/big-talk-go-design-patterns/tree/master/03_proxy](https://github.com/h1z3y3/big-talk-go-design-patterns/tree/master/03_proxy)