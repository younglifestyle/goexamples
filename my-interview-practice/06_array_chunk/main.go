package main

import (
	"fmt"
)

// 将传入的数组，按照分块数进行切分
func arrayChunk(arr []int, chunk int) (res [][]int) {

	arrTmp := []int{}
	index := 0
	data := 0

	for index, data = range arr {
		arrTmp = append(arrTmp, data)
		if (index+1)%chunk == 0 {
			res = append(res, arrTmp)
			arrTmp = []int{}
		}
	}

	if (index+1)%chunk != 0 {
		res = append(res, arrTmp)
	}

	return
}

// solution_better 先给整个的group添加array，再像 array 中添加item
func arrayChunk2(intArray []int, chunk int) (result [][]int) {
	for i, num := range intArray {
		// 0余其他数都为0
		if i%chunk == 0 /*|| len(result) == 0*/ { // 只有当result为空，或者是达到要增加array的上限时，才增加一个array
			result = append(result, []int{})
		}
		result[len(result)-1] = append(result[len(result)-1], num) // 给 result 的最后一个array添加 item
	}
	return
}

// 769. Max Chunks To Make Sorted  https://leetcode.com/problems/max-chunks-to-make-sorted/
// 题目意思：划分数组，分开进行排序，最多能分多少个，输入[4,3,2,1,0]，至多分一个
// 输入[1,0,2,3,4], 能分出[1, 0], [2], [3], [4]
// 注意：分开是将元素按位置分开，然后合并能是已排序好的。
func maxChunksToSorted1(arr []int) int {
	var res int
	m := 0
	for i := 0; i < len(arr); i++ {
		m = max(m, arr[i])
		if m == i {
			res++

			fmt.Println("qiege", m, i)
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	//[1,2,0,3]
	arr := []int{0, 2, 1}

	res := arrayChunk(arr, 2)
	fmt.Println(res)

	res1 := arrayChunk2(arr, 2)
	fmt.Println(res1)

	fmt.Println(maxChunksToSorted(arr))
	fmt.Println(maxChunksToSorted1(arr))
}
