---
title: 纯代码高仿网易新闻客户端两个scrollView联动（二）：实现界面逻辑变换
date: 2015-11-08 11:37:09
tags: [iOS]
---
上一篇已经实现了页面的布局，这一篇我们来实现界面的逻辑变换。主要用到的是scollView的两个代理方法。我们先看一下效果。

![效果预览图](https://p2.ssl.qhimg.com/t015616f03de7b014b5.gif)


<!--more-->

## 1.实现点击顶部按钮切换底部详情页面

### 给顶部items的按钮添加target
在`- (void) addFirstScrollViewOnView `方法中的`for循环`中，为每个button添加一个target，用于监听点击按钮的事件。
	
	  [itemButton addTarget:self action:@selector(itemButtonClicked:) forControlEvents:UIControlEventTouchUpInside];
	  
### 添加监听方法
我们应该知道当点击顶部item按钮时，底部详情页应该做相应的切换，同时按钮样式也改变。
大家可能注意到了需要用到按钮的tag值，所以我们同样在`上一步的方法`的`for循环`中为每个按钮设置tag值。  
	
	itemButton.tag = 100 + i;//设置button的tag值 
	
同时，需要添加一个实例变量：  
	 
	 //记录当前被点击的按钮tag
	@property (nonatomic, assign) NSInteger currentButtonTag;

**(重要)**并在`上一步方法的最后`为它设置初始值为`100`，也就是当前点击的按钮为第一个。

#### tips：不从1开始设置tag值的原因是因为前100可能有系统的控件占用

scrollView详情页切换：获取当前按钮的tag值，计算偏移量


	- (void) itemButtonClicked:(UIButton *)button {
		//**1.偏移底部详情scrollView
	    NSInteger buttonTag = button.tag - 100;//获取点击按钮的tag
	    self.secondScrollView.contentOffset = CGPointMake(buttonTag * KSCREEN_WIDTH, 0);//设置底部scrollview的内容偏移量
	    
	    //**2.恢复前一个被点击的按钮的样式
	    UIButton *preClickedButton = (UIButton *)[self.view viewWithTag:self.currentButtonTag];
	    preClickedButton.titleLabel.font = [UIFont systemFontOfSize:14.0f];
	    [preClickedButton setTitleColor:[UIColor colorWithRed:0.4 green:0.4 blue:0.4 alpha:1.0] forState:UIControlStateNormal];//标题颜色
	
	    //**3.设置当前点击按钮样式
	    button.titleLabel.font = [UIFont systemFontOfSize:18.0f];
	    [button setTitleColor:[UIColor colorWithRed:1.0 green:0.3 blue:0.3 alpha:1.0f] forState:UIControlStateNormal];
	
		//**4.改变当前点击按钮的tag值
	    self.currentButtonTag = buttonTag + 100;
	}

完成这一步，当你点击顶部按钮的时候，你就可以看到顶部按钮样式的改变和底部详情页面的切换。


## 2.实现滚动底部详情页顶部按钮字体切换

### 添加scrollView的代理 
添加代理
	
	@interface ViewController () <UIScrollViewDelegate>
	
在`- (void) addSecondScrollViewOnView `方法中设置`secondScrollView`的代理为self。这步不要忘记了。

	secondScrollView.delegate = self;
	
### 实现scrollView 的代理方法 - (void) scrollViewDidScroll:(UIScrollView *)scrollView；

在这一步，我们要实现滚动时，顶部items的按钮字体大小和颜色的改变。

		//正在滑动调用的代理方法
		- (void) scrollViewDidScroll:(UIScrollView *)scrollView {
		    //获取当前第二个scrollView的偏移量
		    CGFloat secondScrollViewContentOffsetX = scrollView.contentOffset.x;
		    //获取选中按钮的序号
		    int buttonTag = (secondScrollViewContentOffsetX) / KSCREEN_WIDTH;
		    //计算手指滑动了多少距离
		    CGFloat fingerDistance = secondScrollViewContentOffsetX - KSCREEN_WIDTH * buttonTag;
		    //获取到下一个按钮，并改变其字体大小和颜色，逐渐放大（根据手指滑动的距离动态改变）
		    UIButton *buttonNext = (UIButton *)[self.view viewWithTag:(buttonTag + 100 + 1)];
		    buttonNext.titleLabel.font = [UIFont systemFontOfSize:(14.0 + fingerDistance * 4 / (KSCREEN_WIDTH))];
		    [buttonNext setTitleColor:[UIColor colorWithRed:(0.4f + 3 * fingerDistance / (KSCREEN_WIDTH * 5)) green:0.3 blue:0.3 alpha:1.0] forState:UIControlStateNormal];
		    //同样方法获取到当前按钮，并改变其字体大小和颜色恢复回原来样式，逐渐缩小（根据手指滑动的距离动态改变）
		    UIButton *buttonCurr = (UIButton *)[self.view viewWithTag:(buttonTag + 100)];
		    buttonCurr.titleLabel.font = [UIFont systemFontOfSize:(18.0 - fingerDistance * 4 / (KSCREEN_WIDTH))];
		    [buttonCurr setTitleColor:[UIColor colorWithRed:(1.0f - 3 * fingerDistance / (KSCREEN_WIDTH * 5)) green:0.3 blue:0.3 alpha:1.0] forState:UIControlStateNormal];
		}

这里可以看到下一个按钮的字体动态改变，计算方法其实非常简单。
设手指滑动距离为x，字体大小为y。那么我们有两个值 *(0,14)*, *(KSCREEN_WIDTH, 18)*。
然后列二元一次方程组，就可以得到 *a* 和 *b*。
我这里解得 *a = 4/KSCREEN_WIDTD* , *b = 14*

当前按钮的改变亦然。只是两个点变为 *(KSCREEN_WIDTH, 14)*, *(0, 18)*
计算得到 *a = －(4/KSCREEN_WIDTH)*, *b = 18*

颜色大家计算方式相似，不再赘述。


## 3.使顶部scrollView随底部scrollView滑动而滚动

我们注意看网易新闻客户端，它底部scrollView滚动之后，顶部的scrollView也会随之滚动，并且除了开头或者末尾的几个按钮，它当前所在的新闻类型始终在屏幕中间。并且顶部的滚动总是在底部滑动结束之后，所以我们实现scrollView的代理方法	`- (void) scrollViewDidEndDecelerating:(UIScrollView *)scrollView `

	//滑动结束调用的代理方法
	- (void) scrollViewDidEndDecelerating:(UIScrollView *)scrollView {
	    //**1.获取选中详情页对应的顶部button
	    //移动顶部选中的按钮
	    CGFloat secondScrollViewContentOffsetX = scrollView.contentOffset.x;
	    //获取选中按钮的序号
	    int buttonTag = (secondScrollViewContentOffsetX) / KSCREEN_WIDTH;
	    //根据按钮号码获取到顶部的按钮
	    UIButton *buttonCurr = (UIButton *)[self.view viewWithTag:(buttonTag + 100)];
	    //**2.(重要)设置当前选中的按钮号。如若不写，将导致滑动后再点击顶部按钮，上一个按钮颜色，字体不会改变
	    self.currentButtonTag = buttonTag + 100;
	    
	    //**3.始终保持顶部选中按钮在中间位置
	    //注意一：开始的几个按钮，和末尾的几个按钮并不需要一直保持中间。
	    //注意二：对于已经放置在firstScrollView中的按钮，它的center是相对于scrollView的content而言的，注意并不是相对于self.view的bounds而言的。也就是说，放置好按钮，它的center就不会再改变
	    
	    //如果是顶部scrollView即将到末尾的几个按钮，设置偏移量，直接return
	    if (buttonCurr.center.x + KSCREEN_WIDTH * 0.5 > self.firstScrollView.contentSize.width) {
	        [UIView animateWithDuration:0.3 animations:^{
	            self.firstScrollView.contentOffset = CGPointMake(self.firstScrollView.contentSize.width - KSCREEN_WIDTH, 0);
	          }];
	        return;
	    }
	    //如果是顶部scrollView开头的几个按钮，设置偏移量，直接return
	    if (buttonCurr.center.x < KSCREEN_WIDTH * 0.5) {
	        [UIView animateWithDuration:0.3 animations:^{
	            self.firstScrollView.contentOffset = CGPointMake(0, 0);
	        }];
	        return;
	    }
	    
	    //如果是中间几个按钮的情况
	    if (buttonCurr.center.x > (KSCREEN_WIDTH * 0.5)) {
	        [UIView animateWithDuration:0.3 animations:^{
	            self.firstScrollView.contentOffset = CGPointMake(buttonCurr.center.x - self.view.center.x, 0);
	        }];
	    }
	
	}

这里需要注意的是我注释里面的**第2步**和**第3步**。
第2步一定要写，否则导致先滑动底部scrollView，再点击顶部scrollView的button，出现之前的那个button样式不恢复的情况。
第3步，注意前几个按钮和后几个按钮位置的判断，如果一味保持按钮在中间，就会出现顶部的offset过多，而后面出现空白，大家可以简单尝试一下。

整个效果就已经完成了～☺️



