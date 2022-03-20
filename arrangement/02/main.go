/**
 * @Author: tonychen@sqqmall.com
 * @Description: 题目，有10个台阶，一个人可以一步一步走，也可以两步两步走，一共有多少中走法？
 * 解法串说一：
 * 解题思路：思路重点第一步，大问题拆分为小问题，
 * 1、大问题分为小问题：10个台阶所有走法=9个台阶所有走法+8个台阶所有走法——————这样的说话不容易理解？
 * 2、递归公式：f(n)=f(n-1) + f(n-2), n表示10
 * 3、终止条件:f(1)=1 ;f(2)=2;
 ***************************************************************************************
 * 解法串说二：
 * 如果你上10个台阶，可以分解成下面两种情况：
 * ● 上9个台阶，最后上1个台阶。假设这种情况下，上前面9个台阶的方法数为m。
 * ● 上8个台阶，最后上2个台阶。假设这种情况下，上前面8个台阶的方法数为n。
 * 所以，上10个台阶的方法数，其实就是 m + n。
 * 一般人看到这里就会有点蒙蔽了，为什么？
 * 答：按照最后一步来分走法的话，可分为：最后走两步和最后走一步两种情况。如果知道这两类走法的总和就是最终解。
 * 以下为递归公式，应该用** 归纳法 **来证明该公式。
 * f(n) = f(n - 1) + f(n - 2),   n > 2
 * f(n) = 2 ,   n = 2
 * f(n) = 1 ,   n = 1
 * @File:
 * @Version: 1.0.0
 * @Date: 2022年3月20日
 */

package main

import "fmt"

func main() {
	var maxStep = 7 // 最大台阶数
	var recursion func(remainStep int) int
	recursion = func(remainStep int) int {
		//终止条件1
		if remainStep == 1 {
			return 1
		}
		//终止条件2
		if remainStep == 2 {
			return 2
		}
		return recursion(remainStep-1) + recursion(remainStep-2)
	}
	fmt.Println(recursion(maxStep))
}
