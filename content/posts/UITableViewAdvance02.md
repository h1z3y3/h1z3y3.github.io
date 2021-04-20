---
title: UITableView使用简单进阶(二):索引条
date: 2016-03-19 21:32:22
tags: [iOS, UITableView]
---

前文已经介绍了如何给UITableView添加搜索栏，这次要给UITableView进一步添加索引条。
基本思路：

1. 获取总统名字的首字母组成一个索引字母表储存在数组中；
2. 修改TableView的代理方法实现section的显示，section的数量应为索引字母表的元素的个数；
3. 把索引条添加到TableView 中，用TableView的代理方法即可实现。

<!--more-->


由于接着上文，所以本文中代码将不会全部展示，代码的重复部分将用省略号代替。

### Step1. 添加属性

在ViewController.h中添加属性：NSMutableArray *alphabetArray;
	
	#import <UIKit/UIKit.h>

	@interface ViewController : UIViewController
	
	...	
	
	@property (nonatomic, strong) NSMutableArray *alphabetArray;
	
	...			
	
	@end
	
### Step2. 修改ViewController.m
	
	//
	//  ViewController.m
	//  UITableViewAdvanced01
	//
	//  Created by mungo on 19/03/16.
	//  Copyright © 2016 mungo. All rights reserved.
	//
	
	#import "ViewController.h"
	#import "President.h"
	
	@interface ViewController () <UITableViewDataSource, UITableViewDelegate, UISearchBarDelegate, UISearchResultsUpdating, UISearchDisplayDelegate>
	
	@end
	
	@implementation ViewController
	
	- (void)viewDidLoad {
	    [super viewDidLoad];
	    
	    //初始化数据
	    ...
	    
	    //创建tableview
	   	...
	 
	    //创建searchController
	  	...
	  	
	    //设置tableview的搜索栏
		...
	    
	    //设置字母表
	    self.alphabetArray = [self getAlphetSortedArray];
	    
	}
	
	/**
	 * 新添加方法：
	 * 获取字母表
	 * @return MSMutableArray* 已经排序的字母表数组
	 */
	- (NSMutableArray *) getAlphetSortedArray {
	    self.alphabetArray = [[NSMutableArray alloc] init];
	    for (int i = 0; i < [self.presidents count]; i ++) {
	        //获取名字的第一个字母
	        President *president = [self.presidents objectAtIndex:i];
	        char letter = [president.firstName characterAtIndex:0];
	        NSString *uniqueChar = [NSString stringWithFormat:@"%c", letter];
	        //将字母加入到字母表中
	        if (![self.alphabetArray containsObject:uniqueChar]) {
	            [self.alphabetArray addObject:uniqueChar];
	        }
	    }
	    
	    //对字母表进行排序
	    NSSortDescriptor *sortDescriptor = [[NSSortDescriptor alloc] initWithKey:@"self" ascending:YES];
	    NSArray *sortDescirptors = [NSArray arrayWithObject:sortDescriptor];
	    NSArray *sortedArray;
	    sortedArray = [self.alphabetArray sortedArrayUsingDescriptors:sortDescirptors];
	    NSMutableArray *alphabetArray = [[NSMutableArray alloc] initWithArray:sortedArray copyItems:YES];
	    
	    return alphabetArray;
	}
	
	#pragma mark - tableView Delegate
	
	- (NSInteger)tableView:(UITableView *)tableView numberOfRowsInSection:(NSInteger)section {
	    if (self.searchController.active) {
	        ...
	    } else {
	      	//根据section筛选以当前字母开头的总统数组
      		NSArray *tmpArray = [self getAlphabetArrayWithIndex:section];
	        return [tmpArray count];
	    }
	}
	

	#pragma mark - Indexing UITableView
	//索引条的字母表
	- (NSArray<NSString *> *)sectionIndexTitlesForTableView:(UITableView *)tableView {
	    return self.alphabetArray;
	}
	
	- (NSInteger) tableView:(UITableView *)tableView sectionForSectionIndexTitle:(NSString *)title atIndex:(NSInteger)index {
	    return index;
	}
	
	
	#pragma mark - tableView Datasource
	
	- (NSInteger) numberOfSectionsInTableView:(UITableView *)tableView {
	    if (self.searchController.active) {
	        self.alphabetArray = nil;//搜索时不显示section
	        return 1;
	    } else {
	        self.alphabetArray = [self getAlphetSortedArray];//停止搜索恢复section显示
	        return [self.alphabetArray count];
	    }
	    
	}
	
	//tableView section的标题
	- (NSString *) tableView:(UITableView *)tableView titleForHeaderInSection:(NSInteger)section {
	    return [self.alphabetArray objectAtIndex: section];
	}
	
	- (UITableViewCell *)tableView:(UITableView *)tableView cellForRowAtIndexPath:(NSIndexPath *)indexPath {
	   
	   	...
	   	    
	    President *president;
	    if (self.searchController.active) {
	       	...
	    } else {
	        //根据section筛选以当前字母开头的总统数组
	        NSArray *tmpArray = [self getAlphabetArrayWithIndex:indexPath.section];
	        if ([tmpArray count]) {
	            president = [tmpArray objectAtIndex:indexPath.row];
	        }
	        
	    }
	  
	    if (president) {
	        cell.textLabel.text = [NSString stringWithFormat:@"%@ %@", president.firstName, president.lastName];
	    }
	   
	    return cell;
	    
	}
	
	/*
	 * 新添加方法：
	 * 根据首字母对所有总统进行筛选
	 * @return NSArray* 于当前首字母相同的总统数组
	 */
	
	- (NSArray *) getAlphabetArrayWithIndex:(NSInteger)index{
	    
	    NSString *alphabet = [self.alphabetArray objectAtIndex:index];
	    NSPredicate *president = [NSPredicate predicateWithFormat:@"firstName BEGINSWITH [cd] %@", alphabet];
	    NSArray *tmpArray = [self.presidents filteredArrayUsingPredicate:president];
	    
	    return tmpArray;
	}
	
	#pragma mark - SearchController delegate
	- (void)updateSearchResultsForSearchController:(UISearchController *)searchController {
	    ...
	    
	}
	
	@end


#### 注意：

numberOfSectionsInTableView中， 开始搜索的时候要将TableView的section数量设置为1，并且要把字母表设置为空；不搜索时要恢复section的数量。

### 实现效果

![实现效果](https://p4.ssl.qhimg.com/t0162f8e32897367e8f.gif)



