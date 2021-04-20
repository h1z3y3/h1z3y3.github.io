---
title: JS实现密码加密(base64, md5, sha1)
date: 2016-03-17 12:32:54
tags: [md5, base64, sha1, javascript, 加密]

---

### 简介

在编写Web程序时，表单的提交若密码使用明文提交会十分不安全，因此在浏览器端也要对密码进行加密处理。但是若只是在浏览器端处理了，而服务器没有再一次加密，也是不妥当的，因为"中间人"只要获取了浏览器端加密的密码，不需要进行处理也能进行登录。所以我一般的做法是在前端加密一次，在服务器再加密一次。密码加盐（salt）的问题等我先研究下再写一下。而浏览器端加密一般我都用javascript进行加密后再提交。下面是用javascript编写的base64加密，md5加密和sha1加密。使用方法也极其简单，只要在页面内引入相应js文件即可。

<!--more-->
### base64加解密

	<!DOCTYPE HTML>
	<html>
	<head>
		<meta charset="utf-8">
		<title>base64加解密</title>
		<script type="text/javascript" src="base64.js"></script>
		<script>
			var base64 = new Base64();
			//加密
			var base64encodeStg = base64.encode("hello world!");
			alert("base64encode:" + base64encodeStg);
			
			//解密
			var base64decodeStg = base64.decode(base64encodeStg);
			alert("base64decode:" + base64decodeStg);
		</script>
	</head>
	</html>
	
### md5加密

	<!DOCTYPE HTML>
	<html>
	<head>
		<meta charset="utf-8">
		<title>md5加密</title>
		<script type="text/javascript" src="md5.js"></script>
		<script>
			var hash = hex_md5("hello world!");
			alert(hash);			
		</script>
	</head>
	</html>


### sha1加密
相对于前两个，sha1加密可能更安全
	
	<!DOCTYPE HTML>
	<html>
	<head>
		<meta charset="utf-8">
		<title>sha1加密</title>
		<script type="text/javascript" src="sha1.js"></script>
		<script>
			var sha1 = hex_sha1("hello world!");
			alert(sha1);
		</script>
	</head>
	</html>
	
### 相关文件下载

[base64.js | md5.js | sha1.js 百度网盘分享](http://pan.baidu.com/s/1i3TPnCx)

