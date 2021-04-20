---
title: tableView 两种重用cell的方法区别
date: 2015-10-18 16:21:04
tags: [iOS]
---

今天在学习iOS开发的时候，偶然发现tableView有两种重用cell的方法。先整理一下方便日后查阅。  

第一种：`[tableView dequeueReusableCellWithIdentifier:identifier]` （SDK 6.0之前）   
第二种：`[tableView dequeueReusableCellWithIdentifier:identifier forIndexPath:indexPath]]`  （SDK 6.0之后）

<!--more-->

区别：
第一种：必须判断cell是否定义，未定义则手动创建，代码如下：
	
	const NSString *identifier = @"cell";
	UITableView *cell = [tableView dequeueReusableCellWithIndetifier:identifier];
	if(!cell) {
		cell = [[UITableView alloc] initWithStyle:UITableViewCellStyleDefault];
	}
	/*
		设置cell
		cell.textLable.text = @"第...行";
	*/
	return cell;


第二种是SDK 6.0开始新添加的方法。用于你**已经用Nib定义了一个Cell**，然后就可以省下上面那些代码，只用一行就可以解决：

	 UITableViewCell *cell = [tableView dequeueReusableCellWithIdentifier:@"cell" forIndexPath:indexPath];
	  /*
		设置cell
		cell.textLable.text = @"第...行";
	  */

所以如果有这个错误：

	reason: 'unable to dequeue a cell with identifier friendCell - must register a nib or a class for the identifier or connect a prototype cell in a storyboard'
	
应该想一下自己**是不是为cell创建的Nib**。
