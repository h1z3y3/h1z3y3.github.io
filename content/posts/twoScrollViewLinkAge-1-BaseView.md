---
title: 纯代码高仿网易新闻客户端两个scrollView联动（一）：设置基本的界面布局
date: 2015-11-07 23:31:46
tags: [iOS]
---

自己再开发app过程中遇到了这样那样的需求，其中有一项是新闻页面，需要两个scrollView联动，上面的scrollView是新闻类型，下面scrollView是tableView新闻标题。开发过程中我全部都是用代码布局的，因为自己是新手，不知道到底是用代码写比较方便还是用storyBoard更方便，但是感觉手写代码适应屏幕上更容易上手吧。需求实现之后现在拿出来简单整理一下，方便自己日后查看，也给后来者以参考。因为刚刚开始学，用到的都是些简单的方法，也可能会出错，如果有什么不足，请留言给我指出。谢谢～
<!--more-->
现在开始正题：
本文是第一篇，先是简单的页面布局，首先我们看一下布局之后的效果：

![twoScrollViewLinkage_baseViewLayout](https://p4.ssl.qhimg.com/t010500cb2fa4fe823f.gif)

### 1.宏定义

	//屏幕的宽和高
	#define KSCREEN_HEIGHT [UIScreen mainScreen].bounds.size.height
	#define KSCREEN_WIDTH [UIScreen mainScreen].bounds.size.width
	//最顶部状态栏的高度
	#define KSTATUS_HEIGHT 20
	//顶部items的scrollView的高度
	#define KFIRST_SCROLLVIEW_HEIGHT 30
	//根据计算，得到底部详情scrollView的高度
	#define KSECOND_SCROLLVIEW_HEIGHT (KSCREEN_HEIGHT - KSTATUS_HEIGHT - KFIRST_SCROLLVIEW_HEIGHT)
	//顶部scrollView每个item按钮的宽度
	#define KFIRST_SCROLLVIEW_ITEM_WIDTH 55

### 2.首先定义需要用到的变量

	@interface ViewController ()

	//顶部items的scrollView
	@property (nonatomic, weak) UIScrollView *firstScrollView;
	//底部详情scrollView
	@property (nonatomic, weak) UIScrollView *secondScrollView;
	//item类型的数组
	@property (nonatomic, strong) NSArray *itemsTitlesArray;
	//为了方便，定义一个包含颜色的NSArray（自行取舍）
	@property (nonatomic, strong) NSArray *colorArray;

	@end
	
### 3.懒加载

	#pragma mark - lazy load
	//定义items的标题
	- (NSArray *) itemsTitlesArray {
    	if (!_itemsTitlesArray) {
        
    	    _itemsTitlesArray = [NSArray arrayWithObjects:@"热点新闻", @"新闻快讯", @"通知公告", @"院系通知", @"人物风采", nil];
    	}
    
    	return _itemsTitlesArray;
	}
	//颜色数组
	- (NSArray *) colorArray {
    	if (!_colorArray) {
        	_colorArray = [NSArray arrayWithObjects:
                       [UIColor colorWithRed:240.0/255.0 green:89.0/255.0 blue:136.0/255.0 alpha:1.0],
                       [UIColor colorWithRed:0.0/255.0   green:179.0/255.0 blue:155.0/255.0 alpha:1.0],
                       [UIColor colorWithRed:244.0/255.0 green:131.0/255.0 blue:69.0/255.0 alpha:1.0],
                       [UIColor colorWithRed:241.0/255.0 green:90.0/255.0 blue:102.0/255.0 alpha:1.0],
                       [UIColor colorWithRed:0.0/255.0 green:179.0/255.0 blue:155.0/255.0 alpha:1.0],
                       [UIColor colorWithRed:255.0/255.0 green:223.0/255.0 blue:104.0/255.0 alpha:1.0],
                       nil];
    	}
    
    	return _colorArray;
	}

### 4.添加顶部items的scrollView

	//添加第一个scrollview
	- (void) addFirstScrollViewOnView {
	    //**1.设置顶部类型scrollView
	    UIScrollView *firstScrollView = [[UIScrollView alloc] initWithFrame:CGRectMake(0, KSTATUS_HEIGHT, KSCREEN_WIDTH, KFIRST_SCROLLVIEW_HEIGHT)];
	    firstScrollView.bounces = NO;//禁止反弹
	    firstScrollView.showsHorizontalScrollIndicator = NO;//禁止显示水平滚动条
	    //设置scrollView的内容大小，宽度为 宏定义顶部按钮的宽度 ＊ items数组的数量
	    firstScrollView.contentSize = CGSizeMake(self.itemsTitlesArray.count * KFIRST_SCROLLVIEW_ITEM_WIDTH, KFIRST_SCROLLVIEW_HEIGHT);
	    self.firstScrollView = firstScrollView;
	    [self.view addSubview:firstScrollView];//添加到self.view
	
	    //**2.为第一个scrollView添加buttons
	    for (int i = 0; i < self.itemsTitlesArray.count; i ++) {
	        UIButton *itemButton = [[UIButton alloc] initWithFrame:CGRectMake(i * KFIRST_SCROLLVIEW_ITEM_WIDTH, 0, KFIRST_SCROLLVIEW_ITEM_WIDTH, KFIRST_SCROLLVIEW_HEIGHT)];//注意左边距的写法
	        [itemButton setTitle:[self.itemsTitlesArray objectAtIndex:i] forState:UIControlStateNormal];//标题
	        itemButton.backgroundColor = [UIColor colorWithRed:0.97 green:0.97 blue:0.97 alpha:1.0];//背景颜色
	        [itemButton setTitleColor:[UIColor colorWithRed:0.4 green:0.4 blue:0.4 alpha:1.0] forState:UIControlStateNormal];//标题颜色
	        itemButton.titleLabel.font = [UIFont systemFontOfSize:14.0];
	        [firstScrollView addSubview:itemButton];//添加到第一个scrollView
	        
	        //定义第一个顶部item的按钮样式
	        if (i == 0) {
	            itemButton.titleLabel.font = [UIFont systemFontOfSize:18.0f];
	            [itemButton setTitleColor:[UIColor colorWithRed:1.0 green:0.3 blue:0.3 alpha:1.0f] forState:UIControlStateNormal];
	        }
	        
	    }
	    
	}
### 5.添加底部详情scrollView
	
	//添加第二个scrollview
	- (void) addSecondScrollViewOnView {
	    //**1.设置底部详情scrollView
	    UIScrollView *secondScrollView = [[UIScrollView alloc] initWithFrame:CGRectMake(0, KSTATUS_HEIGHT + KFIRST_SCROLLVIEW_HEIGHT, KSCREEN_WIDTH, KSECOND_SCROLLVIEW_HEIGHT)];
	    secondScrollView.pagingEnabled = YES;//分页
	    secondScrollView.bounces = NO;//禁止反弹
	    secondScrollView.delegate = self;
	    //设置内容大小，宽度为 屏幕的宽度 * items数组的数量
	    secondScrollView.contentSize = CGSizeMake(self.itemsTitlesArray.count * KSCREEN_WIDTH, KSECOND_SCROLLVIEW_HEIGHT);
	    self.secondScrollView = secondScrollView;
	    [self.view addSubview:secondScrollView];
	    //**2.添加Views
	    for (int i = 0; i < self.itemsTitlesArray.count; i ++) {
	        UIView *bottomView = [[UIView alloc] initWithFrame:CGRectMake(i * KSCREEN_WIDTH, 0, KSCREEN_WIDTH, KSECOND_SCROLLVIEW_HEIGHT)];
	        bottomView.backgroundColor = [self.colorArray objectAtIndex:(i % 6)];//颜色数组取余
	        //为了方便观察，放置一个label，如果需要其它控件自行替换（如若新闻，用tableView）
	        UILabel *flagLabel = [[UILabel alloc] initWithFrame:bottomView.bounds];
	        flagLabel.font = [UIFont systemFontOfSize:50.0f];
	        flagLabel.textColor = [UIColor whiteColor];
	        flagLabel.textAlignment = NSTextAlignmentCenter;
	        flagLabel.text = [NSString stringWithFormat:@"%@", [self.itemsTitlesArray objectAtIndex:i]];
	        [bottomView addSubview:flagLabel];
	        
	        [secondScrollView addSubview:bottomView];
	    }
	}

### 6.在-(void)viewDidLoad方法中调用两个添加scrollView的方法

	- (void)viewDidLoad {
    	[super viewDidLoad];
    	[self addFirstScrollViewOnView];//添加顶部items的scrollView
    	[self addSecondScrollViewOnView];//添加底部详情的scrollView
	}



### 7.声明

本文涉及的一切代码，都将上传到我的github仓库:[https://github.com/Gitzhaoyang/iOSWithOC/tree/master/twoScrollViewLinkage](https://github.com/Gitzhaoyang/iOSWithOC/tree/master/twoScrollViewLinkage)，请自行查看。





