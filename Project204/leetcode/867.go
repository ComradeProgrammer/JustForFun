package leetcode

/*867 easy*/
func transpose867(A [][]int) [][]int {
	var row,col int=len(A),len(A[0])
	var tmp[][]int=make([][]int,col)
	for i:=0;i<col;i++{
		tmp[i]=make([]int,row)
	}
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			tmp[j][i]=A[i][j];
		}

	}
	return tmp
}
