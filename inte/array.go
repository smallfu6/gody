package main

import (
	"fmt"
	"time"
)

// 由内存和性能，在函数间传递数组是个开销很大的操作
// 在函数间传递变量时，总是以值的方式传递，如果这个变量是一个数组,
// 意味着整个数组，不管有多长，都会完整复制，并传递给函数

// 创建一个包含100万个int类型元素的数组,在64位架构上, 将需要800万
// 字节， 即 8MB 的内存
// 注: golang 的 int 类型是不确定的，与具体的平台有关系
// int 在 32 位系统中是4字节，在64位系统中是8字节

// 每次 foo 被调用时， 必须在栈上分配8MB的内存，之后，整个数组的
// 值(8MB的内存) 被复制到刚分配的内存里, 虽然go底层会处理这个复制
// 操作, 但可以只传入指向数组的指针， 只需要复制8字节的数据而不是
// 8MB 的内存数据到栈上。

// 内存地址按位寻址，64位系统中字长为64，即一个内存地址的位数为64，
// 即8个字节， 传int类型数组的指针， 即数组第一个元素所占内存(8个字节)
// 地址的第一位的地址即为传入 foo1  函数的地址， 而这个地址需要用64位的
// 一个数来表示， 即8个字节。

// todo: 内存寻址相关

func main() {
	// 声明一个需要 8MB 的数组
	// 1int= 64bit = 8byte(64位平台)
	// 1000000 * 8 byte = 8000000 byte = 8MB
	var array [1e6]int

	startTime := time.Now()
	// fmt.Println("--------")
	// 将数组传递给函数 foo
	// foo(array)
	// fmt.Println(time.Since(startTime))
	// fmt.Println("--------")
	fmt.Println("--------")
	// 将数组的地址传递给函数foo1
	foo1(&array)
	fmt.Println(time.Since(startTime))
	fmt.Println("--------")
}

func foo(array [1e8]int) {
	//...
}

func foo1(array *[1e6]int) {
	//...
}

// todo: 对一个 [1e6]int 类型的变量取地址的逻辑是什么???

//  传入指针会更有效的利用内存，性能也更好，但
// 如果改变指针指向的值，会改变共享的内存(上的值?),
// 使用切片能更好的处理这类共享问题?(todo)

// foo 传数组
// 1e6
//--------
//16.343µs
//--------
//

// 1e8
//--------
//137.501092ms
//--------

// foo 传指针
// 1e6
//--------
//9.735µs
//--------
//
