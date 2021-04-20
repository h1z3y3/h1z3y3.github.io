---
title: 最简单方法解决使用Shiro后URL中JSESSIONID的问题
date: 2017-03-30 22:02:19
tags: [J2EE, Shiro]
---

**在J2EE项目中使用Shiro进行权限验证后，每次部署跳转到登录界面总会在链接后面多出`;JSESSION=xxxx`，查了很多，大概有下面几种方法：**  

1. 在web.xml中添加以下代码：  

		<session-config>
	     	<tracking-mode>COOKIE</tracking-mode>
		</session-config>
	
	具体请参考：  
	[http://stackoverflow.com/questions/11327631/remove-jsessionid-from-url]()
	<!--more-->
2. 使用Filter对URL进行rewrite  
	具体请参考：  
	[http://dr-yanglong.github.io/2015/07/07/del-jeesessionid/]()    

3. 有些同学重写了shiro重定向时需要调用的方法`encodeRedirectURL()`和`toEncoded()`  

	具体请参考：  
	[http://dwangel.iteye.com/blog/2275899]()  
	[http://alex233.blog.51cto.com/8904951/1856155]()

4. 后来去看了源码，发现了一个最最简单的方法，如下：  

	shiro源码：  
 	[https://github.com/apache/shiro/pull/31/commits/7f37394c6048d8c8a214eabd312721ddb51adc9b]()  
 
 阅读源码之后，可以发现`DefaultWebSessionManager.java`文件中添加了新属性`private boolean sessionIdUrlRewritingEnabled;`, 顾名思义, 是用来控制是否重写URL添加SESSIONID的，只要修改shiro的sessionManager配置如下即可：  
 

		<bean id="sessionManager" class="org.apache.shiro.web.session.mgt.DefaultWebSessionManager">
			...
			<property name="sessionIdUrlRewritingEnabled" value="true"/>
			...
		</bean>


#### More Source：  

[https://fralef.me/tomcat-disable-jsessionid-in-url.html]()



