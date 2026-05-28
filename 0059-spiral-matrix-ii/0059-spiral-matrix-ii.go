func generateMatrix(n int) [][]int {
	up := 0
	left := 0
	right := n - 1
	bot := n - 1

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	k := 1

	for k <= n*n {
		// 1. 上边：从左到右
		for j := left; j <= right && k <= n*n; j++ {
			res[up][j] = k
			k++
		}
		up++

		// 2. 右边：从上到下
		for i := up; i <= bot && k <= n*n; i++ {
			res[i][right] = k
			k++
		}
		right--

		// 3. 下边：从右到左
		for j := right; j >= left && k <= n*n; j-- {
			res[bot][j] = k
			k++
		}
		bot--

		// 4. 左边：从下到上
		for i := bot; i >= up && k <= n*n; i-- {
			res[i][left] = k
			k++
		}
		left++
	}

	return res
}