---
title: 在Github上面搭建Hexo博客（四）:使用不同电脑维护
date: 2015-10-14 12:13:01
tags: [Hexo, Github]
---

这一篇是本系列的最后一篇，到目前为止，我们搭建的Hexo博客已经能满足我们日常的需求。可是有没有想过我们其实只能在这一台电脑上更新我们的博客？如果有一天我换了一台电脑，或者公司和家里不同的电脑都想更新博客，应该怎么办呢？

这里只给类似我这样的git新人做参考，git的很多用法我也不太熟练。如果有错误还请各位指正。

Note: **这里需要注意的是，当我们执行`$ hexo deploy`，部署到[Github](https://github.com)
上面的是hexo给我们生成的静态页面，并不是整个hexo博客工程文件，所以并不能简单的在不同PC的更新博客。**

其实有两种办法，第一就是**_把整个目录备份到云盘_**，然后开启云盘同步该文件夹，虽然操作简单，但是同步很麻烦，需要手动将文件夹进行覆盖。第二种就是**_使用Git的第三方服务_**
，只要配置完成，不管到哪里，用哪台电脑都能简单实现更新博客（当然需要Node.js和Git运行环境，在我们[本系列的第一篇](/2015/10/12/create-hexo-on-github-1/)有讲解）。

在本文中详细讲解如何使用第三方git服务进行博客的备份。可以[Github](https://github.com)放在共有仓库，如果你担心泄密，可以用[Github](https://github.com)
的私有仓库（收费），国内除了[Github](https://github.com)
还有许多知名的git服务商，如：gitcafe，bitbucket，oschina，coding等，据我了解，oschina的私有仓库是免费的，而且可以和[Github](https://github.com)
进行同步。因为我们的博客放在了[Github](https://github.com)，所以我们不妨就把我们博客程序也同步到[Github](https://github.com)的仓库中。

下面我们就来讲解如何实现不同电脑同时更新博客。

<!-- more -->

# 将我们的博客目录备份到[Github](https://github.com)，实现多PC维护

## 在[Github](https://github.com)网站创建一个新的repository

我们在这里给新创建的repository命名为blog；

不会创建方法的朋友可以参考[在Github上面搭建Hexo博客（一）](/2015/10/12/create-hexo-on-github-1/)中的创建方法。

## 在A电脑中从本地上传Hexo到[Github](https://github.com)仓库

Note: **A电脑指的是建立Hexo博客的电脑。**

## 初始化仓库

在Hexo博客的根目录运行Git Bash并输入以下命令：

````
> git init
> git remote add origin <server address>
````

这里`<server>`指的是在线仓库的地址，比如在这里我的就应该是`https://github.com/Gitzhaoyang/blog.git`，如果你用其它git仓库服务，填写对应仓库地址即可。
`origin`是本地分支,`remote add`会将本地仓库映射到[Github](https://github.com)仓库

## 把本地文件同步到[Github](https://github.com)上面

分别输入执行以下命令：

```
> git add .  #添加所有目录，注意add后面有个点`.`
> git commit -m "add to Github"  #添加提交说明，每次提交都需要
> git push -u origin master      #把更新推送到云端
```

这时可以登录[Github](https://github.com)账户查看刚创建的`blog仓库`中是否上传成功

windows平台可能push过程中会提示输入[Github](https://github.com)的用户名和[Github](https://github.com)的密码，输入正确便是。

### 注意

为了在另一台电脑上配置更加方便，严重建议把Hexo博客目录下`_config.yml`文件复制粘贴一份，并重命名为`hexo_config.yml`；把themes目录下你用到主题目录下的`_config.yml`
文件也复制一份，并粘贴到**博客根目录，注意，是'博客根目录'**，并命名为`theme_config.yml`。原因是我们上传的时候，我们自己安装的`themes`
如：`[NexT](http://theme-next.iissnan.com)`，它的`'next'`目录并**不能上传**，所以我们需要把这两个配置文件都保存下来在进行同步工作。

## 在B电脑中从[Github](https://github.com)仓库取回Hexo到本地

Note: **B电脑指的是另一台电脑，如果没有另一台电脑也可以找地方新建一个文件夹尝试。**

## 安装Git和Node.js

值得注意的是新电脑也需要安装Git和Node.js环境，参考[本系列的第一篇](/2015/10/12/create-hexo-on-github-1/)中安装方法。

## 把文件取回本地

安装环境完成后，在新文件夹下运行Git Bash并分别执行以下几条命令：

````
> git init $ git remote add origin <server>
> git fetch --all $ git reset --hard origin/master
````

这里`<server>`仍然是你的Giuhub地址。`fetch`是将仓库中的内容取出来。`reset`则是不做任何合并（merge）处理，直接把取出的内容保存。

运行完`reset`后你会发现文件夹中就会出现刚刚上传的内容。但是配置并没有完成，请继续往下看。

## 配置新的Hexo

Note: **如果是新PC，不要忘记我们本机并没有安装Hexo博客**

首先，在刚才的目录下执行以下命令以在新机器中安装Hexo

````
   $ npm install hexo --save
````

初始化Hexo并安装相应依赖包

````
> hexo init
> npm intall
````

记得在第一篇中讲过，新安装的Hexo是没有`hexo-deployer-git`依赖包的，需要手动安装

````
> npm install hexo-deployer-git --save
````

如果你在A机器上设置了`订阅（feed）`，那么你需要重新烧制feed，需要重新安装依赖包，没有设置feed的可以略过

````
> npm install hexo-generation-feed --save
````

安装主题，我在上文中提到新安装的主题并不能被上传，所以也需要重新手动安装(以[NexT](http://theme-next.iissnan.com)主题为例)

````
> git clone https://github.com/iissnan/hexo-theme-next themes/next
````

这里要注意的是：`themes/next`是**主题保存目录**。

我们之前备份的两个配置文件`hexo_config.yml`和`theme_config.yml`有用了,`hexo_config.yml`重命名为`_config.yml`
   覆盖根目录下的同名文件，而`theme_config.yml`也重命名为`_config.yml`覆盖主题目录下的`config.yml`文件。注意文件名前面的下划线'_'。

输入命令 `$ hexo generate` 和命令 `$ hexo server` ，
然后在浏览器输入 `localhost:4000` 中进行预览。
如果没有问题那么我们在B电脑上就配置成功了。

## 在B电脑上更新博客

现在在B电脑上也可以像在A电脑上一样更新博客了，同样是`$ hexo new post "my_new_post"`
编辑完文章，然后执行`$ hexo generate`和`$ hexo deploy`就可以成功发表了。
这里`$ hexo deploy`命令是将我们的博客文章发表到我们的[Github](https://github.com)
上的Hexo博客，并不是前文新建的`blog仓库`，新建的`blog仓库`用来保存我们的Hexo程序。

## 把B电脑上的Hexo从本地同步到[Github](https://github.com)仓库

当发表完文章，我们还需要把Hexo程序同步到我们[Github](https://github.com)的`blog仓库`。执行下面指令：

$ git add .

这是可以输入命令`$ git status`查看状态，回显示刚才编辑过的文件的信息。

之后分别输入下面指令完成上传：

$ git commit -m "commit from PC_B"
$ git push -u origin master

成功后，我们再次把程序同步更新到了我们的[Github](https://github.com)仓库`blog`。 如果再想用A电脑更新我们的博客，只需要在执行添加文章之前先把程序从`blog仓库`拉取下来便可。输入命令：

````
> git pull https://github.com/Gitzhaoyang/blog.git
````

即可完成。

## 注意事项

我们每次更新博客时，为了保持我们每次用到的程序都是最新的。

每次更新博客之前都需要执行`$ git pull https://github.com/xxxx/xxx.git`保持本地最新；

每次更新博客之后都需要执行
`$ git add .` , `$ git commit -m "message"` , `$ git push -u origin master`
以保持 [Github](https://github.com) 仓库程序最新。

好了，现在我们就能实现在不同电脑都能对我们的Hexo博客进行维护了。😎😘
