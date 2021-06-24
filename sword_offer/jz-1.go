package sword_offer
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix)==0 {
		return false
	}

	a,b:=len(matrix),len(matrix[0])
	for i:=0;i<a;i++{
		for j:=0;j<b;j++{
			if matrix[i][j]==target{
				return true
			}
			if matrix[i][j]>target{
				break
			}
		}
	}
	return false
}