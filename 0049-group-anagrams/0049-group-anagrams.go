func groupAnagrams(strs []string) [][]string {
    /*
    count the frequency
    or use sort?
    */
    group:=make(map[[26]int][]string)
    for _, s := range strs{
        var count [26]int
        for _, ch :=range s{
            count[ch-'a']++
        }
        group[count]=append(group[count],s)
    }

    result:=[][]string{}
    for _, value := range group{
        result=append(result,value)
    }
    return result
}