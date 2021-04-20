---
title: stretchableImageWithLeftCapWith:(NSInteger)leftCapWidth topCapHeight:(NSInteger)topCapHeight

date: 2016-03-19 14:35:39
tags: [iOS]
---

	- (UIImage *) stretchableImageWithLeftCapWidth:(NSInteger)leftCapWidth topCapHeight:(NSInteger)topCapHeight

这个函数是UIImage的一个类函数，它的功能是创建一个内容可拉伸，而边角不拉伸的图片。两个参数的含义分别为：不拉伸区域的宽度、不拉伸区域的高度。

根据设置的宽度和高度，在像素点`((leftCapWidth+1), (topCapWidth+1))`开始左右扩展和上下拉伸。

### 注意：

- 可拉伸范围是距离leftCapWidth+1的那一列像素和topCapHeight+1的那一横排像素。
如果设置参数为10和5，那么，图片左边10个像素和上边5个像素区域内不会被拉伸。而从(11, 5)开始扩展和拉伸。

- 只是对一个像素进行复制到一定宽度，而图像后面的剩余部分也不会被拉伸。

		//原始图片
    	UIImage *image = [UIImage imageNamed:@"yoububble.png"] ;
    	UIImageView *imageView1 = [[UIImageView alloc] initWithImage:image];
    	[imageView1 setFrame:CGRectMake(10, 60, 40, 40)];
    	[self.view addSubview:imageView1];
    
    	//拉伸后图片
    	UIImage *strechImage = [image stretchableImageWithLeftCapWidth:10 topCapHeight:10];
    	UIImageView *imageView2 = [[UIImageView alloc] initWithImage:strechImage];
    	[imageView2 setFrame:CGRectMake(10, 120, 300, 100)];
    	[self.view addSubview:imageView2];
    	
    	
    	
  
  
![图片效果](https://p3.ssl.qhimg.com/t012520479c8ccb0d67.png)

