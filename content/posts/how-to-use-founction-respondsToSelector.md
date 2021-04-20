---
title: 函数respondsToSelector的使用
date: 2015-10-15 00:44:48
tags: [IOS]
---
-(BOOL)respondsToSelector: selector
//用来判断是否有以某个名字命名的方法（被封装在一个selector的对象里传递）

<!--more-->

---

	@protocol appViewDelegate <NSObject>
		@optional
		- (void) theButtonOnClicked:(UIButton *)button;
	@end
	
---

	if([self.delegate respondsToSelector:@selector(theButtonOnClicked:)]) {
		self.delegate theButtonOnClicked:sender;
	}
	
---

	#pragma mark - appViewDelegate
	
	- (void) theButtonOnClicked:(UIButton *)button {
		button.enabled = NO;
		[]button setTitle:@"已下载" ForState:UIControlStateDisabled];
	}
	
---

