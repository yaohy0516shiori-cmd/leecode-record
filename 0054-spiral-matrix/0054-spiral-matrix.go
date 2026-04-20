func spiralOrder(matrix [][]int) []int {
    up,down,left,right:=0,len(matrix)-1,0,len(matrix[0])-1
    res:=[]int{}
    for up<=down && left<=right{
        res=rightwards( matrix, left, up, right, res)
        up++
        res=downwards( matrix, up, right, down, res)
        right--
        if up<=down{
            res=leftwards(matrix, right, down, left, res)
            down--
        }
        if left<=right{
            res=upwards(matrix, down, left, up, res)
            left++
        }

    }
    return res
}

func rightwards(matrix [][] int, start int, row , boundary int, res []int) []int{
    for i:=start; i<= boundary; i++{
        res=append(res,matrix[row][i])
    }
    return res
}

func leftwards(matrix [][] int, start int, row , boundary int, res []int) []int{
    for i:=start; i>=boundary; i--{
        res=append(res,matrix[row][i])
    }
    return res
}

func downwards(matrix [][] int, start int, col , boundary int, res []int) []int{
    for i:=start; i<=boundary; i++{
        res=append(res,matrix[i][col])
    }
    return res
}

func upwards(matrix [][] int, start int, col , boundary int, res []int) []int{
    for i:=start; i>=boundary; i--{
        res=append(res,matrix[i][col])
    }
    return res
}
