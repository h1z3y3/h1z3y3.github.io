---
title: go module 使用 gitlab 私有仓库
date: 2021-05-21 12:33:52
tags: [Golang]
---

# Go Module 使用 gitlab 私有仓库

包管理是 Go 一直被诟病做的不好的功能。在 1.11 之前，go get 缺乏对依赖包版本管理和 reproducible build 的支持。
当时在 Go 社区当时诞生了许多好用的工具，比如 glide，dep 等。
在 1.11 版本之后， Go 引入了 Go Module，再也没有 GOPATH 的限制，你可以随意在任何路径写项目，但是此时对私有仓库的支持还不是很好。 
而在 1.13 版本之后， Go 对 Go Module 又进行了优化，支持了 `GOPRIVATE` 环境变量，可以指定私有仓库的地址，使用十分便捷。
大家在使用过程中，或多或少地会遇到一些问题，下面我针对自己遇到的问题进行总结。

## go get

如果在没有进行任何设置的情况下直接执行 `go get your.gitlab.com/pkg/example`，你很可能会遇到以下错误：

```
go get: module your.gitlab.com/pkg/example: git ls-remote -q origin in /go/pkg/mod/cache/vcs/a39fc2dbfb0a9645950d24df5d7e922bb7a6a877aecfe2b20f74b96385a83109: exit status 128:
	fatal: could not read Username for 'https://your.gitlab.com': terminal prompts disabled
Confirm the import path was entered correctly.
If this is a private repository, see https://golang.org/doc/faq#git_https for additional information.
```

其实错误提示已经把解决方案给到我们了，我们只需要点击 [https://golang.org/doc/faq#git_https](https://golang.org/doc/faq#git_https) 查看即可。

下面是原文：

> #### Why does "go get" use HTTPS when cloning a repository?
> 
> Companies often permit outgoing traffic only on the standard TCP ports 80 (HTTP) and 443 (HTTPS), blocking outgoing traffic on other ports, including TCP port 9418 (git) and TCP port 22 (SSH). 
> When using HTTPS instead of HTTP, git enforces certificate validation by default, providing protection against man-in-the-middle, eavesdropping and tampering attacks. 
> The go get command therefore uses HTTPS for safety.
> 
> Git can be configured to authenticate over HTTPS or to use SSH in place of HTTPS. To authenticate over HTTPS, you can add a line to the $HOME/.netrc file that git consults:
> ``` 
> machine github.com login USERNAME password APIKEY
> ```
> For GitHub accounts, the password can be a [personal access token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/).
> Git can also be configured to use SSH in place of HTTPS for URLs matching a given prefix. For example, to use SSH for all GitHub access, add these lines to your ~/.gitconfig:
> ```
> [url "ssh://git@github.com/"]
>       insteadOf = https://github.com/
> ```
 
大概意思是，HTTPS 更安全，所以 `go get` 命令使用 HTTPS。

如果你要用 HTTPS，那你就需要配置 HTTPS 的用户名和密码：

``` 
machine github.com login USERNAME password APIKEY
```

当然也可以使用 ssh，需要修改你的 git 配置，

修改当前用户的 `~/.gitconfig`，添加：

```
[url "ssh://git@your.gitlab.com/"]
      insteadOf = https://your.gitlab.com/
```

另外执行下面的命令也能达到同样的效果：

```bash
git config --global url."git@your.gitlab.com/".insteadof "https://your.gitlab.com/"
```

操作完之后，我们就可以使用 `go get` 了，使用 `go get -v` 可以展示执行日志。

## GONOPROXY

众所周知，国内用户大多设置代理，我们在 Go 1.12 之前如果使用 `GOPROXY` 这个环境变量设置代理，并使用私有仓库，很有可能会遇到下面的错误：

```
go get your.gitlab.com/pkg/example: module your.gitlab.com/pkg/example: 
reading https://goproxy.cn/your.gitlab.com/pkg/example/@v/list: 404 Not Found
```

这是因为代理服务不可能访问到我们的私有代码仓库，所以报错 404。而且，就算使用上文提到的 `ssh` 鉴权也不行。

**Go 1.13 后可以设置 `GONOPROXY` 这个环境变量来指定不使用代理的域名，支持逗号分隔多个值。**

## GONOSUMDB

go mod 需要对下载后的依赖包进行 checksum 校验，当你的 git 仓库是开放的话没问题，但是如果是不可访问的私有仓库，甚至在公司内网。
很可能出现校验失败的错误：
     
```bash
get "your.gitlab.com/pkg/example": found meta tag get.metaImport{Prefix:"your.gitlab.com/pkg/example", VCS:"git", RepoRoot:"https://your.gitlab.com/pkg/example.git"} at //your.gitlab.com/pkg/example?go-get=1
  verifying your.gitlab.com/pkg/example@v0.0.0: your.gitlab.com/pkg/example@v0.0.0: reading https://sum.golang.org/lookup/your.gitlab.com/pkg/example@v0.0.0: 410 Gone
```

和代理一样，我们的私有仓库对 sum.golang.org 是不可见的，所以肯定没办法执行安全校验。

**同样的在 Go 1.13 后可以设置 `GONOSUMDB` 环境变量指定跳过校验的的域名，支持逗号分割多个值。**

## GOPRIVATE

最后 Go 1.13 还引入的 `GOPRIVATE` 环境变量，可以说设置后一劳永逸，能自动跳过 proxy server 和 校验检查，
这个变量值也支持逗号分割，可以填写多个值，如：

```bash
GOPRIVATE=*.corp.example.com,your.gitlab.com
```

当然，设置 `GOPRIVATE` 之后，还可以在通过 `GONOPROXY` 和 `GONOSUMDB` 来单独进行控制，

不过需要注意下 `GOPRIVATE` 失效的问题，

举个例子，如果公司内部有私有仓库：`your.corp.com`，如果这样设置：

```bash 
GOPRIVATE=your.corp.com
GOPROXY=https://goproxy.cn 
GONOPROXY=none   
```

因为 `GONOPROXY` 的值是 `none`，那么用户还是会从 `GOPROXY` 的地址下载所有私有和共有的仓库，
此时可能还是会报错，`GONOSUMDB` 同理，大家注意一下这个问题。








 


