---
title: 译：HTTP/3 来了
date: 2018-12-06 14:43:12
tags: [HTTP]
---

英文原文链接：[https://www.zdnet.com/article/http-over-quic-to-be-renamed-http3/](https://www.zdnet.com/article/http-over-quic-to-be-renamed-http3/)

据 IETF 官方人员透露，HTTP-over-QUIC 实验方案将会被命名为 HTTP/3，并将成为 HTTP 协议的第三个官方版本。    
在谷歌将 SPDY 技术发展成为 HTTP/2 协议之后，这是谷歌第二次将实验技术发展成为 HTTP 的官方协议版本。        
HTTP-over-QUIC 协议是 HTTP 协议的升级，谷歌使用 QUIC 取代 TCP (Transmission Control Protocal) 作为 HTTP 的基础技术。    
QUIC 全称 Quick UDP Internet Connections，是谷歌将 TCP 协议重写为一种结合了HTTP/2、TCP、UDP 和 TLS 的改进技术。   

<!--more-->
 
谷歌希望 QUIC 能逐渐取代 TCP 和 UDP 成为在因特网传输二进制数据协议的新选择，而使用它的更好的理由，是 QUIC 的加密方案实现已经被测试证明更快而且更安全 (目前 HTTP-over-QUIC 协议草案使用的是 TLS1.3 协议)。    

![Google](https://p4.ssl.qhimg.com/t01f147b97634e732d0.png)
	
QUIC 被提议作为2015年 IETF 的标准草案，而 HTTP-over-QUIC 这个基于 QUIC 而不是 TCP 重写 HTTP 的协议则在2016年7月被提议。          
在那之后，HTTP-over-QUIC 在 Chrome 29、Opera 16 被支持，当然还有一些低性能的浏览器。最初，只有谷歌的服务器支持 HTTP-over-QUIC，今年，Facebook也开始采用这项技术。    
在2018年10月份的邮件讨论里，IETF HTTP 和 QUIC 工作组的主席 Mark Nottingham 提出了将 HTTP-over-QUIC 重命名为 HTTP/3 的官方申请，并希望将 QUIC 工作组的开发工作转递给 HTTP 工作组。    
	经过几天的讨论，Nottingham 的提议被 IETF 的成员接受并给出了官方认可，将 HTTP-over-QUIC 作为 HTTP 协议的下一个版本 —— HTTP/3，用于完善优化当今的网络。    
	根据 W3Techs 的统计，截止2018年11月，全球访问量最高的1千万个网站中，已经有31.2%的网站支持了 HTTP/2，只有1.2%的网站支持了 QUIC。
	
*以下是自己的一点思考：*  
可以看到，HTTP/3 协议带来的最大改变是协议底层将采用 UDP 协议，而不再是 TCP 协议，那这样的好处可以说是更低时延和更好的拥塞控制，还有更高效率的多路复用，可以说谷歌真的很厉害了，要知道 HTTP/2 也是谷歌的 SPDY 标准化之后的协议。而且这次 QUIC 发音同 quick ，上次 SPDY 发音同 speedy，是巧合还是有意为之呢。：）

真是强者制定规则：

> They are in control of future web protocol development.

