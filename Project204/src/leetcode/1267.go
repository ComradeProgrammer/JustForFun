package leetcode

/*
You are given a map of a server center, represented as a m * n integer matrix grid, where 1 means that on that cell there is a server and 0 means that it is no server. Two servers are said to communicate if they are on the same row or on the same column.

Return the number of servers that communicate with any other server.

Input: grid = [[1,0],[0,1]]
Output: 0
Explanation: No servers can communicate with others.

Input: grid = [[1,0],[1,1]]
Output: 3
Explanation: All three servers can communicate with at least one other server.

Input: grid = [[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]
Output: 4
Explanation: The two servers in the first row can communicate with each other. The two servers in the third column can communicate with each other. The server at right bottom corner can't communicate with any other server.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-servers-that-communicate
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func countServers1267(grid [][]int) int {
	var m,n int=len(grid),len(grid[0])
	var row=make([]int,m)
	var col=make([]int,n)
	var total int=0
	for i:=0;i<m;i++{
		sum:=0
		for j:=0;j<n;j++{
			sum+=grid[i][j]
			total+=grid[i][j]
		}
		row[i]=sum
	}
	for j:=0;j<n;j++{
		sum:=0
		for i:=0;i<m;i++{
			sum+=grid[i][j]
		}
		col[j]=sum
	}

	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if row[i]==1&&col[j]==1&&grid[i][j]==1{
				total--;
			}
		}
	}
	return total

}
