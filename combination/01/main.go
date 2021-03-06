package main

import "fmt"

/********************************
抽奖系统。需要依次从 100 个人中，抽取三等奖 10 名，二等奖 3 名和一等奖 1 名。请列出所有可能的组合，需要注意的每人最多只能被抽中 1 次。
排列和组合，组合没有去掉了顺序要求，解题思路应该和排列是一致的，但需要保存中奖的组合，进行去重
思路简化分析一：简单6人参与抽奖，3个人3等奖，1个人二等奖，1个人1等奖
1、把大问题分为小问题：一等奖所有组合、二等奖所有组合、三等奖所有组合，三个得奖数据中中奖人不能重复。
2、公式
3、终止条件

*/

func main() {
	var people = []int{1, 2, 3, 4, 5, 6}
	var recursion func(leftNums, tmpNums []int)
	recursion = func(leftNums, tmpNums []int) {
		if len(tmpNums) == 5 {
			fmt.Printf("tmpNums:%v\r\n", tmpNums)
			return
		}
		for _, pv := range people {
			newTmpRes := append(tmpNums, pv)
			newLeftNums := make([]int, len(leftNums)-1)
			count := 0
			for _, pv1 := range people {
				if pv1 != pv {
					newLeftNums[count] = pv1
					count++
				}
			}
			recursion(leftNums, newTmpRes)
		}
	}
	recursion(people, make([]int, 0))
}
