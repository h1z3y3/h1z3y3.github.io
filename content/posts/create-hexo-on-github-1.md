---
title: 在Github上面搭建Hexo博客（一）：部署到Github
date: 2015-10-12 18:57:44
tags: [Hexo, Github]
---

# 什么是[Hexo](https://hexo.io/)

[Hexo](https://hexo.io/)是一个基于Node.js的静态博客程序，可以方便的生成静态网页托管在[Github](https://github.com/)和Heroku上。并且有很多人为其制作了很多优秀的主题（theme），你可以根据自己的喜好进行设置。主题的设置将在后面的章节中介绍。

这个是[Hexo](https://hexo.io/)官方网站介绍：

> Hexo is a fast, simple and powerful blog framework.
> You write posts in [Markdown](http://daringfireball.net/projects/markdown/)
> (or other languages) and Hexo generates
> static files with a beautiful theme in seconds.

翻译过来就是：

> Hexo 是一个快速、简洁且高效的博客框架。Hexo 使用 Markdown（或其他渲染引擎）解析文章，
> 在几秒内，即可利用靓丽的主题生成静态网页。

# 怎么在[Github](https://github.com/)上搭建一个hexo博客

<!-- more -->

我用了一天时间研究和搭建了一个[Github](https://github.com/)博客(GitHub Pages site)，过程中遇到一些小问题，现在写一篇教程，方便和我一样爱折腾但是是新手的人。

## 注意

因为题主在搭建时在Windows平台，所以讲解为Windows版本，但是各个平台大同小异，所以实践起来并没有很大的差别。

以下为教程正文：

## 安装Git

前往[Git官网](http://git-scm.com/)下载Windows版本压缩包，下载完成后解压安装。

## 安装Node.js

前往[Node.js](https://nodejs.org/en/)官方下载网站，下载Node.js官方安装包，下载完成后同样解压安装。

## 安装[Hexo](https://hexo.io/)

到目前为止，安装[Hexo](https://hexo.io/)所需要的环境已将安装完成，下一步只需要安装[Hexo](https://hexo.io/)便可以了。
点击鼠标右键，看是否有**Git bash Here**选项。如果没有可以前往Git安装根目录，启动**git-base.exe**也可以。 在命令行中输入：

$ npm install -g hexo-cli

[Hexo](https://hexo.io/) 便安装完成了

## 创建[Hexo](https://hexo.io/)文件夹

找到想要放置博客的文件夹，比如（`F:\Hexo`），在该目录下鼠标右击打开Gitbash工具，（右键菜单中没有该选项的可以用cmd命令`cd`等进入该文件夹）。执行下面的语句，会在`F:\Hexo`文件夹下创建`node_modules`
文件夹：

````
> hexo init
````

这里 **init** 后面可以跟文件目录，比如我想在`F:\Hexo`下创建博客文件夹，那么可以用下面的命令：

$ hexo init F:\Hexo

## 安装依赖包

在[Hexo](https://hexo.io/)目录下，执行以下命令，你会发现`F:\Hexo\node_modules`目录下多了好多文件夹

````
> npm install
````

## 本地调试

目前为止，已经搭建好自己的[Hexo](https://hexo.io/)博客了，但是只能在本机上查看。
执行以下两个命令（在`F:\Hexo`目录下），然后在浏览器中输入 `localhost:4000` 就可以看到自己的博客了

````
> hexo generate
> hexo serrver
````

但是只能在本地查看，如果想让别人也能访问，那么就需要部署到[Github](https://github.com/) 上面，下面，我们就部署上去。

## 注册[Github](https://github.com/)账户

前往[Github](https://github.com/)网站，注册一个新用户。用邮箱注册的一定前往邮箱去验证邮件。要不然之后可能会有小问题。

## 创建一个新的repository

在自己的[Github](https://github.com/)主页右下角,创建一个新的`repository`。
比如我的Github用户名为`Gitzhaoyang`，那么我创建的repository的名字应该是 `gitzhaoyang.github.io` 。

![添加reponsitories](https://p3.ssl.qhimg.com/t01eb3bf66c3bc3ad0e.jpg)

## 这里严重注意

一定要以`你的Github用户名.github.io`创建。假如我没有用`gitzhaoyang.github.io`而是用了`mungo.github.io`
，那么当我浏览器访问博客的时候会出现404错误。这里并不是没有部署成功，而是把它部署在了这里:`http://gitzhaoyang.github.io/mungo.github.io`。所以，如果想直接`gitzhaoyang.github.io`访问，那么就需要和用户名保持一致。题主在这里吃了不小的苦头，最后给Github客服发邮件才知道原因。

创建好如下图：

![一定要保持一致](https://p2.ssl.qhimg.com/t010ca5a2c0a78b950e.jpg)

## 将本地的文件部署（上传）到[Github](https://github.com/)账户中

编辑本地[Hexo](https://hexo.io/)目录下文件`_comfig.yml`，在最后添加如下代码（在你修改时，把  `gitzhaoyang` 要替换成你自己的用户名）

````
deploy:
type: git
repository: http://github.com/Gitzhaoyang/gitzhaoyang.github.io.git
branch: master
````

.yml文件对格式规范要求很严格，`type:` `repository:` `branch:` 前面有两个空格，冒号后面都有一个空格。

执行以下指令即可完成部署（如果提示错误，可以看下面**注意**）：

````
> hexo generate
> hexo deploy
````

## 注意事项

* 有些用户没有设置Github的SSH，会导致上面两句失败。SSH的介绍和设置方法可以查看
   [官方教程](https://help.github.com/articles/generating-ssh-keys),配置起来很简单。
   如果英文看不明白或者过程中出现小问题， 可以查看我写的
   [SSH设置教程](/2015/10/13/how-to-config-ssh-on-github/index.html) ，是对官方教程的解释和扩展，
   针对配置过程中的小问题都有解决办法。
* 每次修改本地文件，都需要命令`$ hexo generate`才能保存。而且每次使用命令都必须在 [Hexo](https://hexo.io/) 根目录下使用。
* 如果你在执行`$ hexo deloy`,如果提示 `ERROR Deployer not found: git`，那说明你没有安装`hexo-deployer-git`依赖包，进入`F:\Hexo\node_modules`
   发现真的没有`hexo-deployer-git`，不用担心，只需要输入下面命令创建`hexo-deployer-git`依赖包，然后再执行`hexo deploy`就能上传成功了

   ````
   > npm install hexo-deployer-git --save
   ````

* 如果你是windows用户，那么当你执行`$ hexo deploy`命令的时候，
   可能会先后出现提示框让你输入你的**Github用户名**和**Github密码**，只要输入正确，上传就没有问题。

好了，现在我们的博客已经在Github上面部署成功了，可以在浏览器访问`gitzhaoyang.github.io`试试了。

## 提示

现在Hexo支持更加简单的命令格式了，比如：

````
hexo s == hexo server
hexo g == hexo generate
hexo d == hexo deploy
hexo n == hexo new
````

后续我会把如何配置博客信息，发表文章，设置博客主题，不同电脑间进行同时更新自己的Blog的方法等更新上来,感兴趣的人可以关注








































