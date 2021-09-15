package leet

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// LeetMain
func LeetMain() {
	/*
		1.求和
		求切片内两个数只和等于目标数并且返回下标切片
	*/
	fmt.Println(twoSum2([]int{2, 3, 8, 7}, 9))

	/*
		2.求和
		给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

		请你将两个数相加，并以相同形式返回一个表示和的链表。

		你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
	*/
	addTwoNumbers1(
		&ListNode{
			Val: 1, Next: &ListNode{
				Val: 0, Next: &ListNode{
					Val: 3, Next: &ListNode{
						Val: 7,
					},
				},
			},
		}, &ListNode{
			Val: 4, Next: &ListNode{
				Val: 7, Next: &ListNode{
					Val: 9,
				},
			},
		},
	)
	/*
		3.给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数
		//[2,3]
	*/
	fmt.Println(
		findMedianSortedArrays(
			[]int{1000, 1000, 3000, 4000, 5000, 5000, 5002, 5003, 5009, 5100, 11000, 12000}, []int{
				1000, 2000, 2000, 2001, 2200, 5000, 6000, 7000, 9000, 65000, 78000, 990000, 202000,
			},
		),
	)
	fmt.Println(
		findAveragArrays(
			[]int{1000, 1000, 3000, 4000, 5000, 5000, 5002, 5003, 5009, 5100, 11000, 12000}, []int{
				1000, 2000, 2000, 2001, 2200, 5000, 6000, 7000, 9000, 65000, 78000, 990000, 202000,
			},
		),
	)
	/*
		4.给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
		如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
		假设环境不允许存储 64 位整数（有符号或无符号）
	*/
	fmt.Println(reverse(2147483647))

	/*
		11.给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
		说明：你不能倾斜容器。
	*/
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

}

// 迭代比较每两个点的面积
func maxArea(height []int) int {
	// 初始化面积 宽*长
	area := 0
	for i := 0; i <= len(height)-1; i++ {
		for j := i + 1; j <= len(height)-1; j++ {
			width := 0
			long := j - i
			// 选择一个更低的 当心漏水
			if height[j] < height[i] {
				width = height[j]
			} else {
				width = height[i]
			}
			tmpArea := width * long
			// 去最大的面积
			if tmpArea > area {
				area = tmpArea
			}
		}
	}
	return area
}

// 4321 ==>1234
func reverse(x int32) int32 {
	var initB bool
	if x > 0 {
		initB = true
	}
	// 创建中间变量
	tmp := int32(0)
	// 当x==0时候退出循环
	// 循环1 十位 循环2 百位 循环3千位(每次都是上一次的值乘10+当前取余下来的数)
	for x != 0 {
		// 通过去目标数值10的余数得到倒序的每一位
		remainder := x % 10
		tmp = tmp*10 + remainder
		// 每次循环将x赋值为 4321 4321/10==432/10==43/10==4/10==0
		x = x / 10
	}
	// 判断是够超出精度
	if initB {
		if tmp < 0 {
			return 0
		}
	} else {
		if tmp > 0 {
			return 0
		}
	}
	return tmp
}

// 平均数
func findAveragArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	// sort.Ints(nums1)
	// sort.Ints(nums2)
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)
	num := 0
	for _, v := range nums1 {
		num += v
	}

	return float64(num / len(nums1))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	// sort.Ints(nums1)
	// sort.Ints(nums2)
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)
	indexI := 0
	indexJ := 0
	business := len(nums1) / 2
	indexI = business
	if len(nums1)%2 == 0 {
		indexJ = business - 1
	} else {
		return float64(nums1[indexI])
	}
	return (float64(nums1[indexI]) + float64(nums1[indexJ])) / 2
}

// 求切片内两个数只和等于目标数并且返回下标切片
func twoSum1(nums []int, target int) []int {
	result := make([]int, 0)
	for i, num := range nums {
		for i2 := len(nums) - 1; i2 >= 0; i2-- {
			if num+nums[i2] == target {
				result = append(result, i, i2)
			}
		}
	}
	return result
}

// 求切片内两个数只和等于目标数并且返回下标切片
func twoSum2(nums []int, target int) []int {
	result := make([]int, 0)
	tmpMap := make(map[int]int)
	for i, num := range nums {
		if v, ok := tmpMap[num]; ok {
			result = append(result, v, i)
			return result
		}
		tmpMap[target-num] = i
	}
	return result
}

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
[5,6,4]
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	num1s := make([]int, 0)
	num1s = append(num1s, l1.Val)
	for l1.Next != nil {
		num1s = append(num1s, l1.Next.Val)
		l1 = l1.Next
	}

	num2s := make([]int, 0)
	num2s = append(num2s, l2.Val)
	for l2.Next != nil {
		num2s = append(num2s, l2.Next.Val)
		l2 = l2.Next
	}
	num1tmp := int(0)
	for i := len(num1s) - 1; i >= 0; i-- {
		num1tmp = num1tmp + num1s[i]*int(math.Pow10(i))
	}
	fmt.Println(num1tmp)
	num2tmp := int(0)
	for i := len(num2s) - 1; i >= 0; i-- {
		num2tmp = num2tmp + num2s[i]*int(math.Pow10(i))
	}
	fmt.Println(num2tmp)
	numResult := num1tmp + num2tmp
	fmt.Println(numResult)
	numResultStr := strconv.FormatInt(int64(numResult), 10)
	strings := make([]string, 0)
	for _, v := range numResultStr {
		strings = append(strings, string(rune(v)))
	}
	fmt.Println(strings)

	listNode := ListNode{}

	nodeResult := ll(strings, &listNode)
	return nodeResult
}

func ll(nums []string, head *ListNode) *ListNode {
	fnode := head
	for i := len(nums) - 1; i >= 0; i-- {
		parseInt, _ := strconv.ParseInt(nums[i], 10, 64)
		temp := ListNode{Val: int(parseInt)}
		head.Next = &temp
		head = &temp
	}
	return fnode.Next
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	// 初始化初始值
	nodeRoot := &ListNode{
		Val: 0,
	}
	// 游标: cursor 十
	var nodeCursor *ListNode
	nodeCursor = nodeRoot
	// 十位
	tenNum := 0
	for l1 != nil || l2 != nil || tenNum > 0 {
		// 获取第一节点值
		var l1Val int
		if l1 != nil {
			l1Val = l1.Val
		}
		var l2Val int
		if l2 != nil {
			l2Val = l2.Val
		}
		// 计算两两值
		sumNum := l1Val + l2Val + tenNum
		// 判断当前值是够需要进位
		tenNum = sumNum / 10
		// 获取进位的数字 如果需要进位获取当前取余的值
		carryNum := sumNum % 10
		// 创建中间node变量
		nodeTmp := ListNode{
			Val: carryNum,
		}
		// 指向游标的下一个节点为以上
		nodeCursor.Next = &nodeTmp
		// 重置游标为当前节点
		nodeCursor = &nodeTmp
		// 重新初始化下个要检查的数据l1 l2
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return nodeRoot.Next
}
