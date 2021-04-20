---
title: Grep, Sed, Awk 日常使用
date: 2016-08-19 22:03:44
tags: [Linux]
---

## grep

### 概述

在给出文件列表或标准输入后, grep会对匹配一个或多个正则表达式的文本进行搜索, 并输出匹配（或者不匹配）的行或文本.

### 使用格式

	grep [options] PATTERN [FILE...]

### 常用选项

* -i   忽略字符大小写
* -v   显示未被模式匹配到的行或串
* -o   只显示匹配到的串而不是整行
* -n   显示匹配的行及行号
* -E   使用扩展的正则表达式
* -A n 显示出匹配到的行和后n行
* -B n 显示出匹配到的行和前n行
* -C n 显示出匹配到的行和前后各n行

**扩展的正则表达式**

扩展正则表达式与基础正则表达式的唯一区别在于: `() {} ? +` 这几个字符.

基础正则表达式中, `() {} ? +` 表示特殊含义，使用时需要将他们转义

而扩展正则表达式中, `() {} ? +` 不表示特殊含义, 你需要将他们转义.

转义符号, 都是一样的: 反斜线 `\` .

所谓特殊含义, 就是正则表达式中的含义. 非特殊含义, 就是这个符号本身.

<!--more-->

### 正则表达式

#### 元字符

|||
|:--------|:-----------|
| . | 匹配任意单个字符 |
| [] | 匹配[]指定范围内的任意单个字符 |
| [^] | 匹配[]指定范围外的任意单个字符 |

#### 字符集

|||
|-|-|
| [:digit:] | 数字 |
| [:lower:] | 小写字母 |
| [:upper:] | 大写字母 |
| [:punct:] | 标点符号 |
| [:space:] | 空白字符 |
| [:alpha:] | 所有字母 |
| [:alnum:] | 所有字母和数字, 非标点 |

**例1.1** 

	$ grep '[[:punct:]]' grep_example  # 所有有标点的行

![](http://p2.qhimg.com/t01e24acc78b267c1db.png)

#### 匹配次数

|||
|-|-|
| .* | 匹配任意字符串 |
| * | 前面的字符重复任意次数, 0次1次或多次 |
| + | 前面的字符重复1次或多次|
| ? | 前面的字符重复0次或者多次|
| {m, n} | 匹配前面的字符至少m次, 至多n次 |
| {m, } | 匹配前面的字符至少m次|

#### 位置锚定

|||
| ^ | ^后面的字符必须出现在行首 |
| $ | $前面的字符必须出现在行尾 |
| ^$ | 空白行 |
| < | <后面的字符必须出现在词首 |
| > | >前面的字符必须出现在词尾 |
| () | ()中的当做一个整体匹配 | 

**例1.2** 

	$ grep -E '.*(d+)' grep_example
	$ grep -E '.*(d+)$' grep_example

![](http://p7.qhimg.com/t01376509bbf9bd0068.png)


## sed

### 概述

sed (意为流编辑器, 源自英语"Stream Editor"的缩写)

sed用来把文档或字符串里面的文字经过一系列编辑命令转换为另一种格式输出.

sed通常用来匹配一个或多个正则表达式的文本进行处理.

**模式空间**

模式空间: 把当前处理的行存储在临时的缓冲区, 称为"模式空间"(Pattern Space)

模式空间就是读入行所在的缓存，sed对文本行进行的处理都是在这个缓存中进行的。

### 使用格式

	sed 'Address Command' [FILE...]

### 常用选项

* -n 不显示模式空间中的内容
* -i 直接修改原文件
* -r 使用扩展的正则表达式

### Address

	1. StartLine, EndLine: 
	# 开始行, 结束行
	# 如: 1, 100 表示 1 到 100 行

	2. /RegExp/ 
	# 正则表达式

	3. /pattern1/, /pattern2/
	# 从匹配到/pattern1/开始, 一直到匹配到/pattern2/结束中间所有的行

	4. LineNumber 
	# 指定的某一行

	5. StartLine, +N
	# 从StartLine开始后面N行

### Command

1. d   删除满足模式的行
2. p   显示满足模式的行
3. a \string  在满足模式的行后添加新行, 内容为string串的内容
4. i \string  在满足模式的行前添加新行, 内容为string串的内容
5. r file：将指定的文件的内容添加至符合条件的行后
6. w file：将地址指定范围内的行另存至指定的文件中
7. s \pattern\string\修饰符    将满足pattern的串替换为string
   修饰符:
   \g  全局替换, 即将该行中所有匹配项替换, 默认只替换第一个匹配项
   \i  忽略大小写
   \&  引用模式匹配的整个串
   \1  表示匹配到模式中的第一个 ( 开始的串
   \2  表示匹配到模式中的第二个 ( 开始的串

**例2.1**

	#从以hezhaoyang开头的行开始匹配, 一直到以I开头的行, 显示中间所有行
	$ sed -n '/^hezhaoyang/,/^I/p' sed_example  

![](http://p1.qhimg.com/t016c7b865c3597bb7c.png)

**例2.2**

	#将sed_insert中的内容插入到以hezhaoyang开头的行后
	$ sed '/hezhaoyang/r sed_insert' sed_example

![](http://p2.qhimg.com/t01c35938e797df3318.png)

## awk

### 概述

AWK是一种处理文本文件的语言. 它将文件作为记录序列处理. 
在一般情况下，文件内容的每行都是一个记录. 每行内容都会被分割成一系列的域, 
因此, 我们可以认为一行的第一个词为第一个域,第二个词为第二个, 以此类推.
AWK程序是由一些处理特定模式的语句块构成的. AWK一次可以读取一个输入行. 
对每个输入行, AWK解释器会判断它是否符合程序中出现的各个模式, 并执行符合的模式所对应的动作。

——阿尔佛雷德·艾侯，[The A-Z of Programming Languages: AWK](http://www.computerworld.com.au/index.php/id;1726534212;pp;2)

### 使用格式

	awk [options] 'BEGIN{ } pattern{ } END{ }'

### 常用选项

* -F fs          : 指定分隔符为fs
* -v var=value   : 定义可传递给awk的变量
* -f scriptfile  : 从脚本文件中读取awk命令

### 基本结构

一个awk脚本通常由: BEGIN语句块、能够使用模式匹配的通用语句块、END语句块3部分组成, 这三个部分是可选的. 
任意一个部分都可以不出现在脚本中, 脚本通常是被单引号或双引号中.


### awk 的工作过程

   第一步: 执行BEGIN{ commands }语句块中的语句; 
   第二步: 从文件或标准输入(stdin)读取一行, 然后执行pattern{ commands }语句块, 它逐行扫描文件, 从第一行到最后一行重复这个过程, 直到文件全部被读取完毕.
   第三步: 当读至输入流末尾时, 执行END{ commands }语句块.

   BEGIN语句块在awk开始从输入流中读取行之前被执行, 这是一个可选的语句块, 比如变量初始化、打印输出表格的表头等语句通常可以写在BEGIN语句块中.
   END语句块在awk从输入流中读取完所有的行之后即被执行, 比如打印所有行的分析结果这类信息汇总都是在END语句块中完成, 它也是一个可选语句块.
   pattern语句块中的通用命令是最重要的部分, 它也是可选的. {} 相当于一个循环体, 它会对文件中的每一行进行迭代.


### 输出 print / printf

#### print

1. 当要输出多个变量时, 应该用逗号','分隔

		$ echo | awk '{ var1="v1"; var2="v2"; var3="v3"; print var1,var2,var3; }'

2. 只写 `print` 相当于 `print $0`, 如果想要输出空行, 应该写做: `print ''`
3. `$1`, `$2`, `$3` ... 分别对应该行的第1个, 第2个, 第3个字段, 以此类对

#### printf

与 print 的区别在于 printf 需要格式化输出.

|||
|-|-|
| %c     | 字符类型 |
| %d, %i | 十进制   |
| %e, %E | 科学计数法输出数值 |
| %f | 浮点数类型 |
| %g, %G | 科学计数法或浮点数格式显示 |
| %s | 字符串 |
| %u | 无符号整数 |
| %% | %本身 |

##### 修饰符

* N: 宽度
* -: 左对齐
* +: 显示数值符号

**注意**: 
printf 默认不换行, 需要使用`"\n"`

**例3.1:**

	#输出用户名和使用的shell类型
	$ head -10 /etc/passwd | awk -F: '{printf "%-15s%s\n", $1, $7}'

![](http://p3.qhimg.com/t0171ee814b38a60d04.png)

### awk内置变量

#### 记录变量

* FS: field separator  字段分隔符, 默认为空白字符
* RS: record separator 记录分隔符, 默认为换行符
* OFS: output field separator  输出字段分隔符
* ORS: output record separator 输出记录分隔符

#### 数据变量

* NR: number of input rows 处理的行数
* NF: number of fields     该行字段的个数
* ARGV: 数组, 保存awk命令行. 
  如: $ awk '{ print $0 }' a.txt b.txt中ARGV[0]=awk, ARGV[1]=a.txt
* ARGC: awk命令参数的个数
* FILENAME: 所处理文件的名称
* IGNORECASE: IGNORECASE=1时忽略大小写

		$ echo "TEST HHHH" | awk '{IGNORECASE=1; if($1 == "test"){print "ok";}}'

### 赋值

	$ awk -v var="value" '{print var}'
	$ awk 'BEGIN{ var="value"; print var; }'

### 操作符

	+ - * / ^ ** %

### 比较运算符

		< <= >= > >= == !=
		x~y
		# True if the string x matches the regexp denoted by y
    	# 如果字符串x被正则表达式y匹配, 则为真
		x!~y

### 数组

#### 定义

	arr[1] = 'first'
	arr[2] = 'second'
	arr['year'] = 2016
	arr['month'] = 8

#### 长度

	length() 获取字符串，数组的长度
	split()  分割数组， 返回长度
	asort()  排序数组， 返回长度

**需要注意**：

判断数组是否包含某键值

	#错误写法
	$ echo | awk 'BEGIN{a["a"]=1;a["b"]=2;  if(a["c"]!=1){print "ok"}; for(i in a) {print i, a[i];} }'

	#正确判断
	$ echo | awk 'BEGIN{a["a"]=1;a["b"]=2;  if("c" in a){print "ok"}; for(i in a) {print i, a[i];} }'

awk 中的数组是关联数组, 只要是通过他的数组引用过key, 就会自动创建该序列.

#### 多维数组

awk的多维数组本质上是一维数组, 更确切地说, awk是不支持以为多维数组的. 
但是awk提供了逻辑上模拟多维数组的方式. 在awk中, `array[1, 2]=1` , 这样的写法是允许的.
他是使用了一个特殊字符串`SUBSEP(\034)`分隔键, **[2, 4]**实际上是**[2SUBSEP4]**.

访问可以使用`for( (i, j) in array ){}`但是必须用圆括号.
但是单独访问时, 必须使用split()将数组键值分开.
即: 

	arr[1,1] = 11;
	arr[1,2] = 12;
	arr[2,1] = 21;
	arr[2,2] = 22;
	for (a in arr) {
	  split(a, tmp, SUBSEP);
	  print tmp[1], tmp[2];z
	}


### 控制语句

	1. if-else 
	if (condition) {} else {}
	
	2. while
	while (condition) {}
	
	3. do-while
	do {} while (condition)
	
	4. for 
	
	for ( ; ; ) {}
	for (item in array) {}
	
	5. switch
	switch (expression) { case x: ; case y: ;}
	
	6. break, continue
	
	7. next # 提前结束本行 

### 内置函数

	split(string, array [,fieldsep])
	# 使用fieldsep分隔string, 并存储在名为array的数组中

	#实例:
	$ awk '{ split( "20:18:00", time, ":" ); print time[2] }'
	# 上例把时间按冒号分割到time数组内，并显示第二个数组元素18。
	# 返回值为分割数组的长度

	length(string)

	substr(string, start [, length])

	system(command) : 执行命令并返回给awk

	systime() : 返回当前时间时间戳

	tolower(string) 

	toupper(string)

	match(string, regular expression)
	# match函数返回在字符串中正则表达式位置的索引,
	# 如果找不到指定的正则表达式则返回0, 找到返回1。
	# match函数会设置内建变量RSTART为字符串中子字符串的开始位置, RLENGTH为到子字符串末尾的字符个数.
	# substr可利于这些变量来截取字符串.

	#实例：

	$ awk '{start=match("this is a test",/[a-z]+$/); print start}'
	$ awk '{start=match("this is a test",/[a-z]+$/); print start, RSTART, RLENGTH }'
	# 第一个实例打印以连续小写字符结尾的开始位置, 这里是11.
	# 第二个实例还打印RSTART和RLENGTH变量, 这里是11(start), 11(RSTART), 4(RLENGTH).


### 内建数学函数

	1. cos(x)  余弦函数
	2. exp(x)  求幂
	3. int(x)  取整
	4. log(x)  自然对数, 过程没有舍入
	5. rand()  产生一个大于等于0而小于1的随机数
	6. sin(x)  正弦
	7. sqrt(x) 平方根
	8. srand(x)  x是rand()函数的种子


### 自定义函数

    function F_NAME([variable])
    {
         statements
    }
    # 函数还可以使用return语句返回值，格式为“return value”
 