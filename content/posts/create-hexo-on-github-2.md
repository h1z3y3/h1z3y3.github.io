---
title: 在Github上面搭建Hexo博客（二）：配置和发表文章
date: 2015-10-13 23:42:28
tags: [Hexo, Github]
---

# 如何配置Hexo

上一节中我们已经在本地和Github上搭建起了自己的博客，但是博客的配置都是默认值，如果我们想个性化自己的博客，我们应该做什么呢。这一节中我们一起来配置自己的博客的基本信息和介绍如何写博客和发表博客。下一节我们将一起为自己的博客安装新的主题。

博客的主要配置用到根目录下的`_config.yml`文件，我在下面给出文件和解释，你只需要根据自己需求作出简单更改即可：

````
# Hexo Configuration
## Docs: http://hexo.io/docs/configuration.html
## Source: https://github.com/hexojs/hexo/

# Site
title: Mungo's Note##站点的名字
subtitle:##站点的副标题
description: 日常技术分享 ##站点介绍，对站点进行描述
author: Mungo##站点文章的作者
email: gmzhaoyang@gmail.com##你的邮箱地址
language:##语言，默认中文，不填写即可
timezone:##时区，默认即可

# URL
## If your site is put in a subdirectory,
## set url as 'http://yoursite.com/child' and root as '/child/'
url: http://mungo.space##站点的域名，如果奇怪我为什么可以用自己的域名，可以看后续更新
root: /
permalink: :year/:month/:day/:title/
permalink_defaults:

## 对资源文件夹的配置，如资源文件夹名称，标签云名称，分类页面名称
# Directory
source_dir: source##资源文件夹,当执行`$ hexo deloy`命令，上传的即是该文件夹里面的内容
public_dir: public##公共文件夹，当执行`$ hexo generate`命令，生成的文件都在里面
tag_dir: tags##标签云文件夹，需要自己生成，详情见下一节如何配置主题
archive_dir: archives##归档文件夹，需要自己生成
category_dir: categories##分类文件夹，需要自己生成
code_dir: downloads/code##代码存放区
i18n_dir: :lang
skip_render:

##此处时配置博客文章内容格式的，可以保持默认，不做修改
# Writing
new_post_name: :title.md # File name of new posts
default_layout: post
titlecase: false # Transform title into titlecase
external_link: true # Open external links in new tab
filename_case: 0
render_drafts: false
post_asset_folder: false
relative_link: false
future: true
highlight:
  enable: true
  line_number: true
  auto_detect: true
  tab_replace:

#分类和标签云的配置，可以不做修改，默认即可
# Category & Tag
default_category: uncategorized
category_map:
tag_map:

#日期和时间格式的配置，可以不做修改，默认即可
# Date / Time format
## Hexo uses Moment.js to parse and display date
## You can customize the date format as defined in
## http://momentjs.com/docs/#/displaying/format/
date_format: YYYY-MM-DD
time_format: HH:mm:ss

#用来配置每页显示的文章数目，可以根据自己需求自行修改
# Pagination
## Set per_page to 0 to disable pagination
per_page: 10
pagination_dir: page

#插件和主题配置，在这里可以修改自己的主题
# Extensions
## Plugins: http://hexo.io/plugins/
## Themes: http://hexo.io/themes/
theme: next

#上传配置，上一节中我们已经配置完成了，在这里不需要再次修改
# Deployment
## Docs: http://hexo.io/docs/deployment.html
deploy:
 type: git
  repository: https://github.com/Gitzhaoyang/gitzhaoyang.github.io.git
  branch: master

````

当配置完成并保存后，就可以执行`$ hexo generate`生成静态文件，然后执行`$ hexo server`后就可以打开浏览器输入`localhost:4000`进行本地预览了

## tips

如果`_config.yml`文件打开有乱码应该是用到的编辑器的原因，我用的是Sublime2，所以一般不会出现乱码。如果乱码，那么需要把文件格式转换为UTF-8，转化方法我在这里就不再赘述了。

## 如何添加和发表文章

### 新建文章

在Git Bash中输入

$ hexo new post "my_first_post"

### 编辑文章

第一步的命令会在`\Hexo\source\_posts`文件夹下创建一个后缀`.md`文件，你可以在里面添加任何字符串。

这其实是一个`markdown`类型的文件，使用`markdown`语言编写，我这篇博文就是用`markdown`编写的，如果不了解的，可以看我的后续更新，我会把`markdown`的基本使用方法进行整理。

### 生成和上传

* 在GitBash中输入`$ hexo generate`对文件进行生成；

* 生成完成后，可以输入`$ hexo server`，然后在浏览器输入`localhost:4000`进行预览；

* 预览没有问题后，接着输入`$ hexo deploy`，windows平台下会提示输入Github的用户名，然后提示输入Github的登录密码。如果输入正确，等待几秒便能上传成功；

* 现在可以在浏览器中输入`xxxx.github.io`进行访问了。

### 注意

可能不会立即生效，只要等待几分钟或者清空一下浏览器缓存基本就能解决。如果仍然看不到，说明前面步骤操作有错误，重新生成和上传就可以了。

如果实在不行，可以在Git Bash中输入`$ hexo clean`或者手动删除`.deploy_git`文件夹和`db.json`文件再重新生成和上传。

到目前为止，我们已经搭建起自己的博客，可以进行基本的配置，也可以发表文章，后面会有更高阶的设置，如：如何配置主题，如何在不同电脑上都可以更新自己的博客 etc.感兴趣的人可以关注。😘
