---
title: UISearchBar和UISearchDisplayController实现搜索栏
date: 2016-03-19 17:36:43
tags: [iOS, UITableView, 搜索栏]
---

要实现tableview的搜索栏，实现方法有两种：第一种是UISearchBar和UIDisplayController结合起来实现，另一种是通过UISearchController实现。这里只介绍第一种：

注意: UISearchBar和UISearchDisplay只推荐iOS8.0之前使用。
![](https://p1.ssl.qhimg.com/t014836c37890cb46c7.jpg)

关于UISearchController的使用请跳转至：
[UITableView使用进阶(一):搜索栏](http://mungo.space/2016/03/19/UITableViewAdvanced01/)

<!--more-->
### 使用的协议
需要使用四个协议，分别是
	
<UITableViewDelegate, UITableViewDataSource, UISearchBarDelegate, UISearchDisplayDelegate>


### 头文件 ViewController.h

	#import <UIKit/UIKit.h>

	@interface ViewController : UIViewController

	@property (nonatomic, strong) UITableView *mTableView;
	//搜索结果
	@property (nonatomic, strong) NSArray *filterData;
	//全部数据
	@property (nonatomic, strong) NSMutableArray *allData;
	@property (nonatomic, retain) UISearchDisplayController *searchDisplayController;

	@end

### 实现文件 ViewController.m

	#import "ViewController.h"

	@interface ViewController () <UITableViewDelegate, UITableViewDataSource, UISearchBarDelegate, UISearchDisplayDelegate>
	
	@end
	
	@implementation ViewController
	
	- (void)viewDidLoad {
	    [super viewDidLoad];
	    
	    //初始化数据
	    int count = 100;
	    self.allData = [NSMutableArray arrayWithCapacity:count];
	    for (int i = 0; i < count; i ++) {
	        [self.allData addObject:[NSString stringWithFormat:@"第%d行", i]];
	    }
	    
	    
	    //定义tableview
	    CGRect appFrame = [[UIScreen mainScreen] bounds];
	    self.mTableView = [[UITableView alloc] initWithFrame:CGRectMake(0, 20, appFrame.size.width, appFrame.size.height - 20) style:UITableViewStylePlain];
	    self.mTableView.delegate = self;
	    self.mTableView.dataSource = self;
	    [self.view addSubview:self.mTableView];
	    
	    //定义UISearchBar
	    
	    UISearchBar *mySearchBar = [[UISearchBar alloc] init];
	    mySearchBar.delegate = self;
	    [mySearchBar setAutocapitalizationType:UITextAutocapitalizationTypeNone];
	    [mySearchBar sizeToFit];
	    self.mTableView.tableHeaderView = mySearchBar;
	    
	    //定义UISearchDisplayController
	    self.searchDisplayController = [[UISearchDisplayController alloc] initWithSearchBar:mySearchBar contentsController:self];
	    self.searchDisplayController.delegate = self;
	    [self.searchDisplayController setSearchResultsDataSource:self];
	    
	    [self setSearchDisplayController:self.searchDisplayController];
	    
	    
	}
	
	
	#pragma mark - UITableView Delegate
	- (NSInteger) numberOfSectionsInTableView:(UITableView *)tableView {
	    return 1;
	}
	
	- (NSInteger) tableView:(UITableView *)tableView numberOfRowsInSection:(NSInteger)section {
	    if (tableView == self.searchDisplayController.searchResultsTableView) {
	        return [self.filterData count];
	    } else {
	        return [self.allData count];
	    }
	}
	
	#pragma mark - UITableView DataSource
	- (UITableViewCell *) tableView:(UITableView *)tableView cellForRowAtIndexPath:(NSIndexPath *)indexPath {
	    static NSString *cellid = @"cellid";
	    UITableViewCell *cell = [tableView dequeueReusableCellWithIdentifier:cellid];
	    if (!cell) {
	        cell = [[UITableViewCell alloc] initWithStyle:UITableViewCellStyleDefault reuseIdentifier:cellid];
	    }
	    
	    if (tableView == self.searchDisplayController.searchResultsTableView) {
	        cell.textLabel.text = [self.filterData objectAtIndex:indexPath.row];
	    } else {
	        cell.textLabel.text = [self.allData objectAtIndex:indexPath.row];
	    }
	    
	    return cell;
	}
	
	#pragma mark - UISearchDisplayController 
	- (BOOL) searchDisplayController:(UISearchDisplayController *)controller shouldReloadTableForSearchString:(NSString *)searchString {
	    NSPredicate *predicate = [NSPredicate predicateWithFormat:@"SELF CONTAINS[cd] %@", searchString];
	    self.filterData = [self.allData filteredArrayUsingPredicate:predicate];
	    
	    return YES;
	}
	
	- (void)didReceiveMemoryWarning {
	    [super didReceiveMemoryWarning];
	    // Dispose of any resources that can be recreated.
	}
	
	@end


### 实现效果

![实现效果](https://p1.ssl.qhimg.com/t01136391b0c01f4d40.gif)




