package leetcode
/*A majority element is an element that makes up more than half of the items in an array.
Given a positive integers array, find the majority element.
If there is no majority element, return -1. Do this in O(N) time and O(1) space.

Example 1:
Input: [1,2,5,9,5,9,5,5,5]
Output: 5

Example 2:
Input: [3,2]
Output: -1

Example 3:
Input: [2,2,1,1,1,2,2]
Output: 2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-majority-element-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func majorityElement1710(nums []int) int {
	if len(nums)==0{
		return -1
	}
	var major,cnt int=nums[0],1
	for i:=1;i<len(nums);i++{
		if nums[i]==major{
			cnt++
		}else{
			cnt--
			if cnt<0{
				major=nums[i]
				cnt=1
			}
		}
	}
	if cnt<=0{
		return -1
	}
	var check int=0
	for i:=0;i<len(nums);i++{
		if major==nums[i]{
			check++
		}
	}
	if 2*check>len(nums){
		return major
	}
	return -1
}