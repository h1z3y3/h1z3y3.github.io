---
title: 根据字数自适应高度的UILabel
date: 2016-03-17 15:36:59
tags: [iOS, UILabel]
---

	- (UILabel *) getLabelHeightFixedWithText: (NSString *) text {
    
    	UIFont *font = [UIFont boldSystemFontOfSize: 12.0f];
    	int width = 225, height = 10000;
	    NSMutableDictionary *attrs = [[NSMutableDictionary alloc] init];
    	[attrs setObject: font forKey: NSFontAttributeName];
	    CGRect size = [text boundingRectWithSize:CGSizeMake(width, height) options: NSStringDrawingUsesLineFragmentOrigin attributes: attrs context: nil];
    	UILabel *label = [[UILabel alloc] initWithFrame: CGRectMake(0, 0, size.size.width, size.size.height)];
    	label.numberOfLines = 0;//一定要设置行数为0
    	label.font = font;
    	label.lineBreakMode = NSLineBreakByWordWrapping;
    	label.text = (text ? text : @"");
    	label.backgroundColor = [UIColor clearColor];
    	label.textColor = [UIColor blackColor];
    
    	return label;
    
	}
