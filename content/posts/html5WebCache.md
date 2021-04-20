---
title: HTML5应用缓存简单使用
date: 2016-03-20 14:53:00
tags: [HTML5, Cache]
---

1. 什么是应用缓存？

	HTML5引入了应用缓存概念，意味着在没有因特网连接时也可以进行访问。

2. 使用应用缓存好处：

	- 离线浏览，没有因特网的情况下依然可以进行访问
	- 访问速度提升，已经缓存的资源加载更快
	- 减少服务器负载，浏览器只需要下载更新过的页面资源
	
3. 实现方法：
	
	如果需要使用应用缓存，需要在页面`<html>`标签中包含 `manifest` 属性，而manifest文件建议使用文件扩展名`.appcache`。
	
	
4. Manifest文件功能：
	
	- CACHE: 在此标题下列出的文件会在首次访问加载之后进行缓存；
	- NETWORK: 在此标题下列出的文件需要与服务器连接、且不会被缓存；
	- FALLBACK: 在此标题下列出的文件规定当页面无法访问时的退回页面（如404页面）
	
<!--more-->

### 功能实现

#### index.html  请注意`<html>`标签
	
	<!DOCTYPE html>
	<html lang="en" manifest="index.appcache">
	<head>
	    <meta charset="UTF-8">
	    <title></title>
	    <link rel="stylesheet" href="style.css">
	    <script src="index.js"></script>
	</head>
	<body>
	    <h1 class="h1">HELLO HMTL5</h1>
	</body>
	</html>
	
#### style.css

	.h1 {
    	color: red;
    	background: blue;
	}
	
#### index.js

	/*空文件*/
	
#### index.appcache

	CACHE MANIFEST

	CACHE:
	index.html
	style.css
	index.js
	
	NETWORK:
	
	FALLBACK:
	
#### 测试

- 开启本地服务器，在Chrome输入`localhost/webCache`
- 使用Chrome审查元素功能，切换到Resources功能标签，点击左侧`Application Cache`功能标签，可以观察到我们设置缓存的三个文件已经缓存成功

	![缓存列表](https://p1.ssl.qhimg.com/t01cc408c49a3c3052e.jpg)
	
- 断开本地服务器，重新刷新页面，会发现页面样式仍然保持，说明缓存起作用了。



