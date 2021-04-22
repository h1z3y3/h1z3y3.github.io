---
title: Golang 逃逸分析
date: 2021-04-22 13:12:00
tags: [Golang]
---

# Golang 逃逸分析

Golang 的垃圾回收机制可以进行自动内存管理让我们的代码更简洁，同时发生内存泄漏的可能性更小。
然而，GC 会定期停止并收集未使用的对象，因此还是会增加程序的开销。
Go 的编译器十分聪明，比如决定变量需要分配在堆上还是栈上，和分配在堆上不同的是在栈上的变量在声明它的函数结束之后就会被回收。
那对于 GC 来说，分配在栈上的变量不会带来额外的开销，在函数 return 之后，函数的整个调用栈都会被销毁。

那到底变量应该分配在堆上还是栈上，Golang 如何决定呢，这就要说到 Golang 的逃逸分析。
判断逃逸的基本规则是如果一个函数的返回值是本函数内声明的某个变量的引用，那么就称这个变量从这个函数中逃逸了。
作为本函数的返回值，它还能被函数外的其他程序修改，所以它必须分配在堆上，而不能分配在那个函数的栈上。

因此，如果在编译过程中，能对变量的逃逸情况作分析，可以提高我们程序的性能。
首先最大的好处就是能够减少垃圾回收的压力，没有逃逸的变量分配在栈上，函数返回就能直接回收资源；
其次逃逸分析完之后，我们能确定哪些变量其实可以分配在栈上，栈的分配比堆快，性能更好；
还有可以进行同步消除，如果定义变量的函数有同步锁，但是运行时却只有一个线程访问，那此时逃逸分析后的机器码，会去掉同步锁运行。

## 开启 Go 编译时的逃逸分析日志

在编译时，添加 `-gcflags '-m'` 参数即可查看go编译过程中详细的逃逸分析日志。
但为了不让 Go 编译时自动内联函数，会加上 `-l` 参数，最终为 `-gcflags '-m -l'`。

Example 0:

```
package main

type S struct {}

func main() {
  var x S
  _ = identity(x)
}

func identity(x S) S {
  return x
}
```

Output:

```
$ go run -gcflags '-m -l' escape.go
$
$
```

可以看到没有任何输出，我们知道 Go 在调用函数的时候，会采用按值传递，那么在 `main` 函数中声明的 `x`，
会被拷贝到 `identity()` 函数的栈中。**通常，没有引用的代码总是使用栈分配，因此没有逃逸分析日志的输出。**

那如果稍微改一下代码：

Example 1:

```
package main

type S struct {}

func main() {
  var x S
  y := &x
  _ = *identity(y)
}

func identity(z *S) *S {
  return z
}
```

Output:

```
$ go run -gcflags '-m -l' escape.go
# command-line-arguments
./escape.go:11:22: leaking param: z to result ~r1 level=0
./escape.go:7:8: main &x does not escape
```

第一行是 `z` 变量是流经某个函数的意思，仅作为函数的输入，并且直接返回，
在 `identity()` 中也没有使用到 `z` 的引用，所以变量没有逃逸。

第二行，`x` 在 `main()` 函数中声明，所以是在 `main()` 函数中的栈中的，也没有逃逸。

Example 2:

```
package main

type S struct {}

func main() {
  var x S
  _ = *ref(x)
}

func ref(z S) *S {
  return &z
}
```

Output:

```
$ go run -gcflags '-m -l' escape.go
# command-line-arguments
./escape.go:11:10: &z escapes to heap
./escape.go:10:16: moved to heap: z
```

可以看到发生了逃逸。
`ref()` 的参数 `z` 是通过值传递的，所以 `z` 是 `main()` 函数中 `x` 的一个值拷贝，
而 `ref()` 返回了 `z` 的引用，所以 `z` 不能放在`ref()`的栈中， 实际上被分配到了堆上。

实际上，我们发现 `main()` 函数中没有直接使用 `ref()` 返回的引用，这种情况其实 `z` 可以分配到`ref()`的栈上，
但是Go的逃逸分析并没有复杂到来识别出这种情况，它只看输入还有返回的变量的流程。
**值得注意的是，如果我们没有加 `-l` 参数，其实 `ref()` 会被编译器内联到 `main()` 使用。**

如果引用被赋值到结构体的成员呢？

Example 3:

```
package main

type S struct {
  M *int
}

func main() {
  var i int
  refStruct(i)
}

func refStruct(y int) (z S) {
  z.M = &y
  return z
}
```

Output:

```
$ go run -gcflags '-m -l' escape.go
# command-line-arguments
./escape.go:13:9: &y escapes to heap
./escape.go:12:26: moved to heap: y
```

可以发现，即使这个引用是结构体的一个成员，Go的逃逸分析可以跟踪引用的。
当结构体 `refStruct` 返回时，`y` 一定是从 `refStruct()` 中逃逸的。

可以再和下面例子比较一下：

Example 4:

```
package main

type S struct {
  M *int
}

func main() {
  var i int
  refStruct(&i)
}

func refStruct(y *int) (z S) {
  z.M = y
  return z
}
```

Output:

```
$ go run -gcflags '-m -l' escape.go
# command-line-arguments
./escape.go:12:27: leaking param: y to result z level=0
./escape.go:9:13: main &i does not escape
```

那这个 `y` 没有逃逸的原因是，`main()` 中带着 `i` 的引用调用了 `refStruct()` 并直接返回了，
从来没有超过 `main()` 函数的调用栈，原因和 Example 1 实际上是一样的。

另外要说明一点，Example 4 要比 Example 3 更高效：

在 Example 3 中，`i` 必须在`main()` 的栈中申请一块栈空间，经过`refStruct()`后，`y`还要在堆上再申请一块空间；
而在 Example 4 中，实际上只有`i`申请了一次空间，然后它的引用经过了 `refStruct()` 而已。

一个更复杂的例子：

Example 5:

```
package main

type S struct {
  M *int
}

func main() {
  var x S
  var i int
  ref(&i, &x)
}

func ref(y *int, z *S) {
  z.M = y
}
```

Output:

```
$ go run -gcflags '-m -l' escape.go
# command-line-arguments
./escape.go:13:21: leaking param: y
./escape.go:13:21: ref z does not escape
./escape.go:10:7: &i escapes to heap
./escape.go:9:7: moved to heap: i
./escape.go:10:11: main &x does not escape
```

`y` 和 `z` 没有逃逸很好理解，但问题在于 `y` 还被赋值到函数 `ref()` 的输入 `z` 的成员了，
而Go的逃逸分析不能跟踪变量之间的关系，不知道 `i` 变成了 `x` 的一个成员，
分析结果说 `i` 是逃逸的，但本质上 `i`是没逃逸的， 这个时候Go的逃逸分析实际上是有问题的。

这里还有好多因为Go逃逸分析的不足而导致被分配到堆的变态例子：
[https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw/preview](https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw/preview)

其实说这么多其实就是为了说明如果想要减少垃圾回收的时间，提高程序性能，
那就要尽量避免在堆上分配空间，之后在写程序的时候可以多考虑一下这方面的问题。;)

### 参考

[《Golang escape analysis》](http://www.agardner.me/golang/garbage/collection/gc/escape/analysis/2015/10/18/go-escape-analysis.html)

[《Go Escape Analysis Flaws》](https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw/preview)
