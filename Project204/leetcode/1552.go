package leetcode

import (
	"sort"
)

type sortarray []int

func (this sortarray) Len() int {
	return len(this)
}

func (this sortarray) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this sortarray) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

/*
In universe Earth C-137, Rick discovered a special form of magnetic force between two balls if they are put in his new invented basket. Rick has n empty baskets, the ith basket is at position[i], Morty has m balls and needs to distribute the balls into the baskets such that the minimum magnetic force between any two balls is maximum.

Rick stated that magnetic force between two different balls at positions x and y is |x - y|.

Given the integer array position and the integer m. Return the required force.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/magnetic-force-between-two-balls
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func MaxDistance(position []int, m int) int {
	sort.Sort(sortarray(position))
	var minNum, maxNum int = position[0], position[len(position)-1]
	//fmt.Println(position)
	var x, y, last int = 1, int((int64(maxNum - minNum)) / (int64(m - 1))), 1
	//fmt.Println(x,y)
	for x <= y {
		tmp := (x + y) / 2
		if testNumber(position, m, tmp) {
			x = tmp + 1
			last = tmp
		} else {
			y = tmp - 1
		}
	}
	return last
}

func testNumber(position []int, m int, l int) bool {
	m--
	var former int = position[0]
	//fmt.Println(position)
	for i := 1; i < len(position) && m > 0; i++ {
		if position[i]-former >= l {
			m--
			former = position[i]
		}
	}
	return m == 0
}

/*思路 在给定区间的整数域里二分查找*/
