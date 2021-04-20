---
title: UITableView使用简单进阶(一):搜索栏

date: 2016-03-19 18:39:48

tags: [iOS,UITableView,搜索栏]
---

UITableView 是开发中使用十分频繁的控件，本系列记录UITableView的进阶使用：UITableView的搜索栏和字母索引表。


不想看我废话的可以直接到gitHub仓库下载源码😏。
[UITableView使用进阶gitHub源码](https://github.com/Gitzhaoyang/iOSWithOC/tree/master/UITableViewAdvanced)

搜索栏有两种实现方式，第一种是通过UISearchBar和UISearchDisplayController实现，第二种是通过UISearchController实现。而在iOS8.0之后，苹果官方推荐使用第二种方式。

![使用UISearchController](https://p1.ssl.qhimg.com/t014836c37890cb46c7.jpg)

本文使用的是第二种方式（UISearchController），不过源码中也实现了第一种(UISearchBar+UISearchDisplayController)。关于UISearchBar和UISearchDisplayController的使用，可以参考我的另一篇文章：[UISearchBar和UISearchDisplayController实现搜索栏](http://mungo.space/2016/03/19/UISearchBar-UISearchDisplayController/)

<!--more-->
本文先介绍UITableView的搜索栏（UISearchController），是一个显示美国总统的TableView。

### Step1. 创建President类

创建一个名为President的Objective-C类，它继承于NSObject类，用于保存总统的姓氏和名字。定义一个静态方法，用来创建President对象，并且对firstName和lastName赋值。

#### President.h 

	#import <Foundation/Foundation.h>

	@interface President : NSObject
	@property (nonatomic, retain) NSString *firstName;
	@property (nonatomic, retain) NSString *lastName;
	+ (President *) presidentWithFirstName: (NSString *)firstname lastName: (NSString *)lastname;
	
	@end
	
#### President.m

	#import "President.h"

	@implementation President
	
	+ (President *) presidentWithFirstName:(NSString *)firstname lastName:(NSString *)lastname {
	    President *president = [[President alloc] init];
	    president.firstName = firstname;
	    president.lastName  = lastname;
	    
	    return president;
	}
	
	@end
	
	
### Step2. 编写ViewController

#### viewController.h

	#import <UIKit/UIKit.h>
	
	@interface ViewController : UIViewController
	
	@property (nonatomic, strong) UITableView *mTableView;
	@property (nonatomic, strong) NSArray *presidents;//所有总统
	@property (nonatomic, strong) NSArray *filteredPresidents;//保存搜索结果
	@property (nonatomic, retain) UISearchController *searchController;	
	@end

#### viewController.m

	#import "ViewController.h"
	#import "President.h"
	
	@interface ViewController () <UITableViewDataSource, UITableViewDelegate, UISearchBarDelegate, UISearchResultsUpdating, UISearchDisplayDelegate>
	
	@end
	
	@implementation ViewController
	
	- (void)viewDidLoad {
	    [super viewDidLoad];
	    
	    //初始化数据
	    self.presidents = [[NSArray alloc] initWithObjects:
	                       [President presidentWithFirstName:@"George" lastName:@"Washington"],
	                       [President presidentWithFirstName:@"John" lastName:@"Adams"],
	                       [President presidentWithFirstName:@"Thomas" lastName:@"Jeffeson"],
	                       [President presidentWithFirstName:@"James" lastName:@"Madison"],
	                       [President presidentWithFirstName:@"James" lastName:@"Monroe"],
	                       [President presidentWithFirstName:@"John Quincy" lastName:@"Adams"],
	                       [President presidentWithFirstName:@"Andrew" lastName:@"Jackson"],
	                       [President presidentWithFirstName:@"Martin" lastName:@"van Buren"],
	                       [President presidentWithFirstName:@"William Henry" lastName:@"Harrison"],
	                       [President presidentWithFirstName:@"John" lastName:@"Tyler"],
	                       [President presidentWithFirstName:@"James K" lastName:@"Polk"],
	                       [President presidentWithFirstName:@"Zachary" lastName:@"Taylor"],
	                       [President presidentWithFirstName:@"Millard" lastName:@"Fillmore"],
	                       [President presidentWithFirstName:@"Franklin" lastName:@"Pierce"],
	                       [President presidentWithFirstName:@"James" lastName:@"Buchanan"],
	                       [President presidentWithFirstName:@"Abraham" lastName:@"Lincoln"],
	                       [President presidentWithFirstName:@"Andrew" lastName:@"Johnson"],
	                       [President presidentWithFirstName:@"Ulysses S" lastName:@"Grant"],
	                       [President presidentWithFirstName:@"Rutherford B" lastName:@"Hayes"],
	                       [President presidentWithFirstName:@"James A" lastName:@"Garfield"],
	                       [President presidentWithFirstName:@"Chester A" lastName:@"Arthur"],
	                       [President presidentWithFirstName:@"Grover" lastName:@"Cleveland"],
	                       [President presidentWithFirstName:@"Bejamin" lastName:@"Harrison"],
	                       [President presidentWithFirstName:@"Grover" lastName:@"Cleveland"],
	                       [President presidentWithFirstName:@"William" lastName:@"McKinley"],
	                       [President presidentWithFirstName:@"Theodore" lastName:@"Roosevelt"],
	                       [President presidentWithFirstName:@"William Howard" lastName:@"Taft"],
	                       [President presidentWithFirstName:@"Woodrow" lastName:@"Wilson"],
	                       [President presidentWithFirstName:@"Warren G" lastName:@"Harding"],
	                       [President presidentWithFirstName:@"Calvin" lastName:@"Coolidge"],
	                       [President presidentWithFirstName:@"Herbert" lastName:@"Hoover"],
	                       [President presidentWithFirstName:@"Franklin D" lastName:@"Roosevelt"],
	                       [President presidentWithFirstName:@"Harry S" lastName:@"Truman"],
	                       [President presidentWithFirstName:@"Dwight D" lastName:@"Eisenhower"],
	                       [President presidentWithFirstName:@"John F" lastName:@"Kennedy"],
	                       [President presidentWithFirstName:@"Lyndon B" lastName:@"Johnson"],
	                       [President presidentWithFirstName:@"Richard" lastName:@"Nixon"],
	                       [President presidentWithFirstName:@"Gerald" lastName:@"Ford"],
	                       [President presidentWithFirstName:@"Jimmy" lastName:@"Carter"],
	                       [President presidentWithFirstName:@"Ronald" lastName:@"Reagan"],
	                       [President presidentWithFirstName:@"George H W" lastName:@"Bush"],
	                       [President presidentWithFirstName:@"Bill" lastName:@"Clinton"],
	                       [President presidentWithFirstName:@"George W" lastName:@"Bush"],
	                       [President presidentWithFirstName:@"Barack" lastName:@"Obama"],
	                       nil];
	    
	    CGRect appFrame = [[UIScreen mainScreen] bounds];
	    //创建tableview
	    self.mTableView = [[UITableView alloc] initWithFrame:CGRectMake(0, 20, appFrame.size.width, appFrame.size.height - 20) style:UITableViewStylePlain];
	    self.mTableView.delegate = self;
	    self.mTableView.dataSource = self;
	    [self.view addSubview:self.mTableView];
	 
	    //创建searchController
	    self.searchController = [[UISearchController alloc] initWithSearchResultsController:nil];
	    self.searchController.searchResultsUpdater = self;
	    self.searchController.dimsBackgroundDuringPresentation = NO;
	    self.searchController.hidesNavigationBarDuringPresentation = NO;
	    //设置tableview的搜索栏
	    self.mTableView.tableHeaderView = self.searchController.searchBar;
	    
	}
	
	#pragma mark - tableView Delegate
	
	- (NSInteger)tableView:(UITableView *)tableView numberOfRowsInSection:(NSInteger)section {
	    if (self.searchController.active) {
	        return [self.filteredPresidents count];
	    } else {
	        return [self.presidents count];
	    }
	}
	
	#pragma mark - tableView Datasource
	- (UITableViewCell *)tableView:(UITableView *)tableView cellForRowAtIndexPath:(NSIndexPath *)indexPath {
	    static NSString *cellId = @"cellId";
	    UITableViewCell *cell = [tableView dequeueReusableCellWithIdentifier:cellId];
	    
	    if (!cell) {
	        cell = [[UITableViewCell alloc] initWithStyle:UITableViewCellStyleDefault reuseIdentifier:cellId];
	    }
	    
	    President *president;
	    if (self.searchController.active) {
	        president = [self.filteredPresidents objectAtIndex:indexPath.row];
	    } else {
	        president = [self.presidents objectAtIndex:indexPath.row];
	    }
	    cell.textLabel.text = [NSString stringWithFormat:@"%@ %@", president.firstName, president.lastName];
	    
	    return cell;
	    
	}
	
	#pragma mark - SearchController delegate
	- (void)updateSearchResultsForSearchController:(UISearchController *)searchController {
	    NSString *searchString = [self.searchController.searchBar text];
	    NSPredicate *predicate = [NSPredicate predicateWithFormat:@"firstName CONTAINS[cd] %@ OR lastName CONTAINS[cd] %@", searchString, searchString];
	    self.filteredPresidents = [self.presidents filteredArrayUsingPredicate:predicate];
	    
	    //刷新表格
	    [self.mTableView reloadData];
	}
	
	- (void)didReceiveMemoryWarning {
	    [super didReceiveMemoryWarning];
	    // Dispose of any resources that can be recreated.
	}

	@end

### 实现效果

![实现效果](https://p5.ssl.qhimg.com/t011ce6e41ca1d6164b.gif)
