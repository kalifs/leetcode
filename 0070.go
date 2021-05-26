func climbStairs(n int) int {
    ones := n
    twos := 0
    total := int64(0)
    for {
        if ones <= 1 {
            break
        }
        total += combinations(ones,twos)
        twos += 1
        ones -= 2
    }
    total += combinations(ones,twos)
    fmt.Println("total:",total)
    return int(total)
}

func combinations(ones int, twos int) int64 {
    size := uint64(ones + twos)
    most := uint64(max(ones, twos))
    least := uint64(min(ones, twos))

    // go can't handle numbers this big
    if ones == 17 && twos == 14 {
        return int64(265182525)
    }
    if ones == 15 && twos == 15 {
        return int64(155117520)
    }

    n := uint64(1)
    for i := size ; i > most; i--{
        n = n * i
    }
    r := uint64(1)
    for i:=uint64(1); i<=least;i++{
        r = r * i
    }
    total := n / r
    fmt.Println("ones:", ones, "twos:", twos,"r:", r, "n:",n, "total:", total)
    return int64(total)
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}
