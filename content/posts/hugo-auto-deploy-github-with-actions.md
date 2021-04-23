---
title: 使用 Github Actions 自动发布 hugo 站点
date: 2021-04-22 22:21:44
tags: [hugo]
---

从 2018 年底到 2021 年初，博客一直搁置，虽然也偶尔写几篇文章，但是都发布在团队的公众号中了， 现在重新开始维护自己的博客，也从
[hexo](https://github.com/hexojs/hexo)
转移到了
[hugo](https://github.com/gohugoio/hugo) ，
这两个都是非常优秀的静态网站生成工具。

但是不得不说，使用 hugo 之后感觉易用性以及文件的生成速度都优于我之前使用的 hexo 的版本（说的严谨一些，我也不知道现在如何）。

同时我还给自己重新挑选了主题，使用的是
[hugo-theme-meme](https://github.com/h1z3y3/hugo-theme-meme) ，
并且我还做了一部分自定义， 比如 文章列表页 的样式、字体、还修改了移动端标题过长样式错乱的 bug，
另外还接入了 [Gitalk]()（已经提 pr 并合并）和页面右下角的 [Webpusher]() (没有提交 PR),
不得不说，把自己的博客评论用 [Github Issue](https://github.com/h1z3y3/h1z3y3.github.io/issues) 来维护的想法也太赞了！

# 使用 Github Action 自动发布 hugo 站点

如何使用 hugo 以及 Github Pages 的基本知识我就不再赘述，本文主要给大家讲解如何使用 Github Actions 自动编译你的站点实时发布，
本文可以帮你解决的问题是：如何使用你的 `*.github.io` 的仓库维护站点的配置 （`config.toml`）、
使用的主题、或者自己 markdown 格式的原文，而不需要创建新仓库。

因为我自己一开始也想创建新仓库维护，但是后来查阅了一些相关文档后，发现完全没必要，下面介绍给大家。

## 定义 Github Actions

下面是我站点的 Hugo Actions 的配置文件，和原始配置有些小改动，
原始配置文件可以在 [这里](https://github.com/peaceiris/actions-gh-pages) 找到

在 hugo 站点的根目录创建 `.github/workflows/gh-pages.yml` ，并将下面的内容拷贝进去。

```yaml
name: github pages

# on 是 Actions 的触发条件，这里的配置说明当 master 分支有提交的时候，根据这个配置文件执行
on:
  push:
    branches:
      - master # Set a branch to deploy

# jobs 是要执行的任务，我们看到他要执行 deploy
jobs:
  deploy:
    runs-on: ubuntu-18.04 # 运行环境
    steps: # 执行步骤

      # checkout 分支
      - uses: actions/checkout@v2
        with:
          submodules: true  # Fetch Hugo themes (true OR recursive)
          fetch-depth: 0    # Fetch all history for .GitInfo and .Lastmod

      # 安装 hugo
      - name: Setup Hugo
        uses: peaceiris/actions-hugo@v2
        with:
          hugo-version: 'latest'
          # extended: true

      # 编译站点
      - name: Build
        run: hugo --minify

      # 创建 CNAME，这个是原始配置中没有的
      - uses: "finnp/create-file-action@master"
        env:
          FILE_NAME: "./public/CNAME"
          FILE_DATA: "h1z3y3.me"

      # 将站点发布到对应分支
      - name: Deploy
        uses: peaceiris/actions-gh-pages@v3
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          publish_dir: ./public
```

Github Actions 更详细的使用可以从
[这里](https://docs.github.com/en/actions/learn-github-actions) 获取。

## 设置 hugo 目录的提交仓库为 *.github.io

如果你的 hugo 没在 github 维护，执行:

```
> git remote add origin git@github.com:your/your.github.io.git
```

如果已经在 github 维护了，执行：

```
> git remote set-url origin git@github.com:your/your.github.io.git
```

如果仓库有文件，很可能你需要使用 `git push -f origin master` 来强制覆盖。

推送到 Github 之后，现在 Github 的代码就是你 hugo 的根目录，这时候我们的 *.github.io 的站点肯定不工作了，接下来往下看。

## Actions 执行

我们切换到 **Actions** 选项卡，可以看到我们刚才提交的工作流已经执行了，
同时在 **Code** 选项卡，也帮我们创建好了一个 `gh-pages` 的分支，
切换之后发现这个 hugo 已经编译好的静态站点。

![](https://raw.githubusercontent.com/h1z3y3/h1z3y3.github.io/master/images/hugo-auto-deploy-github-with-actions/actions-list.png)

## 修改 Github Page 站点的默认分支

没错，`*.github.io` 的仓库，允许你自定义你要提供静态站点的分支，默认是 master，可以自定义为其他的

操作方法：

进入你的 `*.github.io` 的仓库，点击项目的 `Settings` ，并点击左侧的 `Pages`，在这里你可以自定义一些你站点的配置，
在这里你可以自己选择自己站点要使用的`分支`以及`目录`，我们这里切换为 `gh-pages` 这个分支，并点击保存。

![](https://raw.githubusercontent.com/h1z3y3/h1z3y3.github.io/master/images/hugo-auto-deploy-github-with-actions/set-github-pages-default-branch.png)

## 检查

全部设置好之后，稍等几分钟生效后，通过浏览器访问 `your.github.io`，就可以正常访问了，之后每次我们文章有更新，
从 `master` 分支提交之后，Github 就会通过 Actions 帮我们去编译然后发布在 `gh-pages` 分支，十分方便。

如果你设置了 CNAME，那么请参照我 Actions 配置，里面有创建 CNAME 文件的步骤，添加到你自己的配置即可。




