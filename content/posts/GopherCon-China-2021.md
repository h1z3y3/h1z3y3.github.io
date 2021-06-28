---
title: 2021 北京 Gopher 大会
date: 2021-06-28 09:01:11
tags: [Golang]
---

![](https://raw.githubusercontent.com/h1z3y3/h1z3y3.github.io/master/images/GopherCon-China-2021/logo.png)

托 [团队 leader](https://fukun.org/) 的福，白嫖了今年 GopherCon China 的票，和历届大会相比，第二天分了两个分会场，
分享不同的内容，参会者可以根据个人兴趣自行选择参加。

然而不得不说，本次和之前几次的分享安排相比，多了一些针对自己公司产品的推广。
但不管怎么说，收获还是有的，印象比较深的 talk 有 PingCAP CEO 黄东旭分享的《全链路可观测性：从应用到 Go Runtime》，比较有趣；还有曹大曹春晖分享的《Go 语言的抢式调度》，内容确实很卷很深入。

大会提到的很多技术其实我们团队也早已经实践使用，比如鸟窝在《深入探究 Go Module》中分享如何使用 GOPRIVATE 从而使用自己的私有仓库；
Grab 公司在分享《Improving Go Backend Developer Experience In Grab》中提到的 CI/CD 我们也已经在我们大多数项目中使用。

我总结了讲师提到最多的关键词，一方面也能体现业内目前在流行什么：

* **K8S**：毋庸置疑，云原生时代的王者
* **API Gateway**、**L7 Banlancer**：处理南北向流量，两者定位不同又有一部分功能重合，APISIX、BFE 最近两年大家讨论比较多的开源项目
* **Service Mesh**：处理东西向流量，Envoy、xDS协议
* **微服务**：不再赘述
* **Tracing**、**Metrics**、**Logging**：及时发现服务问题
* **CI/CD**：加快应用的部署速度，解放开发者双手不可或缺的步骤
* **Rust**：没错，Gopher 大会确实多次提到了 Rust 语言

我针对我自己参加的几个分享谈一谈我自己的一些感受和思考。

大会的 PPT，可以在这里获取：[https://mp.weixin.qq.com/s/734ac0JeQtSrzcmZ1-so8w](https://mp.weixin.qq.com/s/734ac0JeQtSrzcmZ1-so8w)

##《Generic in go》

大会第一场是跟老外连线，但是直播断断续续，我这蹩脚听力更听不懂了，体验不太好。但是针对泛型这个主题，在社区中有很多讨论，
也有一些争议，主要是泛型的支持者，和认为加入泛型破坏 golang 简单特性的反对者，
不过今年官方正式提出将泛型特性加入 Go 的 
proposal [https://blog.golang.org/generics-proposal](https://blog.golang.org/generics-proposal) [https://github.com/golang/go/issues/43651](https://github.com/golang/go/issues/43651) 。
而且，golang 官方已经在 master 分支实现了泛型，鸟窝也在早些时间给出了 Go 泛型尝鲜的方法：https://colobu.com/2021/03/22/try-go-generic/

##《MOSN在云原生的探索和实践》

MOSN 是一款网络代理软件，可以与任何支持 xDS API 的 Service Mesh 集成，也具备南北向流量代理的功能。
不得不佩服 MOSN 对 CGO 的调研和使用深度，其中介绍的 MOE（Mosn on Envoy）使得 MOSN 底层网络具备 C 同等处理能力的同时，
又能使上层业务可高效复用 MOSN 的处理能力及 Golang 高效的开发效率。


##《浅谈全链路可观测性：从应用到Go Runtime》

我个人认为很不错的分享，由浅入深带领大家如何 Trace In Go Runtime，PPT 风格还有讲解风格都很有趣。
最终的实现方式我也没想到和我最近翻译的这篇文章相关
（ [深入剖析 Golang Pprof 标签](https://h1z3y3.me/posts/demysitifying-pprof-labels-with-go/) ），
翻译时也没想到还能这么玩，当时听了确实也令人眼前一亮。

![](https://raw.githubusercontent.com/h1z3y3/h1z3y3.github.io/master/images/GopherCon-China-2021/profile-label.png)

## 《Improving Go Backend Developer Experience in Grab》

算是 Grab 公司的经验分享，关于整个软件的设计、开发和交付流程。和 Go 的关系并不是很大，很多经验经验也很难进行复制。

![](https://raw.githubusercontent.com/h1z3y3/h1z3y3.github.io/master/images/GopherCon-China-2021/build-times.png)

## 《利用夜莺扩展能力打造全方位监控体系》

夜莺也是国内应用很火的监控系统，和 Prometheus 相比，有一些优势吧，比如有自身的告警平台，
我们团队内部使用的基于 Prometheus 又另外开发的一套报警系统，目前也是开源，地址是 [https://github.com/Qihoo360/doraemon](https://github.com/Qihoo360/doraemon) 。
更多夜莺和Prometheus的对比可以看这篇分析文章，[《夜莺与Prometheus的对比》](https://www.yuque.com/ictc/manual/nr798n) 。
但是干货太少，全程介绍产品功能。

## 《如何构建易于拆分的单体应用》

由于项目的特殊性，我们组内大多数应用都是单体应用，也在尝试做一些微服务的拆分工作。这个分享虽然没讲 DDD，
但从现实例子出发，分析一个应用如何设计、开发，同时实战讲解了如何使用 go-kit 构建一个简单的应用，还是有些收获。

## 《深入探索Go Module: 实践、技巧和陷阱》 

内容充实还不错，一些比较常见的用法我们团队也一直在用，我们的也是去年开始从 dep 全部换到了 go module，
所以积累了一些 go module 的使用经验，所以整个分享听下来也比较轻松。我之前也写了一篇文章介绍如何让 go module 
使用 gitlab 私有仓库：[https://h1z3y3.me/posts/go-private-git-repository/](https://h1z3y3.me/posts/go-private-git-repository/) 。
鸟窝还介绍了一些之前基本没用过的子命令，比如 `go mod graph`、`go mod why`，收获最大的是终于弄明白了 go module 拉取依赖包的策略。

![](https://raw.githubusercontent.com/h1z3y3/h1z3y3.github.io/master/images/GopherCon-China-2021/go-module-history.png)
![](https://raw.githubusercontent.com/h1z3y3/h1z3y3.github.io/master/images/GopherCon-China-2021/go-module-xxx.png)

## 《深入理解 BFE》

也算是推广分享，介绍 BFE 产品，没什么干货。

我们 19 年也开源了一个七层转发的代理 [「HTTPS Layer」](https://github.com/Qihoo360/HTTPSLayer) ，
和 BEF 相比，功能基本类似，甚至可以说 80% 相同。最大的区别就是 BFE 基于 Go 生态，我们基于 Openresty 生态。
或许缺少了维护和宣传，我们并没有获得很高的关注，比较可惜。其实 HTTPS Layer 在我们公司内部也得到了生产环境的验证，
做了一些分布式部署之后，目前我们单机房峰值 QPS 也能达到 10w QPS，当然还预留了一些 buffer 的情况下。

![](https://raw.githubusercontent.com/h1z3y3/h1z3y3.github.io/master/images/GopherCon-China-2021/why-bfe.png)

## 《Go 语言的抢占式调度》

曹大曹春晖的分享，很卷，讲解了 GMP 模型以及 Golang 的抢占式调度，干货又太干了，有点吸收不了。
主要内容是 golang 1.14 版本前后两种不同的抢占式调度模型，深入源码，通过编译后的汇编进行讲解，挺有深度的，之后会再仔细研究这方面的知识。

![](https://raw.githubusercontent.com/h1z3y3/h1z3y3.github.io/master/images/GopherCon-China-2021/GMP.png)

## 《K8S私有云建设实践》

公司的经验分享，不可复制，因为太客制化了。为了迎合公司现有平台，抛弃了很多 k8s 的原生用法，我个人觉得很鸡肋。
针对私有云建设，360 其实走在比较前面，之前还开源了 k8s 的多集群管理平台 
[Wayne](https://github.com/Qihoo360/wayne) ，开源初期获得很多关注，但是后续也不没有持续更新，也比较可惜。

## 《Go 如何助力企业进行微服务转型》

go-zero 的维护者万俊峰分享，虽然经验之谈多一些，关于 Go 的讨论少，但是整体分享是好分享，收获很多，
也有些之后可借鉴的点。不多赘述，PPT 可以下载下来看，里面挺详细。

![](https://raw.githubusercontent.com/h1z3y3/h1z3y3.github.io/master/images/GopherCon-China-2021/monolith-to-microservice.png)

## 总结

从整体来看，本次 Gopher 大会还算成功，有些小插曲，但是瑕不掩瑜。只是希望之后能少一些对于产品的推广和介绍，
对于产品功能完全可以去看各自产品的文档，我能想到他们和 Go 唯一相关的就是这个产品是 Go 写的。
作为技术人花了钱到这里想看到的还是对于 Go 的一些可借鉴和可复制的实践经验，这个其实看下国外的 gopher 大会就能体会到了。

不管怎么说，虽然周末两天都没睡懒觉，7点早起就往会场赶，但是最后还是心满意足。我看大会全程有录像，
大家可以参考我的思考，去翻翻对应的 PPT 或者视频，相信也能有所收获！