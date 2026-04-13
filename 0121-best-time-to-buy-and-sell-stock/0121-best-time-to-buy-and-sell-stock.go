func maxProfit(prices []int) int {
    profit:=0
    buy:=prices[0] // recording the lowest price
    for i:=1;i<len(prices);i++{
        if prices[i]<buy{
            buy=prices[i]
        }else if prices[i]-buy>profit{
            profit=prices[i]-buy
        }
    }
    return profit
}