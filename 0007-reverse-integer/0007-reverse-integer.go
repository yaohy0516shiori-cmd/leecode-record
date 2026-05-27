func reverse(x int) int {
    const MaxInt32 = 1<<31 - 1
    const MinInt32 = -1 << 31

    res := 0

    for x != 0 {
        // go 可以直接处理负数取值
        digit := x % 10
        x /= 10

        // 正数溢出检查
        if res > MaxInt32/10 || (res == MaxInt32/10 && digit > 7) {
            return 0
        }

        // 负数溢出检查
        if res < MinInt32/10 || (res == MinInt32/10 && digit < -8) {
            return 0
        }

        res = res*10 + digit
    }

    return res
}