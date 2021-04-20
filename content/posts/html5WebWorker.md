---
title: HTML5 WebWorker 简单使用
date: 2016-03-20 15:18:33
tags: [HTML5, WebWorker]
---

1. 什么是WebWorker
	
	web worker 是运行在后台的 JavaScript，独立于其他脚本，不会影响页面的性能。我们知道页面的展示放在主线程，如果让主线程进行一系列复杂的操作，那么页面就会变得非常卡，用户体验会很差。这是我们可以使用web worker进行复杂操作的实现，然后将处理结果返回给页面，页面进行更新即可，这样就不会影响用户主页面展示的执行。
	
2. 方法：
	postMessage() : 用于向HTML页面返回消息
	terminate() : 终止web worker， 并且释放资源
	
<!--more-->

#### 实现方法：

Demo: 数字从0开始累加

#### index.html

	<!DOCTYPE html>
	<html lang="en">
	<head>
	    <meta charset="UTF-8">
	    <title></title>
	    <script src="index.js"></script>
	</head>
	<body>
	<div id="numDiv">0</div>
	</body>
	</html>
	
#### index.js

	var numDiv;
	
	window.onload = function() {
	    numDiv = document.getElementById("numDiv");
	    var worker = new Worker("webWorker.js");//创建Worker对象
	    worker.onmessage = function (e) {
	        numDiv.innerHTML = e.data;
	    }
	}
	
#### webWorker.js

	var countNum = 0;
	
	function count () {
	    postMessage(countNum);//给html页面返回数据
	    countNum ++;//数字累加
	    setTimeout(count, 1000);//一秒执行一次
	}
	
	count();//调用函数执行


