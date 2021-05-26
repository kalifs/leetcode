func convert(s string, numRows int) string {
    lines := make([]string,numRows)
    size := len(s)
    rowMod := numRows - 1
    firstRowMod := rowMod * 2
    if rowMod == 0 {
        return s
    }
    for i:=0;i<size;i++{
        letter := s[i:(i+1)]
        rem := i % firstRowMod
        if rem <= rowMod {
            lines[rem] += letter
        } else {
            lines[firstRowMod - rem] += letter
        }
    }
    return strings.Join(lines, "")
}
