---
title: NSDictionary含有null导致写文件(writeToFile)失败(豆瓣电影Api返回Json数据中含有null)

date: 2015-11-10 22:43:40
tags: [iOS]
---

之前从接口获取数据后转换为NSDictionary然后执行writeToFile就可以写入文件成功，昨天在使用[豆瓣电影Api](http://developers.douban.com/wiki/?title=api_v2)时，出现了前几条请求的数据写入缓存成功，从某一次请求开始就总是写入不成功。因为我是用gcd方式请求数据的，我还以为是因为线程竞争的问题，但是写文件的原子操作我设置的YES啊。然后写进程锁试一下，然而仍然写入失败，后来我不用gcd以为能好，结果还是写入失败。后来翻资料，说好像如果NSDictionary中有自定义的object类型是不能写入文件的，比如像**null**，呵呵。但是我找了半天也没找到错误所在，然后也不太影响正常使用，就暂时放在一边了。

结果今天出问题了，之前获取的是电影的列表信息，今天要写电影的详情页信息。写好后不断测试，结果有一条电影，一点进入详情就会崩溃。好了，设置好断点，一步步排查，最后NSLog导演头像的urlString时，找到了问题所在，导演的头像图片的url为**null**！终于找到了问题所在。那么可以解决了...

我获取了该条目电影的id号，然后在浏览器中获取了他的json数据，果不其然，就是这个电影!《火云端》!

![含有null](https://p3.ssl.qhimg.com/t01e3e13585b1cbbbe5.jpg)

然而我并没有想到有这么多**null**。

经过百度，好像要写文件，NSDictionary里面的object必须是NSString，NSData，NSNumber，NSDate，NSArray，NSDictionary中的数据类型。不过我知道，含有**null**是万万不能的，混蛋。

## 解决方案：NSDictionary -> NSDate

NSDictionary写文件之前，可以把它转换成NSData类型的数据，再执行写文件操作。
从文件读取时把读取出的NSData转换为NSDictionary就可以了。

### 用到的方法

NSDictionary -> NSData:

	NSData *data = [NSKeyedArchiver archivedDataWithRootObject:dictionary];
	
NSData -> NSDictionary:

	NSDictionary *dictionary = (NSDictionary *)[NSKeyedUnarchiver unarchiveObjectWithData:data];
	
	
昨天的问题解决了，这样写时候，每一次都能写入缓存文件成功。

今天的涉及到NSDictionary值的获取的部分，我都加了一个判断 `if([theValue isKindOfClass:[NSNull class]])`，然后进行相应处理就可以了。

这个时iOS的一个归档方法，不仅仅能归档`null`，`自定义的类型`也是可以的。具体可以参见[小白猪jianjian的博客-使用NSKeyedArchiver归档](http://www.cnblogs.com/xiaobaizhu/p/4011332.html)

<!--more-->

## 感想

生活中或者程序开发中，遇到了问题不一定非要立马解决，如果死钻牛角尖儿可能一直都解决不了，那时只能浪费时间。如果暂时放放，休息一会，再回来想解决办法可能就能想到解决方案。就像今天我解决我的问题一样，完全是机缘巧合。同时我们也应该发现，如果想做出好的软件，前期开发的投入很重要，但是后期的测试也很重要。功能做完之后我也是点击测试了不说几百次，几十次也是有了才发现了这个问题。

最后给大家看下我做的**电影影讯**，嘿嘿，使用了[豆瓣电影Api](http://developers.douban.com/wiki/?title=api_v2)。

![威海影讯列表](https://p3.ssl.qhimg.com/t010921a0625be5aab8.gif)
![电影详情页](https://p4.ssl.qhimg.com/t018dd1c22d5ce572a7.gif)





