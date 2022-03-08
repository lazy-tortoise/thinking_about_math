/**
 * @Author: tonychen@sqqmall.com
 * @Description: 田忌赛马问题，考察两个方面，全排列：https://www.cnblogs.com/gyyyl/p/13967517.html
 * @Problem: 心得一：map 这个数据结构 生成遍历每次输出的顺序都不一致
 * 心得二：递归 第一步先要把大问题分解为小问题，第二步形成公式 ，第三步 终止条件
 * @Version: 1.0.0
 * @Date: 2022/3/6 18:44
 */

package main

import (
	"fmt"
)

func GetQHorses() []int {
	var q []int //马匹号 与全程需要时间
	q = []int{1, 3, 5}
	return q
}

func GetTHorses() []int {
	var t []int
	t = []int{2, 4, 6}
	return t
}

//compare 齐王匹战田忌当前排序马匹结果
func compare(q []int, t []int) {
	fmt.Printf("齐王对战田忌,田忌马匹排序:%v,齐王马匹排序:%v\r\n",
		t, q)
	var tWinCnt = 0
	for k, v := range t {
		if v < q[k] {
			tWinCnt++
		}
	}

	if tWinCnt > len(t)/2 {
		fmt.Printf("齐王对战田忌,结果<<<<<田忌>>>>>win\r\n")
	} else {
		fmt.Printf("齐王对战田忌,结果齐王win\r\n")
	}
}

func permute(nums []int, qHorse []int) {
	tmp := make([]int, 0)
	var recursion func(leftNums, tmpRes []int)
	recursion = func(leftNums, tmpRes []int) {
		if len(leftNums) == 0 {
			compare(qHorse, tmpRes)
			return
		}
		for index, value := range leftNums {
			newTmpRes := append(tmpRes, value)
			newLeftNums := make([]int, len(leftNums)-1)
			count := 0
			for i, v := range leftNums {
				if i != index {
					newLeftNums[count] = v
					count++
				}
			}
			// newLeftNums := append(leftNums[:index], leftNums[index+1:]...)
			// 不使用上面这行代码的原因是append会改变leftNums底层数组的值
			recursion(newLeftNums, newTmpRes)
		}
	}
	recursion(nums, tmp)
}

func main() {
	permute(GetTHorses(), GetQHorses())
}
