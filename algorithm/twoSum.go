package main

import "fmt"

/*
给定一个整数数组 nums 和一个目标值 target，
请你在该数组中找出和为目标值的那 两个 整数，
并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

func main (){
	var nums = [...]int{2,7,11,15}
	var keys []int
	target := 13
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			sum := nums[i] + nums[j]
			if(sum == target){
				keys = append(keys,i)
				keys = append(keys,j)
				fmt.Println(keys)
				break
			}
		}
		if(len(keys)>0){
			break
		}

	}
}