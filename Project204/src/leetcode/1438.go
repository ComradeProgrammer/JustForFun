package leetcode

/*Given an array of integers nums and an integer limit,
return the size of the longest non-empty subarray such that the absolute difference between any two elements of this subarray is less than or equal to limit.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func LongestSubarray1438(nums []int, limit int) int {
	minDeque := make([]int, 1, len(nums))
	maxDeque := make([]int, 1, len(nums))
	var i, j int = 0, 1
	var res = 0
	minDeque[0], maxDeque[0] = 0, 0
	for i < j && j < len(nums) {
		for len(minDeque) > 0 && nums[minDeque[len(minDeque)-1]] > nums[j] {
			minDeque = minDeque[:len(minDeque)-1]
		}
		minDeque = append(minDeque, j)
		for len(maxDeque) > 0 && nums[maxDeque[len(maxDeque)-1]] < nums[j] {
			maxDeque = maxDeque[:len(maxDeque)-1]
		}
		maxDeque = append(maxDeque, j)
		for (i < j) && (nums[maxDeque[0]]-nums[minDeque[0]] > limit) {
			i++
			if maxDeque[0] < i {
				maxDeque = maxDeque[1:]
			}
			if minDeque[0] < i {
				minDeque = minDeque[1:]
			}
		}
		//fmt.Println(j,i,j-1)
		if (j - i) > res {
			res = j - i
		}
		j++
	}
	return res
}

/*
这道题目的思路是滑动窗口+双端队列维护区间内的最大值和最小值，非常经典
滑动窗口是因为如下原因：
	假设现在窗口是A-B是一个可行解，而A-B+1不再可行，那么对于A开头的数组我们已经找出了最优解
	对于A+1来说，如果A+1-B不是可行的，那么A+1开头的最优解的长度一定没有A-B大所以舍弃
	如果A+1-B是可行的那么继续向右移动B寻找最优解

双端队列是可以以O(n)维护一个滑动区间最大值最小值的方法 以最大值为例具体方法如下
	队列里面存放的是[当前最大值的下标，最大值右边的次大值下标，次大值右边的（！！！）次次大值下标.....]
	当有新的元素被从右侧纳入的时候，如果队尾对于的值小于要加入的元素，就吐掉队尾（因为队尾不是次大值了），一直吐，吐到队尾比要加入的元素大为止
	然后把新元素下标放进去

	如果有元素从左侧被排除，就检查队首元素是不是已经被排除了，如果是就吐掉（此时次大值也就是目前最大值自动顶上）
	神奇吧
*/
