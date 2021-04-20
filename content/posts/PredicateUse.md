---
title: NSPredicate(谓词) 的使用
date: 2016-03-22 17:41:22
tags: [iOS]
---

> 谓词（NSPredicate）提供在Cocoa中指定查询的普通解决方案。NSPredicate类用于定义逻辑条件以限制或筛选获取结果。

### NSPredicate 的基本使用

### 定义：
	
	NSPredicate *predict = [NSPredicate predicateWithFormat: @"SELF CONTAINS[cd] %@", SOMESTRING];
	
### 常用方法：

| 关键字 | 效果 |
|:---|:----------|
|**比较运算符**|
|>,<,==,>=,<=,!=|左侧满足比较运算符|
|**字符相关**|
| BEGINSWITH|左侧表达式 以 右侧表达式 开始|
| ENDSWITH| 左侧表达式 以 右侧表达式 结束|
| CONTAINS| 左侧表达式包含右侧表达式|
|**范围相关**|
| BETWEEN|左侧表达式在右侧表达式范围内|
| IN|左侧表达式在右侧表达式内|
|**正则表达式**|
| MATCHES|左侧表达式满足右侧表达式，右侧为正则表达式|
|**通配符**|
| LIKE|左侧表达式等于右侧表达式: 允许将 **?** 和 **\*** 用作通配符,其中 **?** 匹配一个字符,而 **\*** 匹配零个或多个字符|

### 字符串

	NSString *originStr = @"I like shangha!";
    NSString *str = @"shanghai";
    NSPredicate *predict = [NSPredicate predicateWithFormat: @"SELF CONTAINS %@", str];
    if ([predict evaluateWithObject:originStr]) {
        NSLog(@"含有字符串%@", str);
    }
    

其中 `CONTAINS` 也可以替换为 `BEGINSWITH`,`ENDSWITH`。
在它们后面，可以在方括号中添加`c`、`d` 或 `cd`, 如`CONTAINS[cd]`,其中`c`代表不区分大小写, `d`代表不区分音调符号。

### 数组和字典

谓词常用于数组的筛选:
	
	 NSArray *mArray = [NSArray arrayWithObjects:@"beijing", @"shanghai", @"shenzhen", @"guangzhou", nil];
	 
    //IN 用法
    NSPredicate *predicate1 = [NSPredicate predicateWithFormat:@"SELF IN {'shanghai', 'guangzhou'}"];
    NSArray *filteredArray1 = [mArray filteredArrayUsingPredicate: predicate1];
    NSLog(@"筛选后数组为:%@", filteredArray1);
    
    //LIKE 用法
    NSPredicate *predicate2 = [NSPredicate predicateWithFormat:@"SELF LIKE '*h*'"];
    NSArray *filteredArray2 = [mArray filteredArrayUsingPredicate: predicate2];
    NSLog(@"筛选后数组为:%@", filteredArray2);
	
	
	
谓词的条件也可以通过字典的占位符实现:

	NSPredicate *p1 = [NSPredicate predicateWithFormat:@"name==$NAME AND price==$PRICE"];
    NSDictionary *dic = [NSDictionary dictionaryWithObjectsAndKeys:@"name5", @"NAME", @"5000", @"PRICE", nil];
    NSPredicate *p2 = [p1 predicateWithSubstitutionVariables:dic];
    
    //表示从cars数组中筛选满足 name='name5' AND price='5000' 条件的元素
    NSArray *filteredCars = [cars filteredArrayUsingPredicate:p2];


	

	