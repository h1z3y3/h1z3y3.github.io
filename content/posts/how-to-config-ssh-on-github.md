---
title: 如何配置Github的SSH key
date: 2015-10-13 16:02:34
tags: [Github, SSH]
---
## 什么是SSH

SSH为Secure Shell的缩写，是建立在应用层和传输层基础上的协议。SSH是目前较可靠，专为远程登录会话和其他网络服务提供安全性的协议。利用SSH协议可以有效防止远程管理中额信息泄漏问题。我理解的就是给数据进行加密，然后防止中间人进行盗取，能使你的数据安全可靠的传输到目的方。在这里就是为了保证你电脑和[Github](https://github.com)仓库之间通信的安全。

## 如何配置[Github](https://github.com)的SSH

<!-- more -->
题主这里用的依然是windows平台。

### Step 1: 检查SSH keys

首先，我们要查看在你电脑上已经存在的SSH keys。运行Git Bash 然后输入：

	$ ls -al ~/.ssh
	
如果你已经有SSH公钥了。那么你将会看到下面格式的文件名字：

	id_dsa.pub
	id_esdsa.pub
	id_ed25519.pub
	id_rsa.pub

如果你已经存在公钥了，那么可以跳过**Step 2**直接去**Step 3**了。如果没有也不要担心，我们将在**Step 2** 会创建公钥。

### Step 2: 生成SSH key
1. #### 在Git Bash中输入下面命令，引号内一定是你的[Github](https://github.com)注册邮箱地址

		$ ssh-keygen -t rsa -b 4096 -C "your_github_email@example.com" 
		#这句作用是生成一个新的SSH key
		
2. #### 等待几秒，当提示让你输入保存地址时，官方特别推荐放在默认位置就可以了。所以这里直接输入**回车**，提示如下：

		Enter file in which to save the key (/Users/you/.ssh/id_rsa):[直接输入回车]
		
3. #### 将会提示你输入一个密码串**(这里输入密码时不会显示在屏幕上的，只要输入正确按回车就好)**:
		
		Enter passphrase （empty for no passphrase）: [输入你想设置的密码]
		Enter same passphrase again：[在输入一遍密码]
		＃虽然说这里可以设置为空，但是推荐用一个更加安全的密码
		
4. #### 输完密码之后，你将会得到你的SSH的指纹(fingerprint)或者id。他看起来如下图：

	![fingerprint](https://p3.ssl.qhimg.com/t0161534b42efa8b260.jpg)
	
### Step 3: 把你的SSH key添加到ssh-agent

1. #### 输入如下命令
	
		$ ssh-agent -s
		
	会响应：
	
		echo Agent pid [端口号]
		
2. #### 加下来输入如下命令，把你的SSH key添加到ssh-agent
	
		$ ssh-add ~/.ssh/id_rsa
		
	如果他提示如下，说明不能打开您身份验证的代理
	
		Could not open a connection to your authentication agent.
		
	只需要输入如下命令即可解决：
	
		ssh-agnet bash
		
	更多关于ssh-agent的细节，可以用man ssh-agent 来查看
	
### Step 4: 把你的SSH key添加到你的[Github](https://github.com)账户

首先你应该把你的 SSH key 复制到你的剪贴板，输入命令即可完成把 SSH key 复制到你的剪贴板：

	$ clip < ~/.ssh/id_ras.pub

添加到你的[Github](https://github.com)账户：
	
1. #### 浏览器登陆你的[Github](https://github.com)账户，点击右上角你的头像，然后点击**Settings**
	
	![点击Settings](https://p4.ssl.qhimg.com/t012de386332fd4275a.jpg)
	
2. #### 进入**Settings**，点击侧栏选项**SSH key**
 
	![点击SSH key](https://p2.ssl.qhimg.com/t0139f709bd2e05334f.jpg)
	
3. #### 单击右边 Add SSH key 按钮
 
	![点击Add key](https://p3.ssl.qhimg.com/t01b1d4cf5d05ee2de6.jpg)

4. #### 在下面输入标题（Title，这个可以自定义）和SHH Key（直接 Ctrl＋V 粘贴就可以）

5. #### 点击下面的`Add key`按钮便可以添加成功了


### Step 5: 测试是否连接成功

1. #### 在Git Bash中输入：
	
		$ ssh -T git@github.com
		# ssh尝试连接到GitHub
		
2. #### 你可能看到下面的警告：

		The authenticity of host 'github.com(207.97.227.239)' can't be established.
		RSA key fingerprint is SHA256:nJKJFKDnDLFJDndndnfkdfldjfldldfjld.
		Are you sure you want to continue connecting (yes/no)?
		
	确定提示信息里的指纹（fingerprint）是否匹配，如果匹配就键入｀yes｀，将得到：
	
		Hi [你的用户名]! You've successfully authenticated, but GitHub does not 		provide shell access.
		
3. #### 如果提示信息中你的用户名是你的，那么你就成功建立了SSH key！😎😎

### TIPS：如果遇到其他问题，可以参考[官方文档](https://help.github.com/articles/generating-ssh-keys/)，也可以给我留言 





	

	

	
	


	

		
		
		
	
		
	
