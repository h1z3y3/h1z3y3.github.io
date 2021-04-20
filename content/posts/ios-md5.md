---
title: iOS NSString字符串MD5加密
date: 2016-03-13 18:59:02
tags: [iOS, 加密, MD5]
---
为了使保存的密码更安全， 我们应该实现一个NSString的分类，为密码创建一个MD5的哈希值，而且并把它保存在keychain中；keychain是在设备中保存关键数据的唯一安全的地方。

### step1.  新建文件

新建文件，在模板列表中选择Objective-C Category项。单击Next按钮，输入“MD5”作为这个创建的分类的名字，然后在CategoryOn下拉菜单中选择NSString，表明是NSString的Category。

### step2.  头文件

	#import <Foundation/Foundation.h>
	
	@interface NSString (MD5)
	
	- (NSString *) MD5;
	
	@end
		

### step3.  源文件

	#import "NSString+MD5.h"
	#import <CommonCrypto/CommonDigest.h>
	
	@implementation NSString (MD5) 
	
	- (NSString *) MD5{
		
		//转化为UTF8格式字符串
		const char *ptr = [self UTF8String];
		
		//开辟一个16字节数组
		unsigned char md5Buffer[CC_MD5_DIGEST_LENGTH];
		
		//调用官方封装的加密方法, 将ptr开始的字符串存储到md5Buffer[]中
		CC_MD5(ptr, strlen(ptr), md5Buffer);
		
		//转换位NSString并返回
		NSMutableString *output = [NSMutableString stringWithCapacity: CC_MD5_DIGEST_LENGTH * 2];
		for (int i = 0; i < CC_MD5_DIGEST_LENGTH; i ++) {
			[output appendFormat: @"%02x", md5Buffer[i]];
		}
		
		/*
		 * `x`表示十六进制，%02 意思是不足两位用0补齐，如果多于2位则不影响
		 * NSLog("%02x", 0x666);   //输出 `666`
		 * NSLog("%02x", 0x3);		//输出 `03` 
		 */
		
		return output;
	
	}
	
	@end