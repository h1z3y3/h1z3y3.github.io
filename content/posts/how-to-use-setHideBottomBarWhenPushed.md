---
title: UITabBarController嵌套UINavigationController后，关于tabBar的问题以及解决方法
date: 2015-10-24 20:03:13
tags: [iOS]
---

一开始自己将tabBar全部自定义，将系统tabbar设置为`self.tabBar.hidden=YES`，隐藏系统的tabbar。但是自定义的tabbar在push出新页面的时候，tabbar并不会自动隐藏。之后在新页面`viewWillAppear`中设置自动隐藏，但是pop回之前页面（在`viewDidAppear`中设置tabbar显示）又出现不能及时出现，会有时间延迟，看上去特别不舒服。而且，边缘返回旧页面的时候不能及时显示tabbar也，所以打算用系统默认的tabbar。记录下**系统tabBar**样式的简单定义。

<!--more-->

## 简单自定义系统的tabBar

	- (void) initTabBarView {
    	NSArray *titles = @[@"工大威海", @"校园应用", @"校园生活", @"更多功能"];
    	//tabBarItem选中图片简单颜色变化时，自定义选中后的颜色
    	[self.tabBar setTintColor:[UIColor redColor]];
    	int i = 0;
    	for (UITabBarItem *item in self.tabBar.items) {
    	
    		//自定义tabBarItem的图片
        	item.image = [[UIImage imageNamed:
        	[NSString stringWithFormat:@"home_tab_icon_%d", i + 1]] imageWithRenderingMode:UIImageRenderingModeAutomatic];
        	//定义选中图片,上面imageWithRederingMode设置为Automatic，只是简单颜色变换，不需要设置,不再赘述
        	//item.selectedImage = ....
        	//设置tabBarItem的标题
        	item.title = titles[i];
        	i ++;
    	}
	}
	
这是运行后的样式：

![样式](https://p0.ssl.qhimg.com/t01dc370efcf696663c.jpg)


## push新页面tabBar自动隐藏，pop回显示tabBar

上面说如果用完全自定义的tabbar，我只能设置到等pop回的页面全部出现之后再让tabBar显示出来，这样十分不友好，我们如果实现手机微信类似这样的效果我们应该如何设置呢？注意红框部分

![微信边缘拖动返回时页面演示](https://p5.ssl.qhimg.com/t01b76fc70985541e13.jpg)

我们应该在**要push出新页面的那个页面**设置`viewDidAppear`和`viewWillDisappear`方法：

	- (void) viewDidAppear:(BOOL)animated {
	    [super viewWillAppear:animated];
	    [self setHidesBottomBarWhenPushed:YES];
	}

	- (void) viewWillDisappear:(BOOL)animated {
	    [self setHidesBottomBarWhenPushed:NO];
	    [super viewDidDisappear:animated];
	}

	
边缘拖动返回的时候，即可实现如下效果：

![我想要的效果](https://p1.ssl.qhimg.com/t01485a1ce59a7eb0e7.jpg)

完成~ 😏😏😏
