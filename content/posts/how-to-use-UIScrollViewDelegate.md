---
title: UIScrollView的基本属性设置和常用代理方法
date: 2015-10-17 17:54:31
tags: [iOS]
---

罗列了UIscrollView的一些基本设置和常用方法

<!--more-->


	- (void)viewDidLoad {
    	[super viewDidLoad];
		//    [self initScrollView];
		//    [self initPageCtrl];
		//    [self addTimer];
	    self.scrollView = [[UIScrollView alloc]initWithFrame:CGRectMake(10, 20, 320, 460)];
	    self.scrollView.backgroundColor = [UIColor blueColor];
	    //是否支持滑动到最顶端
	    self.scrollView.scrollsToTop = NO;
	    //设置UIScrollView的代理
	    self.scrollView.delegate = self;
	    //设置内容大小
	    self.scrollView.contentSize = CGSizeMake(320 * 10, 460);
	    //是否反弹
	    self.scrollView.bounces = NO;
	    //是否滚动
	    self.scrollView.scrollEnabled = YES;
	    //是否分页
	    self.scrollView.pagingEnabled = NO;
	    //设置indecator风格
	    self.scrollView.indicatorStyle = UIScrollViewIndicatorStyleWhite;
	    //设置内容边缘和Indicator边缘
	    self.scrollView.contentInset = UIEdgeInsetsMake(0, 50, 50, 0);
	    self.scrollView.scrollIndicatorInsets = UIEdgeInsetsMake(0, 50, 0, 0);
	    //是否同时运动
	    self.scrollView.directionalLockEnabled = NO;
	    
	    [self.view addSubview:self.scrollView];
	}

	//是否支持滑动到顶部
	- (BOOL) scrollViewShouldScrollToTop:(UIScrollView *)scrollView {
	    return YES;
	}
	
	//滑动到顶部时调用该方法
	- (void) scrollViewDidScrollToTop:(UIScrollView *)scrollView {
	    NSLog(@"滑动到顶部了");
	}
	
	//已经滑动(正在滑动也会调用该方法)
	- (void) scrollViewDidScroll:(UIScrollView *)scrollView {
	    NSLog(@"已经开始滑动");
	}
	
	//开始拖动
	- (void) scrollViewWillBeginDragging:(UIScrollView *)scrollView {
	    NSLog(@"开始拖动");
	}
	
	//结束拖动
	- (void) scrollViewDidEndDragging:(UIScrollView *)scrollView willDecelerate:(BOOL)decelerate {
	    NSLog(@"结束拖动");
	}
	
	//开始减速
	- (void) scrollViewWillBeginDecelerating:(UIScrollView *)scrollView {
	    NSLog(@"开始减速");
	}
	
	//减速停止
	- (void) scrollViewDidEndDecelerating:(UIScrollView *)scrollView {
	    NSLog(@"减速停止");
	}
	
	//结束滚动动画
	- (void) scrollViewDidEndScrollingAnimation:(UIScrollView *)scrollView {
	    NSLog(@"结束滚动动画");
	}
	
	- (void)didReceiveMemoryWarning {
	    [super didReceiveMemoryWarning];
	    // Dispose of any resources that can be recreated.
	}