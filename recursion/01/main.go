/**
 * @Author: tonychen@sqqmall.com
 * @Description:
 * @File:最大公约数
 * @Version: 1.0.0
 * @Date: 2022/3/26 15:21
 */

package main

import "fmt"

func main() {
	var m = 138
	var n = 420

	var recursion func(m, n int) int
	recursion = func(m int, n int) int {
		if n == 0 {
			fmt.Printf("最大公约数:%v\r\n", m)
			return m
		}
		fmt.Printf("m:%v,n:%v\r\n", m, n)
		return recursion(n, m%n) //没有多余的操作，是尾递归 https://www.zhihu.com/question/20761771
	}
	recursion(m, n)
}
