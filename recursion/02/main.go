/**
 * @Author: tonychen@sqqmall.com
 * @Description:计算2的几次方
 * @File:是否可以改为尾递归
 * @Version: 1.0.0
 * @Date: 2022/3/26 15:44
 */

package main

import (
	"fmt"
)

func main() {
	//start := time.Nanoseconds()
	fmt.Printf("last %v\r\n", foo(50))
	//elapsed := time.Since(t)
	//fmt.Println("app elapsed:", elapsed)
}

func foo(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Printf("n:%v\r\n", n)
	return 2 * foo(n-1)
}

//fooUp TODO改进版
func fooUp(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Printf("n:%v\r\n", n)

	return 2 * foo(n-1)
}
