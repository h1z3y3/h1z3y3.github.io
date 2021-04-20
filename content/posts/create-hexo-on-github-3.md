---
title: 在Github上面搭建Hexo博客（三）：更换Hexo主题
date: 2015-10-14 11:18:22
tags: [Hexo, Github]
---

更换主题很简单，首先感谢[@iissnan](https://github.com/iissnan)，因为博主用的是[@iissnan](https://github.com/iissnan)的[NexT](http://theme-next.iissnan.com)主题，所以本文就以[NexT](http://theme-next.iissnan.com)主题为例讲解如何更换博客主题。在文末我会给出更多优秀主题的Github地址。

<!-- more -->

# 为自己的Hexo博客更换主题

我这里讲到的也是参考了[NexT](http://theme-next.iissnan.com)主题作者的[使用文档](http://theme-next.iissnan.com)，更详细的设置读者们可以直接去参考该[使用文档](http://theme-next.iissnan.com)。

## 下载[NexT](http://theme-next.iissnan.com)主题包

仍然在博客根目录下运行Git bash命令行工具，输入下面一条指令用以克隆最新版本：

````
> git clone https://github.com/iissnan/hexo-theme-next theme/next
````

## 启用主题

下载完成后，打开博客配置文件`_config.theme`，修改里面的`theme`字段，将其值设置为`next`

## 验证主题是否启用成功

在Git Bash中输入`$ npm server`启用本地服务，然后在浏览器中输入`localhost:4000`进行预览即可

更多详细的设置，不再赘述，可以参考[NexT](http://theme-next.iissnan.com)的[使用文档](http://theme-next.iissnan.com)

# 其他优秀的主题推荐：

* [iissnan/hexo-theme-next](https://github.com/iissnan/hexo-theme-next.git)
* [TryGhost/Casper](https://github.com/TryGhost)
* [daleanthony/uno](https://github.com/daleanthony)
* [orderedlist/modernist](https://github.com/orderedlist)
* [litten/hexo-theme-yilia](https://github.com/litten/hexo-theme-yilia)
* [A-limon/pacman](https://github.com/A-limon)
* [AlxMedia/hueman](https://github.com/AlxMedia)
* [kathyqian/crisp-ghost-theme](https://github.com/kathyqian)

更多优秀的主题可以参考知乎回答：[有哪些好看的hexo主题?－家顺张的回答](http://www.zhihu.com/question/24422335)



