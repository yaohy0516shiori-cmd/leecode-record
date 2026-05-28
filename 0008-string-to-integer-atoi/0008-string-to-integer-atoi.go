func myAtoi(s string) int {
	const INT_MAX = 1<<31 - 1
	const INT_MIN = -1 << 31

	n := len(s)
	i := 0

	// 1. 跳过前导空格
	for i < n && s[i] == ' ' {
		i++
	}

	// 2. 处理符号
	sign := 1
	if i < n && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// 3. 读取连续数字
	res := 0

	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// 4. 提前判断溢出
		if sign == 1 {
			if res > INT_MAX/10 || (res == INT_MAX/10 && digit > 7) {
				return INT_MAX
			}
		} else {
			if -res < INT_MIN/10 || (-res == INT_MIN/10 && digit > 8) {
				return INT_MIN
			}
		}

		res = res*10 + digit
		i++
	}

	return sign * res
}