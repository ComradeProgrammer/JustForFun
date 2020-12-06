package leetcode

/*
Given an integer array arr of distinct integers and an integer k.

A game will be played between the first two elements of the array (i.e. arr[0] and arr[1]). In each round of the game, we compare arr[0] with arr[1], the larger integer wins and remains at position 0 and the smaller integer moves to the end of the array. The game ends when an integer wins k consecutive rounds.

Return the integer which will win the game.

It is guaranteed that there will be a winner of the game.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-the-winner-of-an-array-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func getWinner1535(arr []int, k int) int {
	var count int = 0
	for count < k {
		if count > len(arr)-1 {
			return arr[0]
		}
		if arr[0] >= arr[1] {
			count++
			arr = append(arr, arr[1])
			arr[1] = arr[0]
			arr = arr[1:]
		} else {
			count = 1
			arr = append(arr, arr[0])
			arr = arr[1:]
		}
	}
	return arr[0]

}
