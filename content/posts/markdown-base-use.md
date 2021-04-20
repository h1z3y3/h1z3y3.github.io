---
title: MarkDown基本使用方法
date: 2015-10-17 18:31:15
tags: [Markdown, Hexo]
---

## Markdown 介绍
> Markdown是一种可以使用普通文本编辑器编写的**标记语言**，通过简单的标记语法，它可以使普通文本内容具有一定的格式。Markdown的使用十分简单，常用的**标记**其实就十几个，所以相对于HTML这种复杂的标记语言来说，Markdown是十分轻量的标记语言。

<!--more-->

## Markdown的基本使用
1. ### 加粗和斜体
		
	`*斜体* _斜体_`
	
	只要用一个`_`或者一个`#`包裹文字，即可实现文字斜体。
	  
	`**加粗文字**  __加粗文字__`
	
	只要用两个`_`或者两个`#`包裹文字，即可实现文字粗体。
	
	![粗体和斜体](https://p5.ssl.qhimg.com/t014be129baf06de9da.jpg)

2. ### 链接与图片
	
	插入链接和图片的格式很相似，只有一个`!`的区别
	
	链接格式为：`[链接文字](链接地址)`
	
	图片格式为：`![图片名字](图片地址)`
	
	在Markdown中插入图片地址要用到图床，我用到的是[围脖图床修复计划](http://weibotuchuang.sinaapp.com)和[云图图床](http://www.pic100.net)，上传图片就可以生成URL
	
	![链接和图片](https://p0.ssl.qhimg.com/t01846fd89958a14640.jpg)
	

3. ### 标题

	标题是文章中最常用的文本格式，在Markdown中只要在一行文字前添加`#`，即可标记为标题格式。
	
	`＃ 一级标题`
	
	`## 二级标题`
	
	`### 三级标题`
	
	以此类推，Markdown一共有六级标题，六级标题只需要加上六个`######`即可。
	
	![标题](https://p0.ssl.qhimg.com/t013492ed250532d583.jpg)
	

4. ### 列表
	列表有**有序列表**和**无序列表**
	
	有序列表格式：
		
		1. 红色
		2. 蓝色
		3. 黑色
		
	无序列表格式：
		
		* 红色
		* 蓝色
		* 黑色
		
	![列表](https://p4.ssl.qhimg.com/t012e847b97fdfde44f.jpg)

5. ### 引用
	
	引用只需要在文字前面加上 `>` 就可以了。你可以联合其它的标记符一起使用。  
	
		> * 引用中列表
		> * 列表
		> ### 引用中三级标题
		>> 二级引用
	
	![引用](https://p3.ssl.qhimg.com/t0163d79fe042bb581a.jpg)
	
6. ### 行内代码和代码块
	行内代码格式： `｀这里写代码｀`  
	代码块格式： 只要比上一行进行右缩进即可。按键盘`tab`键可以实现缩进
	
	![行内代码和代码块](https://p1.ssl.qhimg.com/t013a7cbc965a07325f.jpg)
	
7. ### 表格
	
	表格应该是Markdown中最难的标记了
	
	最简单的表格：
	
		| First Header | Second Header | Third Header |
		| ------------ | ------------- | ------------ |
		| Content Cell | Content Cell  | Content Cell |
		| Content Cell | Content Cell  | Content Cell |

	你也可以设置文字的对齐方式

		First Header | Second Header | Third Header
		:----------- | :-----------: | -----------:
		Left         | Center        | Right
		Left         | Center        | Right

	![表格](https://p1.ssl.qhimg.com/t017070396e44e5c1ef.jpg)

8. ### 水平分隔线
	
	使用 `---` 或者 `***`即可。
	
	![水平线](https://p2.ssl.qhimg.com/t0192d850b3c0576e4e.jpg)
	
9. ### 删除线

	删除线格式：  
	
		~~删除线~~
	
	![删除线](https://p0.ssl.qhimg.com/t012a8f4294c4d33c0b.jpg)
	
10. ### 换行

	只要在文字每行文字后面加上两个或两个以上的`空格`即可实现换行
	
## 注意
1. 标记符号后面一定加上**空格**
2. 标记语言必须使用英文符号
3. 如果使用正确而不起作用，换一行再试一次
	
## Markdown编辑工具
本人使用的工具是Mac OS下的**Mou**,支持**实时预览**，文章中的截图即为Mou的截图。你也可以在Github上搜索其它主题样式进行安装。  
Mac下Markdown编辑器：[Mou ¥0](http://25.io/mou/)、[Byword ¥40](http://itunes.apple.com/cn/app/byword/id482063361?mt=8)、[iA Writer ¥128](https://itunes.apple.com/cn/app/ia-writer-pro/id775737590?mt=12)、[Ulysses ¥283](https://itunes.apple.com/cn/app/ulysses-iii/id623795237?mt=12) etc.  
Windows下Markdown编辑器：[MarkdownPad](http://markdownpad.com/)、[MarkPad](http://markpad.fluid.impa.br/) etc.
	
## 官方文档

[英文Markdown文档](http://daringfireball.net/projects/markdown/syntax)  
[中文Markdown文档](http://wowubuntu.com/markdown/)

	
	
	
	
	
	
	
	
	
	